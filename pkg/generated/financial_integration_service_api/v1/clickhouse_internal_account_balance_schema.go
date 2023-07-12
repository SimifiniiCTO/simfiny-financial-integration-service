package financial_integration_service_apiv1

import (
	time "time"

	"github.com/uptrace/go-clickhouse/ch"
)

type AccountBalanceHistoryInternal struct {
	ch.CHModel      `ch:"partition:toYYYYMM(time)"`
	Time            time.Time `ch:"type:DateTime"`
	AccountId       string    `ch:"type:String"`
	IsoCurrencyCode string    `ch:"type:String"`
	Balance         float64   `ch:"type:Float64"`
	UserId          uint64    `ch:"type:Uint64"`
	Sign            uint32    `ch:"type:Uint32"`
}

func (source *AccountBalanceHistoryInternal) ConvertToORM() *AccountBalanceHistoryORM {
	return &AccountBalanceHistoryORM{
		AccountId:       source.AccountId,
		Balance:         source.Balance,
		IsoCurrencyCode: source.IsoCurrencyCode,
		Time:            &source.Time,
		UserId:          source.UserId,
	}
}

func (source *AccountBalanceHistoryORM) ConvertToInternal() *AccountBalanceHistoryInternal {
	res := &AccountBalanceHistoryInternal{
		AccountId:       source.AccountId,
		IsoCurrencyCode: source.IsoCurrencyCode,
		Balance:         source.Balance,
		UserId:          source.UserId,
	}

	if source.Time != nil {
		res.Time = *source.Time
	}

	return res
}
