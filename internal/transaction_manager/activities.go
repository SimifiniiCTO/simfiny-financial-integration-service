package transactionmanager

import (
	"context"
)

// DeleteBankAccountActivity implements WorkflowActivityManager
func (tx *TransactionManager) DeleteBankAccountActivity(ctx context.Context, bankAccountID uint64) error {
	if tx.MetricsEnabled && tx.TelemetrySDK != nil {
		txn := tx.TelemetrySDK.GetTraceFromContext(ctx)
		span := tx.TelemetrySDK.StartSegment(txn, "temporal-delete-bank-account-activity")
		defer span.End()
	}

	if err := tx.DatabaseConn.DeleteBankAccount(ctx, bankAccountID); err != nil {
		return err
	}

	return nil
}

// DeleteProfileActivity implements WorkflowActivityManager
func (tx *TransactionManager) DeleteProfileActivity(ctx context.Context, userID uint64) error {
	if tx.MetricsEnabled && tx.TelemetrySDK != nil {
		txn := tx.TelemetrySDK.GetTraceFromContext(ctx)
		span := tx.TelemetrySDK.StartSegment(txn, "temporal-delete-profile-activity")
		defer span.End()
	}

	if err := tx.DatabaseConn.DeleteUserProfileByUserID(ctx, userID); err != nil {
		return err
	}

	return nil
}
