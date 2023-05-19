// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: api/v1/clickhouse.proto

package apiv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId              string                   `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Amount                 float64                  `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	IsoCurrencyCode        string                   `protobuf:"bytes,3,opt,name=iso_currency_code,json=isoCurrencyCode,proto3" json:"iso_currency_code,omitempty"`
	UnofficialCurrencyCode string                   `protobuf:"bytes,4,opt,name=unofficial_currency_code,json=unofficialCurrencyCode,proto3" json:"unofficial_currency_code,omitempty"`
	Category               []string                 `protobuf:"bytes,5,rep,name=category,proto3" json:"category,omitempty"`
	CategoryId             string                   `protobuf:"bytes,6,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	CheckNumber            string                   `protobuf:"bytes,7,opt,name=check_number,json=checkNumber,proto3" json:"check_number,omitempty"`
	Date                   string                   `protobuf:"bytes,8,opt,name=date,proto3" json:"date,omitempty"`
	Datetime               string                   `protobuf:"bytes,9,opt,name=datetime,proto3" json:"datetime,omitempty"`
	AuthorizedDate         string                   `protobuf:"bytes,10,opt,name=authorized_date,json=authorizedDate,proto3" json:"authorized_date,omitempty"`
	AuthorizedDatetime     string                   `protobuf:"bytes,11,opt,name=authorized_datetime,json=authorizedDatetime,proto3" json:"authorized_datetime,omitempty"`
	Location               *Transaction_Location    `protobuf:"bytes,12,opt,name=location,proto3" json:"location,omitempty"`
	Name                   string                   `protobuf:"bytes,13,opt,name=name,proto3" json:"name,omitempty"`
	MerchantName           string                   `protobuf:"bytes,14,opt,name=merchant_name,json=merchantName,proto3" json:"merchant_name,omitempty"`
	PaymentMeta            *Transaction_PaymentMeta `protobuf:"bytes,15,opt,name=payment_meta,json=paymentMeta,proto3" json:"payment_meta,omitempty"`
	PaymentChannel         string                   `protobuf:"bytes,16,opt,name=payment_channel,json=paymentChannel,proto3" json:"payment_channel,omitempty"`
	Pending                bool                     `protobuf:"varint,17,opt,name=pending,proto3" json:"pending,omitempty"`
	PendingTransactionId   string                   `protobuf:"bytes,18,opt,name=pending_transaction_id,json=pendingTransactionId,proto3" json:"pending_transaction_id,omitempty"`
	AccountOwner           string                   `protobuf:"bytes,19,opt,name=account_owner,json=accountOwner,proto3" json:"account_owner,omitempty"`
	TransactionId          string                   `protobuf:"bytes,20,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	TransactionCode        string                   `protobuf:"bytes,21,opt,name=transaction_code,json=transactionCode,proto3" json:"transaction_code,omitempty"`
	Id                     uint64                   `protobuf:"varint,22,opt,name=id,proto3" json:"id,omitempty"`
	UserId                 uint64                   `protobuf:"varint,23,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_clickhouse_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_clickhouse_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_api_v1_clickhouse_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Transaction) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Transaction) GetIsoCurrencyCode() string {
	if x != nil {
		return x.IsoCurrencyCode
	}
	return ""
}

func (x *Transaction) GetUnofficialCurrencyCode() string {
	if x != nil {
		return x.UnofficialCurrencyCode
	}
	return ""
}

func (x *Transaction) GetCategory() []string {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *Transaction) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *Transaction) GetCheckNumber() string {
	if x != nil {
		return x.CheckNumber
	}
	return ""
}

func (x *Transaction) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Transaction) GetDatetime() string {
	if x != nil {
		return x.Datetime
	}
	return ""
}

func (x *Transaction) GetAuthorizedDate() string {
	if x != nil {
		return x.AuthorizedDate
	}
	return ""
}

func (x *Transaction) GetAuthorizedDatetime() string {
	if x != nil {
		return x.AuthorizedDatetime
	}
	return ""
}

func (x *Transaction) GetLocation() *Transaction_Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Transaction) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Transaction) GetMerchantName() string {
	if x != nil {
		return x.MerchantName
	}
	return ""
}

func (x *Transaction) GetPaymentMeta() *Transaction_PaymentMeta {
	if x != nil {
		return x.PaymentMeta
	}
	return nil
}

func (x *Transaction) GetPaymentChannel() string {
	if x != nil {
		return x.PaymentChannel
	}
	return ""
}

func (x *Transaction) GetPending() bool {
	if x != nil {
		return x.Pending
	}
	return false
}

func (x *Transaction) GetPendingTransactionId() string {
	if x != nil {
		return x.PendingTransactionId
	}
	return ""
}

func (x *Transaction) GetAccountOwner() string {
	if x != nil {
		return x.AccountOwner
	}
	return ""
}

func (x *Transaction) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *Transaction) GetTransactionCode() string {
	if x != nil {
		return x.TransactionCode
	}
	return ""
}

func (x *Transaction) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Transaction) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Transaction_Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address     string  `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	City        string  `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Region      string  `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	PostalCode  string  `protobuf:"bytes,4,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	Country     string  `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	Lat         float64 `protobuf:"fixed64,6,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon         float64 `protobuf:"fixed64,7,opt,name=lon,proto3" json:"lon,omitempty"`
	StoreNumber string  `protobuf:"bytes,8,opt,name=store_number,json=storeNumber,proto3" json:"store_number,omitempty"`
}

func (x *Transaction_Location) Reset() {
	*x = Transaction_Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_clickhouse_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction_Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction_Location) ProtoMessage() {}

func (x *Transaction_Location) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_clickhouse_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction_Location.ProtoReflect.Descriptor instead.
func (*Transaction_Location) Descriptor() ([]byte, []int) {
	return file_api_v1_clickhouse_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Transaction_Location) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Transaction_Location) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Transaction_Location) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Transaction_Location) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *Transaction_Location) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Transaction_Location) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Transaction_Location) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

func (x *Transaction_Location) GetStoreNumber() string {
	if x != nil {
		return x.StoreNumber
	}
	return ""
}

type Transaction_PaymentMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ByOrderOf        string `protobuf:"bytes,1,opt,name=by_order_of,json=byOrderOf,proto3" json:"by_order_of,omitempty"`
	Payee            string `protobuf:"bytes,2,opt,name=payee,proto3" json:"payee,omitempty"`
	Payer            string `protobuf:"bytes,3,opt,name=payer,proto3" json:"payer,omitempty"`
	PaymentMethod    string `protobuf:"bytes,4,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"`
	PaymentProcessor string `protobuf:"bytes,5,opt,name=payment_processor,json=paymentProcessor,proto3" json:"payment_processor,omitempty"`
	PpdId            string `protobuf:"bytes,6,opt,name=ppd_id,json=ppdId,proto3" json:"ppd_id,omitempty"`
	Reason           string `protobuf:"bytes,7,opt,name=reason,proto3" json:"reason,omitempty"`
	ReferenceNumber  string `protobuf:"bytes,8,opt,name=reference_number,json=referenceNumber,proto3" json:"reference_number,omitempty"`
}

func (x *Transaction_PaymentMeta) Reset() {
	*x = Transaction_PaymentMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_clickhouse_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction_PaymentMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction_PaymentMeta) ProtoMessage() {}

func (x *Transaction_PaymentMeta) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_clickhouse_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction_PaymentMeta.ProtoReflect.Descriptor instead.
func (*Transaction_PaymentMeta) Descriptor() ([]byte, []int) {
	return file_api_v1_clickhouse_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Transaction_PaymentMeta) GetByOrderOf() string {
	if x != nil {
		return x.ByOrderOf
	}
	return ""
}

func (x *Transaction_PaymentMeta) GetPayee() string {
	if x != nil {
		return x.Payee
	}
	return ""
}

func (x *Transaction_PaymentMeta) GetPayer() string {
	if x != nil {
		return x.Payer
	}
	return ""
}

func (x *Transaction_PaymentMeta) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

func (x *Transaction_PaymentMeta) GetPaymentProcessor() string {
	if x != nil {
		return x.PaymentProcessor
	}
	return ""
}

func (x *Transaction_PaymentMeta) GetPpdId() string {
	if x != nil {
		return x.PpdId
	}
	return ""
}

func (x *Transaction_PaymentMeta) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Transaction_PaymentMeta) GetReferenceNumber() string {
	if x != nil {
		return x.ReferenceNumber
	}
	return ""
}

var File_api_v1_clickhouse_proto protoreflect.FileDescriptor

var file_api_v1_clickhouse_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f,
	0x72, 0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc,
	0x0a, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73, 0x6f, 0x5f, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x69, 0x73, 0x6f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x38, 0x0a, 0x18, 0x75, 0x6e, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x5f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x16, 0x75, 0x6e, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52,
	0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x0f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12,
	0x34, 0x0a, 0x16, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x14, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x13, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x04, 0x42, 0x03, 0x92, 0x41, 0x00, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x23, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x17, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x0a, 0x92, 0x41, 0x00, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0xd2, 0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73,
	0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x87, 0x02, 0x0a, 0x0b,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0b, 0x62,
	0x79, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x62, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4f, 0x66, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x61, 0x79, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x79, 0x65,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x79, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x61, 0x79, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x2b,
	0x0a, 0x11, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x15, 0x0a, 0x06, 0x70,
	0x70, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x70, 0x64,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x42, 0xbc, 0x01,
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x43, 0x6c,
	0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x64, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x69, 0x6d, 0x69,
	0x66, 0x69, 0x6e, 0x69, 0x69, 0x43, 0x54, 0x4f, 0x2f, 0x73, 0x69, 0x6d, 0x66, 0x69, 0x6e, 0x79,
	0x2d, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x69, 0x61, 0x6c, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b,
	0x61, 0x70, 0x69, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x41, 0x70,
	0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x06, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x12,
	0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x07, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_clickhouse_proto_rawDescOnce sync.Once
	file_api_v1_clickhouse_proto_rawDescData = file_api_v1_clickhouse_proto_rawDesc
)

func file_api_v1_clickhouse_proto_rawDescGZIP() []byte {
	file_api_v1_clickhouse_proto_rawDescOnce.Do(func() {
		file_api_v1_clickhouse_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_clickhouse_proto_rawDescData)
	})
	return file_api_v1_clickhouse_proto_rawDescData
}

var file_api_v1_clickhouse_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_v1_clickhouse_proto_goTypes = []interface{}{
	(*Transaction)(nil),             // 0: api.v1.Transaction
	(*Transaction_Location)(nil),    // 1: api.v1.Transaction.Location
	(*Transaction_PaymentMeta)(nil), // 2: api.v1.Transaction.PaymentMeta
}
var file_api_v1_clickhouse_proto_depIdxs = []int32{
	1, // 0: api.v1.Transaction.location:type_name -> api.v1.Transaction.Location
	2, // 1: api.v1.Transaction.payment_meta:type_name -> api.v1.Transaction.PaymentMeta
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1_clickhouse_proto_init() }
func file_api_v1_clickhouse_proto_init() {
	if File_api_v1_clickhouse_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_clickhouse_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_clickhouse_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction_Location); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_clickhouse_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction_PaymentMeta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_clickhouse_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_clickhouse_proto_goTypes,
		DependencyIndexes: file_api_v1_clickhouse_proto_depIdxs,
		MessageInfos:      file_api_v1_clickhouse_proto_msgTypes,
	}.Build()
	File_api_v1_clickhouse_proto = out.File
	file_api_v1_clickhouse_proto_rawDesc = nil
	file_api_v1_clickhouse_proto_goTypes = nil
	file_api_v1_clickhouse_proto_depIdxs = nil
}
