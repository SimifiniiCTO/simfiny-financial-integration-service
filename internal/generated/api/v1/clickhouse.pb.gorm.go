package apiv1

import (
	context "context"
	fmt "fmt"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm "github.com/jinzhu/gorm"
	pq "github.com/lib/pq"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	strings "strings"
)

type ReOccuringTransactionORM struct {
	AccountId                       string
	AverageAmount                   string
	AverageAmountIsoCurrencyCode    string
	Category                        pq.StringArray `gorm:"type:text[]"`
	CategoryId                      string
	Description                     string
	FirstDate                       string
	Frequency                       string
	Id                              uint64
	IsActive                        bool
	LastAmount                      string
	LastAmountIsoCurrencyCode       string
	LastDate                        string
	LinkId                          uint64
	MerchantName                    string
	PersonalFinanceCategoryDetailed string
	PersonalFinanceCategoryPrimary  string
	Status                          string
	StreamId                        string
	TransactionIds                  pq.StringArray `gorm:"type:text[]"`
	UpdatedTime                     string
	UserId                          uint64
}

// TableName overrides the default tablename generated by GORM
func (ReOccuringTransactionORM) TableName() string {
	return "re_occuring_transactions"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *ReOccuringTransaction) ToORM(ctx context.Context) (ReOccuringTransactionORM, error) {
	to := ReOccuringTransactionORM{}
	var err error
	if prehook, ok := interface{}(m).(ReOccuringTransactionWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.AccountId = m.AccountId
	to.StreamId = m.StreamId
	if m.Category != nil {
		to.Category = make(pq.StringArray, len(m.Category))
		copy(to.Category, m.Category)
	}
	to.CategoryId = m.CategoryId
	to.Description = m.Description
	to.MerchantName = m.MerchantName
	to.PersonalFinanceCategoryPrimary = m.PersonalFinanceCategoryPrimary
	to.PersonalFinanceCategoryDetailed = m.PersonalFinanceCategoryDetailed
	to.FirstDate = m.FirstDate
	to.LastDate = m.LastDate
	to.Frequency = ReOccuringTransactionsFrequency_name[int32(m.Frequency)]
	if m.TransactionIds != nil {
		to.TransactionIds = make(pq.StringArray, len(m.TransactionIds))
		copy(to.TransactionIds, m.TransactionIds)
	}
	to.AverageAmount = m.AverageAmount
	to.AverageAmountIsoCurrencyCode = m.AverageAmountIsoCurrencyCode
	to.LastAmount = m.LastAmount
	to.LastAmountIsoCurrencyCode = m.LastAmountIsoCurrencyCode
	to.IsActive = m.IsActive
	to.Status = ReOccuringTransactionsStatus_name[int32(m.Status)]
	to.UpdatedTime = m.UpdatedTime
	to.UserId = m.UserId
	to.LinkId = m.LinkId
	to.Id = m.Id
	if posthook, ok := interface{}(m).(ReOccuringTransactionWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *ReOccuringTransactionORM) ToPB(ctx context.Context) (ReOccuringTransaction, error) {
	to := ReOccuringTransaction{}
	var err error
	if prehook, ok := interface{}(m).(ReOccuringTransactionWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.AccountId = m.AccountId
	to.StreamId = m.StreamId
	if m.Category != nil {
		to.Category = make(pq.StringArray, len(m.Category))
		copy(to.Category, m.Category)
	}
	to.CategoryId = m.CategoryId
	to.Description = m.Description
	to.MerchantName = m.MerchantName
	to.PersonalFinanceCategoryPrimary = m.PersonalFinanceCategoryPrimary
	to.PersonalFinanceCategoryDetailed = m.PersonalFinanceCategoryDetailed
	to.FirstDate = m.FirstDate
	to.LastDate = m.LastDate
	to.Frequency = ReOccuringTransactionsFrequency(ReOccuringTransactionsFrequency_value[m.Frequency])
	if m.TransactionIds != nil {
		to.TransactionIds = make(pq.StringArray, len(m.TransactionIds))
		copy(to.TransactionIds, m.TransactionIds)
	}
	to.AverageAmount = m.AverageAmount
	to.AverageAmountIsoCurrencyCode = m.AverageAmountIsoCurrencyCode
	to.LastAmount = m.LastAmount
	to.LastAmountIsoCurrencyCode = m.LastAmountIsoCurrencyCode
	to.IsActive = m.IsActive
	to.Status = ReOccuringTransactionsStatus(ReOccuringTransactionsStatus_value[m.Status])
	to.UpdatedTime = m.UpdatedTime
	to.UserId = m.UserId
	to.LinkId = m.LinkId
	to.Id = m.Id
	if posthook, ok := interface{}(m).(ReOccuringTransactionWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type ReOccuringTransaction the arg will be the target, the caller the one being converted from

// ReOccuringTransactionBeforeToORM called before default ToORM code
type ReOccuringTransactionWithBeforeToORM interface {
	BeforeToORM(context.Context, *ReOccuringTransactionORM) error
}

// ReOccuringTransactionAfterToORM called after default ToORM code
type ReOccuringTransactionWithAfterToORM interface {
	AfterToORM(context.Context, *ReOccuringTransactionORM) error
}

// ReOccuringTransactionBeforeToPB called before default ToPB code
type ReOccuringTransactionWithBeforeToPB interface {
	BeforeToPB(context.Context, *ReOccuringTransaction) error
}

// ReOccuringTransactionAfterToPB called after default ToPB code
type ReOccuringTransactionWithAfterToPB interface {
	AfterToPB(context.Context, *ReOccuringTransaction) error
}

type TransactionORM struct {
	AccountId              string
	AccountOwner           string
	Amount                 float64
	AuthorizedDate         string
	AuthorizedDatetime     string
	Category               string
	CategoryId             string
	CheckNumber            string
	Date                   string
	Datetime               string
	Id                     uint64
	IsoCurrencyCode        string
	LinkId                 uint64
	MerchantName           string
	Name                   string
	PaymentChannel         string
	Pending                bool
	PendingTransactionId   string
	TransactionCode        string
	TransactionId          string
	UnofficialCurrencyCode string
	UserId                 uint64
}

// TableName overrides the default tablename generated by GORM
func (TransactionORM) TableName() string {
	return "transactions"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Transaction) ToORM(ctx context.Context) (TransactionORM, error) {
	to := TransactionORM{}
	var err error
	if prehook, ok := interface{}(m).(TransactionWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.AccountId = m.AccountId
	to.Amount = m.Amount
	to.IsoCurrencyCode = m.IsoCurrencyCode
	to.UnofficialCurrencyCode = m.UnofficialCurrencyCode
	to.Category = m.Category
	to.CategoryId = m.CategoryId
	to.CheckNumber = m.CheckNumber
	to.Date = m.Date
	to.Datetime = m.Datetime
	to.AuthorizedDate = m.AuthorizedDate
	to.AuthorizedDatetime = m.AuthorizedDatetime
	to.Name = m.Name
	to.MerchantName = m.MerchantName
	to.PaymentChannel = m.PaymentChannel
	to.Pending = m.Pending
	to.PendingTransactionId = m.PendingTransactionId
	to.AccountOwner = m.AccountOwner
	to.TransactionId = m.TransactionId
	to.TransactionCode = m.TransactionCode
	to.Id = m.Id
	to.UserId = m.UserId
	to.LinkId = m.LinkId
	if posthook, ok := interface{}(m).(TransactionWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *TransactionORM) ToPB(ctx context.Context) (Transaction, error) {
	to := Transaction{}
	var err error
	if prehook, ok := interface{}(m).(TransactionWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.AccountId = m.AccountId
	to.Amount = m.Amount
	to.IsoCurrencyCode = m.IsoCurrencyCode
	to.UnofficialCurrencyCode = m.UnofficialCurrencyCode
	to.Category = m.Category
	to.CategoryId = m.CategoryId
	to.CheckNumber = m.CheckNumber
	to.Date = m.Date
	to.Datetime = m.Datetime
	to.AuthorizedDate = m.AuthorizedDate
	to.AuthorizedDatetime = m.AuthorizedDatetime
	to.Name = m.Name
	to.MerchantName = m.MerchantName
	to.PaymentChannel = m.PaymentChannel
	to.Pending = m.Pending
	to.PendingTransactionId = m.PendingTransactionId
	to.AccountOwner = m.AccountOwner
	to.TransactionId = m.TransactionId
	to.TransactionCode = m.TransactionCode
	to.Id = m.Id
	to.UserId = m.UserId
	to.LinkId = m.LinkId
	if posthook, ok := interface{}(m).(TransactionWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Transaction the arg will be the target, the caller the one being converted from

// TransactionBeforeToORM called before default ToORM code
type TransactionWithBeforeToORM interface {
	BeforeToORM(context.Context, *TransactionORM) error
}

// TransactionAfterToORM called after default ToORM code
type TransactionWithAfterToORM interface {
	AfterToORM(context.Context, *TransactionORM) error
}

// TransactionBeforeToPB called before default ToPB code
type TransactionWithBeforeToPB interface {
	BeforeToPB(context.Context, *Transaction) error
}

// TransactionAfterToPB called after default ToPB code
type TransactionWithAfterToPB interface {
	AfterToPB(context.Context, *Transaction) error
}

// DefaultCreateReOccuringTransaction executes a basic gorm create call
func DefaultCreateReOccuringTransaction(ctx context.Context, in *ReOccuringTransaction, db *gorm.DB) (*ReOccuringTransaction, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ReOccuringTransactionORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ReOccuringTransactionORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type ReOccuringTransactionORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ReOccuringTransactionORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadReOccuringTransaction(ctx context.Context, in *ReOccuringTransaction, db *gorm.DB) (*ReOccuringTransaction, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == 0 {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(ReOccuringTransactionORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &ReOccuringTransactionORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ReOccuringTransactionORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := ReOccuringTransactionORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(ReOccuringTransactionORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type ReOccuringTransactionORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ReOccuringTransactionORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ReOccuringTransactionORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteReOccuringTransaction(ctx context.Context, in *ReOccuringTransaction, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(ReOccuringTransactionORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&ReOccuringTransactionORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(ReOccuringTransactionORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type ReOccuringTransactionORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ReOccuringTransactionORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteReOccuringTransactionSet(ctx context.Context, in []*ReOccuringTransaction, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	var err error
	keys := []uint64{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == 0 {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&ReOccuringTransactionORM{})).(ReOccuringTransactionORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&ReOccuringTransactionORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&ReOccuringTransactionORM{})).(ReOccuringTransactionORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type ReOccuringTransactionORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*ReOccuringTransaction, *gorm.DB) (*gorm.DB, error)
}
type ReOccuringTransactionORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*ReOccuringTransaction, *gorm.DB) error
}

// DefaultStrictUpdateReOccuringTransaction clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateReOccuringTransaction(ctx context.Context, in *ReOccuringTransaction, db *gorm.DB) (*ReOccuringTransaction, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateReOccuringTransaction")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &ReOccuringTransactionORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(ReOccuringTransactionORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(ReOccuringTransactionORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ReOccuringTransactionORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type ReOccuringTransactionORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ReOccuringTransactionORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ReOccuringTransactionORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchReOccuringTransaction executes a basic gorm update call with patch behavior
func DefaultPatchReOccuringTransaction(ctx context.Context, in *ReOccuringTransaction, updateMask *field_mask.FieldMask, db *gorm.DB) (*ReOccuringTransaction, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj ReOccuringTransaction
	var err error
	if hook, ok := interface{}(&pbObj).(ReOccuringTransactionWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadReOccuringTransaction(ctx, &ReOccuringTransaction{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(ReOccuringTransactionWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskReOccuringTransaction(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(ReOccuringTransactionWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateReOccuringTransaction(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(ReOccuringTransactionWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type ReOccuringTransactionWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *ReOccuringTransaction, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type ReOccuringTransactionWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *ReOccuringTransaction, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type ReOccuringTransactionWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *ReOccuringTransaction, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type ReOccuringTransactionWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *ReOccuringTransaction, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetReOccuringTransaction executes a bulk gorm update call with patch behavior
func DefaultPatchSetReOccuringTransaction(ctx context.Context, objects []*ReOccuringTransaction, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*ReOccuringTransaction, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*ReOccuringTransaction, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchReOccuringTransaction(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskReOccuringTransaction patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskReOccuringTransaction(ctx context.Context, patchee *ReOccuringTransaction, patcher *ReOccuringTransaction, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*ReOccuringTransaction, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"AccountId" {
			patchee.AccountId = patcher.AccountId
			continue
		}
		if f == prefix+"StreamId" {
			patchee.StreamId = patcher.StreamId
			continue
		}
		if f == prefix+"Category" {
			patchee.Category = patcher.Category
			continue
		}
		if f == prefix+"CategoryId" {
			patchee.CategoryId = patcher.CategoryId
			continue
		}
		if f == prefix+"Description" {
			patchee.Description = patcher.Description
			continue
		}
		if f == prefix+"MerchantName" {
			patchee.MerchantName = patcher.MerchantName
			continue
		}
		if f == prefix+"PersonalFinanceCategoryPrimary" {
			patchee.PersonalFinanceCategoryPrimary = patcher.PersonalFinanceCategoryPrimary
			continue
		}
		if f == prefix+"PersonalFinanceCategoryDetailed" {
			patchee.PersonalFinanceCategoryDetailed = patcher.PersonalFinanceCategoryDetailed
			continue
		}
		if f == prefix+"FirstDate" {
			patchee.FirstDate = patcher.FirstDate
			continue
		}
		if f == prefix+"LastDate" {
			patchee.LastDate = patcher.LastDate
			continue
		}
		if f == prefix+"Frequency" {
			patchee.Frequency = patcher.Frequency
			continue
		}
		if f == prefix+"TransactionIds" {
			patchee.TransactionIds = patcher.TransactionIds
			continue
		}
		if f == prefix+"AverageAmount" {
			patchee.AverageAmount = patcher.AverageAmount
			continue
		}
		if f == prefix+"AverageAmountIsoCurrencyCode" {
			patchee.AverageAmountIsoCurrencyCode = patcher.AverageAmountIsoCurrencyCode
			continue
		}
		if f == prefix+"LastAmount" {
			patchee.LastAmount = patcher.LastAmount
			continue
		}
		if f == prefix+"LastAmountIsoCurrencyCode" {
			patchee.LastAmountIsoCurrencyCode = patcher.LastAmountIsoCurrencyCode
			continue
		}
		if f == prefix+"IsActive" {
			patchee.IsActive = patcher.IsActive
			continue
		}
		if f == prefix+"Status" {
			patchee.Status = patcher.Status
			continue
		}
		if f == prefix+"UpdatedTime" {
			patchee.UpdatedTime = patcher.UpdatedTime
			continue
		}
		if f == prefix+"UserId" {
			patchee.UserId = patcher.UserId
			continue
		}
		if f == prefix+"LinkId" {
			patchee.LinkId = patcher.LinkId
			continue
		}
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListReOccuringTransaction executes a gorm list call
func DefaultListReOccuringTransaction(ctx context.Context, db *gorm.DB) ([]*ReOccuringTransaction, error) {
	in := ReOccuringTransaction{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ReOccuringTransactionORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &ReOccuringTransactionORM{}, &ReOccuringTransaction{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ReOccuringTransactionORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []ReOccuringTransactionORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ReOccuringTransactionORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*ReOccuringTransaction{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type ReOccuringTransactionORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ReOccuringTransactionORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ReOccuringTransactionORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]ReOccuringTransactionORM) error
}

// DefaultCreateTransaction executes a basic gorm create call
func DefaultCreateTransaction(ctx context.Context, in *Transaction, db *gorm.DB) (*Transaction, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type TransactionORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadTransaction(ctx context.Context, in *Transaction, db *gorm.DB) (*Transaction, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == 0 {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &TransactionORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := TransactionORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(TransactionORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type TransactionORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteTransaction(ctx context.Context, in *Transaction, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&TransactionORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type TransactionORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteTransactionSet(ctx context.Context, in []*Transaction, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	var err error
	keys := []uint64{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == 0 {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&TransactionORM{})).(TransactionORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&TransactionORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&TransactionORM{})).(TransactionORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type TransactionORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*Transaction, *gorm.DB) (*gorm.DB, error)
}
type TransactionORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*Transaction, *gorm.DB) error
}

// DefaultStrictUpdateTransaction clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateTransaction(ctx context.Context, in *Transaction, db *gorm.DB) (*Transaction, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateTransaction")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &TransactionORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(TransactionORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type TransactionORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchTransaction executes a basic gorm update call with patch behavior
func DefaultPatchTransaction(ctx context.Context, in *Transaction, updateMask *field_mask.FieldMask, db *gorm.DB) (*Transaction, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj Transaction
	var err error
	if hook, ok := interface{}(&pbObj).(TransactionWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadTransaction(ctx, &Transaction{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(TransactionWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskTransaction(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(TransactionWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateTransaction(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(TransactionWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type TransactionWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *Transaction, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type TransactionWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *Transaction, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type TransactionWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *Transaction, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type TransactionWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *Transaction, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetTransaction executes a bulk gorm update call with patch behavior
func DefaultPatchSetTransaction(ctx context.Context, objects []*Transaction, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*Transaction, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*Transaction, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchTransaction(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskTransaction patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskTransaction(ctx context.Context, patchee *Transaction, patcher *Transaction, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*Transaction, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	var updatedLocation bool
	var updatedPaymentMeta bool
	for i, f := range updateMask.Paths {
		if f == prefix+"AccountId" {
			patchee.AccountId = patcher.AccountId
			continue
		}
		if f == prefix+"Amount" {
			patchee.Amount = patcher.Amount
			continue
		}
		if f == prefix+"IsoCurrencyCode" {
			patchee.IsoCurrencyCode = patcher.IsoCurrencyCode
			continue
		}
		if f == prefix+"UnofficialCurrencyCode" {
			patchee.UnofficialCurrencyCode = patcher.UnofficialCurrencyCode
			continue
		}
		if f == prefix+"Category" {
			patchee.Category = patcher.Category
			continue
		}
		if f == prefix+"CategoryId" {
			patchee.CategoryId = patcher.CategoryId
			continue
		}
		if f == prefix+"CheckNumber" {
			patchee.CheckNumber = patcher.CheckNumber
			continue
		}
		if f == prefix+"Date" {
			patchee.Date = patcher.Date
			continue
		}
		if f == prefix+"Datetime" {
			patchee.Datetime = patcher.Datetime
			continue
		}
		if f == prefix+"AuthorizedDate" {
			patchee.AuthorizedDate = patcher.AuthorizedDate
			continue
		}
		if f == prefix+"AuthorizedDatetime" {
			patchee.AuthorizedDatetime = patcher.AuthorizedDatetime
			continue
		}
		if !updatedLocation && strings.HasPrefix(f, prefix+"Location.") {
			if patcher.Location == nil {
				patchee.Location = nil
				continue
			}
			if patchee.Location == nil {
				patchee.Location = &Transaction_Location{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"Location."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.Location, patchee.Location, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"Location" {
			updatedLocation = true
			patchee.Location = patcher.Location
			continue
		}
		if f == prefix+"Name" {
			patchee.Name = patcher.Name
			continue
		}
		if f == prefix+"MerchantName" {
			patchee.MerchantName = patcher.MerchantName
			continue
		}
		if !updatedPaymentMeta && strings.HasPrefix(f, prefix+"PaymentMeta.") {
			if patcher.PaymentMeta == nil {
				patchee.PaymentMeta = nil
				continue
			}
			if patchee.PaymentMeta == nil {
				patchee.PaymentMeta = &Transaction_PaymentMeta{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"PaymentMeta."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.PaymentMeta, patchee.PaymentMeta, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"PaymentMeta" {
			updatedPaymentMeta = true
			patchee.PaymentMeta = patcher.PaymentMeta
			continue
		}
		if f == prefix+"PaymentChannel" {
			patchee.PaymentChannel = patcher.PaymentChannel
			continue
		}
		if f == prefix+"Pending" {
			patchee.Pending = patcher.Pending
			continue
		}
		if f == prefix+"PendingTransactionId" {
			patchee.PendingTransactionId = patcher.PendingTransactionId
			continue
		}
		if f == prefix+"AccountOwner" {
			patchee.AccountOwner = patcher.AccountOwner
			continue
		}
		if f == prefix+"TransactionId" {
			patchee.TransactionId = patcher.TransactionId
			continue
		}
		if f == prefix+"TransactionCode" {
			patchee.TransactionCode = patcher.TransactionCode
			continue
		}
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"UserId" {
			patchee.UserId = patcher.UserId
			continue
		}
		if f == prefix+"LinkId" {
			patchee.LinkId = patcher.LinkId
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListTransaction executes a gorm list call
func DefaultListTransaction(ctx context.Context, db *gorm.DB) ([]*Transaction, error) {
	in := Transaction{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &TransactionORM{}, &Transaction{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []TransactionORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Transaction{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type TransactionORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]TransactionORM) error
}
