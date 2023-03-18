package transform

import (
	"github.com/plaid/plaid-go/plaid"
	"google.golang.org/protobuf/types/known/timestamppb"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func NewTransactionFromPlaid(input plaid.Transaction) (*schema.Transaction, error) {
	return &schema.Transaction{
		AccountId:              input.GetAccountId(),
		Amount:                 float64(input.Amount),
		IsoCurrencyCode:        input.GetIsoCurrencyCode(),
		UnofficialCurrencyCode: input.GetUnofficialCurrencyCode(),
		Category:               input.GetCategory(),
		CategoryId:             input.GetCategoryId(),
		CheckNumber:            input.GetCheckNumber(),
		Date:                   input.GetDate(),
		Datetime:               &timestamppb.Timestamp{},
		AuthorizedDate:         input.GetAuthorizedDate(),
		AuthorizedDatetime:     &timestamppb.Timestamp{Seconds: 0, Nanos: 0},
		Location: &schema.TransactionLocation{
			Address:     *input.GetLocation().Address.Get(),
			City:        *input.GetLocation().City.Get(),
			Region:      *input.GetLocation().Region.Get(),
			PostalCode:  *input.GetLocation().PostalCode.Get(),
			Country:     *input.GetLocation().Country.Get(),
			Lat:         float64(*input.GetLocation().Lat.Get()),
			Lon:         float64(*input.GetLocation().Lon.Get()),
			StoreNumber: *input.GetLocation().StoreNumber.Get(),
		},
		Name:         input.GetName(),
		MerchantName: input.GetMerchantName(),
		PaymentMeta: &schema.PaymentMeta{
			ByOrderOf:        input.PaymentMeta.GetByOrderOf(),
			Payee:            input.PaymentMeta.GetPayee(),
			Payer:            input.PaymentMeta.GetPayer(),
			PaymentMethod:    input.PaymentMeta.GetPaymentMethod(),
			PaymentProcessor: input.PaymentMeta.GetPaymentProcessor(),
			PpdId:            input.PaymentMeta.GetPpdId(),
			Reason:           input.PaymentMeta.GetReason(),
			ReferenceNumber:  input.PaymentMeta.GetReferenceNumber(),
		},
		PaymentChannel:       input.GetPaymentChannel(),
		Pending:              input.GetPending(),
		PendingTransactionId: input.GetPendingTransactionId(),
		AccountOwner:         input.GetAccountOwner(),
		TransactionId:        input.GetTransactionId(),
		TransactionCode:      string(input.GetTransactionCode()),
	}, nil
}
