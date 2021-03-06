// Code generated by protoc-gen-go.
// source: github.com/appcelerator/amp/api/rpc/stats/stats.proto
// DO NOT EDIT!

/*
Package stats is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/api/rpc/stats/stats.proto

It has these top-level messages:
	MetricsEntry
	MetricsCPUEntry
	MetricsIOEntry
	MetricsMemEntry
	MetricsNetEntry
	StatsRequest
	StatsReply
*/
package stats

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

type MetricsEntry struct {
	Timestamp          string            `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	ContainerId        string            `protobuf:"bytes,2,opt,name=container_id,json=containerId" json:"container_id,omitempty"`
	ContainerName      string            `protobuf:"bytes,3,opt,name=container_name,json=containerName" json:"container_name,omitempty"`
	ContainerShortName string            `protobuf:"bytes,4,opt,name=container_short_name,json=containerShortName" json:"container_short_name,omitempty"`
	ContainerState     string            `protobuf:"bytes,5,opt,name=container_state,json=containerState" json:"container_state,omitempty"`
	ServiceName        string            `protobuf:"bytes,6,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	ServiceId          string            `protobuf:"bytes,7,opt,name=service_id,json=serviceId" json:"service_id,omitempty"`
	TaskId             string            `protobuf:"bytes,8,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	StackName          string            `protobuf:"bytes,9,opt,name=stack_name,json=stackName" json:"stack_name,omitempty"`
	NodeId             string            `protobuf:"bytes,10,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	TimeId             string            `protobuf:"bytes,11,opt,name=time_id,json=timeId" json:"time_id,omitempty"`
	Labels             map[string]string `protobuf:"bytes,12,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Group              string            `protobuf:"bytes,13,opt,name=group" json:"group,omitempty"`
	Sgroup             string            `protobuf:"bytes,14,opt,name=sgroup" json:"sgroup,omitempty"`
	Cpu                *MetricsCPUEntry  `protobuf:"bytes,15,opt,name=cpu" json:"cpu,omitempty"`
	Io                 *MetricsIOEntry   `protobuf:"bytes,16,opt,name=io" json:"io,omitempty"`
	Mem                *MetricsMemEntry  `protobuf:"bytes,17,opt,name=mem" json:"mem,omitempty"`
	Net                *MetricsNetEntry  `protobuf:"bytes,18,opt,name=net" json:"net,omitempty"`
}

func (m *MetricsEntry) Reset()                    { *m = MetricsEntry{} }
func (m *MetricsEntry) String() string            { return proto.CompactTextString(m) }
func (*MetricsEntry) ProtoMessage()               {}
func (*MetricsEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MetricsEntry) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *MetricsEntry) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *MetricsEntry) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

func (m *MetricsEntry) GetContainerShortName() string {
	if m != nil {
		return m.ContainerShortName
	}
	return ""
}

func (m *MetricsEntry) GetContainerState() string {
	if m != nil {
		return m.ContainerState
	}
	return ""
}

func (m *MetricsEntry) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *MetricsEntry) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *MetricsEntry) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *MetricsEntry) GetStackName() string {
	if m != nil {
		return m.StackName
	}
	return ""
}

func (m *MetricsEntry) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *MetricsEntry) GetTimeId() string {
	if m != nil {
		return m.TimeId
	}
	return ""
}

func (m *MetricsEntry) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *MetricsEntry) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *MetricsEntry) GetSgroup() string {
	if m != nil {
		return m.Sgroup
	}
	return ""
}

func (m *MetricsEntry) GetCpu() *MetricsCPUEntry {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *MetricsEntry) GetIo() *MetricsIOEntry {
	if m != nil {
		return m.Io
	}
	return nil
}

func (m *MetricsEntry) GetMem() *MetricsMemEntry {
	if m != nil {
		return m.Mem
	}
	return nil
}

func (m *MetricsEntry) GetNet() *MetricsNetEntry {
	if m != nil {
		return m.Net
	}
	return nil
}

type MetricsCPUEntry struct {
	TotalUsage        float64 `protobuf:"fixed64,1,opt,name=total_usage,json=totalUsage" json:"total_usage,omitempty"`
	UsageInKernelMode float64 `protobuf:"fixed64,2,opt,name=usage_in_kernel_mode,json=usageInKernelMode" json:"usage_in_kernel_mode,omitempty"`
	UsageInUserMode   float64 `protobuf:"fixed64,3,opt,name=usage_in_user_mode,json=usageInUserMode" json:"usage_in_user_mode,omitempty"`
}

func (m *MetricsCPUEntry) Reset()                    { *m = MetricsCPUEntry{} }
func (m *MetricsCPUEntry) String() string            { return proto.CompactTextString(m) }
func (*MetricsCPUEntry) ProtoMessage()               {}
func (*MetricsCPUEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MetricsCPUEntry) GetTotalUsage() float64 {
	if m != nil {
		return m.TotalUsage
	}
	return 0
}

func (m *MetricsCPUEntry) GetUsageInKernelMode() float64 {
	if m != nil {
		return m.UsageInKernelMode
	}
	return 0
}

func (m *MetricsCPUEntry) GetUsageInUserMode() float64 {
	if m != nil {
		return m.UsageInUserMode
	}
	return 0
}

type MetricsIOEntry struct {
	Read  int64 `protobuf:"varint,1,opt,name=read" json:"read,omitempty"`
	Write int64 `protobuf:"varint,2,opt,name=write" json:"write,omitempty"`
	Total int64 `protobuf:"varint,3,opt,name=total" json:"total,omitempty"`
}

func (m *MetricsIOEntry) Reset()                    { *m = MetricsIOEntry{} }
func (m *MetricsIOEntry) String() string            { return proto.CompactTextString(m) }
func (*MetricsIOEntry) ProtoMessage()               {}
func (*MetricsIOEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MetricsIOEntry) GetRead() int64 {
	if m != nil {
		return m.Read
	}
	return 0
}

func (m *MetricsIOEntry) GetWrite() int64 {
	if m != nil {
		return m.Write
	}
	return 0
}

func (m *MetricsIOEntry) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type MetricsMemEntry struct {
	Failcnt  int64   `protobuf:"varint,1,opt,name=failcnt" json:"failcnt,omitempty"`
	Limit    int64   `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	Maxusage int64   `protobuf:"varint,3,opt,name=maxusage" json:"maxusage,omitempty"`
	Usage    int64   `protobuf:"varint,4,opt,name=usage" json:"usage,omitempty"`
	UsageP   float64 `protobuf:"fixed64,5,opt,name=usage_p,json=usageP" json:"usage_p,omitempty"`
}

func (m *MetricsMemEntry) Reset()                    { *m = MetricsMemEntry{} }
func (m *MetricsMemEntry) String() string            { return proto.CompactTextString(m) }
func (*MetricsMemEntry) ProtoMessage()               {}
func (*MetricsMemEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MetricsMemEntry) GetFailcnt() int64 {
	if m != nil {
		return m.Failcnt
	}
	return 0
}

func (m *MetricsMemEntry) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *MetricsMemEntry) GetMaxusage() int64 {
	if m != nil {
		return m.Maxusage
	}
	return 0
}

func (m *MetricsMemEntry) GetUsage() int64 {
	if m != nil {
		return m.Usage
	}
	return 0
}

func (m *MetricsMemEntry) GetUsageP() float64 {
	if m != nil {
		return m.UsageP
	}
	return 0
}

type MetricsNetEntry struct {
	TotalBytes int64 `protobuf:"varint,1,opt,name=total_bytes,json=totalBytes" json:"total_bytes,omitempty"`
	RxBytes    int64 `protobuf:"varint,2,opt,name=rx_bytes,json=rxBytes" json:"rx_bytes,omitempty"`
	RxDropped  int64 `protobuf:"varint,3,opt,name=rx_dropped,json=rxDropped" json:"rx_dropped,omitempty"`
	RxErrors   int64 `protobuf:"varint,4,opt,name=rx_errors,json=rxErrors" json:"rx_errors,omitempty"`
	RxPackets  int64 `protobuf:"varint,5,opt,name=rx_packets,json=rxPackets" json:"rx_packets,omitempty"`
	TxBytes    int64 `protobuf:"varint,6,opt,name=tx_bytes,json=txBytes" json:"tx_bytes,omitempty"`
	TxDropped  int64 `protobuf:"varint,7,opt,name=tx_dropped,json=txDropped" json:"tx_dropped,omitempty"`
	TxErrors   int64 `protobuf:"varint,8,opt,name=tx_errors,json=txErrors" json:"tx_errors,omitempty"`
	TxPackets  int64 `protobuf:"varint,9,opt,name=tx_packets,json=txPackets" json:"tx_packets,omitempty"`
}

func (m *MetricsNetEntry) Reset()                    { *m = MetricsNetEntry{} }
func (m *MetricsNetEntry) String() string            { return proto.CompactTextString(m) }
func (*MetricsNetEntry) ProtoMessage()               {}
func (*MetricsNetEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MetricsNetEntry) GetTotalBytes() int64 {
	if m != nil {
		return m.TotalBytes
	}
	return 0
}

func (m *MetricsNetEntry) GetRxBytes() int64 {
	if m != nil {
		return m.RxBytes
	}
	return 0
}

func (m *MetricsNetEntry) GetRxDropped() int64 {
	if m != nil {
		return m.RxDropped
	}
	return 0
}

func (m *MetricsNetEntry) GetRxErrors() int64 {
	if m != nil {
		return m.RxErrors
	}
	return 0
}

func (m *MetricsNetEntry) GetRxPackets() int64 {
	if m != nil {
		return m.RxPackets
	}
	return 0
}

func (m *MetricsNetEntry) GetTxBytes() int64 {
	if m != nil {
		return m.TxBytes
	}
	return 0
}

func (m *MetricsNetEntry) GetTxDropped() int64 {
	if m != nil {
		return m.TxDropped
	}
	return 0
}

func (m *MetricsNetEntry) GetTxErrors() int64 {
	if m != nil {
		return m.TxErrors
	}
	return 0
}

func (m *MetricsNetEntry) GetTxPackets() int64 {
	if m != nil {
		return m.TxPackets
	}
	return 0
}

type StatsRequest struct {
	StatsCpu                 bool   `protobuf:"varint,1,opt,name=stats_cpu,json=statsCpu" json:"stats_cpu,omitempty"`
	StatsMem                 bool   `protobuf:"varint,2,opt,name=stats_mem,json=statsMem" json:"stats_mem,omitempty"`
	StatsIo                  bool   `protobuf:"varint,3,opt,name=stats_io,json=statsIo" json:"stats_io,omitempty"`
	StatsNet                 bool   `protobuf:"varint,4,opt,name=stats_net,json=statsNet" json:"stats_net,omitempty"`
	Group                    string `protobuf:"bytes,5,opt,name=group" json:"group,omitempty"`
	FilterContainerId        string `protobuf:"bytes,6,opt,name=filter_container_id,json=filterContainerId" json:"filter_container_id,omitempty"`
	FilterContainerName      string `protobuf:"bytes,7,opt,name=filter_container_name,json=filterContainerName" json:"filter_container_name,omitempty"`
	FilterContainerShortName string `protobuf:"bytes,8,opt,name=filter_container_short_name,json=filterContainerShortName" json:"filter_container_short_name,omitempty"`
	FilterContainerState     string `protobuf:"bytes,9,opt,name=filter_container_state,json=filterContainerState" json:"filter_container_state,omitempty"`
	FilterServiceName        string `protobuf:"bytes,10,opt,name=filter_service_name,json=filterServiceName" json:"filter_service_name,omitempty"`
	FilterServiceId          string `protobuf:"bytes,11,opt,name=filter_service_id,json=filterServiceId" json:"filter_service_id,omitempty"`
	FilterTaskId             string `protobuf:"bytes,12,opt,name=filter_task_id,json=filterTaskId" json:"filter_task_id,omitempty"`
	FilterStackName          string `protobuf:"bytes,13,opt,name=filter_stack_name,json=filterStackName" json:"filter_stack_name,omitempty"`
	FilterNodeId             string `protobuf:"bytes,14,opt,name=filter_node_id,json=filterNodeId" json:"filter_node_id,omitempty"`
	Since                    string `protobuf:"bytes,15,opt,name=since" json:"since,omitempty"`
	Until                    string `protobuf:"bytes,16,opt,name=until" json:"until,omitempty"`
	Period                   string `protobuf:"bytes,17,opt,name=period" json:"period,omitempty"`
	TimeGroup                string `protobuf:"bytes,18,opt,name=time_group,json=timeGroup" json:"time_group,omitempty"`
	TimeZone                 string `protobuf:"bytes,19,opt,name=time_zone,json=timeZone" json:"time_zone,omitempty"`
	Avg                      bool   `protobuf:"varint,20,opt,name=avg" json:"avg,omitempty"`
	AllowsInfra              bool   `protobuf:"varint,21,opt,name=allows_infra,json=allowsInfra" json:"allows_infra,omitempty"`
}

func (m *StatsRequest) Reset()                    { *m = StatsRequest{} }
func (m *StatsRequest) String() string            { return proto.CompactTextString(m) }
func (*StatsRequest) ProtoMessage()               {}
func (*StatsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *StatsRequest) GetStatsCpu() bool {
	if m != nil {
		return m.StatsCpu
	}
	return false
}

func (m *StatsRequest) GetStatsMem() bool {
	if m != nil {
		return m.StatsMem
	}
	return false
}

func (m *StatsRequest) GetStatsIo() bool {
	if m != nil {
		return m.StatsIo
	}
	return false
}

func (m *StatsRequest) GetStatsNet() bool {
	if m != nil {
		return m.StatsNet
	}
	return false
}

func (m *StatsRequest) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *StatsRequest) GetFilterContainerId() string {
	if m != nil {
		return m.FilterContainerId
	}
	return ""
}

func (m *StatsRequest) GetFilterContainerName() string {
	if m != nil {
		return m.FilterContainerName
	}
	return ""
}

func (m *StatsRequest) GetFilterContainerShortName() string {
	if m != nil {
		return m.FilterContainerShortName
	}
	return ""
}

func (m *StatsRequest) GetFilterContainerState() string {
	if m != nil {
		return m.FilterContainerState
	}
	return ""
}

func (m *StatsRequest) GetFilterServiceName() string {
	if m != nil {
		return m.FilterServiceName
	}
	return ""
}

func (m *StatsRequest) GetFilterServiceId() string {
	if m != nil {
		return m.FilterServiceId
	}
	return ""
}

func (m *StatsRequest) GetFilterTaskId() string {
	if m != nil {
		return m.FilterTaskId
	}
	return ""
}

func (m *StatsRequest) GetFilterStackName() string {
	if m != nil {
		return m.FilterStackName
	}
	return ""
}

func (m *StatsRequest) GetFilterNodeId() string {
	if m != nil {
		return m.FilterNodeId
	}
	return ""
}

func (m *StatsRequest) GetSince() string {
	if m != nil {
		return m.Since
	}
	return ""
}

func (m *StatsRequest) GetUntil() string {
	if m != nil {
		return m.Until
	}
	return ""
}

func (m *StatsRequest) GetPeriod() string {
	if m != nil {
		return m.Period
	}
	return ""
}

func (m *StatsRequest) GetTimeGroup() string {
	if m != nil {
		return m.TimeGroup
	}
	return ""
}

func (m *StatsRequest) GetTimeZone() string {
	if m != nil {
		return m.TimeZone
	}
	return ""
}

func (m *StatsRequest) GetAvg() bool {
	if m != nil {
		return m.Avg
	}
	return false
}

func (m *StatsRequest) GetAllowsInfra() bool {
	if m != nil {
		return m.AllowsInfra
	}
	return false
}

type StatsReply struct {
	Entries []*MetricsEntry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
}

func (m *StatsReply) Reset()                    { *m = StatsReply{} }
func (m *StatsReply) String() string            { return proto.CompactTextString(m) }
func (*StatsReply) ProtoMessage()               {}
func (*StatsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *StatsReply) GetEntries() []*MetricsEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func init() {
	proto.RegisterType((*MetricsEntry)(nil), "stats.MetricsEntry")
	proto.RegisterType((*MetricsCPUEntry)(nil), "stats.MetricsCPUEntry")
	proto.RegisterType((*MetricsIOEntry)(nil), "stats.MetricsIOEntry")
	proto.RegisterType((*MetricsMemEntry)(nil), "stats.MetricsMemEntry")
	proto.RegisterType((*MetricsNetEntry)(nil), "stats.MetricsNetEntry")
	proto.RegisterType((*StatsRequest)(nil), "stats.StatsRequest")
	proto.RegisterType((*StatsReply)(nil), "stats.StatsReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Stats service

type StatsClient interface {
	StatsQuery(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsReply, error)
}

type statsClient struct {
	cc *grpc.ClientConn
}

func NewStatsClient(cc *grpc.ClientConn) StatsClient {
	return &statsClient{cc}
}

func (c *statsClient) StatsQuery(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsReply, error) {
	out := new(StatsReply)
	err := grpc.Invoke(ctx, "/stats.Stats/StatsQuery", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Stats service

type StatsServer interface {
	StatsQuery(context.Context, *StatsRequest) (*StatsReply, error)
}

func RegisterStatsServer(s *grpc.Server, srv StatsServer) {
	s.RegisterService(&_Stats_serviceDesc, srv)
}

func _Stats_StatsQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).StatsQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.Stats/StatsQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).StatsQuery(ctx, req.(*StatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Stats_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stats.Stats",
	HandlerType: (*StatsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StatsQuery",
			Handler:    _Stats_StatsQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/appcelerator/amp/api/rpc/stats/stats.proto",
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/api/rpc/stats/stats.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 1079 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x56, 0xcd, 0x6e, 0x23, 0x45,
	0x10, 0x96, 0xed, 0xf8, 0xaf, 0xec, 0x75, 0x36, 0x1d, 0x27, 0x19, 0x92, 0x5d, 0x6d, 0xb0, 0x58,
	0x11, 0x2d, 0x22, 0x86, 0x00, 0x02, 0x16, 0x71, 0x21, 0xac, 0x90, 0x81, 0x84, 0x30, 0x21, 0x17,
	0x2e, 0xa3, 0xce, 0x4c, 0xc7, 0xdb, 0xca, 0xcc, 0xf4, 0xd0, 0xd3, 0xce, 0xda, 0x1c, 0x39, 0x73,
	0x40, 0xe2, 0xce, 0xd3, 0xf0, 0x06, 0xbc, 0x02, 0x0f, 0xc1, 0x11, 0x75, 0x55, 0x8f, 0x67, 0xec,
	0xe4, 0x12, 0x4d, 0x55, 0x7d, 0xdf, 0x57, 0xdd, 0xed, 0xee, 0xaf, 0x02, 0x9f, 0x4c, 0xa5, 0x79,
	0x3d, 0xbb, 0x3e, 0x0e, 0x55, 0x32, 0xe6, 0x59, 0x16, 0x8a, 0x58, 0x68, 0x6e, 0x94, 0x1e, 0xf3,
	0x24, 0x1b, 0xf3, 0x4c, 0x8e, 0x75, 0x16, 0x8e, 0x73, 0xc3, 0x4d, 0x4e, 0x7f, 0x8f, 0x33, 0xad,
	0x8c, 0x62, 0x4d, 0x0c, 0xf6, 0x9f, 0x4c, 0x95, 0x9a, 0xc6, 0x02, 0x81, 0x3c, 0x4d, 0x95, 0xe1,
	0x46, 0xaa, 0xd4, 0x81, 0x46, 0x7f, 0x37, 0xa1, 0x7f, 0x26, 0x8c, 0x96, 0x61, 0xfe, 0x2a, 0x35,
	0x7a, 0xc1, 0x9e, 0x40, 0xd7, 0xc8, 0x44, 0xe4, 0x86, 0x27, 0x99, 0x57, 0x3b, 0xac, 0x1d, 0x75,
	0xfd, 0x32, 0xc1, 0xde, 0x86, 0x7e, 0xa8, 0x52, 0xc3, 0x65, 0x2a, 0x74, 0x20, 0x23, 0xaf, 0x8e,
	0x80, 0xde, 0x32, 0x37, 0x89, 0xd8, 0x73, 0x18, 0x94, 0x90, 0x94, 0x27, 0xc2, 0x6b, 0x20, 0xe8,
	0xd1, 0x32, 0x7b, 0xce, 0x13, 0xc1, 0x3e, 0x80, 0x61, 0x09, 0xcb, 0x5f, 0x2b, 0x6d, 0x08, 0xbc,
	0x81, 0x60, 0xb6, 0xac, 0x5d, 0xda, 0x12, 0x32, 0xde, 0x85, 0xcd, 0x0a, 0xc3, 0x70, 0x23, 0xbc,
	0x26, 0x82, 0xcb, 0x7e, 0x97, 0x36, 0x6b, 0x17, 0x99, 0x0b, 0x7d, 0x27, 0x43, 0x41, 0x92, 0x2d,
	0x5a, 0xa4, 0xcb, 0xa1, 0xd6, 0x53, 0x80, 0x02, 0x22, 0x23, 0xaf, 0x4d, 0xdb, 0x74, 0x99, 0x49,
	0xc4, 0xf6, 0xa0, 0x6d, 0x78, 0x7e, 0x6b, 0x6b, 0x1d, 0xac, 0xb5, 0x6c, 0x38, 0x89, 0x90, 0x67,
	0x78, 0x78, 0x4b, 0xc2, 0x5d, 0xc7, 0xb3, 0x19, 0x94, 0xdd, 0x83, 0x76, 0xaa, 0x22, 0xd4, 0x04,
	0xe2, 0xd9, 0xd0, 0x09, 0xca, 0x04, 0x0b, 0x3d, 0x27, 0x28, 0x13, 0x5b, 0xf8, 0x14, 0x5a, 0x31,
	0xbf, 0x16, 0x71, 0xee, 0xf5, 0x0f, 0x1b, 0x47, 0xbd, 0x93, 0x67, 0xc7, 0xf4, 0x13, 0x56, 0x7f,
	0x93, 0xe3, 0xef, 0x11, 0x81, 0xdf, 0xbe, 0x83, 0xb3, 0x21, 0x34, 0xa7, 0x5a, 0xcd, 0x32, 0xef,
	0x11, 0xea, 0x51, 0xc0, 0x76, 0xa1, 0x95, 0x53, 0x7a, 0x40, 0x6d, 0x28, 0x62, 0x47, 0xd0, 0x08,
	0xb3, 0x99, 0xb7, 0x79, 0x58, 0x3b, 0xea, 0x9d, 0xec, 0xae, 0xf6, 0x38, 0xbd, 0xb8, 0x22, 0x69,
	0x0b, 0x61, 0xcf, 0xa1, 0x2e, 0x95, 0xf7, 0x18, 0x81, 0x3b, 0xab, 0xc0, 0xc9, 0x0f, 0x84, 0xab,
	0x4b, 0x65, 0x05, 0x13, 0x91, 0x78, 0x5b, 0x0f, 0x09, 0x9e, 0x89, 0xc4, 0x09, 0x26, 0x22, 0xb1,
	0xc8, 0x54, 0x18, 0x8f, 0x3d, 0x84, 0x3c, 0x17, 0xc6, 0x21, 0x53, 0x61, 0xf6, 0x3f, 0x87, 0x5e,
	0x65, 0xa7, 0xec, 0x31, 0x34, 0x6e, 0xc5, 0xc2, 0xdd, 0x41, 0xfb, 0x69, 0xf7, 0x7c, 0xc7, 0xe3,
	0x99, 0x70, 0xd7, 0x8e, 0x82, 0x97, 0xf5, 0xcf, 0x6a, 0xa3, 0x3f, 0x6a, 0xb0, 0xb9, 0xb6, 0x1d,
	0xf6, 0x0c, 0x7a, 0x46, 0x19, 0x1e, 0x07, 0xb3, 0x9c, 0x4f, 0x05, 0xea, 0xd4, 0x7c, 0xc0, 0xd4,
	0x95, 0xcd, 0xb0, 0x31, 0x0c, 0xb1, 0x14, 0xc8, 0x34, 0xb8, 0x15, 0x3a, 0x15, 0x71, 0x90, 0xa8,
	0x88, 0xd4, 0x6b, 0xfe, 0x16, 0xd6, 0x26, 0xe9, 0x77, 0x58, 0x39, 0x53, 0x91, 0x60, 0xef, 0x01,
	0x5b, 0x12, 0x66, 0xb9, 0xd0, 0x04, 0x6f, 0x20, 0x7c, 0xd3, 0xc1, 0xaf, 0x72, 0xa1, 0x2d, 0x78,
	0x74, 0x01, 0x83, 0xd5, 0x73, 0x63, 0x0c, 0x36, 0xb4, 0xe0, 0x11, 0xae, 0xa4, 0xe1, 0xe3, 0xb7,
	0xdd, 0xd2, 0x1b, 0x2d, 0x0d, 0x35, 0x6d, 0xf8, 0x14, 0xd8, 0x2c, 0xae, 0x13, 0xb5, 0x1b, 0x3e,
	0x05, 0xa3, 0xdf, 0xcb, 0x4d, 0x16, 0x47, 0xcc, 0x3c, 0x68, 0xdf, 0x70, 0x19, 0x87, 0xa9, 0x71,
	0xb2, 0x45, 0x68, 0x35, 0x62, 0x99, 0x48, 0x53, 0x28, 0x63, 0xc0, 0xf6, 0xa1, 0x93, 0xf0, 0x39,
	0x9d, 0x08, 0x89, 0x2f, 0x63, 0xcb, 0xa0, 0xc2, 0x06, 0x31, 0x28, 0xbb, 0x07, 0x6d, 0xda, 0x74,
	0x86, 0xcf, 0xad, 0xe6, 0xb7, 0x30, 0xbc, 0x18, 0xfd, 0x55, 0x5f, 0x2e, 0xa7, 0xf8, 0x1d, 0xcb,
	0x33, 0xbf, 0x5e, 0x18, 0x91, 0xbb, 0x25, 0xd1, 0x99, 0x7f, 0x65, 0x33, 0xec, 0x2d, 0xe8, 0xe8,
	0xb9, 0xab, 0xd2, 0xc2, 0xda, 0x7a, 0x4e, 0xa5, 0xa7, 0x00, 0x7a, 0x1e, 0x44, 0x5a, 0x65, 0x99,
	0x88, 0xdc, 0xe2, 0xba, 0x7a, 0xfe, 0x35, 0x25, 0xd8, 0x01, 0x74, 0xf5, 0x3c, 0x10, 0x5a, 0x2b,
	0x9d, 0xbb, 0x15, 0x76, 0xf4, 0xfc, 0x15, 0xc6, 0x8e, 0x9b, 0xf1, 0xf0, 0x56, 0x98, 0x1c, 0xd7,
	0x89, 0xdc, 0x0b, 0x4a, 0xd8, 0xae, 0xa6, 0xe8, 0xda, 0xa2, 0xae, 0xa6, 0xec, 0x6a, 0xca, 0xae,
	0x6d, 0x62, 0x9a, 0x6a, 0x57, 0xb3, 0xec, 0xda, 0xa1, 0xae, 0xa6, 0xd2, 0xd5, 0x94, 0x5d, 0xbb,
	0x05, 0xd7, 0x75, 0x1d, 0xfd, 0xd7, 0x84, 0xbe, 0x75, 0xa4, 0xdc, 0x17, 0xbf, 0xcc, 0x44, 0x6e,
	0xac, 0x18, 0x5e, 0xff, 0xc0, 0xbe, 0x45, 0x7b, 0x36, 0x1d, 0xbf, 0x83, 0x89, 0xd3, 0x6c, 0x56,
	0x16, 0xed, 0xbb, 0xaa, 0x57, 0x8a, 0x67, 0x22, 0xb1, 0x1b, 0xa0, 0xa2, 0x54, 0x78, 0x32, 0x1d,
	0xbf, 0x8d, 0xf1, 0x44, 0x95, 0x3c, 0xfb, 0xca, 0x36, 0x2a, 0xbc, 0x73, 0x61, 0x4a, 0x97, 0x68,
	0x56, 0x5d, 0xe2, 0x18, 0xb6, 0x6f, 0x64, 0x6c, 0x84, 0x0e, 0x56, 0xcc, 0x9c, 0x7c, 0x72, 0x8b,
	0x4a, 0xa7, 0x15, 0x4b, 0x3f, 0x81, 0x9d, 0x7b, 0x78, 0x34, 0x40, 0x32, 0xce, 0xed, 0x35, 0x06,
	0x5a, 0xe1, 0x97, 0x70, 0x70, 0x8f, 0x53, 0xb1, 0x79, 0xb2, 0x55, 0x6f, 0x8d, 0x59, 0x9a, 0xfd,
	0xc7, 0xb0, 0x7b, 0x9f, 0x8e, 0x9e, 0x4f, 0xa6, 0x3b, 0x5c, 0x67, 0xa2, 0xf3, 0x97, 0x1b, 0x5b,
	0x19, 0x00, 0x50, 0xdd, 0xd8, 0x65, 0x65, 0x0c, 0xbc, 0x80, 0xad, 0x35, 0xfc, 0xd2, 0xa0, 0x37,
	0x57, 0xd0, 0x93, 0x88, 0xbd, 0x03, 0x03, 0x87, 0x2d, 0x46, 0x43, 0x1f, 0x81, 0x7d, 0xca, 0xfe,
	0x44, 0x03, 0xa2, 0xa2, 0x58, 0xce, 0x89, 0x47, 0x2b, 0x8a, 0xcb, 0x69, 0x51, 0x2a, 0x16, 0x43,
	0x63, 0x50, 0x55, 0x3c, 0xa7, 0xd1, 0x31, 0x84, 0x66, 0x2e, 0xd3, 0x50, 0xa0, 0x79, 0x77, 0x7d,
	0x0a, 0xf0, 0xad, 0xa6, 0x46, 0xc6, 0xe8, 0xd4, 0x5d, 0x9f, 0x02, 0x6b, 0xff, 0x99, 0xd0, 0x52,
	0x45, 0x68, 0xcc, 0x5d, 0xdf, 0x45, 0x78, 0x51, 0xed, 0xf8, 0xa1, 0xbb, 0xc0, 0xca, 0xa9, 0xfe,
	0x0d, 0xde, 0x87, 0x03, 0x9a, 0xf9, 0xc1, 0xaf, 0x2a, 0x15, 0xde, 0x36, 0x56, 0x3b, 0x36, 0xf1,
	0xb3, 0x4a, 0x85, 0xb5, 0x61, 0x7e, 0x37, 0xf5, 0x86, 0x78, 0xb3, 0xec, 0xa7, 0x9d, 0xaf, 0x3c,
	0x8e, 0xd5, 0x9b, 0x3c, 0x90, 0xe9, 0x8d, 0xe6, 0xde, 0x0e, 0x96, 0x7a, 0x94, 0x9b, 0xd8, 0xd4,
	0xe8, 0x0b, 0x00, 0x77, 0xf3, 0xb3, 0x78, 0xc1, 0xde, 0x87, 0xb6, 0x48, 0x8d, 0x96, 0xe8, 0x08,
	0x76, 0xca, 0x6d, 0x3f, 0x30, 0xe5, 0xfc, 0x02, 0x73, 0x72, 0x09, 0x4d, 0x24, 0xb3, 0x6f, 0x9d,
	0xca, 0x8f, 0x33, 0xa1, 0x17, 0xac, 0x20, 0x55, 0x9f, 0xd4, 0xfe, 0xd6, 0x6a, 0x32, 0x8b, 0x17,
	0xa3, 0xe1, 0x6f, 0xff, 0xfc, 0xfb, 0x67, 0x7d, 0x30, 0xea, 0x8e, 0xef, 0x3e, 0xa4, 0x7f, 0x88,
	0x5e, 0xd6, 0x5e, 0x5c, 0xb7, 0xf0, 0xff, 0x9d, 0x8f, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x20,
	0x62, 0xe6, 0x00, 0x4d, 0x09, 0x00, 0x00,
}
