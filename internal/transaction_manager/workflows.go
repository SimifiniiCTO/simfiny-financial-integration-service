package transactionmanager

import (
	"github.com/pborman/uuid"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// DeleteBankAccountWorkflow implements WorkflowManager
// The bulk of the deletion logic occurs asynchrnously. The transaction manager witholds
// a set of workers who process jobs from a redis queue. The workflow is responsible for
// scheduling the deletion of the bank account and all associated records. The asynq library process (TO BE IMPLEMENTED)
// is responsible for performing the actual deletion of the bank account and all associated records as a background job.
// we do this because we don't want to block the server thread who's invoking this workflow from performing other operations while the deletion is in progress.
// additionally, deletion can fail for a number of reasons, leaving the backend with half deleted records.
//
// Queue persistence enable us to retry the deletion of the bank account and all associated records in the event of a failure.
// ref: https://github.com/hibiken/asynq (redis based task queue)
// ref: https://github.com/temporalio/temporal-ecommerce (temporal deletion workflow)
func (tx *TransactionManager) DeleteBankAccountWorkflow(ctx workflow.Context, bankAccountID uint64) error {
	options := workflow.ActivityOptions{
		TaskQueue:              tx.ServiceTaskQueue,
		ScheduleToCloseTimeout: tx.RPCTimeout,
		StartToCloseTimeout:    tx.RPCTimeout,
		HeartbeatTimeout:       0,
		WaitForCancellation:    false,
		ActivityID:             uuid.New(),
		RetryPolicy:            &temporal.RetryPolicy{InitialInterval: *tx.RetryPolicy.RetryInitialInterval, BackoffCoefficient: tx.RetryPolicy.RetryBackoffCoefficient, MaximumInterval: tx.RetryPolicy.MaximumInterval, MaximumAttempts: int32(tx.RetryPolicy.MaximumAttempts), NonRetryableErrorTypes: []string{}},
		DisableEagerExecution:  false,
	}

	ctx = workflow.WithActivityOptions(ctx, options)
	// delete bank account and all referenced associated to it
	// NOTE: this is performing deletion of all associated records within the postgres database
	err := workflow.ExecuteActivity(ctx, tx.DeleteBankAccountActivity, bankAccountID).Get(ctx, nil)
	if err != nil {
		return err
	}

	// TODO: delete transactions associated to the bank account in mongodb
	/*
	 *  // Dispatch this operation to the async queue process for the asynq worker pool to pick up and process
	 *  err := workflow.ExecuteActivity(ctx, tx.DeleteBankAccountTransactionsActivity, bankAccountID).Get(ctx, nil)
	 *	if err != nil {
	 *		return err
	 *	}
	 */

	// TODO: delete item via plaid if bank account is linked to a plaid item
	// ensure that the item is only deleted if there are no other accounts associated with the profile
	/*
	 * 	// Dispatch this operation to the async queue process for the asynq worker pool to pick up and process
	 *  err := workflow.ExecuteActivity(ctx, tx.DeletePlaidItemForBankAccount, profileID).Get(ctx, nil)
	 *	if err != nil {
	 *		return err
	 *	}
	 */

	return nil
}

// DeleteProfileWorkflow perform deletion operation as a DTX (distributed transaction)
//
// TODO: given there will be a lot of data to delete, we should probably
// do this in a background job. We should also perform this entire transaction as a
// distributed transaction that is atomic and consistent.
//
// DTX:
// 1. delete the user profile record and all account associations
// 2. delete the user profile from the context of stripe (background job)
// 3. delete the user profile from the context of plaid (background job)
// 4. delete the tx records (background job)
//
// ref: https://github.com/hibiken/asynq (redis based task queue)
// ref: https://github.com/temporalio/temporal-ecommerce (temporal deletion workflow)
func (tx *TransactionManager) DeleteProfileWorkflow(ctx workflow.Context, userID uint64) error {
	options := workflow.ActivityOptions{
		TaskQueue:              tx.ServiceTaskQueue,
		ScheduleToCloseTimeout: tx.RPCTimeout,
		StartToCloseTimeout:    tx.RPCTimeout,
		HeartbeatTimeout:       0,
		WaitForCancellation:    false,
		ActivityID:             uuid.New(),
		RetryPolicy:            &temporal.RetryPolicy{InitialInterval: *tx.RetryPolicy.RetryInitialInterval, BackoffCoefficient: tx.RetryPolicy.RetryBackoffCoefficient, MaximumInterval: tx.RetryPolicy.MaximumInterval, MaximumAttempts: int32(tx.RetryPolicy.MaximumAttempts), NonRetryableErrorTypes: []string{}},
		DisableEagerExecution:  false,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	// delete bank account and all referenced associated to it
	// NOTE: this is performing deletion of all associated records within the postgres database
	err := workflow.ExecuteActivity(ctx, tx.DeleteProfileActivity, userID).Get(ctx, nil)
	if err != nil {
		return err
	}

	// TODO: delete profile from the context of stripe
	/*
	 *  err := workflow.ExecuteActivity(ctx, tx.DeleteProfileRecordInStripe, userID).Get(ctx, nil)
	 *	if err != nil {
	 *		return err
	 *	}
	 */

	// TODO: delete profile items via plaid
	/*
	 *  err := workflow.ExecuteActivity(ctx, tx.DeletePlaidItemsForProfile, userID).Get(ctx, nil)
	 *	if err != nil {
	 *		return err
	 *	}
	 */

	// TODO: delete all transactions associated to all bank accounts associated to the profile
	/*
	 *  err := workflow.ExecuteActivity(ctx, tx.DeeleteProfileTransactions, userID).Get(ctx, nil)
	 *	if err != nil {
	 *		return err
	 *	}
	 */

	return nil
}
