package transformer

import (
	"fmt"
	"time"

	"github.com/plaid/plaid-go/v14/plaid"
	"google.golang.org/protobuf/types/known/timestamppb"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

// NewTransactionFromPlaid converts a Plaid transaction to our own transaction interface.
func NewTransactionFromPlaid(input *plaid.Transaction) (*schema.Transaction, error) {
	if input == nil {
		return nil, fmt.Errorf("invalid input argument. transaction cannot be nil")
	}

	location := input.GetLocation()
	paymentMeta := input.GetPaymentMeta()
	personalFinanceCategory := input.GetPersonalFinanceCategory()
	layout := "2006-01-02" // replace with your time format
	str := ""
	if input.GetDate() != "" {
		str = input.GetDate()
	} else {
		str = input.GetAuthorizedDate()
	}

	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
	}
	timestamp := timestamppb.New(t)

	tx := &schema.Transaction{
		AccountId:                       input.GetAccountId(),
		Amount:                          float64(input.GetAmount()),
		IsoCurrencyCode:                 input.GetIsoCurrencyCode(),
		UnofficialCurrencyCode:          input.GetUnofficialCurrencyCode(),
		CategoryId:                      input.GetCategoryId(),
		CheckNumber:                     input.GetCheckNumber(),
		CurrentDate:                     input.GetDate(),
		CurrentDatetime:                 input.GetDate(),
		AuthorizedDate:                  input.GetAuthorizedDate(),
		AuthorizedDatetime:              input.GetAuthorizedDate(),
		Name:                            input.GetName(),
		MerchantName:                    input.GetMerchantName(),
		PaymentChannel:                  input.GetPaymentChannel(),
		Pending:                         input.GetPending(),
		PendingTransactionId:            input.GetPendingTransactionId(),
		AccountOwner:                    input.GetAccountOwner(),
		TransactionId:                   input.GetTransactionId(),
		TransactionCode:                 string(input.GetTransactionCode()),
		Id:                              "",
		UserId:                          0,
		LinkId:                          0,
		Sign:                            1,
		PersonalFinanceCategoryPrimary:  personalFinanceCategory.GetPrimary(),
		PersonalFinanceCategoryDetailed: personalFinanceCategory.GetDetailed(),
		LocationAddress:                 location.GetAddress(),
		LocationCity:                    location.GetCity(),
		LocationRegion:                  location.GetRegion(),
		LocationPostalCode:              location.GetPostalCode(),
		LocationCountry:                 location.GetCountry(),
		LocationLat:                     location.GetLat(),
		LocationLon:                     location.GetLon(),
		LocationStoreNumber:             location.GetStoreNumber(),
		PaymentMetaByOrderOf:            paymentMeta.GetByOrderOf(),
		PaymentMetaPayee:                paymentMeta.GetPayee(),
		PaymentMetaPayer:                paymentMeta.GetPayer(),
		PaymentMetaPaymentMethod:        paymentMeta.GetPaymentMethod(),
		PaymentMetaPaymentProcessor:     paymentMeta.GetPaymentProcessor(),
		PaymentMetaPpdId:                paymentMeta.GetPpdId(),
		PaymentMetaReason:               paymentMeta.GetReason(),
		PaymentMetaReferenceNumber:      paymentMeta.GetReferenceNumber(),
		Time:                            timestamp,
		Categories:                      input.GetCategory(),
	}

	return tx, nil
}
