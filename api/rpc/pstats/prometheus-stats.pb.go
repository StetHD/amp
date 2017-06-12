// Code generated by protoc-gen-go.
// source: github.com/appcelerator/amp/api/rpc/pstats/prometheus-stats.proto
// DO NOT EDIT!

/*
Package pstats is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/api/rpc/pstats/prometheus-stats.proto

It has these top-level messages:
	LoadSourcesRequest
	LoadSourcesReply
	PrometheusStatsRequest
	PrometheusStatsReply
*/
package pstats

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LoadSourcesRequest struct {
}

func (m *LoadSourcesRequest) Reset()                    { *m = LoadSourcesRequest{} }
func (m *LoadSourcesRequest) String() string            { return proto.CompactTextString(m) }
func (*LoadSourcesRequest) ProtoMessage()               {}
func (*LoadSourcesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type LoadSourcesReply struct {
}

func (m *LoadSourcesReply) Reset()                    { *m = LoadSourcesReply{} }
func (m *LoadSourcesReply) String() string            { return proto.CompactTextString(m) }
func (*LoadSourcesReply) ProtoMessage()               {}
func (*LoadSourcesReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type PrometheusStatsRequest struct {
}

func (m *PrometheusStatsRequest) Reset()                    { *m = PrometheusStatsRequest{} }
func (m *PrometheusStatsRequest) String() string            { return proto.CompactTextString(m) }
func (*PrometheusStatsRequest) ProtoMessage()               {}
func (*PrometheusStatsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type PrometheusStatsReply struct {
}

func (m *PrometheusStatsReply) Reset()                    { *m = PrometheusStatsReply{} }
func (m *PrometheusStatsReply) String() string            { return proto.CompactTextString(m) }
func (*PrometheusStatsReply) ProtoMessage()               {}
func (*PrometheusStatsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*LoadSourcesRequest)(nil), "pstats.LoadSourcesRequest")
	proto.RegisterType((*LoadSourcesReply)(nil), "pstats.LoadSourcesReply")
	proto.RegisterType((*PrometheusStatsRequest)(nil), "pstats.PrometheusStatsRequest")
	proto.RegisterType((*PrometheusStatsReply)(nil), "pstats.PrometheusStatsReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PrometheusStats service

type PrometheusStatsClient interface {
	ReloadSources(ctx context.Context, in *LoadSourcesRequest, opts ...grpc.CallOption) (*LoadSourcesReply, error)
	Stats(ctx context.Context, in *PrometheusStatsRequest, opts ...grpc.CallOption) (*PrometheusStatsReply, error)
}

type prometheusStatsClient struct {
	cc *grpc.ClientConn
}

func NewPrometheusStatsClient(cc *grpc.ClientConn) PrometheusStatsClient {
	return &prometheusStatsClient{cc}
}

func (c *prometheusStatsClient) ReloadSources(ctx context.Context, in *LoadSourcesRequest, opts ...grpc.CallOption) (*LoadSourcesReply, error) {
	out := new(LoadSourcesReply)
	err := grpc.Invoke(ctx, "/pstats.PrometheusStats/ReloadSources", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prometheusStatsClient) Stats(ctx context.Context, in *PrometheusStatsRequest, opts ...grpc.CallOption) (*PrometheusStatsReply, error) {
	out := new(PrometheusStatsReply)
	err := grpc.Invoke(ctx, "/pstats.PrometheusStats/Stats", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PrometheusStats service

type PrometheusStatsServer interface {
	ReloadSources(context.Context, *LoadSourcesRequest) (*LoadSourcesReply, error)
	Stats(context.Context, *PrometheusStatsRequest) (*PrometheusStatsReply, error)
}

func RegisterPrometheusStatsServer(s *grpc.Server, srv PrometheusStatsServer) {
	s.RegisterService(&_PrometheusStats_serviceDesc, srv)
}

func _PrometheusStats_ReloadSources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadSourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrometheusStatsServer).ReloadSources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pstats.PrometheusStats/ReloadSources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrometheusStatsServer).ReloadSources(ctx, req.(*LoadSourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrometheusStats_Stats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrometheusStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrometheusStatsServer).Stats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pstats.PrometheusStats/Stats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrometheusStatsServer).Stats(ctx, req.(*PrometheusStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PrometheusStats_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pstats.PrometheusStats",
	HandlerType: (*PrometheusStatsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReloadSources",
			Handler:    _PrometheusStats_ReloadSources_Handler,
		},
		{
			MethodName: "Stats",
			Handler:    _PrometheusStats_Stats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/appcelerator/amp/api/rpc/pstats/prometheus-stats.proto",
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/api/rpc/pstats/prometheus-stats.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4c, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x2c, 0x28, 0x48, 0x4e, 0xcd, 0x49, 0x2d, 0x4a,
	0x2c, 0xc9, 0x2f, 0xd2, 0x4f, 0xcc, 0x2d, 0xd0, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0x2a, 0x48, 0xd6,
	0x2f, 0x28, 0x2e, 0x49, 0x2c, 0x29, 0xd6, 0x2f, 0x28, 0xca, 0xcf, 0x4d, 0x2d, 0xc9, 0x48, 0x2d,
	0x2d, 0xd6, 0x05, 0x0b, 0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xb1, 0x41, 0xa4, 0xa5, 0x64,
	0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xc1, 0xba, 0x12, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32,
	0xf3, 0xf3, 0xa0, 0xaa, 0x94, 0x44, 0xb8, 0x84, 0x7c, 0xf2, 0x13, 0x53, 0x82, 0xf3, 0x4b, 0x8b,
	0x92, 0x53, 0x8b, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x94, 0x84, 0xb8, 0x04, 0x50, 0x44,
	0x0b, 0x72, 0x2a, 0x95, 0x24, 0xb8, 0xc4, 0x02, 0xe0, 0x36, 0x05, 0x83, 0x8c, 0x86, 0xa9, 0x16,
	0xe3, 0x12, 0xc1, 0x90, 0x29, 0xc8, 0xa9, 0x34, 0xba, 0xcf, 0xc8, 0xc5, 0x8f, 0x26, 0x21, 0x94,
	0xc2, 0xc5, 0x1b, 0x94, 0x9a, 0x83, 0x30, 0x5b, 0x48, 0x4a, 0x0f, 0xe2, 0x4e, 0x3d, 0x4c, 0x67,
	0x48, 0x49, 0x60, 0x95, 0x03, 0x39, 0x46, 0xb6, 0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0xe2, 0x42, 0xa2,
	0xfa, 0x65, 0x86, 0xfa, 0x90, 0x60, 0x40, 0x36, 0x34, 0x8d, 0x8b, 0x15, 0x62, 0x9d, 0x1c, 0xcc,
	0x04, 0xec, 0x4e, 0x97, 0x92, 0xc1, 0x29, 0x0f, 0xb2, 0x45, 0x1e, 0x6c, 0x8b, 0xa4, 0x92, 0x08,
	0xc8, 0x16, 0xf4, 0x60, 0xb6, 0x62, 0xd4, 0x4a, 0x62, 0x03, 0x07, 0xa2, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x92, 0x44, 0x10, 0xb0, 0xaf, 0x01, 0x00, 0x00,
}