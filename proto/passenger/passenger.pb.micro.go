// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/passenger/passenger.proto

package passenger

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Api Endpoints for PassengerSrv service

func NewPassengerSrvEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PassengerSrv service

type PassengerSrvService interface {
	ReportGeo(ctx context.Context, in *ReportPassengerRequest, opts ...client.CallOption) (*ReportPassengerResponse, error)
	QueryGeo(ctx context.Context, in *QueryPassengerRequest, opts ...client.CallOption) (*QueryPassengerResponse, error)
}

type passengerSrvService struct {
	c    client.Client
	name string
}

func NewPassengerSrvService(name string, c client.Client) PassengerSrvService {
	return &passengerSrvService{
		c:    c,
		name: name,
	}
}

func (c *passengerSrvService) ReportGeo(ctx context.Context, in *ReportPassengerRequest, opts ...client.CallOption) (*ReportPassengerResponse, error) {
	req := c.c.NewRequest(c.name, "PassengerSrv.ReportGeo", in)
	out := new(ReportPassengerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passengerSrvService) QueryGeo(ctx context.Context, in *QueryPassengerRequest, opts ...client.CallOption) (*QueryPassengerResponse, error) {
	req := c.c.NewRequest(c.name, "PassengerSrv.QueryGeo", in)
	out := new(QueryPassengerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PassengerSrv service

type PassengerSrvHandler interface {
	ReportGeo(context.Context, *ReportPassengerRequest, *ReportPassengerResponse) error
	QueryGeo(context.Context, *QueryPassengerRequest, *QueryPassengerResponse) error
}

func RegisterPassengerSrvHandler(s server.Server, hdlr PassengerSrvHandler, opts ...server.HandlerOption) error {
	type passengerSrv interface {
		ReportGeo(ctx context.Context, in *ReportPassengerRequest, out *ReportPassengerResponse) error
		QueryGeo(ctx context.Context, in *QueryPassengerRequest, out *QueryPassengerResponse) error
	}
	type PassengerSrv struct {
		passengerSrv
	}
	h := &passengerSrvHandler{hdlr}
	return s.Handle(s.NewHandler(&PassengerSrv{h}, opts...))
}

type passengerSrvHandler struct {
	PassengerSrvHandler
}

func (h *passengerSrvHandler) ReportGeo(ctx context.Context, in *ReportPassengerRequest, out *ReportPassengerResponse) error {
	return h.PassengerSrvHandler.ReportGeo(ctx, in, out)
}

func (h *passengerSrvHandler) QueryGeo(ctx context.Context, in *QueryPassengerRequest, out *QueryPassengerResponse) error {
	return h.PassengerSrvHandler.QueryGeo(ctx, in, out)
}
