package clickhousedatabase

import (
	"context"
	"testing"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/stretchr/testify/assert"
)

func TestDb_AddReocurringTransaction(t *testing.T) {
	type args struct {
		ctx    context.Context
		userId *uint64
		tx     *schema.ReOccuringTransaction
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"[success] - add transaction",
			args{
				ctx:    context.Background(),
				userId: generateRandomId(),
				tx:     generateRandomReOccurringTransaction(),
			},
			false,
		},
		{
			"[failure] - add transaction with nil user id",
			args{
				ctx:    context.Background(),
				userId: nil,
				tx:     generateRandomReOccurringTransaction(),
			},
			true,
		},
		{
			"[failure] - add transaction with nil transaction",
			args{
				ctx:    context.Background(),
				userId: generateRandomId(),
				tx:     nil,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := conn.AddReOccurringTransaction(tt.args.ctx, tt.args.userId, tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("conn.AddTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDb_AddReOccurringTransactions(t *testing.T) {
	type args struct {
		ctx    context.Context
		userId *uint64
		txs    []*schema.ReOccuringTransaction
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"[success] - add transactions",
			args{
				ctx:    context.Background(),
				userId: generateRandomId(),
				txs:    generateMultipleReOcurringTransactions(10),
			},
			false,
		},
		{
			"[failure] - add transactions with nil user id",
			args{
				ctx:    context.Background(),
				userId: nil,
				txs:    generateMultipleReOcurringTransactions(10),
			},
			true,
		},
		{
			"[failure] - add transactions with nil transactions",
			args{
				ctx:    context.Background(),
				userId: generateRandomId(),
				txs:    nil,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := conn.AddReOccurringTransactions(tt.args.ctx, tt.args.userId, tt.args.txs); (err != nil) != tt.wantErr {
				t.Errorf("conn.AddTransactions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDb_DeleteReOcurringTransaction(t *testing.T) {
	type args struct {
		ctx          context.Context
		precondition func(ctx context.Context, t *testing.T, arg *args) *string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"[success] - delete transaction",
			args{
				ctx: context.Background(),
				precondition: func(ctx context.Context, t *testing.T, arg *args) *string {
					tx := generateRandomReOccurringTransaction()
					userId := generateRandomId()
					txId, err := conn.AddReOccurringTransaction(ctx, userId, tx)
					if err != nil {
						t.Errorf("conn.AddTransaction() error = %v", err)
					}

					return txId
				},
			},
			false,
		},
		{
			"[failure] - delete transaction with nil transaction id",
			args{
				ctx: context.Background(),
				precondition: func(ctx context.Context, t *testing.T, arg *args) *string {
					return nil
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txId := tt.args.precondition(tt.args.ctx, t, &tt.args)
			if err := conn.DeleteReOccuringTransaction(tt.args.ctx, txId); (err != nil) != tt.wantErr {
				t.Errorf("conn.DeleteTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if _, err := conn.GetReOcurringTransactionById(tt.args.ctx, txId); err == nil {
					t.Errorf("conn.GetTransactionById() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestDb_DeleteReOcurringTransactionsByIds(t *testing.T) {
	type args struct {
		ctx             context.Context
		precondition    func(ctx context.Context, t *testing.T, arg *args) []string
		numTransactions int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"[success] - delete transaction",
			args{
				ctx: context.Background(),
				precondition: func(ctx context.Context, t *testing.T, arg *args) []string {
					userId := generateRandomId()
					idset := make([]string, 0)

					for i := 0; i < arg.numTransactions; i++ {
						tx := generateRandomReOccurringTransaction()
						txId, err := conn.AddReOccurringTransaction(ctx, userId, tx)
						if err != nil {
							t.Errorf("conn.AddTransaction() error = %v", err)
						}

						idset = append(idset, *txId)
					}

					return idset
				},
				numTransactions: 10,
			},
			false,
		},
		{
			"[failure] - delete transaction with nil transaction id",
			args{
				ctx: context.Background(),
				precondition: func(ctx context.Context, t *testing.T, arg *args) []string {
					return nil
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txIds := tt.args.precondition(tt.args.ctx, t, &tt.args)
			if err := conn.DeleteReOccurringTransactionsByIds(tt.args.ctx, txIds); (err != nil) != tt.wantErr {
				t.Errorf("conn.DeleteTransactionsByIds() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDb_GetReOcurringTransactions(t *testing.T) {
	type args struct {
		ctx          context.Context
		precondition func(ctx context.Context, t *testing.T, arg *args) *uint64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"[success] - delete transaction",
			args{
				ctx: context.Background(),
				precondition: func(ctx context.Context, t *testing.T, arg *args) *uint64 {
					tx := generateRandomReOccurringTransaction()
					userId := generateRandomId()
					if _, err := conn.AddReOccurringTransaction(ctx, userId, tx); err != nil {
						t.Errorf("conn.AddTransaction() error = %v", err)
					}

					return &tx.UserId
				},
			},
			false,
		},
		{
			"[failure] - delete transaction with nil transaction id",
			args{
				ctx: context.Background(),
				precondition: func(ctx context.Context, t *testing.T, arg *args) *uint64 {
					return nil
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userId := tt.args.precondition(tt.args.ctx, t, &tt.args)
			got, nextPageNumber, err := conn.GetReOcurringTransactions(tt.args.ctx, userId, 1, 10)
			if (err != nil) != tt.wantErr {
				t.Errorf("conn.GetTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, nextPageNumber)
				if len(got) == 0 {
					t.Errorf("conn.GetTransactions() = %v, want %v", got, 1)
				}
			}
		})
	}
}

func TestDb_UpdateReOccurringTransaction(t *testing.T) {
	type args struct {
		ctx          context.Context
		precondition func(ctx context.Context, t *testing.T, arg *args) (*string, *uint64, *schema.ReOccuringTransaction)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"[success] - update transaction",
			args{
				ctx: context.Background(),
				precondition: func(ctx context.Context, t *testing.T, arg *args) (*string, *uint64, *schema.ReOccuringTransaction) {
					tx := generateRandomReOccurringTransaction()
					userId := generateRandomId()
					txId, err := conn.AddReOccurringTransaction(ctx, userId, tx)
					if err != nil {
						t.Errorf("conn.AddTransaction() error = %v", err)
					}

					// query for the tx
					tx, err = conn.GetReOcurringTransactionById(ctx, txId)
					if err != nil {
						t.Errorf("conn.GetTransaction() error = %v", err)
					}

					return txId, &tx.UserId, tx
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txId, userId, tx := tt.args.precondition(tt.args.ctx, t, &tt.args)
			if err := conn.UpdateReOccurringTransaction(tt.args.ctx, userId, txId, tx); (err != nil) != tt.wantErr {
				t.Errorf("conn.UpdateTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDb_AddReOccurringTransaction(t *testing.T) {
	type args struct {
		ctx    context.Context
		userId *uint64
		tx     *schema.ReOccuringTransaction
	}
	tests := []struct {
		name    string
		db      *Db
		args    args
		wantErr bool
	}{
		{
			"[success] - add transaction",
			conn,
			args{
				ctx:    context.Background(),
				userId: generateRandomId(),
				tx:     generateRandomReOccurringTransaction(),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.AddReOccurringTransaction(tt.args.ctx, tt.args.userId, tt.args.tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.AddReOccurringTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if _, err := tt.db.GetReOcurringTransactionById(tt.args.ctx, got); err != nil {
					t.Errorf("Db.GetReOcurringTransactionById() error = %v", err)
				}
			}
		})
	}
}

func TestDb_DeleteReOccuringTransaction(t *testing.T) {
	type args struct {
		ctx          context.Context
		precondition func(ctx context.Context, t *testing.T, arg *args) *string
	}
	tests := []struct {
		name    string
		db      *Db
		args    args
		wantErr bool
	}{
		{
			"[success] - delete transaction",
			conn,
			args{
				ctx: context.Background(),
				precondition: func(ctx context.Context, t *testing.T, arg *args) *string {
					tx := generateRandomReOccurringTransaction()
					userId := generateRandomId()
					txId, err := conn.AddReOccurringTransaction(ctx, userId, tx)
					if err != nil {
						t.Errorf("conn.AddTransaction() error = %v", err)
					}

					return txId
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txId := tt.args.precondition(tt.args.ctx, t, &tt.args)
			if err := tt.db.DeleteReOccuringTransaction(tt.args.ctx, txId); (err != nil) != tt.wantErr {
				t.Errorf("Db.DeleteReOccuringTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDb_DeleteReOccurringTransactionsByIds(t *testing.T) {
	type args struct {
		ctx          context.Context
		precondition func(ctx context.Context, t *testing.T, arg *args) []string
	}
	tests := []struct {
		name    string
		db      *Db
		args    args
		wantErr bool
	}{
		{
			"[success] - delete transaction",
			conn,
			args{
				ctx: context.Background(),
				precondition: func(ctx context.Context, t *testing.T, arg *args) []string {
					userId := generateRandomId()
					idset := make([]string, 0)

					for i := 0; i < 10; i++ {
						tx := generateRandomReOccurringTransaction()
						txId, err := conn.AddReOccurringTransaction(ctx, userId, tx)
						if err != nil {
							t.Errorf("conn.AddTransaction() error = %v", err)
						}

						idset = append(idset, *txId)
					}

					return idset
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txIds := tt.args.precondition(tt.args.ctx, t, &tt.args)
			if err := tt.db.DeleteReOccurringTransactionsByIds(tt.args.ctx, txIds); (err != nil) != tt.wantErr {
				t.Errorf("Db.DeleteReOccurringTransactionsByIds() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDb_DeleteReOcurringTransactionsByLinkId(t *testing.T) {
	type args struct {
		ctx          context.Context
		precondition func(ctx context.Context, t *testing.T, arg *args) *uint64
	}
	tests := []struct {
		name    string
		db      *Db
		args    args
		wantErr bool
	}{
		{
			"[success] - delete transaction",
			conn,
			args{
				ctx: context.Background(),
				precondition: func(ctx context.Context, t *testing.T, arg *args) *uint64 {
					tx := generateRandomReOccurringTransaction()
					userId := generateRandomId()
					_, err := conn.AddReOccurringTransaction(ctx, userId, tx)
					if err != nil {
						t.Errorf("conn.AddTransaction() error = %v", err)
					}

					return &tx.LinkId
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			linkId := tt.args.precondition(tt.args.ctx, t, &tt.args)
			if err := tt.db.DeleteReOcurringTransactionsByLinkId(tt.args.ctx, linkId); (err != nil) != tt.wantErr {
				t.Errorf("Db.DeleteReOcurringTransactionsByLinkId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDb_DeleteUserReOcurringTransactons(t *testing.T) {
	type args struct {
		ctx          context.Context
		userId       *uint64
		precondition func(ctx context.Context, t *testing.T, arg *args)
	}
	tests := []struct {
		name    string
		db      *Db
		args    args
		wantErr bool
	}{
		{
			"[success] - delete transaction",
			conn,
			args{
				ctx:    context.Background(),
				userId: generateRandomId(),
				precondition: func(ctx context.Context, t *testing.T, arg *args) {
					for i := 0; i < 10; i++ {
						tx := generateRandomReOccurringTransaction()
						_, err := conn.AddReOccurringTransaction(ctx, arg.userId, tx)
						if err != nil {
							t.Errorf("conn.AddTransaction() error = %v", err)
						}
					}
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.precondition(tt.args.ctx, t, &tt.args)
			if err := tt.db.DeleteUserReOcurringTransactons(tt.args.ctx, tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("Db.DeleteUserReOcurringTransactons() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDb_UpdateReOccurringTransactions(t *testing.T) {
	type args struct {
		ctx          context.Context
		userId       *uint64
		precondition func(ctx context.Context, t *testing.T, arg *args) []*schema.ReOccuringTransaction
	}
	tests := []struct {
		name    string
		db      *Db
		args    args
		wantErr bool
	}{
		{
			"[success] - update transaction",
			conn,
			args{
				ctx:    context.Background(),
				userId: generateRandomId(),
				precondition: func(ctx context.Context, t *testing.T, arg *args) []*schema.ReOccuringTransaction {
					txs := make([]*schema.ReOccuringTransaction, 0)
					for i := 0; i < 10; i++ {
						tx := generateRandomReOccurringTransaction()
						txId, err := conn.AddReOccurringTransaction(ctx, arg.userId, tx)
						if err != nil {
							t.Errorf("conn.AddTransaction() error = %v", err)
						}

						tx, err = conn.GetReOcurringTransactionById(ctx, txId)
						if err != nil {
							t.Errorf("conn.GetTransaction() error = %v", err)
						}

						txs = append(txs, tx)
					}

					for _, tx := range txs {
						tx.AverageAmount = helper.GenerateRandomString(20)
					}

					return txs
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txs := tt.args.precondition(tt.args.ctx, t, &tt.args)
			if err := tt.db.UpdateReOccurringTransactions(tt.args.ctx, tt.args.userId, txs); (err != nil) != tt.wantErr {
				t.Errorf("Db.UpdateReOccurringTransactions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDb_GetReOcurringTransactionById(t *testing.T) {
	type args struct {
		ctx          context.Context
		precondition func(ctx context.Context, t *testing.T, arg *args) *string
	}
	tests := []struct {
		name    string
		db      *Db
		args    args
		wantErr bool
	}{
		{
			"[success] - get transaction",
			conn,
			args{
				ctx: context.Background(),
				precondition: func(ctx context.Context, t *testing.T, arg *args) *string {
					tx := generateRandomReOccurringTransaction()
					userId := generateRandomId()
					txId, err := conn.AddReOccurringTransaction(ctx, userId, tx)
					if err != nil {
						t.Errorf("conn.AddTransaction() error = %v", err)
					}

					return txId
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txId := tt.args.precondition(tt.args.ctx, t, &tt.args)
			got, err := tt.db.GetReOcurringTransactionById(tt.args.ctx, txId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetReOcurringTransactionById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func TestDb_GetUserReOccurringTransactions(t *testing.T) {
	type args struct {
		ctx          context.Context
		userId       *uint64
		precondition func(ctx context.Context, t *testing.T, arg *args)
	}
	tests := []struct {
		name    string
		db      *Db
		args    args
		wantErr bool
	}{
		{
			"[success] - get transaction",
			conn,
			args{
				ctx:    context.Background(),
				userId: generateRandomId(),
				precondition: func(ctx context.Context, t *testing.T, arg *args) {
					for i := 0; i < 10; i++ {
						tx := generateRandomReOccurringTransaction()
						_, err := conn.AddReOccurringTransaction(ctx, arg.userId, tx)
						if err != nil {
							t.Errorf("conn.AddTransaction() error = %v", err)
						}
					}
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.precondition(tt.args.ctx, t, &tt.args)
			got, err := tt.db.GetUserReOccurringTransactions(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetUserReOccurringTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}
