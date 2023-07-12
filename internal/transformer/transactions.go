package transformer

import (
	"fmt"

	"github.com/plaid/plaid-go/v12/plaid"
	"google.golang.org/protobuf/types/known/anypb"
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
		Time: &timestamppb.Timestamp{
			Seconds: input.GetAuthorizedDatetime().UTC().Unix(),
			Nanos:   int32(input.GetAuthorizedDatetime().Nanosecond()),
		},
		AdditionalProperties: &anypb.Any{},
		Categories:           input.Category,
	}

	return tx, nil
}
