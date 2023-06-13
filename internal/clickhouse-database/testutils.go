package clickhousedatabase

import (
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
)

func generateRandomReOccurringTransaction() *schema.ReOccuringTransaction {
	return &schema.ReOccuringTransaction{
		AccountId:                       helper.GenerateRandomString(20),
		StreamId:                        helper.GenerateRandomString(20),
		Category:                        generateStringList(5),
		CategoryId:                      helper.GenerateRandomString(20),
		Description:                     helper.GenerateRandomString(20),
		MerchantName:                    helper.GenerateRandomString(20),
		PersonalFinanceCategoryPrimary:  helper.GenerateRandomString(20),
		PersonalFinanceCategoryDetailed: helper.GenerateRandomString(20),
		FirstDate:                       helper.GenerateRandomString(20),
		LastDate:                        helper.GenerateRandomString(20),
		Frequency:                       schema.ReOccuringTransactionsFrequency_RE_OCCURING_TRANSACTIONS_FREQUENCY_MONTHLY,
		TransactionIds:                  generateStringList(5),
		AverageAmount:                   helper.GenerateRandomString(20),
		AverageAmountIsoCurrencyCode:    helper.GenerateRandomString(20),
		LastAmount:                      helper.GenerateRandomString(20),
		LastAmountIsoCurrencyCode:       helper.GenerateRandomString(20),
		IsActive:                        false,
		Status:                          schema.ReOccuringTransactionsStatus_RE_OCCURING_TRANSACTIONS_STATUS_EARLY_DETECTION,
		UpdatedTime:                     helper.GenerateRandomString(20),
		UserId:                          *generateRandomId(),
		LinkId:                          *generateRandomId(),
		Id:                              0,
	}
}

func generateRandomTransaction() *schema.Transaction {
	return &schema.Transaction{
		AccountId:              helper.GenerateRandomString(20),
		Amount:                 float64(helper.GenerateRandomId(10000, 3000000)),
		IsoCurrencyCode:        helper.GenerateRandomString(20),
		UnofficialCurrencyCode: helper.GenerateRandomString(20),
		Category:               helper.GenerateRandomString(8),
		CategoryId:             helper.GenerateRandomString(20),
		CheckNumber:            helper.GenerateRandomString(20),
		Date:                   helper.GenerateRandomString(20),
		Datetime:               helper.GenerateRandomString(10),
		AuthorizedDate:         helper.GenerateRandomString(20),
		AuthorizedDatetime:     helper.GenerateRandomString(10),
		Location:               &schema.Transaction_Location{},
		Name:                   helper.GenerateRandomString(20),
		MerchantName:           helper.GenerateRandomString(20),
		PaymentMeta:            &schema.Transaction_PaymentMeta{},
		PaymentChannel:         helper.GenerateRandomString(20),
		Pending:                false,
		PendingTransactionId:   helper.GenerateRandomString(20),
		AccountOwner:           helper.GenerateRandomString(20),
		TransactionId:          helper.GenerateRandomString(20),
		TransactionCode:        helper.GenerateRandomString(20),
		Id:                     0,
		UserId:                 *generateRandomId(),
		LinkId:                 *generateRandomId(),
	}
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
	id := uint64(helper.GenerateRandomId(100, 3000000))
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
