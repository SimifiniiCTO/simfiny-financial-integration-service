package clickhousedatabase

import (
	"time"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func generateRandomInvestmentTransaction() *schema.InvestmentTransaction {
	return &schema.InvestmentTransaction{
		AccountId:               helper.GenerateRandomString(20),
		Ammount:                 helper.GenerateRandomString(20),
		InvestmentTransactionId: helper.GenerateRandomString(20),
		SecurityId:              helper.GenerateRandomString(20),
		CurrentDate:             helper.GenerateRandomString(20),
		Name:                    helper.GenerateRandomString(20),
		Quantity:                float64(*generateRandomId()),
		Amount:                  float64(*generateRandomId()),
		Price:                   float64(*generateRandomId()),
		Fees:                    float64(*generateRandomId()),
		Type:                    helper.GenerateRandomString(20),
		Subtype:                 helper.GenerateRandomString(20),
		IsoCurrencyCode:         helper.GenerateRandomString(20),
		UnofficialCurrencyCode:  helper.GenerateRandomString(20),
		LinkId:                  *generateRandomId(),
		Id:                      "",
		UserId:                  *generateRandomId(),
		CreatedAt:               helper.GenerateRandomString(20),
		Sign:                    1,
	}
}

func generateRandomReOccurringTransaction() *schema.ReOccuringTransaction {
	return &schema.ReOccuringTransaction{
		AccountId:                       helper.GenerateRandomString(20),
		StreamId:                        helper.GenerateRandomString(20),
		CategoryId:                      helper.GenerateRandomString(20),
		Description:                     helper.GenerateRandomString(20),
		MerchantName:                    helper.GenerateRandomString(20),
		PersonalFinanceCategoryPrimary:  helper.GenerateRandomString(20),
		PersonalFinanceCategoryDetailed: helper.GenerateRandomString(20),
		FirstDate:                       helper.GenerateRandomString(20),
		LastDate:                        helper.GenerateRandomString(20),
		Frequency:                       schema.ReOccuringTransactionsFrequency_RE_OCCURING_TRANSACTIONS_FREQUENCY_MONTHLY,
		TransactionIds:                  helper.GenerateRandomString(20),
		AverageAmount:                   helper.GenerateRandomString(20),
		AverageAmountIsoCurrencyCode:    helper.GenerateRandomString(20),
		LastAmount:                      helper.GenerateRandomString(20),
		LastAmountIsoCurrencyCode:       helper.GenerateRandomString(20),
		IsActive:                        false,
		Status:                          schema.ReOccuringTransactionsStatus_RE_OCCURING_TRANSACTIONS_STATUS_EARLY_DETECTION,
		UpdatedTime:                     helper.GenerateRandomString(20),
		UserId:                          *generateRandomId(),
		LinkId:                          *generateRandomId(),
		Id:                              "",
		Flow:                            schema.ReCurringFlow_RE_CURRING_FLOW_INFLOW,
		Sign:                            1,
		Time: &timestamppb.Timestamp{
			Seconds: int64(time.Now().UTC().Second()),
			Nanos:   int32(time.Now().UTC().Nanosecond()),
		},
	}
}

func generateRandomTransaction() *schema.Transaction {
	return &schema.Transaction{
		AccountId:                       helper.GenerateRandomString(20),
		Amount:                          float64(helper.GenerateRandomId(10000, 3000000)),
		IsoCurrencyCode:                 helper.GenerateRandomString(20),
		UnofficialCurrencyCode:          helper.GenerateRandomString(20),
		CategoryId:                      helper.GenerateRandomString(20),
		CheckNumber:                     helper.GenerateRandomString(20),
		CurrentDate:                     time.Now().String(),
		CurrentDatetime:                 time.Now().String(),
		AuthorizedDate:                  time.Now().String(),
		AuthorizedDatetime:              time.Now().String(),
		Name:                            helper.GenerateRandomString(20),
		MerchantName:                    helper.GenerateRandomString(20),
		PaymentChannel:                  helper.GenerateRandomString(20),
		Pending:                         true,
		PendingTransactionId:            helper.GenerateRandomString(20),
		AccountOwner:                    helper.GenerateRandomString(20),
		TransactionId:                   helper.GenerateRandomString(20),
		TransactionCode:                 helper.GenerateRandomString(20),
		Id:                              "",
		UserId:                          *generateRandomId(),
		LinkId:                          *generateRandomId(),
		Sign:                            1,
		PersonalFinanceCategoryPrimary:  helper.GenerateRandomString(20),
		PersonalFinanceCategoryDetailed: helper.GenerateRandomString(20),
		LocationAddress:                 helper.GenerateRandomString(20),
		LocationCity:                    helper.GenerateRandomString(20),
		LocationRegion:                  helper.GenerateRandomString(20),
		LocationPostalCode:              helper.GenerateRandomString(20),
		LocationCountry:                 helper.GenerateRandomString(20),
		LocationLat:                     float64(*generateRandomId()),
		LocationLon:                     float64(*generateRandomId()),
		LocationStoreNumber:             helper.GenerateRandomString(20),
		PaymentMetaByOrderOf:            helper.GenerateRandomString(20),
		PaymentMetaPayee:                helper.GenerateRandomString(20),
		PaymentMetaPayer:                helper.GenerateRandomString(20),
		PaymentMetaPaymentMethod:        helper.GenerateRandomString(20),
		PaymentMetaPaymentProcessor:     helper.GenerateRandomString(20),
		PaymentMetaPpdId:                helper.GenerateRandomString(20),
		PaymentMetaReason:               helper.GenerateRandomString(20),
		PaymentMetaReferenceNumber:      helper.GenerateRandomString(20),
		Time:                            &timestamppb.Timestamp{},
		Categories:                      generateStringList(5),
	}
}

func generateMultipleInvestmentTransaction(numTransactions int) []*schema.InvestmentTransaction {
	var transactions []*schema.InvestmentTransaction
	for i := 0; i < numTransactions; i++ {
		transactions = append(transactions, generateRandomInvestmentTransaction())
	}
	return transactions
}

func generateMultipleTransaction(numTransactions int64) []*schema.Transaction {
	var transactions []*schema.Transaction
	for i := int64(0); i < numTransactions; i++ {
		transactions = append(transactions, generateRandomTransaction())
	}
	return transactions
}

func generateMultipleReOcurringTransactions(numTransactions int) []*schema.ReOccuringTransaction {
	var transactions []*schema.ReOccuringTransaction
	for i := 0; i < numTransactions; i++ {
		transactions = append(transactions, generateRandomReOccurringTransaction())
	}
	return transactions
}

func generateRandomId() *uint64 {
	id := uint64(helper.GenerateRandomId(100000, 3000000))
	return &id
}

func generateStringList(numElements int) []string {
	var list []string
	for i := 0; i < numElements; i++ {
		list = append(list, helper.GenerateRandomString(20))
	}
	return list
}

func PointerUint64(v uint64) *uint64 {
	return &v
}
