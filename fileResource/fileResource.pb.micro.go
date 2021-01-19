// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: fileResource/fileResource.proto

package fileResource

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/anypb"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for FileResource service

func NewFileResourceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for FileResource service

type FileResourceService interface {
	Add(ctx context.Context, in *FileDTO, opts ...client.CallOption) (*Response, error)
	Delete(ctx context.Context, in *DeleteId, opts ...client.CallOption) (*Response, error)
	Update(ctx context.Context, in *FileDTO, opts ...client.CallOption) (*Response, error)
	HistoricalData(ctx context.Context, in *DataType, opts ...client.CallOption) (*Response, error)
}

type fileResourceService struct {
	c    client.Client
	name string
}

func NewFileResourceService(name string, c client.Client) FileResourceService {
	return &fileResourceService{
		c:    c,
		name: name,
	}
}

func (c *fileResourceService) Add(ctx context.Context, in *FileDTO, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "FileResource.Add", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileResourceService) Delete(ctx context.Context, in *DeleteId, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "FileResource.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileResourceService) Update(ctx context.Context, in *FileDTO, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "FileResource.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileResourceService) HistoricalData(ctx context.Context, in *DataType, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "FileResource.HistoricalData", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FileResource service

type FileResourceHandler interface {
	Add(context.Context, *FileDTO, *Response) error
	Delete(context.Context, *DeleteId, *Response) error
	Update(context.Context, *FileDTO, *Response) error
	HistoricalData(context.Context, *DataType, *Response) error
}

func RegisterFileResourceHandler(s server.Server, hdlr FileResourceHandler, opts ...server.HandlerOption) error {
	type fileResource interface {
		Add(ctx context.Context, in *FileDTO, out *Response) error
		Delete(ctx context.Context, in *DeleteId, out *Response) error
		Update(ctx context.Context, in *FileDTO, out *Response) error
		HistoricalData(ctx context.Context, in *DataType, out *Response) error
	}
	type FileResource struct {
		fileResource
	}
	h := &fileResourceHandler{hdlr}
	return s.Handle(s.NewHandler(&FileResource{h}, opts...))
}

type fileResourceHandler struct {
	FileResourceHandler
}

func (h *fileResourceHandler) Add(ctx context.Context, in *FileDTO, out *Response) error {
	return h.FileResourceHandler.Add(ctx, in, out)
}

func (h *fileResourceHandler) Delete(ctx context.Context, in *DeleteId, out *Response) error {
	return h.FileResourceHandler.Delete(ctx, in, out)
}

func (h *fileResourceHandler) Update(ctx context.Context, in *FileDTO, out *Response) error {
	return h.FileResourceHandler.Update(ctx, in, out)
}

func (h *fileResourceHandler) HistoricalData(ctx context.Context, in *DataType, out *Response) error {
	return h.FileResourceHandler.HistoricalData(ctx, in, out)
}
