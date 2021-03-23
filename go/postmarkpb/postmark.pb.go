// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: postmark.proto

package postmarkpb

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type PayloadHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//@inject_tag: json:"Name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"Name"`
	//@inject_tag: json:"Value"
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"Value"`
}

func (x *PayloadHeader) Reset() {
	*x = PayloadHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postmark_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadHeader) ProtoMessage() {}

func (x *PayloadHeader) ProtoReflect() protoreflect.Message {
	mi := &file_postmark_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadHeader.ProtoReflect.Descriptor instead.
func (*PayloadHeader) Descriptor() ([]byte, []int) {
	return file_postmark_proto_rawDescGZIP(), []int{0}
}

func (x *PayloadHeader) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PayloadHeader) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type PayloadAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//@inject_tag: json:"Name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"Name"`
	//@inject_tag: json:"Content"
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"Content"`
	//@inject_tag: json:"ContentType"
	ContentType string `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"ContentType"`
}

func (x *PayloadAttachment) Reset() {
	*x = PayloadAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postmark_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadAttachment) ProtoMessage() {}

func (x *PayloadAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_postmark_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadAttachment.ProtoReflect.Descriptor instead.
func (*PayloadAttachment) Descriptor() ([]byte, []int) {
	return file_postmark_proto_rawDescGZIP(), []int{1}
}

func (x *PayloadAttachment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PayloadAttachment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *PayloadAttachment) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//@inject_tag: protobuf:"varint,1,opt,name=template_id,json=TemplateId,proto3"
	TemplateId int32 `protobuf:"varint,1,opt,name=template_id,json=TemplateId,proto3" json:"template_id,omitempty"`
	//@inject_tag: `protobuf:"bytes,2,opt,name=template_alias,json=TemplateAlias,proto3"
	TemplateAlias string `protobuf:"bytes,2,opt,name=template_alias,json=TemplateAlias,proto3" json:"template_alias,omitempty"`
	//@inject_tag: protobuf:"bytes,3,rep,name=template_model,proto3" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"
	TemplateModel map[string]string `protobuf:"bytes,3,rep,name=template_model,proto3" json:"template_model,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//@inject_tag: protobuf:"varint,4,opt,name=inline_css,json=InlineCss,proto3"
	InlineCss bool `protobuf:"varint,4,opt,name=inline_css,json=InlineCss,proto3" json:"inline_css,omitempty"`
	//@inject_tag: protobuf:"bytes,5,opt,name=from,json=From,proto3"
	From string `protobuf:"bytes,5,opt,name=from,json=From,proto3" json:"from,omitempty"`
	//@inject_tag: protobuf:"bytes,6,opt,name=to,json=To,proto3"
	To string `protobuf:"bytes,6,opt,name=to,json=To,proto3" json:"to,omitempty"`
	//@inject_tag: protobuf:"bytes,7,opt,name=cc,json=Cc,proto3"
	Cc string `protobuf:"bytes,7,opt,name=cc,json=Cc,proto3" json:"cc,omitempty"`
	//@inject_tag: protobuf:"bytes,8,opt,name=bcc,json=Bcc,proto3"
	Bcc string `protobuf:"bytes,8,opt,name=bcc,json=Bcc,proto3" json:"bcc,omitempty"`
	//@inject_tag: protobuf:"bytes,9,opt,name=tag,json=Tag,proto3"
	Tag string `protobuf:"bytes,9,opt,name=tag,json=Tag,proto3" json:"tag,omitempty"`
	//@inject_tag: protobuf:"bytes,10,opt,name=reply_to,json=ReplyTo,proto3"
	ReplyTo string `protobuf:"bytes,10,opt,name=reply_to,json=ReplyTo,proto3" json:"reply_to,omitempty"`
	//@inject_tag: protobuf:"bytes,11,rep,name=headers,json=Headers,proto3"
	Headers []*PayloadHeader `protobuf:"bytes,11,rep,name=headers,json=Headers,proto3" json:"headers,omitempty"`
	//@inject_tag: protobuf:"varint,12,opt,name=track_opens,json=TrackOpens,proto3"
	TrackOpens bool `protobuf:"varint,12,opt,name=track_opens,json=TrackOpens,proto3" json:"track_opens,omitempty"`
	//@inject_tag: protobuf:"bytes,13,opt,name=track_links,json=TrackLinks,proto3"
	TrackLinks string `protobuf:"bytes,13,opt,name=track_links,json=TrackLinks,proto3" json:"track_links,omitempty"`
	//@inject_tag: protobuf:"bytes,14,rep,name=Attachments,json=Attachments,proto3"
	Attachments []*PayloadAttachment `protobuf:"bytes,14,rep,name=Attachments,json=Attachments,proto3" json:"attachments,omitempty"`
	//@inject_tag: protobuf:"bytes,15,rep,name=metadata,json=Metadata,proto3" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"
	Metadata            map[string]string `protobuf:"bytes,15,rep,name=metadata,json=Metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TemplateObjectModel *_struct.Struct   `protobuf:"bytes,16,opt,name=template_object_model,json=templateObjectModel,proto3" json:"template_object_model,omitempty"`
	PostmarkApiToken    string            `protobuf:"bytes,17,opt,name=postmark_api_token,json=postmarkApiToken,proto3" json:"postmark_api_token,omitempty"`
}

func (x *Payload) Reset() {
	*x = Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postmark_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_postmark_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_postmark_proto_rawDescGZIP(), []int{2}
}

func (x *Payload) GetTemplateId() int32 {
	if x != nil {
		return x.TemplateId
	}
	return 0
}

func (x *Payload) GetTemplateAlias() string {
	if x != nil {
		return x.TemplateAlias
	}
	return ""
}

func (x *Payload) GetTemplateModel() map[string]string {
	if x != nil {
		return x.TemplateModel
	}
	return nil
}

func (x *Payload) GetInlineCss() bool {
	if x != nil {
		return x.InlineCss
	}
	return false
}

func (x *Payload) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Payload) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Payload) GetCc() string {
	if x != nil {
		return x.Cc
	}
	return ""
}

func (x *Payload) GetBcc() string {
	if x != nil {
		return x.Bcc
	}
	return ""
}

func (x *Payload) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Payload) GetReplyTo() string {
	if x != nil {
		return x.ReplyTo
	}
	return ""
}

func (x *Payload) GetHeaders() []*PayloadHeader {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Payload) GetTrackOpens() bool {
	if x != nil {
		return x.TrackOpens
	}
	return false
}

func (x *Payload) GetTrackLinks() string {
	if x != nil {
		return x.TrackLinks
	}
	return ""
}

func (x *Payload) GetAttachments() []*PayloadAttachment {
	if x != nil {
		return x.Attachments
	}
	return nil
}

func (x *Payload) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Payload) GetTemplateObjectModel() *_struct.Struct {
	if x != nil {
		return x.TemplateObjectModel
	}
	return nil
}

func (x *Payload) GetPostmarkApiToken() string {
	if x != nil {
		return x.PostmarkApiToken
	}
	return ""
}

var File_postmark_proto protoreflect.FileDescriptor

var file_postmark_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x64, 0x0a, 0x11, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x9b, 0x06, 0x0a, 0x07, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x5f, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x4b, 0x0a,
	0x0e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x72, 0x6b,
	0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x63, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x74, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x63, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x63, 0x63, 0x12, 0x10, 0x0a,
	0x03, 0x62, 0x63, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x63, 0x63, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61,
	0x67, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x12, 0x31, 0x0a, 0x07,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x4f, 0x70, 0x65, 0x6e, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x4c, 0x69, 0x6e, 0x6b,
	0x73, 0x12, 0x3d, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x72,
	0x6b, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x3b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0f, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x2e, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x4b, 0x0a,
	0x15, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x13, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x6f,
	0x73, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x72, 0x6b,
	0x41, 0x70, 0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x40, 0x0a, 0x12, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x3b, 0x70, 0x6f,
	0x73, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_postmark_proto_rawDescOnce sync.Once
	file_postmark_proto_rawDescData = file_postmark_proto_rawDesc
)

func file_postmark_proto_rawDescGZIP() []byte {
	file_postmark_proto_rawDescOnce.Do(func() {
		file_postmark_proto_rawDescData = protoimpl.X.CompressGZIP(file_postmark_proto_rawDescData)
	})
	return file_postmark_proto_rawDescData
}

var file_postmark_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_postmark_proto_goTypes = []interface{}{
	(*PayloadHeader)(nil),     // 0: postmark.PayloadHeader
	(*PayloadAttachment)(nil), // 1: postmark.PayloadAttachment
	(*Payload)(nil),           // 2: postmark.Payload
	nil,                       // 3: postmark.Payload.TemplateModelEntry
	nil,                       // 4: postmark.Payload.MetadataEntry
	(*_struct.Struct)(nil),    // 5: google.protobuf.Struct
}
var file_postmark_proto_depIdxs = []int32{
	3, // 0: postmark.Payload.template_model:type_name -> postmark.Payload.TemplateModelEntry
	0, // 1: postmark.Payload.headers:type_name -> postmark.PayloadHeader
	1, // 2: postmark.Payload.attachments:type_name -> postmark.PayloadAttachment
	4, // 3: postmark.Payload.metadata:type_name -> postmark.Payload.MetadataEntry
	5, // 4: postmark.Payload.template_object_model:type_name -> google.protobuf.Struct
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_postmark_proto_init() }
func file_postmark_proto_init() {
	if File_postmark_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_postmark_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadHeader); i {
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
		file_postmark_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadAttachment); i {
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
		file_postmark_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payload); i {
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
			RawDescriptor: file_postmark_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_postmark_proto_goTypes,
		DependencyIndexes: file_postmark_proto_depIdxs,
		MessageInfos:      file_postmark_proto_msgTypes,
	}.Build()
	File_postmark_proto = out.File
	file_postmark_proto_rawDesc = nil
	file_postmark_proto_goTypes = nil
	file_postmark_proto_depIdxs = nil
}
