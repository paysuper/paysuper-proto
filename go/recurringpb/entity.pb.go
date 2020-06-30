// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: entity.proto

package recurringpb

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type SavedCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Token      string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	ProjectId  string `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	MerchantId string `protobuf:"bytes,4,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	// @inject_tag: json:"pan"
	MaskedPan  string `protobuf:"bytes,5,opt,name=masked_pan,json=maskedPan,proto3" json:"pan"`
	CardHolder string `protobuf:"bytes,6,opt,name=card_holder,json=cardHolder,proto3" json:"card_holder,omitempty"`
	// @inject_tag: json:"expire"
	Expire      *CardExpire          `protobuf:"bytes,7,opt,name=expire,proto3" json:"expire"`
	RecurringId string               `protobuf:"bytes,8,opt,name=recurring_id,json=recurringId,proto3" json:"recurring_id,omitempty"`
	IsActive    bool                 `protobuf:"varint,9,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CreatedAt   *timestamp.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamp.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *SavedCard) Reset() {
	*x = SavedCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavedCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavedCard) ProtoMessage() {}

func (x *SavedCard) ProtoReflect() protoreflect.Message {
	mi := &file_entity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavedCard.ProtoReflect.Descriptor instead.
func (*SavedCard) Descriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{0}
}

func (x *SavedCard) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SavedCard) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SavedCard) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *SavedCard) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

func (x *SavedCard) GetMaskedPan() string {
	if x != nil {
		return x.MaskedPan
	}
	return ""
}

func (x *SavedCard) GetCardHolder() string {
	if x != nil {
		return x.CardHolder
	}
	return ""
}

func (x *SavedCard) GetExpire() *CardExpire {
	if x != nil {
		return x.Expire
	}
	return nil
}

func (x *SavedCard) GetRecurringId() string {
	if x != nil {
		return x.RecurringId
	}
	return ""
}

func (x *SavedCard) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *SavedCard) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *SavedCard) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CardExpire struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"month"
	Month string `protobuf:"bytes,1,opt,name=month,proto3" json:"month"`
	// @inject_tag: json:"year"
	Year string `protobuf:"bytes,2,opt,name=year,proto3" json:"year"`
}

func (x *CardExpire) Reset() {
	*x = CardExpire{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CardExpire) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardExpire) ProtoMessage() {}

func (x *CardExpire) ProtoReflect() protoreflect.Message {
	mi := &file_entity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardExpire.ProtoReflect.Descriptor instead.
func (*CardExpire) Descriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{1}
}

func (x *CardExpire) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

func (x *CardExpire) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

var File_entity_proto protoreflect.FileDescriptor

var file_entity_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x03, 0x0a, 0x09, 0x53, 0x61, 0x76, 0x65,
	0x64, 0x43, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d,
	0x61, 0x73, 0x6b, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x50, 0x61, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61,
	0x72, 0x64, 0x5f, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x61, 0x72, 0x64, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x06, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x52,
	0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x75, 0x72,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72,
	0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x36, 0x0a,
	0x0a, 0x43, 0x61, 0x72, 0x64, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x6f, 0x6e, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x79, 0x65, 0x61, 0x72, 0x42, 0x0d, 0x5a, 0x0b, 0x72, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69,
	0x6e, 0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entity_proto_rawDescOnce sync.Once
	file_entity_proto_rawDescData = file_entity_proto_rawDesc
)

func file_entity_proto_rawDescGZIP() []byte {
	file_entity_proto_rawDescOnce.Do(func() {
		file_entity_proto_rawDescData = protoimpl.X.CompressGZIP(file_entity_proto_rawDescData)
	})
	return file_entity_proto_rawDescData
}

var file_entity_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_entity_proto_goTypes = []interface{}{
	(*SavedCard)(nil),           // 0: entity.SavedCard
	(*CardExpire)(nil),          // 1: entity.CardExpire
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_entity_proto_depIdxs = []int32{
	1, // 0: entity.SavedCard.expire:type_name -> entity.CardExpire
	2, // 1: entity.SavedCard.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: entity.SavedCard.updated_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_entity_proto_init() }
func file_entity_proto_init() {
	if File_entity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavedCard); i {
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
		file_entity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CardExpire); i {
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
			RawDescriptor: file_entity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entity_proto_goTypes,
		DependencyIndexes: file_entity_proto_depIdxs,
		MessageInfos:      file_entity_proto_msgTypes,
	}.Build()
	File_entity_proto = out.File
	file_entity_proto_rawDesc = nil
	file_entity_proto_goTypes = nil
	file_entity_proto_depIdxs = nil
}
