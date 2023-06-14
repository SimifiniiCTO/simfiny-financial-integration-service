package clickhousedatabase

import (
	"context"
	"reflect"
	"testing"

	clickhousedb "github.com/SimifiniiCTO/simfiny-core-lib/database/clickhouse"
	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/dal"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
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
	type fields struct {
		Conn                  *clickhousedb.Client
		logger                *zap.Logger
		instrumentationClient *instrumentation.Client
		queryOperator         *dal.Query
	}
	type args struct {
		ctx    context.Context
		userId *uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*schema.InvestmentTransaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &Db{
				Conn:                  tt.fields.Conn,
				logger:                tt.fields.logger,
				instrumentationClient: tt.fields.instrumentationClient,
				queryOperator:         tt.fields.queryOperator,
			}
			got, err := db.GetAllInvestmentTransactions(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetAllInvestmentTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Db.GetAllInvestmentTransactions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDb_UpdateInvestmentTransactions(t *testing.T) {
	type args struct {
		ctx    context.Context
		userId *uint64
		txs    []*schema.InvestmentTransaction
	}
	tests := []struct {
		name    string
		db      *Db
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.db.UpdateInvestmentTransactions(tt.args.ctx, tt.args.userId, tt.args.txs); (err != nil) != tt.wantErr {
				t.Errorf("Db.UpdateInvestmentTransactions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDb_GetInvestmentTransactionById(t *testing.T) {
	type args struct {
		ctx  context.Context
		txId *uint64
	}
	tests := []struct {
		name    string
		db      *Db
		args    args
		want    *schema.InvestmentTransaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetInvestmentTransactionById(tt.args.ctx, tt.args.txId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetInvestmentTransactionById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Db.GetInvestmentTransactionById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDb_GetInvestmentTransactionsByPlaidTransactionIds(t *testing.T) {
	type args struct {
		ctx   context.Context
		txIds []string
	}
	tests := []struct {
		name    string
		db      *Db
		args    args
		want    []*schema.InvestmentTransaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetInvestmentTransactionsByPlaidTransactionIds(tt.args.ctx, tt.args.txIds)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetInvestmentTransactionsByPlaidTransactionIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Db.GetInvestmentTransactionsByPlaidTransactionIds() = %v, want %v", got, tt.want)
			}
		})
	}
}
