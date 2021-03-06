// Code generated by protoc-gen-go. DO NOT EDIT.
// source: optimus.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PredictSupplierRequest struct {
	Devices *DevicesReply `protobuf:"bytes,1,opt,name=devices" json:"devices,omitempty"`
}

func (m *PredictSupplierRequest) Reset()                    { *m = PredictSupplierRequest{} }
func (m *PredictSupplierRequest) String() string            { return proto.CompactTextString(m) }
func (*PredictSupplierRequest) ProtoMessage()               {}
func (*PredictSupplierRequest) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *PredictSupplierRequest) GetDevices() *DevicesReply {
	if m != nil {
		return m.Devices
	}
	return nil
}

type PredictSupplierReply struct {
	Price    *Price    `protobuf:"bytes,1,opt,name=price" json:"price,omitempty"`
	OrderIDs []*BigInt `protobuf:"bytes,2,rep,name=orderIDs" json:"orderIDs,omitempty"`
}

func (m *PredictSupplierReply) Reset()                    { *m = PredictSupplierReply{} }
func (m *PredictSupplierReply) String() string            { return proto.CompactTextString(m) }
func (*PredictSupplierReply) ProtoMessage()               {}
func (*PredictSupplierReply) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

func (m *PredictSupplierReply) GetPrice() *Price {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *PredictSupplierReply) GetOrderIDs() []*BigInt {
	if m != nil {
		return m.OrderIDs
	}
	return nil
}

func init() {
	proto.RegisterType((*PredictSupplierRequest)(nil), "sonm.PredictSupplierRequest")
	proto.RegisterType((*PredictSupplierReply)(nil), "sonm.PredictSupplierReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for OrderPredictor service

type OrderPredictorClient interface {
	Predict(ctx context.Context, in *BidResources, opts ...grpc.CallOption) (*Price, error)
	PredictSupplier(ctx context.Context, in *PredictSupplierRequest, opts ...grpc.CallOption) (*PredictSupplierReply, error)
}

type orderPredictorClient struct {
	cc *grpc.ClientConn
}

func NewOrderPredictorClient(cc *grpc.ClientConn) OrderPredictorClient {
	return &orderPredictorClient{cc}
}

func (c *orderPredictorClient) Predict(ctx context.Context, in *BidResources, opts ...grpc.CallOption) (*Price, error) {
	out := new(Price)
	err := grpc.Invoke(ctx, "/sonm.OrderPredictor/Predict", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderPredictorClient) PredictSupplier(ctx context.Context, in *PredictSupplierRequest, opts ...grpc.CallOption) (*PredictSupplierReply, error) {
	out := new(PredictSupplierReply)
	err := grpc.Invoke(ctx, "/sonm.OrderPredictor/PredictSupplier", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderPredictor service

type OrderPredictorServer interface {
	Predict(context.Context, *BidResources) (*Price, error)
	PredictSupplier(context.Context, *PredictSupplierRequest) (*PredictSupplierReply, error)
}

func RegisterOrderPredictorServer(s *grpc.Server, srv OrderPredictorServer) {
	s.RegisterService(&_OrderPredictor_serviceDesc, srv)
}

func _OrderPredictor_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidResources)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderPredictorServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.OrderPredictor/Predict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderPredictorServer).Predict(ctx, req.(*BidResources))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderPredictor_PredictSupplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictSupplierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderPredictorServer).PredictSupplier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.OrderPredictor/PredictSupplier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderPredictorServer).PredictSupplier(ctx, req.(*PredictSupplierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderPredictor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonm.OrderPredictor",
	HandlerType: (*OrderPredictorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Predict",
			Handler:    _OrderPredictor_Predict_Handler,
		},
		{
			MethodName: "PredictSupplier",
			Handler:    _OrderPredictor_PredictSupplier_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "optimus.proto",
}

func init() { proto.RegisterFile("optimus.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x5b, 0xff, 0xb3, 0xad, 0x16, 0x17, 0x91, 0x12, 0x3c, 0xd4, 0xe0, 0xa1, 0x07, 0x4d,
	0x20, 0x1e, 0xbd, 0x95, 0x22, 0xf4, 0x20, 0x96, 0xf5, 0xe6, 0x2d, 0xd9, 0x0c, 0x71, 0x68, 0x92,
	0x5d, 0x67, 0x37, 0x4a, 0x3e, 0x84, 0xdf, 0x59, 0xb6, 0x9b, 0x88, 0xf8, 0xe7, 0xb6, 0xef, 0xf7,
	0xde, 0xbe, 0x19, 0x86, 0x1d, 0x2b, 0x6d, 0xb1, 0x6a, 0x4c, 0xa4, 0x49, 0x59, 0xc5, 0xf7, 0x8c,
	0xaa, 0xab, 0x60, 0x9c, 0x61, 0x81, 0xb5, 0xf5, 0x2c, 0x98, 0x60, 0xed, 0x68, 0x8d, 0x69, 0x07,
	0x4e, 0xab, 0x94, 0x36, 0x60, 0x75, 0x99, 0x4a, 0xe8, 0xd0, 0xf8, 0x5d, 0xd1, 0x06, 0xc8, 0xab,
	0xf0, 0x9e, 0x9d, 0xaf, 0x09, 0x72, 0x94, 0xf6, 0xa9, 0xd1, 0xba, 0x44, 0x20, 0x01, 0xaf, 0x0d,
	0x18, 0xcb, 0xaf, 0xd9, 0x61, 0x0e, 0x6f, 0x28, 0xc1, 0x4c, 0x87, 0xb3, 0xe1, 0x7c, 0x94, 0xf0,
	0xc8, 0x75, 0x47, 0x4b, 0x0f, 0x05, 0xe8, 0xb2, 0x15, 0x7d, 0x24, 0x94, 0xec, 0xec, 0x57, 0x8f,
	0x2e, 0x5b, 0x7e, 0xc9, 0xf6, 0x35, 0xa1, 0x84, 0xae, 0x63, 0xe4, 0x3b, 0xd6, 0x0e, 0x09, 0xef,
	0xf0, 0x39, 0x3b, 0x52, 0x94, 0x03, 0xad, 0x96, 0x66, 0xba, 0x33, 0xdb, 0x9d, 0x8f, 0x92, 0xb1,
	0x4f, 0x2d, 0xb0, 0x58, 0xd5, 0x56, 0x7c, 0xb9, 0xc9, 0xc7, 0x90, 0x9d, 0x3c, 0x3a, 0xd1, 0x8d,
	0x52, 0xe4, 0xb6, 0xec, 0x04, 0xe7, 0xfd, 0xaf, 0x5c, 0x80, 0x51, 0x0d, 0x49, 0x30, 0xc1, 0xf7,
	0x79, 0xe1, 0x80, 0x3f, 0xb0, 0xc9, 0x8f, 0x2d, 0xf9, 0x45, 0x9f, 0xf8, 0xeb, 0x08, 0x41, 0xf0,
	0x8f, 0xab, 0xcb, 0x36, 0x1c, 0x2c, 0xae, 0x9e, 0xc3, 0x02, 0xed, 0x4b, 0x93, 0x45, 0x52, 0x55,
	0xb1, 0x4b, 0xde, 0xa0, 0x8a, 0xa5, 0x22, 0x88, 0xb7, 0xd7, 0xbd, 0x73, 0x28, 0x3b, 0xd8, 0xbe,
	0x6f, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x85, 0x6a, 0xa8, 0xc0, 0x01, 0x00, 0x00,
}
