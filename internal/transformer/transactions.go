package transformer

import (
	"github.com/plaid/plaid-go/v12/plaid"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// NewTransactionFromPlaid converts a Plaid transaction to our own transaction interface.
func NewTransactionFromPlaid(input plaid.Transaction) (*schema.Transaction, error) {
	return &schema.Transaction{
		AccountId:                       input.GetAccountId(),
		Amount:                          float64(input.Amount),
		IsoCurrencyCode:                 input.GetIsoCurrencyCode(),
		UnofficialCurrencyCode:          input.GetUnofficialCurrencyCode(),
		CategoryId:                      input.GetCategoryId(),
		CheckNumber:                     input.GetCheckNumber(),
		Date:                            input.GetDate(),
		Datetime:                        input.GetDate(),
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
		Id:                              0,
		UserId:                          0,
		LinkId:                          0,
		Sign:                            0,
		PersonalFinanceCategoryPrimary:  input.PersonalFinanceCategory.Get().Primary,
		PersonalFinanceCategoryDetailed: input.PersonalFinanceCategory.Get().Detailed,
		LocationAddress:                 *input.GetLocation().Address.Get(),
		LocationCity:                    *input.GetLocation().City.Get(),
		LocationRegion:                  *input.GetLocation().Region.Get(),
		LocationPostalCode:              *input.GetLocation().PostalCode.Get(),
		LocationCountry:                 *input.GetLocation().Country.Get(),
		LocationLat:                     *input.GetLocation().Lat.Get(),
		LocationLon:                     *input.GetLocation().Lon.Get(),
		LocationStoreNumber:             *input.GetLocation().StoreNumber.Get(),
		PaymentMetaByOrderOf:            *input.GetPaymentMeta().ByOrderOf.Get(),
		PaymentMetaPayee:                *input.GetPaymentMeta().Payee.Get(),
		PaymentMetaPayer:                *input.GetPaymentMeta().Payer.Get(),
		PaymentMetaPaymentMethod:        *input.GetPaymentMeta().PaymentMethod.Get(),
		PaymentMetaPaymentProcessor:     *input.GetPaymentMeta().PaymentProcessor.Get(),
		PaymentMetaPpdId:                *input.GetPaymentMeta().PpdId.Get(),
		PaymentMetaReason:               *input.GetPaymentMeta().Reason.Get(),
		PaymentMetaReferenceNumber:      *input.GetPaymentMeta().ReferenceNumber.Get(),
	}, nil
}
