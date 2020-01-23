// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: signer.proto

package document_signerpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for DocumentSignerService service

type DocumentSignerService interface {
	CreateSignature(ctx context.Context, in *CreateSignatureRequest, opts ...client.CallOption) (*CreateSignatureResponse, error)
	GetSignatureUrl(ctx context.Context, in *GetSignatureUrlRequest, opts ...client.CallOption) (*GetSignatureUrlResponse, error)
}

type documentSignerService struct {
	c    client.Client
	name string
}

func NewDocumentSignerService(name string, c client.Client) DocumentSignerService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "document_signer"
	}
	return &documentSignerService{
		c:    c,
		name: name,
	}
}

func (c *documentSignerService) CreateSignature(ctx context.Context, in *CreateSignatureRequest, opts ...client.CallOption) (*CreateSignatureResponse, error) {
	req := c.c.NewRequest(c.name, "DocumentSignerService.CreateSignature", in)
	out := new(CreateSignatureResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentSignerService) GetSignatureUrl(ctx context.Context, in *GetSignatureUrlRequest, opts ...client.CallOption) (*GetSignatureUrlResponse, error) {
	req := c.c.NewRequest(c.name, "DocumentSignerService.GetSignatureUrl", in)
	out := new(GetSignatureUrlResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DocumentSignerService service

type DocumentSignerServiceHandler interface {
	CreateSignature(context.Context, *CreateSignatureRequest, *CreateSignatureResponse) error
	GetSignatureUrl(context.Context, *GetSignatureUrlRequest, *GetSignatureUrlResponse) error
}

func RegisterDocumentSignerServiceHandler(s server.Server, hdlr DocumentSignerServiceHandler, opts ...server.HandlerOption) error {
	type documentSignerService interface {
		CreateSignature(ctx context.Context, in *CreateSignatureRequest, out *CreateSignatureResponse) error
		GetSignatureUrl(ctx context.Context, in *GetSignatureUrlRequest, out *GetSignatureUrlResponse) error
	}
	type DocumentSignerService struct {
		documentSignerService
	}
	h := &documentSignerServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&DocumentSignerService{h}, opts...))
}

type documentSignerServiceHandler struct {
	DocumentSignerServiceHandler
}

func (h *documentSignerServiceHandler) CreateSignature(ctx context.Context, in *CreateSignatureRequest, out *CreateSignatureResponse) error {
	return h.DocumentSignerServiceHandler.CreateSignature(ctx, in, out)
}

func (h *documentSignerServiceHandler) GetSignatureUrl(ctx context.Context, in *GetSignatureUrlRequest, out *GetSignatureUrlResponse) error {
	return h.DocumentSignerServiceHandler.GetSignatureUrl(ctx, in, out)
}