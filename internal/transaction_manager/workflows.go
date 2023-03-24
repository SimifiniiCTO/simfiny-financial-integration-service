package transactionmanager

import (
	"github.com/pborman/uuid"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// DeleteBankAccountWorkflow implements WorkflowManager
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
	 *  err := workflow.ExecuteActivity(ctx, tx.DeleteBankAccountTransactionsActivity, bankAccountID).Get(ctx, nil)
	 *	if err != nil {
	 *		return err
	 *	}
	 */

	// TODO: delete item via plaid if bank account is linked to a plaid item
	// ensure that the item is only deleted if there are no other accounts associated with the profile
	/*
	 *  err := workflow.ExecuteActivity(ctx, tx.DeletePlaidItemForBankAccount, profileID).Get(ctx, nil)
	 *	if err != nil {
	 *		return err
	 *	}
	 */

	return nil
}

// DeleteProfileWorkflow implements WorkflowManager
// This operation should be atomic .... all deletion events should be atomic across all services and third party services
// meaning that if one deletion fails, then all deletions should be rolled back (apply saga pattern)
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
