package clickhousedatabase

import (
	"context"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	"github.com/stretchr/testify/assert"
)

func TestDb_AddTransaction(t *testing.T) {
	type args struct {
		ctx    context.Context
		userId *uint64
		tx     *schema.Transaction
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
				tx:     generateRandomTransaction(),
			},
			false,
		},
		{
			"[failure] - add transaction with nil user id",
			args{
				ctx:    context.Background(),
				userId: nil,
				tx:     generateRandomTransaction(),
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
			if _, err := conn.AddTransaction(tt.args.ctx, tt.args.userId, tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("conn.AddTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDb_AddTransactions(t *testing.T) {
	type args struct {
		ctx    context.Context
		userId *uint64
		txs    []*schema.Transaction
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
				txs:    generateMultipleTransaction(10),
			},
			false,
		},
		{
			"[failure] - add transactions with nil user id",
			args{
				ctx:    context.Background(),
				userId: nil,
				txs:    generateMultipleTransaction(10),
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
			if err := conn.AddTransactions(tt.args.ctx, tt.args.userId, tt.args.txs); (err != nil) != tt.wantErr {
				t.Errorf("conn.AddTransactions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDb_DeleteTransaction(t *testing.T) {
	type args struct {
		ctx          context.Context
		precondition func(ctx context.Context, t *testing.T, arg *args) *uint64
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
				precondition: func(ctx context.Context, t *testing.T, arg *args) *uint64 {
					tx := generateRandomTransaction()
					userId := generateRandomId()
					txId, err := conn.AddTransaction(ctx, userId, tx)
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
					txId := generateRandomId()
					return txId
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txId := tt.args.precondition(tt.args.ctx, t, &tt.args)
			if err := conn.DeleteTransaction(tt.args.ctx, txId); (err != nil) != tt.wantErr {
				t.Errorf("conn.DeleteTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if _, err := conn.GetTransactionById(tt.args.ctx, txId); err == nil {
					t.Errorf("conn.GetTransactionById() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestDb_DeleteTransactionsByIds(t *testing.T) {
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

					for i := 0; i < arg.numTransactions; i++ {
						tx := generateRandomTransaction()
						txId, err := conn.AddTransaction(ctx, userId, tx)
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
			if err := conn.DeleteTransactionsByIds(tt.args.ctx, txIds); (err != nil) != tt.wantErr {
				t.Errorf("conn.DeleteTransactionsByIds() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDb_DeleteUserTransactons(t *testing.T) {
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
					tx := generateRandomTransaction()
					userId := generateRandomId()
					if _, err := conn.AddTransaction(ctx, userId, tx); err != nil {
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
			if err := conn.DeleteUserTransactons(tt.args.ctx, userId); (err != nil) != tt.wantErr {
				t.Errorf("conn.DeleteUserTransactons() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDb_GetTransactions(t *testing.T) {
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
					tx := generateRandomTransaction()
					userId := generateRandomId()
					if _, err := conn.AddTransaction(ctx, userId, tx); err != nil {
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
			got, nextPageNumber, err := conn.GetTransactions(tt.args.ctx, userId, 1, 10)
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

func TestDb_UpdateTransaction(t *testing.T) {
	type args struct {
		ctx          context.Context
		precondition func(ctx context.Context, t *testing.T, arg *args) (*uint64, *uint64, *schema.Transaction)
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
				precondition: func(ctx context.Context, t *testing.T, arg *args) (*uint64, *uint64, *schema.Transaction) {
					tx := generateRandomTransaction()
					userId := generateRandomId()
					txId, err := conn.AddTransaction(ctx, userId, tx)
					if err != nil {
						t.Errorf("conn.AddTransaction() error = %v", err)
					}

					// query for the tx
					tx, err = conn.GetTransactionById(ctx, txId)
					if err != nil {
						t.Errorf("conn.GetTransaction() error = %v", err)
					}

					tx.AccountOwner = helper.GenerateRandomString(50)
					return txId, &tx.UserId, tx
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txId, userId, tx := tt.args.precondition(tt.args.ctx, t, &tt.args)
			if err := conn.UpdateTransaction(tt.args.ctx, userId, txId, tx); (err != nil) != tt.wantErr {
				t.Errorf("conn.UpdateTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
