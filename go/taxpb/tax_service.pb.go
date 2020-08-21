// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: tax_service.proto

package taxpb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GeoIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"zip"
	Zip string `protobuf:"bytes,1,opt,name=zip,proto3" json:"zip"`
	// @inject_tag: json:"country"
	Country string `protobuf:"bytes,2,opt,name=country,proto3" json:"country"`
	// @inject_tag: json:"city,omitempty"
	City string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	// @inject_tag: json:"state,omitempty"
	State string `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *GeoIdentity) Reset() {
	*x = GeoIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tax_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoIdentity) ProtoMessage() {}

func (x *GeoIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_tax_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoIdentity.ProtoReflect.Descriptor instead.
func (*GeoIdentity) Descriptor() ([]byte, []int) {
	return file_tax_service_proto_rawDescGZIP(), []int{0}
}

func (x *GeoIdentity) GetZip() string {
	if x != nil {
		return x.Zip
	}
	return ""
}

func (x *GeoIdentity) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *GeoIdentity) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *GeoIdentity) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type TaxRate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	//
	// The unique identifier for the tax rate.
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// @inject_tag: json:"zip"
	//
	// The ZIP code.
	Zip string `protobuf:"bytes,2,opt,name=zip,proto3" json:"zip"`
	// @inject_tag: json:"country"
	//
	// The country's name. Two-letter country code in ISO 3166-1, in uppercase (for instance US).
	Country string `protobuf:"bytes,3,opt,name=country,proto3" json:"country"`
	// @inject_tag: json:"city,omitempty"
	//
	// The city.
	City string `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	// @inject_tag: gorm:"type:varchar(2)"
	//
	// The state code in ISO 3166-2.
	State string `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty" gorm:"type:varchar(2)"`
	// @inject_tag: json:"rate"
	//
	// The tax rate.
	Rate float64 `protobuf:"fixed64,6,opt,name=rate,proto3" json:"rate"`
}

func (x *TaxRate) Reset() {
	*x = TaxRate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tax_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaxRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaxRate) ProtoMessage() {}

func (x *TaxRate) ProtoReflect() protoreflect.Message {
	mi := &file_tax_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaxRate.ProtoReflect.Descriptor instead.
func (*TaxRate) Descriptor() ([]byte, []int) {
	return file_tax_service_proto_rawDescGZIP(), []int{1}
}

func (x *TaxRate) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaxRate) GetZip() string {
	if x != nil {
		return x.Zip
	}
	return ""
}

func (x *TaxRate) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *TaxRate) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *TaxRate) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *TaxRate) GetRate() float64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

type GetRatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"zip"
	Zip string `protobuf:"bytes,1,opt,name=zip,proto3" json:"zip"`
	// @inject_tag: json:"country"
	Country string `protobuf:"bytes,2,opt,name=country,proto3" json:"country"`
	// @inject_tag: json:"city,omitempty"
	City string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	// @inject_tag: json:"state,omitempty"
	State string `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	// @inject_tag: json:"limit,omitempty"
	Limit int32 `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	// @inject_tag: json:"offset,omitempty"
	Offset int32 `protobuf:"varint,6,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetRatesRequest) Reset() {
	*x = GetRatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tax_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRatesRequest) ProtoMessage() {}

func (x *GetRatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tax_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRatesRequest.ProtoReflect.Descriptor instead.
func (*GetRatesRequest) Descriptor() ([]byte, []int) {
	return file_tax_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetRatesRequest) GetZip() string {
	if x != nil {
		return x.Zip
	}
	return ""
}

func (x *GetRatesRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *GetRatesRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *GetRatesRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *GetRatesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetRatesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetRatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"rates"
	Rates []*TaxRate `protobuf:"bytes,1,rep,name=rates,proto3" json:"rates"`
}

func (x *GetRatesResponse) Reset() {
	*x = GetRatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tax_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRatesResponse) ProtoMessage() {}

func (x *GetRatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tax_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRatesResponse.ProtoReflect.Descriptor instead.
func (*GetRatesResponse) Descriptor() ([]byte, []int) {
	return file_tax_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetRatesResponse) GetRates() []*TaxRate {
	if x != nil {
		return x.Rates
	}
	return nil
}

type DeleteRateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteRateResponse) Reset() {
	*x = DeleteRateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tax_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRateResponse) ProtoMessage() {}

func (x *DeleteRateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tax_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRateResponse.ProtoReflect.Descriptor instead.
func (*DeleteRateResponse) Descriptor() ([]byte, []int) {
	return file_tax_service_proto_rawDescGZIP(), []int{4}
}

type DeleteRateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRateRequest) Reset() {
	*x = DeleteRateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tax_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRateRequest) ProtoMessage() {}

func (x *DeleteRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tax_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRateRequest.ProtoReflect.Descriptor instead.
func (*DeleteRateRequest) Descriptor() ([]byte, []int) {
	return file_tax_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteRateRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_tax_service_proto protoreflect.FileDescriptor

var file_tax_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x61, 0x78, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x61, 0x78, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x63, 0x0a, 0x0b, 0x47, 0x65, 0x6f, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x7a, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x7a, 0x69,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x07, 0x54, 0x61, 0x78, 0x52, 0x61, 0x74,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x7a, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x7a, 0x69, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x7a, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x7a, 0x69,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x22, 0x3e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x72, 0x61, 0x74, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x61, 0x78, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x78, 0x52, 0x61, 0x74, 0x65, 0x52, 0x05, 0x72, 0x61,
	0x74, 0x65, 0x73, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x32, 0xa9,
	0x02, 0x0a, 0x0a, 0x54, 0x61, 0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x74, 0x61, 0x78, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x6f, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x1a, 0x14, 0x2e, 0x74, 0x61, 0x78, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x54, 0x61, 0x78, 0x52, 0x61, 0x74, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x52, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x74, 0x61, 0x78, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x61, 0x78, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x74, 0x61, 0x78, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x78, 0x52, 0x61, 0x74, 0x65, 0x1a, 0x14, 0x2e,
	0x74, 0x61, 0x78, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x78, 0x52,
	0x61, 0x74, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x61, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1e, 0x2e, 0x74, 0x61, 0x78, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x61, 0x78, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b,
	0x74, 0x61, 0x78, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tax_service_proto_rawDescOnce sync.Once
	file_tax_service_proto_rawDescData = file_tax_service_proto_rawDesc
)

func file_tax_service_proto_rawDescGZIP() []byte {
	file_tax_service_proto_rawDescOnce.Do(func() {
		file_tax_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_tax_service_proto_rawDescData)
	})
	return file_tax_service_proto_rawDescData
}

var file_tax_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tax_service_proto_goTypes = []interface{}{
	(*GeoIdentity)(nil),        // 0: tax_service.GeoIdentity
	(*TaxRate)(nil),            // 1: tax_service.TaxRate
	(*GetRatesRequest)(nil),    // 2: tax_service.GetRatesRequest
	(*GetRatesResponse)(nil),   // 3: tax_service.GetRatesResponse
	(*DeleteRateResponse)(nil), // 4: tax_service.DeleteRateResponse
	(*DeleteRateRequest)(nil),  // 5: tax_service.DeleteRateRequest
}
var file_tax_service_proto_depIdxs = []int32{
	1, // 0: tax_service.GetRatesResponse.rates:type_name -> tax_service.TaxRate
	0, // 1: tax_service.TaxService.GetRate:input_type -> tax_service.GeoIdentity
	2, // 2: tax_service.TaxService.GetRates:input_type -> tax_service.GetRatesRequest
	1, // 3: tax_service.TaxService.CreateOrUpdate:input_type -> tax_service.TaxRate
	5, // 4: tax_service.TaxService.DeleteRateById:input_type -> tax_service.DeleteRateRequest
	1, // 5: tax_service.TaxService.GetRate:output_type -> tax_service.TaxRate
	3, // 6: tax_service.TaxService.GetRates:output_type -> tax_service.GetRatesResponse
	1, // 7: tax_service.TaxService.CreateOrUpdate:output_type -> tax_service.TaxRate
	4, // 8: tax_service.TaxService.DeleteRateById:output_type -> tax_service.DeleteRateResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tax_service_proto_init() }
func file_tax_service_proto_init() {
	if File_tax_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tax_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoIdentity); i {
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
		file_tax_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaxRate); i {
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
		file_tax_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRatesRequest); i {
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
		file_tax_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRatesResponse); i {
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
		file_tax_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRateResponse); i {
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
		file_tax_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRateRequest); i {
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
			RawDescriptor: file_tax_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tax_service_proto_goTypes,
		DependencyIndexes: file_tax_service_proto_depIdxs,
		MessageInfos:      file_tax_service_proto_msgTypes,
	}.Build()
	File_tax_service_proto = out.File
	file_tax_service_proto_rawDesc = nil
	file_tax_service_proto_goTypes = nil
	file_tax_service_proto_depIdxs = nil
}
