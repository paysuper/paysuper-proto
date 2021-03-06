// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.13.0
// source: repository.proto

package recurringpb

import (
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

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repository_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_repository_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_repository_proto_rawDescGZIP(), []int{0}
}

type FindByStringValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *FindByStringValue) Reset() {
	*x = FindByStringValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repository_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByStringValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByStringValue) ProtoMessage() {}

func (x *FindByStringValue) ProtoReflect() protoreflect.Message {
	mi := &file_repository_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByStringValue.ProtoReflect.Descriptor instead.
func (*FindByStringValue) Descriptor() ([]byte, []int) {
	return file_repository_proto_rawDescGZIP(), []int{1}
}

func (x *FindByStringValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SavedCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string      `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ProjectId   string      `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	MerchantId  string      `protobuf:"bytes,3,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	MaskedPan   string      `protobuf:"bytes,4,opt,name=masked_pan,json=maskedPan,proto3" json:"masked_pan,omitempty"`
	CardHolder  string      `protobuf:"bytes,5,opt,name=card_holder,json=cardHolder,proto3" json:"card_holder,omitempty"`
	RecurringId string      `protobuf:"bytes,6,opt,name=recurring_id,json=recurringId,proto3" json:"recurring_id,omitempty"`
	Expire      *CardExpire `protobuf:"bytes,7,opt,name=expire,proto3" json:"expire,omitempty"`
}

func (x *SavedCardRequest) Reset() {
	*x = SavedCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repository_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavedCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavedCardRequest) ProtoMessage() {}

func (x *SavedCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_repository_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavedCardRequest.ProtoReflect.Descriptor instead.
func (*SavedCardRequest) Descriptor() ([]byte, []int) {
	return file_repository_proto_rawDescGZIP(), []int{2}
}

func (x *SavedCardRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SavedCardRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *SavedCardRequest) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

func (x *SavedCardRequest) GetMaskedPan() string {
	if x != nil {
		return x.MaskedPan
	}
	return ""
}

func (x *SavedCardRequest) GetCardHolder() string {
	if x != nil {
		return x.CardHolder
	}
	return ""
}

func (x *SavedCardRequest) GetRecurringId() string {
	if x != nil {
		return x.RecurringId
	}
	return ""
}

func (x *SavedCardRequest) GetExpire() *CardExpire {
	if x != nil {
		return x.Expire
	}
	return nil
}

type SavedCardList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SavedCards []*SavedCard `protobuf:"bytes,1,rep,name=saved_cards,json=savedCards,proto3" json:"saved_cards,omitempty"`
}

func (x *SavedCardList) Reset() {
	*x = SavedCardList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repository_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavedCardList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavedCardList) ProtoMessage() {}

func (x *SavedCardList) ProtoReflect() protoreflect.Message {
	mi := &file_repository_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavedCardList.ProtoReflect.Descriptor instead.
func (*SavedCardList) Descriptor() ([]byte, []int) {
	return file_repository_proto_rawDescGZIP(), []int{3}
}

func (x *SavedCardList) GetSavedCards() []*SavedCard {
	if x != nil {
		return x.SavedCards
	}
	return nil
}

type DeleteSavedCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *DeleteSavedCardRequest) Reset() {
	*x = DeleteSavedCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repository_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSavedCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSavedCardRequest) ProtoMessage() {}

func (x *DeleteSavedCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_repository_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSavedCardRequest.ProtoReflect.Descriptor instead.
func (*DeleteSavedCardRequest) Descriptor() ([]byte, []int) {
	return file_repository_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteSavedCardRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteSavedCardRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type DeleteSavedCardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteSavedCardResponse) Reset() {
	*x = DeleteSavedCardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repository_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSavedCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSavedCardResponse) ProtoMessage() {}

func (x *DeleteSavedCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_repository_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSavedCardResponse.ProtoReflect.Descriptor instead.
func (*DeleteSavedCardResponse) Descriptor() ([]byte, []int) {
	return file_repository_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteSavedCardResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DeleteSavedCardResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_repository_proto protoreflect.FileDescriptor

var file_repository_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x1a, 0x0c,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x08, 0x0a, 0x06,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x29, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0xf7, 0x01, 0x0a, 0x10, 0x53, 0x61, 0x76, 0x65, 0x64, 0x43, 0x61, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x50, 0x61, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x61, 0x72, 0x64, 0x5f, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c,
	0x72, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12,
	0x2a, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x45, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x22, 0x43, 0x0a, 0x0d, 0x53,
	0x61, 0x76, 0x65, 0x64, 0x43, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x0b,
	0x73, 0x61, 0x76, 0x65, 0x64, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x64,
	0x43, 0x61, 0x72, 0x64, 0x52, 0x0a, 0x73, 0x61, 0x76, 0x65, 0x64, 0x43, 0x61, 0x72, 0x64, 0x73,
	0x22, 0x3e, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x61, 0x76, 0x65, 0x64, 0x43,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x4b, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x61, 0x76, 0x65, 0x64, 0x43,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xc7, 0x02,
	0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x45, 0x0a, 0x0f,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x43, 0x61, 0x72, 0x64, 0x12,
	0x1c, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x53, 0x61, 0x76,
	0x65, 0x64, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x61, 0x76,
	0x65, 0x64, 0x43, 0x61, 0x72, 0x64, 0x12, 0x22, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x61, 0x76, 0x65, 0x64, 0x43,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x72, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x61,
	0x76, 0x65, 0x64, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4b, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x61, 0x76, 0x65, 0x64, 0x43, 0x61,
	0x72, 0x64, 0x73, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x53, 0x61, 0x76, 0x65, 0x64, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x53,
	0x61, 0x76, 0x65, 0x64, 0x43, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x47,
	0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x61, 0x76, 0x65, 0x64, 0x43, 0x61, 0x72, 0x64, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x1d, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x1a, 0x11, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x61, 0x76, 0x65,
	0x64, 0x43, 0x61, 0x72, 0x64, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x3b, 0x72, 0x65,
	0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_repository_proto_rawDescOnce sync.Once
	file_repository_proto_rawDescData = file_repository_proto_rawDesc
)

func file_repository_proto_rawDescGZIP() []byte {
	file_repository_proto_rawDescOnce.Do(func() {
		file_repository_proto_rawDescData = protoimpl.X.CompressGZIP(file_repository_proto_rawDescData)
	})
	return file_repository_proto_rawDescData
}

var file_repository_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_repository_proto_goTypes = []interface{}{
	(*Result)(nil),                  // 0: repository.Result
	(*FindByStringValue)(nil),       // 1: repository.FindByStringValue
	(*SavedCardRequest)(nil),        // 2: repository.SavedCardRequest
	(*SavedCardList)(nil),           // 3: repository.SavedCardList
	(*DeleteSavedCardRequest)(nil),  // 4: repository.DeleteSavedCardRequest
	(*DeleteSavedCardResponse)(nil), // 5: repository.DeleteSavedCardResponse
	(*CardExpire)(nil),              // 6: entity.CardExpire
	(*SavedCard)(nil),               // 7: entity.SavedCard
}
var file_repository_proto_depIdxs = []int32{
	6, // 0: repository.SavedCardRequest.expire:type_name -> entity.CardExpire
	7, // 1: repository.SavedCardList.saved_cards:type_name -> entity.SavedCard
	2, // 2: repository.Repository.InsertSavedCard:input_type -> repository.SavedCardRequest
	4, // 3: repository.Repository.DeleteSavedCard:input_type -> repository.DeleteSavedCardRequest
	2, // 4: repository.Repository.FindSavedCards:input_type -> repository.SavedCardRequest
	1, // 5: repository.Repository.FindSavedCardById:input_type -> repository.FindByStringValue
	0, // 6: repository.Repository.InsertSavedCard:output_type -> repository.Result
	5, // 7: repository.Repository.DeleteSavedCard:output_type -> repository.DeleteSavedCardResponse
	3, // 8: repository.Repository.FindSavedCards:output_type -> repository.SavedCardList
	7, // 9: repository.Repository.FindSavedCardById:output_type -> entity.SavedCard
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_repository_proto_init() }
func file_repository_proto_init() {
	if File_repository_proto != nil {
		return
	}
	file_entity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_repository_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_repository_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByStringValue); i {
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
		file_repository_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavedCardRequest); i {
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
		file_repository_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavedCardList); i {
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
		file_repository_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSavedCardRequest); i {
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
		file_repository_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSavedCardResponse); i {
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
			RawDescriptor: file_repository_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_repository_proto_goTypes,
		DependencyIndexes: file_repository_proto_depIdxs,
		MessageInfos:      file_repository_proto_msgTypes,
	}.Build()
	File_repository_proto = out.File
	file_repository_proto_rawDesc = nil
	file_repository_proto_goTypes = nil
	file_repository_proto_depIdxs = nil
}
