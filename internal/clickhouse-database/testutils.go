package clickhousedatabase

import (
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
)

func generateRandomTransaction() *schema.Transaction {
	return &schema.Transaction{
		AccountId:              helper.GenerateRandomString(20),
		Amount:                 float64(helper.GenerateRandomId(10000, 3000000)),
		IsoCurrencyCode:        helper.GenerateRandomString(20),
		UnofficialCurrencyCode: helper.GenerateRandomString(20),
		Category:               []string{helper.GenerateRandomString(8), helper.GenerateRandomString(8), helper.GenerateRandomString(8)},
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
		UserId:                 0,
	}
}

func generateMultipleTransaction(numTransactions int64) []*schema.Transaction {
	var transactions []*schema.Transaction
	for i := int64(0); i < numTransactions; i++ {
		transactions = append(transactions, generateRandomTransaction())
	}
	return transactions
}

func generateRandomId() *uint64 {
	id := uint64(helper.GenerateRandomId(10000, 3000000))
	return &id
}
