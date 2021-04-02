// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: order.proto

package zbay

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      string                 `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Amount      int64                  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	OrderId     string                 `protobuf:"bytes,4,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Created     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	Annotations map[string]string      `protobuf:"bytes,6,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Account) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Account) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Account) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Account) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Account) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

type PayMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string            `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Kv  map[string]string `protobuf:"bytes,2,rep,name=kv,proto3" json:"kv,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PayMap) Reset() {
	*x = PayMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayMap) ProtoMessage() {}

func (x *PayMap) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayMap.ProtoReflect.Descriptor instead.
func (*PayMap) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{1}
}

func (x *PayMap) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PayMap) GetKv() map[string]string {
	if x != nil {
		return x.Kv
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Snapshot    *Commodity `protobuf:"bytes,12,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
	Groupon     *Groupon   `protobuf:"bytes,15,opt,name=groupon,proto3" json:"groupon,omitempty"`
	Price       *Price     `protobuf:"bytes,13,opt,name=price,proto3" json:"price,omitempty"`
	UserId      string     `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
	Destination *Address   `protobuf:"bytes,4,opt,name=destination,proto3" json:"destination,omitempty"`
	Quantity    uint32     `protobuf:"varint,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Amount      uint64     `protobuf:"varint,6,opt,name=amount,proto3" json:"amount,omitempty"`
	// 待成团|待发货|待收货|待评价|待退款|已退款
	Status      string                 `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Comment     string                 `protobuf:"bytes,8,opt,name=comment,proto3" json:"comment,omitempty"`
	PayInfo     *PayInfo               `protobuf:"bytes,9,opt,name=payInfo,proto3" json:"payInfo,omitempty"`
	Express     *Express               `protobuf:"bytes,14,opt,name=express,proto3" json:"express,omitempty"`
	Annotations map[string]string      `protobuf:"bytes,10,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Created     *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{2}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetSnapshot() *Commodity {
	if x != nil {
		return x.Snapshot
	}
	return nil
}

func (x *Order) GetGroupon() *Groupon {
	if x != nil {
		return x.Groupon
	}
	return nil
}

func (x *Order) GetPrice() *Price {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *Order) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Order) GetDestination() *Address {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *Order) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Order) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Order) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Order) GetPayInfo() *PayInfo {
	if x != nil {
		return x.PayInfo
	}
	return nil
}

func (x *Order) GetExpress() *Express {
	if x != nil {
		return x.Express
	}
	return nil
}

func (x *Order) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *Order) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

type PayInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//alipay|weixin
	Type      string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	PayResult string `protobuf:"bytes,2,opt,name=payResult,proto3" json:"payResult,omitempty"`
}

func (x *PayInfo) Reset() {
	*x = PayInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayInfo) ProtoMessage() {}

func (x *PayInfo) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayInfo.ProtoReflect.Descriptor instead.
func (*PayInfo) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{3}
}

func (x *PayInfo) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PayInfo) GetPayResult() string {
	if x != nil {
		return x.PayResult
	}
	return ""
}

type WechatPayParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid      string `protobuf:"bytes,7,opt,name=appid,proto3" json:"appid,omitempty"`
	Partnerid  string `protobuf:"bytes,1,opt,name=partnerid,proto3" json:"partnerid,omitempty"`
	Prepayid   string `protobuf:"bytes,2,opt,name=prepayid,proto3" json:"prepayid,omitempty"`
	Noncestr   string `protobuf:"bytes,3,opt,name=noncestr,proto3" json:"noncestr,omitempty"`
	Timestamp  string `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Sign       string `protobuf:"bytes,5,opt,name=sign,proto3" json:"sign,omitempty"`
	MwebUrl    string `protobuf:"bytes,6,opt,name=mwebUrl,proto3" json:"mwebUrl,omitempty"`
	OutTradeNo string `protobuf:"bytes,8,opt,name=outTradeNo,proto3" json:"outTradeNo,omitempty"`
}

func (x *WechatPayParams) Reset() {
	*x = WechatPayParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WechatPayParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WechatPayParams) ProtoMessage() {}

func (x *WechatPayParams) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WechatPayParams.ProtoReflect.Descriptor instead.
func (*WechatPayParams) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{4}
}

func (x *WechatPayParams) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *WechatPayParams) GetPartnerid() string {
	if x != nil {
		return x.Partnerid
	}
	return ""
}

func (x *WechatPayParams) GetPrepayid() string {
	if x != nil {
		return x.Prepayid
	}
	return ""
}

func (x *WechatPayParams) GetNoncestr() string {
	if x != nil {
		return x.Noncestr
	}
	return ""
}

func (x *WechatPayParams) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *WechatPayParams) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

func (x *WechatPayParams) GetMwebUrl() string {
	if x != nil {
		return x.MwebUrl
	}
	return ""
}

func (x *WechatPayParams) GetOutTradeNo() string {
	if x != nil {
		return x.OutTradeNo
	}
	return ""
}

type Groupon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderIds []string `protobuf:"bytes,1,rep,name=orderIds,proto3" json:"orderIds,omitempty"`
}

func (x *Groupon) Reset() {
	*x = Groupon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Groupon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Groupon) ProtoMessage() {}

func (x *Groupon) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Groupon.ProtoReflect.Descriptor instead.
func (*Groupon) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{5}
}

func (x *Groupon) GetOrderIds() []string {
	if x != nil {
		return x.OrderIds
	}
	return nil
}

type Express struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Company string `protobuf:"bytes,1,opt,name=company,proto3" json:"company,omitempty"`
	Number  string `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *Express) Reset() {
	*x = Express{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Express) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Express) ProtoMessage() {}

func (x *Express) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Express.ProtoReflect.Descriptor instead.
func (*Express) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{6}
}

func (x *Express) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *Express) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

type ListQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ListQuery) Reset() {
	*x = ListQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQuery) ProtoMessage() {}

func (x *ListQuery) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQuery.ProtoReflect.Descriptor instead.
func (*ListQuery) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{7}
}

func (x *ListQuery) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *ListQuery) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_order_proto protoreflect.FileDescriptor

var file_order_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x7a,
	0x62, 0x61, 0x79, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b,
	0x02, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x40, 0x0a, 0x0b, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x41,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3e, 0x0a, 0x10,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x77, 0x0a, 0x06,
	0x50, 0x61, 0x79, 0x4d, 0x61, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x24, 0x0a, 0x02, 0x6b, 0x76, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x50, 0x61, 0x79, 0x4d,
	0x61, 0x70, 0x2e, 0x4b, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x02, 0x6b, 0x76, 0x1a, 0x35,
	0x0a, 0x07, 0x4b, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc7, 0x04, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x2b, 0x0a, 0x08, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69,
	0x74, 0x79, 0x52, 0x08, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x27, 0x0a, 0x07,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x7a, 0x62, 0x61, 0x79, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x2f, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e,
	0x50, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x70, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x27, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x07, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3e, 0x0a, 0x0b, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x1a,
	0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x3b, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x61, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xe9, 0x01, 0x0a,
	0x0f, 0x57, 0x65, 0x63, 0x68, 0x61, 0x74, 0x50, 0x61, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x70, 0x61, 0x79, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x65, 0x70, 0x61, 0x79, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x72, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x67, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x77, 0x65, 0x62, 0x55, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x77, 0x65, 0x62, 0x55, 0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x75, 0x74, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x75,
	0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x22, 0x25, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22,
	0x3b, 0x0a, 0x07, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x43, 0x0a, 0x09,
	0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x32, 0xb7, 0x02, 0x0a, 0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x03,
	0x41, 0x64, 0x64, 0x12, 0x0b, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x1a, 0x0b, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x00, 0x12,
	0x21, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0b, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x1a, 0x0b, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x22, 0x00, 0x12, 0x24, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x7a,
	0x62, 0x61, 0x79, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x0b, 0x2e, 0x7a, 0x62, 0x61, 0x79,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x0b, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0b, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0b, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x0b, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x22, 0x00, 0x30, 0x01, 0x12, 0x30, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f,
	0x72, 0x42, 0x75, 0x79, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0b, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x22, 0x00, 0x30, 0x01, 0x12, 0x31, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x6f, 0x72, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x7a, 0x62, 0x61, 0x79,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0b, 0x2e, 0x7a, 0x62, 0x61,
	0x79, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x00, 0x30, 0x01, 0x32, 0x86, 0x02, 0x0a, 0x08,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12,
	0x0d, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x0d,
	0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12,
	0x25, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0a, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x1a, 0x0d, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x27, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x0d, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x0d,
	0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12,
	0x26, 0x0a, 0x06, 0x41, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x12, 0x0c, 0x2e, 0x7a, 0x62, 0x61, 0x79,
	0x2e, 0x50, 0x61, 0x79, 0x4d, 0x61, 0x70, 0x1a, 0x0c, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x50,
	0x61, 0x79, 0x4d, 0x61, 0x70, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0e, 0x57, 0x65, 0x63, 0x68, 0x61,
	0x74, 0x48, 0x74, 0x6d, 0x6c, 0x35, 0x50, 0x61, 0x79, 0x12, 0x0c, 0x2e, 0x7a, 0x62, 0x61, 0x79,
	0x2e, 0x50, 0x61, 0x79, 0x4d, 0x61, 0x70, 0x1a, 0x0c, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x50,
	0x61, 0x79, 0x4d, 0x61, 0x70, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0b, 0x57, 0x65, 0x63, 0x68, 0x61,
	0x74, 0x4a, 0x53, 0x50, 0x61, 0x79, 0x12, 0x0c, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x50, 0x61,
	0x79, 0x4d, 0x61, 0x70, 0x1a, 0x0c, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x50, 0x61, 0x79, 0x4d,
	0x61, 0x70, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x67, 0x6f, 0x3b, 0x7a, 0x62, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData = file_order_proto_rawDesc
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_proto_rawDescData)
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_order_proto_goTypes = []interface{}{
	(*Account)(nil),               // 0: zbay.Account
	(*PayMap)(nil),                // 1: zbay.PayMap
	(*Order)(nil),                 // 2: zbay.Order
	(*PayInfo)(nil),               // 3: zbay.PayInfo
	(*WechatPayParams)(nil),       // 4: zbay.WechatPayParams
	(*Groupon)(nil),               // 5: zbay.Groupon
	(*Express)(nil),               // 6: zbay.Express
	(*ListQuery)(nil),             // 7: zbay.ListQuery
	nil,                           // 8: zbay.Account.AnnotationsEntry
	nil,                           // 9: zbay.PayMap.KvEntry
	nil,                           // 10: zbay.Order.AnnotationsEntry
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
	(*Commodity)(nil),             // 12: zbay.Commodity
	(*Price)(nil),                 // 13: zbay.Price
	(*Address)(nil),               // 14: zbay.Address
	(*User)(nil),                  // 15: zbay.User
	(*emptypb.Empty)(nil),         // 16: google.protobuf.Empty
}
var file_order_proto_depIdxs = []int32{
	11, // 0: zbay.Account.created:type_name -> google.protobuf.Timestamp
	8,  // 1: zbay.Account.annotations:type_name -> zbay.Account.AnnotationsEntry
	9,  // 2: zbay.PayMap.kv:type_name -> zbay.PayMap.KvEntry
	12, // 3: zbay.Order.snapshot:type_name -> zbay.Commodity
	5,  // 4: zbay.Order.groupon:type_name -> zbay.Groupon
	13, // 5: zbay.Order.price:type_name -> zbay.Price
	14, // 6: zbay.Order.destination:type_name -> zbay.Address
	3,  // 7: zbay.Order.payInfo:type_name -> zbay.PayInfo
	6,  // 8: zbay.Order.express:type_name -> zbay.Express
	10, // 9: zbay.Order.annotations:type_name -> zbay.Order.AnnotationsEntry
	11, // 10: zbay.Order.created:type_name -> google.protobuf.Timestamp
	15, // 11: zbay.ListQuery.user:type_name -> zbay.User
	2,  // 12: zbay.Orders.Add:input_type -> zbay.Order
	2,  // 13: zbay.Orders.Get:input_type -> zbay.Order
	2,  // 14: zbay.Orders.Update:input_type -> zbay.Order
	2,  // 15: zbay.Orders.Delete:input_type -> zbay.Order
	2,  // 16: zbay.Orders.ListByOrder:input_type -> zbay.Order
	7,  // 17: zbay.Orders.ListForBuyer:input_type -> zbay.ListQuery
	7,  // 18: zbay.Orders.ListForSeller:input_type -> zbay.ListQuery
	0,  // 19: zbay.Accounts.Add:input_type -> zbay.Account
	15, // 20: zbay.Accounts.List:input_type -> zbay.User
	0,  // 21: zbay.Accounts.Total:input_type -> zbay.Account
	1,  // 22: zbay.Accounts.Alipay:input_type -> zbay.PayMap
	1,  // 23: zbay.Accounts.WechatHtml5Pay:input_type -> zbay.PayMap
	1,  // 24: zbay.Accounts.WechatJSPay:input_type -> zbay.PayMap
	2,  // 25: zbay.Orders.Add:output_type -> zbay.Order
	2,  // 26: zbay.Orders.Get:output_type -> zbay.Order
	2,  // 27: zbay.Orders.Update:output_type -> zbay.Order
	16, // 28: zbay.Orders.Delete:output_type -> google.protobuf.Empty
	2,  // 29: zbay.Orders.ListByOrder:output_type -> zbay.Order
	2,  // 30: zbay.Orders.ListForBuyer:output_type -> zbay.Order
	2,  // 31: zbay.Orders.ListForSeller:output_type -> zbay.Order
	0,  // 32: zbay.Accounts.Add:output_type -> zbay.Account
	0,  // 33: zbay.Accounts.List:output_type -> zbay.Account
	0,  // 34: zbay.Accounts.Total:output_type -> zbay.Account
	1,  // 35: zbay.Accounts.Alipay:output_type -> zbay.PayMap
	1,  // 36: zbay.Accounts.WechatHtml5Pay:output_type -> zbay.PayMap
	1,  // 37: zbay.Accounts.WechatJSPay:output_type -> zbay.PayMap
	25, // [25:38] is the sub-list for method output_type
	12, // [12:25] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	file_user_proto_init()
	file_commodity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayMap); i {
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
		file_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayInfo); i {
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
		file_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WechatPayParams); i {
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
		file_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Groupon); i {
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
		file_order_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Express); i {
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
		file_order_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQuery); i {
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
			RawDescriptor: file_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_rawDesc = nil
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}
