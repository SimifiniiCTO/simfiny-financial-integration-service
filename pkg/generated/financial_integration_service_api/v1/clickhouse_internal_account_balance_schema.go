package financial_integration_service_apiv1

import (
	"fmt"
	"strconv"
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
	Id              uint64    `ch:"type:String"`
}

func (source *AccountBalanceHistoryInternal) ConvertToORM() *AccountBalanceHistoryORM {
	return &AccountBalanceHistoryORM{
		AccountId:       source.AccountId,
		Balance:         source.Balance,
		Id:              fmt.Sprintf("%d", source.Id),
		IsoCurrencyCode: source.IsoCurrencyCode,
		Sign:            0,
		Time:            &source.Time,
		UserId:          source.UserId,
	}
}

func (source *AccountBalanceHistoryORM) ConvertToInternal() *AccountBalanceHistoryInternal {
	uint64Value, err := strconv.ParseUint(source.Id, 10, 64)
	if err != nil {
		// Handle the error if the string cannot be converted to uint64
		fmt.Println("Error:", err)
		uint64Value = 0
	}

	res := &AccountBalanceHistoryInternal{
		Time:            time.Time{},
		AccountId:       source.AccountId,
		IsoCurrencyCode: source.IsoCurrencyCode,
		Balance:         source.Balance,
		UserId:          source.UserId,
		Sign:            0,
		Id:              uint64Value,
	}

	if source.Time != nil {
		res.Time = *source.Time
	}

	return res
}
