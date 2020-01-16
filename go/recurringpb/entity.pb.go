// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entity.proto

package recurringpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SavedCard struct {
	// @inject_tag: json:"id"
	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Token      string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	ProjectId  string `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	MerchantId string `protobuf:"bytes,4,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	// @inject_tag: json:"pan"
	MaskedPan  string `protobuf:"bytes,5,opt,name=masked_pan,json=maskedPan,proto3" json:"pan"`
	CardHolder string `protobuf:"bytes,6,opt,name=card_holder,json=cardHolder,proto3" json:"card_holder,omitempty"`
	// @inject_tag: json:"expire"
	Expire               *CardExpire          `protobuf:"bytes,7,opt,name=expire,proto3" json:"expire"`
	RecurringId          string               `protobuf:"bytes,8,opt,name=recurring_id,json=recurringId,proto3" json:"recurring_id,omitempty"`
	IsActive             bool                 `protobuf:"varint,9,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-" bson:"-" structure:"-" validate:"-"`
	XXX_unrecognized     []byte               `json:"-" bson:"-" structure:"-" validate:"-"`
	XXX_sizecache        int32                `json:"-" bson:"-" structure:"-" validate:"-"`
}

func (m *SavedCard) Reset()         { *m = SavedCard{} }
func (m *SavedCard) String() string { return proto.CompactTextString(m) }
func (*SavedCard) ProtoMessage()    {}
func (*SavedCard) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf50d946d740d100, []int{0}
}

func (m *SavedCard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavedCard.Unmarshal(m, b)
}
func (m *SavedCard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavedCard.Marshal(b, m, deterministic)
}
func (m *SavedCard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavedCard.Merge(m, src)
}
func (m *SavedCard) XXX_Size() int {
	return xxx_messageInfo_SavedCard.Size(m)
}
func (m *SavedCard) XXX_DiscardUnknown() {
	xxx_messageInfo_SavedCard.DiscardUnknown(m)
}

var xxx_messageInfo_SavedCard proto.InternalMessageInfo

func (m *SavedCard) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SavedCard) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SavedCard) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *SavedCard) GetMerchantId() string {
	if m != nil {
		return m.MerchantId
	}
	return ""
}

func (m *SavedCard) GetMaskedPan() string {
	if m != nil {
		return m.MaskedPan
	}
	return ""
}

func (m *SavedCard) GetCardHolder() string {
	if m != nil {
		return m.CardHolder
	}
	return ""
}

func (m *SavedCard) GetExpire() *CardExpire {
	if m != nil {
		return m.Expire
	}
	return nil
}

func (m *SavedCard) GetRecurringId() string {
	if m != nil {
		return m.RecurringId
	}
	return ""
}

func (m *SavedCard) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *SavedCard) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *SavedCard) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type CardExpire struct {
	// @inject_tag: json:"month"
	Month string `protobuf:"bytes,1,opt,name=month,proto3" json:"month"`
	// @inject_tag: json:"year"
	Year                 string   `protobuf:"bytes,2,opt,name=year,proto3" json:"year"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-" structure:"-" validate:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-" structure:"-" validate:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-" structure:"-" validate:"-"`
}

func (m *CardExpire) Reset()         { *m = CardExpire{} }
func (m *CardExpire) String() string { return proto.CompactTextString(m) }
func (*CardExpire) ProtoMessage()    {}
func (*CardExpire) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf50d946d740d100, []int{1}
}

func (m *CardExpire) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CardExpire.Unmarshal(m, b)
}
func (m *CardExpire) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CardExpire.Marshal(b, m, deterministic)
}
func (m *CardExpire) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CardExpire.Merge(m, src)
}
func (m *CardExpire) XXX_Size() int {
	return xxx_messageInfo_CardExpire.Size(m)
}
func (m *CardExpire) XXX_DiscardUnknown() {
	xxx_messageInfo_CardExpire.DiscardUnknown(m)
}

var xxx_messageInfo_CardExpire proto.InternalMessageInfo

func (m *CardExpire) GetMonth() string {
	if m != nil {
		return m.Month
	}
	return ""
}

func (m *CardExpire) GetYear() string {
	if m != nil {
		return m.Year
	}
	return ""
}

func init() {
	proto.RegisterType((*SavedCard)(nil), "entity.SavedCard")
	proto.RegisterType((*CardExpire)(nil), "entity.CardExpire")
}

func init() { proto.RegisterFile("entity.proto", fileDescriptor_cf50d946d740d100) }

var fileDescriptor_cf50d946d740d100 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x3d, 0x4f, 0xfb, 0x30,
	0x10, 0xc6, 0x95, 0xbe, 0xe4, 0xdf, 0x5c, 0xfa, 0x67, 0xb0, 0x18, 0xac, 0x22, 0xd4, 0xd2, 0xa9,
	0x62, 0x48, 0x25, 0x90, 0x90, 0x18, 0x0b, 0x42, 0xa2, 0x1b, 0x0a, 0x4c, 0x2c, 0x91, 0x1b, 0x1f,
	0xad, 0x69, 0x13, 0x5b, 0x8e, 0x5b, 0xd1, 0xcf, 0xc1, 0x17, 0x46, 0x7e, 0x09, 0x1d, 0xd9, 0x72,
	0xbf, 0x7b, 0x9e, 0x9c, 0x9f, 0x3b, 0x18, 0x62, 0x6d, 0x84, 0x39, 0x66, 0x4a, 0x4b, 0x23, 0x49,
	0xec, 0xab, 0xd1, 0x78, 0x2d, 0xe5, 0x7a, 0x87, 0x73, 0x47, 0x57, 0xfb, 0x8f, 0xb9, 0x11, 0x15,
	0x36, 0x86, 0x55, 0xca, 0x0b, 0xa7, 0xdf, 0x5d, 0x48, 0x5e, 0xd9, 0x01, 0xf9, 0x23, 0xd3, 0x9c,
	0x9c, 0x41, 0x47, 0x70, 0x1a, 0x4d, 0xa2, 0x59, 0x92, 0x77, 0x04, 0x27, 0xe7, 0xd0, 0x37, 0x72,
	0x8b, 0x35, 0xed, 0x38, 0xe4, 0x0b, 0x72, 0x09, 0xa0, 0xb4, 0xfc, 0xc4, 0xd2, 0x14, 0x82, 0xd3,
	0xae, 0x6b, 0x25, 0x81, 0x2c, 0x39, 0x19, 0x43, 0x5a, 0xa1, 0x2e, 0x37, 0xac, 0x76, 0xfd, 0x9e,
	0xeb, 0x43, 0x8b, 0x96, 0xdc, 0xfa, 0x2b, 0xd6, 0x6c, 0x91, 0x17, 0x8a, 0xd5, 0xb4, 0xef, 0xfd,
	0x9e, 0xbc, 0xb0, 0xda, 0xfa, 0x4b, 0xa6, 0x79, 0xb1, 0x91, 0x3b, 0x8e, 0x9a, 0xc6, 0xde, 0x6f,
	0xd1, 0xb3, 0x23, 0xe4, 0x1a, 0x62, 0xfc, 0x52, 0x42, 0x23, 0xfd, 0x37, 0x89, 0x66, 0xe9, 0x0d,
	0xc9, 0x42, 0x76, 0x9b, 0xe1, 0xc9, 0x75, 0xf2, 0xa0, 0x20, 0x57, 0x30, 0xd4, 0x58, 0xee, 0xb5,
	0x16, 0xf5, 0xda, 0xbe, 0x66, 0xe0, 0xfe, 0x96, 0xfe, 0xb2, 0x25, 0x27, 0x17, 0x90, 0x88, 0xa6,
	0x60, 0xa5, 0x11, 0x07, 0xa4, 0xc9, 0x24, 0x9a, 0x0d, 0xf2, 0x81, 0x68, 0x16, 0xae, 0x26, 0xf7,
	0x00, 0xa5, 0x46, 0x66, 0x90, 0x17, 0xcc, 0x50, 0x70, 0xf3, 0x46, 0x99, 0xdf, 0x6a, 0xd6, 0x6e,
	0x35, 0x7b, 0x6b, 0xb7, 0x9a, 0x27, 0x41, 0xbd, 0x30, 0xd6, 0xba, 0x57, 0xbc, 0xb5, 0xa6, 0x7f,
	0x5b, 0x83, 0x7a, 0x61, 0xa6, 0x77, 0x00, 0xa7, 0x2c, 0xf6, 0x0a, 0x95, 0xac, 0xcd, 0x26, 0x1c,
	0xc6, 0x17, 0x84, 0x40, 0xef, 0x88, 0x4c, 0x87, 0xd3, 0xb8, 0xef, 0x87, 0xff, 0xef, 0xa7, 0x64,
	0x6a, 0xb5, 0x8a, 0xdd, 0x94, 0xdb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x64, 0x4a, 0xd1, 0xcc,
	0x1c, 0x02, 0x00, 0x00,
}