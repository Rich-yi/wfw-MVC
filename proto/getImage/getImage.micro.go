// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/getImage/getImage.proto

package go_micro_srv_getImage

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Client API for GetImage service

type GetImageService interface {
	MicroGetImg(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type getImageService struct {
	c    client.Client
	name string
}

func NewGetImageService(name string, c client.Client) GetImageService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.getImage"
	}
	return &getImageService{
		c:    c,
		name: name,
	}
}

func (c *getImageService) MicroGetImg(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "GetImage.MicroGetImg", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GetImage service

type GetImageHandler interface {
	MicroGetImg(context.Context, *Request, *Response) error
}

func RegisterGetImageHandler(s server.Server, hdlr GetImageHandler, opts ...server.HandlerOption) error {
	type getImage interface {
		MicroGetImg(ctx context.Context, in *Request, out *Response) error
	}
	type GetImage struct {
		getImage
	}
	h := &getImageHandler{hdlr}
	return s.Handle(s.NewHandler(&GetImage{h}, opts...))
}

type getImageHandler struct {
	GetImageHandler
}

func (h *getImageHandler) MicroGetImg(ctx context.Context, in *Request, out *Response) error {
	return h.GetImageHandler.MicroGetImg(ctx, in, out)
}