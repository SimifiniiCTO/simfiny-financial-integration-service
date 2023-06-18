package clickhousedatabase

import (
	"context"
	"fmt"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/stretchr/testify/assert"
)

// TODO: need to extensively test the following functions:
func TestDb_AddInvestmentTransactions(t *testing.T) {
	type args struct {
		ctx    context.Context
		userId *uint64
		tx     []*schema.InvestmentTransaction
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
				tx:     generateMultipleInvestmentTransaction(5),
			},
			false,
		},
		{
			"[failure] - add transaction with nil user id",
			args{
				ctx:    context.Background(),
				userId: nil,
				tx:     generateMultipleInvestmentTransaction(5),
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
			if err := conn.AddInvestmentTransactions(tt.args.ctx, tt.args.userId, tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("conn.AddTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDb_DeleteInvestmentTransactions(t *testing.T) {
	type args struct {
		ctx             context.Context
		precondition    func(ctx context.Context, t *testing.T, arg *args) []uint64
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
				precondition: func(ctx context.Context, t *testing.T, arg *args) []uint64 {
					userId := generateRandomId()
					idset := make([]uint64, 0)

					txs := generateMultipleInvestmentTransaction(arg.numTransactions)
					if err := conn.AddInvestmentTransactions(ctx, userId, txs); err != nil {
						t.Errorf("conn.AddTransaction() error = %v", err)
					}

					currentTxs, _, err := conn.GetInvestmentTransactions(ctx, userId, int64(arg.numTransactions), 1)
					if err != nil {
						t.Errorf("conn.GetTransactions() error = %v", err)
					}

					for _, tx := range currentTxs {
						idset = append(idset, tx.Id)
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
				precondition: func(ctx context.Context, t *testing.T, arg *args) []uint64 {
					return nil
				},
			},
			true,
		},
		{
			"[failure] - delete transaction with non-existent transaction id",
			args{
				ctx: context.Background(),
				precondition: func(ctx context.Context, t *testing.T, arg *args) []uint64 {
					idset := make([]uint64, 0)
					for i := 0; i < arg.numTransactions; i++ {
						id := generateRandomId()
						idset = append(idset, *id)
					}
					return idset
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txIds := tt.args.precondition(tt.args.ctx, t, &tt.args)
			if err := conn.DeleteInvestmentTransactions(tt.args.ctx, txIds...); (err != nil) != tt.wantErr {
				t.Errorf("conn.DeleteTransactionsByIds() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDb_GetInvestmentTransactions(t *testing.T) {
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
					txs := generateMultipleInvestmentTransaction(1)
					userId := generateRandomId()
					if err := conn.AddInvestmentTransactions(ctx, userId, txs); err != nil {
						t.Errorf("conn.AddTransaction() error = %v", err)
					}

					return userId
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
		{
			"[failure] - delete transaction with non-existent transaction id",
			args{
				ctx: context.Background(),
				precondition: func(ctx context.Context, t *testing.T, arg *args) *uint64 {
					return generateRandomId()
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userId := tt.args.precondition(tt.args.ctx, t, &tt.args)
			got, nextPageNumber, err := conn.GetInvestmentTransactions(tt.args.ctx, userId, 1, 1)
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

func TestDb_GetAllInvestmentTransactions(t *testing.T) {
	type args struct {
		ctx                       context.Context
		userId                    *uint64
		numInvestmentTransactions int
		precondition              func(ctx context.Context, t *testing.T, arg *args)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "[success] - get all transactions",
			args: args{
				ctx:                       context.Background(),
				userId:                    generateRandomId(),
				numInvestmentTransactions: 10,
				precondition: func(ctx context.Context, t *testing.T, arg *args) {
					txs := generateMultipleInvestmentTransaction(arg.numInvestmentTransactions)
					if err := conn.AddInvestmentTransactions(ctx, arg.userId, txs); err != nil {
						t.Errorf("conn.AddTransaction() error = %v", err)
					}
				},
			},
			wantErr: false,
		},
		{
			name: "[failure] - get all transactions with nil user id",
			args: args{
				ctx:                       context.Background(),
				userId:                    nil,
				numInvestmentTransactions: 10,
				precondition: func(ctx context.Context, t *testing.T, arg *args) {
					userId := generateRandomId()
					txs := generateMultipleInvestmentTransaction(arg.numInvestmentTransactions)
					if err := conn.AddInvestmentTransactions(ctx, userId, txs); err != nil {
						t.Errorf("conn.AddTransaction() error = %v", err)
					}
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// define precondition here
			tt.args.precondition(tt.args.ctx, t, &tt.args)
			got, err := conn.GetAllInvestmentTransactions(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetAllInvestmentTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// ensure that the number of transactions returned is equal to the number of transactions added
			if !tt.wantErr {
				assert.Equal(t, len(got), tt.args.numInvestmentTransactions)
			}
		})
	}
}

func TestDb_UpdateInvestmentTransactions(t *testing.T) {
	type args struct {
		ctx          context.Context
		userId       *uint64
		txs          []*schema.InvestmentTransaction
		precondition func(ctx context.Context, t *testing.T, arg *args) []*schema.InvestmentTransaction
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "[success] - update transaction",
			args: args{
				ctx:    context.Background(),
				userId: generateRandomId(),
				txs:    generateMultipleInvestmentTransaction(10),
				precondition: func(ctx context.Context, t *testing.T, arg *args) []*schema.InvestmentTransaction {
					if err := conn.AddInvestmentTransactions(ctx, arg.userId, arg.txs); err != nil {
						t.Errorf("conn.AddTransaction() error = %v", err)
					}

					txs, err := conn.GetAllInvestmentTransactions(ctx, arg.userId)
					if err != nil {
						t.Errorf("conn.GetAllTransactions() error = %v", err)
					}

					// update all the transaction merchant names
					for _, tx := range txs {
						tx.Name = fmt.Sprintf("updated_%s_%d", tx.Name, tx.Id)
					}

					return txs
				},
			},
			wantErr: false,
		},
		{
			name: "[failure] - update transaction with nil user id",
			args: args{
				ctx:    context.Background(),
				userId: nil,
				txs:    generateMultipleInvestmentTransaction(10),
				precondition: func(ctx context.Context, t *testing.T, arg *args) []*schema.InvestmentTransaction {
					userId := generateRandomId()
					if err := conn.AddInvestmentTransactions(ctx, userId, arg.txs); err != nil {
						t.Errorf("conn.AddTransaction() error = %v", err)
					}

					txs, err := conn.GetAllInvestmentTransactions(ctx, userId)
					if err != nil {
						t.Errorf("conn.GetAllTransactions() error = %v", err)
					}

					// update all the transaction merchant names
					for _, tx := range txs {
						tx.Name = fmt.Sprintf("updated_%s_%d", tx.Name, tx.Id)
					}

					return txs
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// call precondition
			txs := tt.args.precondition(tt.args.ctx, t, &tt.args)
			if err := conn.UpdateInvestmentTransactions(tt.args.ctx, tt.args.userId, txs); (err != nil) != tt.wantErr {
				t.Errorf("Db.UpdateInvestmentTransactions() error = %v, wantErr %v", err, tt.wantErr)
			}

			// ensure that the transactions were updated
			if !tt.wantErr {
				updatedTxs, err := conn.GetAllInvestmentTransactions(tt.args.ctx, tt.args.userId)
				if err != nil {
					t.Errorf("conn.GetAllTransactions() error = %v", err)
				}

				for _, tx := range updatedTxs {
					assert.Contains(t, tx.Name, "updated")
				}
			}
		})
	}
}

func TestDb_GetInvestmentTransactionById(t *testing.T) {
	type args struct {
		ctx          context.Context
		userId       *uint64
		precondition func(ctx context.Context, t *testing.T, arg *args) *uint64
	}
	tests := []struct {
		name    string
		args    args
		want    *schema.InvestmentTransaction
		wantErr bool
	}{
		{
			name: "[success] - get transaction by id",
			args: args{
				ctx:    context.Background(),
				userId: generateRandomId(),
				precondition: func(ctx context.Context, t *testing.T, arg *args) *uint64 {
					tx := generateRandomInvestmentTransaction()
					if err := conn.AddInvestmentTransactions(ctx, arg.userId, []*schema.InvestmentTransaction{tx}); err != nil {
						t.Errorf("conn.AddTransaction() error = %v", err)
					}

					// get all the transactions for the given userId
					txs, err := conn.GetAllInvestmentTransactions(ctx, arg.userId)
					if err != nil {
						t.Errorf("conn.GetAllTransactions() error = %v", err)
					}

					return &txs[0].Id
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txId := tt.args.precondition(tt.args.ctx, t, &tt.args)
			got, err := conn.GetInvestmentTransactionById(tt.args.ctx, txId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetInvestmentTransactionById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
				assert.True(t, got.Id == *txId)
			}
		})
	}
}

func TestDb_GetInvestmentTransactionsByPlaidTransactionIds(t *testing.T) {
	type args struct {
		ctx                     context.Context
		numTransactionsToCreate int
		precondition            func(ctx context.Context, t *testing.T, arg *args) []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "[success] - get transactions by plaid transaction ids",
			args: args{
				ctx:                     context.Background(),
				numTransactionsToCreate: 10,
				precondition: func(ctx context.Context, t *testing.T, arg *args) []string {
					userId := generateRandomId()
					txs := generateMultipleInvestmentTransaction(arg.numTransactionsToCreate)
					if err := conn.AddInvestmentTransactions(ctx, userId, txs); err != nil {
						t.Errorf("conn.AddTransaction() error = %v", err)
					}

					// get all the transactions for the given userId
					txs, err := conn.GetAllInvestmentTransactions(ctx, userId)
					if err != nil {
						t.Errorf("conn.GetAllTransactions() error = %v", err)
					}

					txIds := make([]string, 0)
					for _, tx := range txs {
						txIds = append(txIds, tx.InvestmentTransactionId)
					}

					return txIds
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txIds := tt.args.precondition(tt.args.ctx, t, &tt.args)
			got, err := conn.GetInvestmentTransactionsByPlaidTransactionIds(tt.args.ctx, txIds)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetInvestmentTransactionsByPlaidTransactionIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
				assert.Equal(t, len(got), len(txIds))
			}
		})
	}
}
