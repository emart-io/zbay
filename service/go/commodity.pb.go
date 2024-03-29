// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.9.0
// source: commodity.proto

package zbay

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type Commodity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string    `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Category    string    `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Content     string    `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Media       []*Medium `protobuf:"bytes,5,rep,name=media,proto3" json:"media,omitempty"`
	City        string    `protobuf:"bytes,6,opt,name=city,proto3" json:"city,omitempty"`
	Prices      []*Price  `protobuf:"bytes,7,rep,name=prices,proto3" json:"prices,omitempty"`
	Inventory   uint32    `protobuf:"varint,15,opt,name=inventory,proto3" json:"inventory,omitempty"`
	ExpressFare uint32    `protobuf:"varint,8,opt,name=expressFare,proto3" json:"expressFare,omitempty"`
	Tags        []string  `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty"`
	OwnerId     string    `protobuf:"bytes,10,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
	// 已下线
	Status      string               `protobuf:"bytes,11,opt,name=status,proto3" json:"status,omitempty"`
	Annotations map[string]string    `protobuf:"bytes,12,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Created     *timestamp.Timestamp `protobuf:"bytes,13,opt,name=created,proto3" json:"created,omitempty"`
	Updated     *timestamp.Timestamp `protobuf:"bytes,14,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *Commodity) Reset() {
	*x = Commodity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commodity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commodity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commodity) ProtoMessage() {}

func (x *Commodity) ProtoReflect() protoreflect.Message {
	mi := &file_commodity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commodity.ProtoReflect.Descriptor instead.
func (*Commodity) Descriptor() ([]byte, []int) {
	return file_commodity_proto_rawDescGZIP(), []int{0}
}

func (x *Commodity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Commodity) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Commodity) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Commodity) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Commodity) GetMedia() []*Medium {
	if x != nil {
		return x.Media
	}
	return nil
}

func (x *Commodity) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Commodity) GetPrices() []*Price {
	if x != nil {
		return x.Prices
	}
	return nil
}

func (x *Commodity) GetInventory() uint32 {
	if x != nil {
		return x.Inventory
	}
	return 0
}

func (x *Commodity) GetExpressFare() uint32 {
	if x != nil {
		return x.ExpressFare
	}
	return 0
}

func (x *Commodity) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Commodity) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *Commodity) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Commodity) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *Commodity) GetCreated() *timestamp.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Commodity) GetUpdated() *timestamp.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ¥0.00
	Single string `protobuf:"bytes,2,opt,name=single,proto3" json:"single,omitempty"`
	Group  string `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commodity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_commodity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_commodity_proto_rawDescGZIP(), []int{1}
}

func (x *Price) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Price) GetSingle() string {
	if x != nil {
		return x.Single
	}
	return ""
}

func (x *Price) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

type Medium struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image   string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Video   string `protobuf:"bytes,2,opt,name=video,proto3" json:"video,omitempty"`
	IsFirst bool   `protobuf:"varint,3,opt,name=isFirst,proto3" json:"isFirst,omitempty"`
}

func (x *Medium) Reset() {
	*x = Medium{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commodity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Medium) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Medium) ProtoMessage() {}

func (x *Medium) ProtoReflect() protoreflect.Message {
	mi := &file_commodity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Medium.ProtoReflect.Descriptor instead.
func (*Medium) Descriptor() ([]byte, []int) {
	return file_commodity_proto_rawDescGZIP(), []int{2}
}

func (x *Medium) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Medium) GetVideo() string {
	if x != nil {
		return x.Video
	}
	return ""
}

func (x *Medium) GetIsFirst() bool {
	if x != nil {
		return x.IsFirst
	}
	return false
}

type Coupon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CommodityId  string               `protobuf:"bytes,3,opt,name=commodityId,proto3" json:"commodityId,omitempty"`
	Denomination uint32               `protobuf:"varint,4,opt,name=denomination,proto3" json:"denomination,omitempty"`
	Count        uint32               `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
	Limit        uint32               `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	Begin        *timestamp.Timestamp `protobuf:"bytes,7,opt,name=begin,proto3" json:"begin,omitempty"`
	End          *timestamp.Timestamp `protobuf:"bytes,8,opt,name=end,proto3" json:"end,omitempty"`
	Owner        string               `protobuf:"bytes,9,opt,name=owner,proto3" json:"owner,omitempty"`
	Annotations  map[string]string    `protobuf:"bytes,10,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Created      *timestamp.Timestamp `protobuf:"bytes,11,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *Coupon) Reset() {
	*x = Coupon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commodity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coupon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coupon) ProtoMessage() {}

func (x *Coupon) ProtoReflect() protoreflect.Message {
	mi := &file_commodity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coupon.ProtoReflect.Descriptor instead.
func (*Coupon) Descriptor() ([]byte, []int) {
	return file_commodity_proto_rawDescGZIP(), []int{3}
}

func (x *Coupon) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Coupon) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Coupon) GetCommodityId() string {
	if x != nil {
		return x.CommodityId
	}
	return ""
}

func (x *Coupon) GetDenomination() uint32 {
	if x != nil {
		return x.Denomination
	}
	return 0
}

func (x *Coupon) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Coupon) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Coupon) GetBegin() *timestamp.Timestamp {
	if x != nil {
		return x.Begin
	}
	return nil
}

func (x *Coupon) GetEnd() *timestamp.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *Coupon) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Coupon) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *Coupon) GetCreated() *timestamp.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CommodityId string               `protobuf:"bytes,2,opt,name=commodityId,proto3" json:"commodityId,omitempty"`
	Content     string               `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Keywords    []string             `protobuf:"bytes,4,rep,name=keywords,proto3" json:"keywords,omitempty"`
	Star        uint32               `protobuf:"varint,5,opt,name=star,proto3" json:"star,omitempty"`
	Owner       string               `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
	Annotations map[string]string    `protobuf:"bytes,7,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Created     *timestamp.Timestamp `protobuf:"bytes,8,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commodity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_commodity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_commodity_proto_rawDescGZIP(), []int{4}
}

func (x *Comment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Comment) GetCommodityId() string {
	if x != nil {
		return x.CommodityId
	}
	return ""
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetKeywords() []string {
	if x != nil {
		return x.Keywords
	}
	return nil
}

func (x *Comment) GetStar() uint32 {
	if x != nil {
		return x.Star
	}
	return 0
}

func (x *Comment) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Comment) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *Comment) GetCreated() *timestamp.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

var File_commodity_proto protoreflect.FileDescriptor

var file_commodity_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x7a, 0x62, 0x61, 0x79, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xba, 0x04, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x7a, 0x62, 0x61,
	0x79, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x52, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x52, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x46, 0x61, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x65, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x46, 0x61, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x42,
	0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0c, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x64, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x1a, 0x3e,
	0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x49,
	0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x4e, 0x0a, 0x06, 0x4d, 0x65, 0x64,
	0x69, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x12,
	0x18, 0x0a, 0x07, 0x69, 0x73, 0x46, 0x69, 0x72, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x69, 0x73, 0x46, 0x69, 0x72, 0x73, 0x74, 0x22, 0xcb, 0x03, 0x0a, 0x06, 0x43, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x64, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65,
	0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x62, 0x65,
	0x67, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x12, 0x2c, 0x0a, 0x03,
	0x65, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x12, 0x3f, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd3, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x74, 0x61, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x7a, 0x62, 0x61,
	0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x1a, 0x3e, 0x0a,
	0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xac, 0x02,
	0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x29, 0x0a,
	0x03, 0x41, 0x64, 0x64, 0x12, 0x0f, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x64, 0x69, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x0f, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79,
	0x1a, 0x0f, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e,
	0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x1a, 0x0f,
	0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x27, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0a, 0x2e, 0x7a, 0x62, 0x61, 0x79,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0f, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3b, 0x0a, 0x06, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x1a, 0x0f, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x69, 0x74, 0x79, 0x22, 0x00, 0x30, 0x01, 0x12, 0x33, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x0f, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69,
	0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x32, 0xd3, 0x01, 0x0a,
	0x07, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x12, 0x23, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12,
	0x0c, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x1a, 0x0c, 0x2e,
	0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x23, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x0c, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x1a, 0x0c, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x22, 0x00, 0x12, 0x26, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x7a,
	0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x1a, 0x0c, 0x2e, 0x7a, 0x62, 0x61,
	0x79, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x0a, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x1a, 0x0c, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x22, 0x00,
	0x30, 0x01, 0x32, 0xe1, 0x01, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x25, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0d, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x0d, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0d, 0x2e,
	0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x0d, 0x2e, 0x7a,
	0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x28, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x0d, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x0d, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x0f, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x69, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x7a, 0x62, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x42, 0x11, 0x5a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x67, 0x6f, 0x3b, 0x7a, 0x62, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_commodity_proto_rawDescOnce sync.Once
	file_commodity_proto_rawDescData = file_commodity_proto_rawDesc
)

func file_commodity_proto_rawDescGZIP() []byte {
	file_commodity_proto_rawDescOnce.Do(func() {
		file_commodity_proto_rawDescData = protoimpl.X.CompressGZIP(file_commodity_proto_rawDescData)
	})
	return file_commodity_proto_rawDescData
}

var file_commodity_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_commodity_proto_goTypes = []interface{}{
	(*Commodity)(nil),            // 0: zbay.Commodity
	(*Price)(nil),                // 1: zbay.Price
	(*Medium)(nil),               // 2: zbay.Medium
	(*Coupon)(nil),               // 3: zbay.Coupon
	(*Comment)(nil),              // 4: zbay.Comment
	nil,                          // 5: zbay.Commodity.AnnotationsEntry
	nil,                          // 6: zbay.Coupon.AnnotationsEntry
	nil,                          // 7: zbay.Comment.AnnotationsEntry
	(*timestamp.Timestamp)(nil),  // 8: google.protobuf.Timestamp
	(*User)(nil),                 // 9: zbay.User
	(*wrappers.StringValue)(nil), // 10: google.protobuf.StringValue
	(*empty.Empty)(nil),          // 11: google.protobuf.Empty
}
var file_commodity_proto_depIdxs = []int32{
	2,  // 0: zbay.Commodity.media:type_name -> zbay.Medium
	1,  // 1: zbay.Commodity.prices:type_name -> zbay.Price
	5,  // 2: zbay.Commodity.annotations:type_name -> zbay.Commodity.AnnotationsEntry
	8,  // 3: zbay.Commodity.created:type_name -> google.protobuf.Timestamp
	8,  // 4: zbay.Commodity.updated:type_name -> google.protobuf.Timestamp
	8,  // 5: zbay.Coupon.begin:type_name -> google.protobuf.Timestamp
	8,  // 6: zbay.Coupon.end:type_name -> google.protobuf.Timestamp
	6,  // 7: zbay.Coupon.annotations:type_name -> zbay.Coupon.AnnotationsEntry
	8,  // 8: zbay.Coupon.created:type_name -> google.protobuf.Timestamp
	7,  // 9: zbay.Comment.annotations:type_name -> zbay.Comment.AnnotationsEntry
	8,  // 10: zbay.Comment.created:type_name -> google.protobuf.Timestamp
	0,  // 11: zbay.Commodities.Add:input_type -> zbay.Commodity
	0,  // 12: zbay.Commodities.Get:input_type -> zbay.Commodity
	0,  // 13: zbay.Commodities.Update:input_type -> zbay.Commodity
	9,  // 14: zbay.Commodities.List:input_type -> zbay.User
	10, // 15: zbay.Commodities.Search:input_type -> google.protobuf.StringValue
	0,  // 16: zbay.Commodities.Delete:input_type -> zbay.Commodity
	3,  // 17: zbay.Coupons.Add:input_type -> zbay.Coupon
	3,  // 18: zbay.Coupons.Get:input_type -> zbay.Coupon
	3,  // 19: zbay.Coupons.Update:input_type -> zbay.Coupon
	3,  // 20: zbay.Coupons.Delete:input_type -> zbay.Coupon
	9,  // 21: zbay.Coupons.List:input_type -> zbay.User
	4,  // 22: zbay.Comments.Add:input_type -> zbay.Comment
	4,  // 23: zbay.Comments.Get:input_type -> zbay.Comment
	4,  // 24: zbay.Comments.Update:input_type -> zbay.Comment
	4,  // 25: zbay.Comments.Delete:input_type -> zbay.Comment
	0,  // 26: zbay.Comments.List:input_type -> zbay.Commodity
	0,  // 27: zbay.Commodities.Add:output_type -> zbay.Commodity
	0,  // 28: zbay.Commodities.Get:output_type -> zbay.Commodity
	0,  // 29: zbay.Commodities.Update:output_type -> zbay.Commodity
	0,  // 30: zbay.Commodities.List:output_type -> zbay.Commodity
	0,  // 31: zbay.Commodities.Search:output_type -> zbay.Commodity
	11, // 32: zbay.Commodities.Delete:output_type -> google.protobuf.Empty
	3,  // 33: zbay.Coupons.Add:output_type -> zbay.Coupon
	3,  // 34: zbay.Coupons.Get:output_type -> zbay.Coupon
	3,  // 35: zbay.Coupons.Update:output_type -> zbay.Coupon
	11, // 36: zbay.Coupons.Delete:output_type -> google.protobuf.Empty
	3,  // 37: zbay.Coupons.List:output_type -> zbay.Coupon
	4,  // 38: zbay.Comments.Add:output_type -> zbay.Comment
	4,  // 39: zbay.Comments.Get:output_type -> zbay.Comment
	4,  // 40: zbay.Comments.Update:output_type -> zbay.Comment
	11, // 41: zbay.Comments.Delete:output_type -> google.protobuf.Empty
	4,  // 42: zbay.Comments.List:output_type -> zbay.Comment
	27, // [27:43] is the sub-list for method output_type
	11, // [11:27] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_commodity_proto_init() }
func file_commodity_proto_init() {
	if File_commodity_proto != nil {
		return
	}
	file_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_commodity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commodity); i {
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
		file_commodity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
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
		file_commodity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Medium); i {
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
		file_commodity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coupon); i {
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
		file_commodity_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
			RawDescriptor: file_commodity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_commodity_proto_goTypes,
		DependencyIndexes: file_commodity_proto_depIdxs,
		MessageInfos:      file_commodity_proto_msgTypes,
	}.Build()
	File_commodity_proto = out.File
	file_commodity_proto_rawDesc = nil
	file_commodity_proto_goTypes = nil
	file_commodity_proto_depIdxs = nil
}
