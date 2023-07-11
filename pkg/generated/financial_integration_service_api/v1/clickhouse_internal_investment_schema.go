package financial_integration_service_apiv1

import (
	context "context"
	"time"

	"github.com/uptrace/go-clickhouse/ch"
)

type InvestmentTransactionInternal struct {
	ch.CHModel              `ch:"InvestmentTransactionInternal,partition:toYYYYMM(time)"`
	AccountId               string    `ch:"AccountId,lc"`
	Amount                  float64   `ch:"Amount"`
	CreatedAt               string    `ch:"CreatedAt,lc"`
	CurrentDate             string    `ch:"CurrentDate,lc"`
	Fees                    float64   `ch:"Fees"`
	ID                      string    `ch:"ID,type:String,default:generateUUIDv4(),pk"`
	InvestmentTransactionId string    `ch:"InvestmentTransactionId,lc"`
	IsoCurrencyCode         string    `ch:"IsoCurrencyCode,lc"`
	LinkId                  uint64    `ch:"LinkId"`
	Name                    string    `ch:"Name,lc"`
	Price                   float64   `ch:"Price"`
	Quantity                float64   `ch:"Quantity"`
	SecurityId              string    `ch:"SecurityId,lc"`
	Sign                    int8      `ch:"Sign"`
	Subtype                 string    `ch:"Subtype,lc"`
	Time                    time.Time `ch:"Time,type:DateTime,pk"`
	Type                    string    `ch:"Type,lc"`
	UnofficialCurrencyCode  string    `ch:"UnofficialCurrencyCode,lc"`
	UserId                  uint64    `ch:"UserId"`
}

func (source *InvestmentTransactionInternal) ConvertToORM() *InvestmentTransactionORM {
	return &InvestmentTransactionORM{
		AccountId:               source.AccountId,
		Amount:                  source.Amount,
		CreatedAt:               source.CreatedAt,
		CurrentDate:             source.CurrentDate,
		Fees:                    source.Fees,
		Id:                      source.ID,
		InvestmentTransactionId: source.InvestmentTransactionId,
		IsoCurrencyCode:         source.IsoCurrencyCode,
		LinkId:                  source.LinkId,
		Name:                    source.Name,
		Price:                   source.Price,
		Quantity:                source.Quantity,
		SecurityId:              source.SecurityId,
		Sign:                    int32(source.Sign),
		Subtype:                 source.Subtype,
		Time:                    &source.Time,
		Type:                    source.Type,
		UnofficialCurrencyCode:  source.UnofficialCurrencyCode,
		UserId:                  source.UserId,
	}
}

func (source *InvestmentTransactionORM) ConvertToInternal() *InvestmentTransactionInternal {
	tx := &InvestmentTransactionInternal{
		AccountId:               source.AccountId,
		Amount:                  source.Amount,
		CreatedAt:               source.CreatedAt,
		CurrentDate:             source.CurrentDate,
		Fees:                    source.Fees,
		ID:                      source.Id,
		InvestmentTransactionId: source.InvestmentTransactionId,
		IsoCurrencyCode:         source.IsoCurrencyCode,
		LinkId:                  source.LinkId,
		Name:                    source.Name,
		Price:                   source.Price,
		Quantity:                source.Quantity,
		SecurityId:              source.SecurityId,
		Sign:                    int8(source.Sign),
		Subtype:                 source.Subtype,
		Type:                    source.Type,
		UnofficialCurrencyCode:  source.UnofficialCurrencyCode,
		UserId:                  source.UserId,
	}

	if source.Time != nil {
		tx.Time = *source.Time
	}

	return tx
}

func (internal *InvestmentTransactionInternal) ConvertToInvestmentTransaction() (*InvestmentTransaction, error) {
	ctx := context.Background()
	ormRec := internal.ConvertToORM()
	tx, err := ormRec.ToPB(ctx)
	if err != nil {
		return nil, err
	}

	return &tx, nil
}

func (internal *InvestmentTransaction) ConvertToInternal() (*InvestmentTransactionInternal, error) {
	ctx := context.Background()
	ormRec, err := internal.ToORM(ctx)
	if err != nil {
		return nil, err
	}

	tx := ormRec.ConvertToInternal()
	return tx, nil
}
