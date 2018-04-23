// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// grpccmd imports
import (
	"io"

	"github.com/spf13/cobra"
	"github.com/sshaman1101/grpccmd"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WorkerRelationshipStatus int32

const (
	WorkerRelationshipStatus_RELATION_UNAPPROVED WorkerRelationshipStatus = 0
	WorkerRelationshipStatus_RELATION_APPROVED   WorkerRelationshipStatus = 1
)

var WorkerRelationshipStatus_name = map[int32]string{
	0: "RELATION_UNAPPROVED",
	1: "RELATION_APPROVED",
}
var WorkerRelationshipStatus_value = map[string]int32{
	"RELATION_UNAPPROVED": 0,
	"RELATION_APPROVED":   1,
}

func (x WorkerRelationshipStatus) String() string {
	return proto.EnumName(WorkerRelationshipStatus_name, int32(x))
}
func (WorkerRelationshipStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

type JoinNetworkRequest struct {
	TaskID    *TaskID `protobuf:"bytes,1,opt,name=taskID" json:"taskID,omitempty"`
	NetworkID string  `protobuf:"bytes,2,opt,name=NetworkID" json:"NetworkID,omitempty"`
}

func (m *JoinNetworkRequest) Reset()                    { *m = JoinNetworkRequest{} }
func (m *JoinNetworkRequest) String() string            { return proto.CompactTextString(m) }
func (*JoinNetworkRequest) ProtoMessage()               {}
func (*JoinNetworkRequest) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *JoinNetworkRequest) GetTaskID() *TaskID {
	if m != nil {
		return m.TaskID
	}
	return nil
}

func (m *JoinNetworkRequest) GetNetworkID() string {
	if m != nil {
		return m.NetworkID
	}
	return ""
}

type DealsReply struct {
	Deal []*Deal `protobuf:"bytes,1,rep,name=deal" json:"deal,omitempty"`
}

func (m *DealsReply) Reset()                    { *m = DealsReply{} }
func (m *DealsReply) String() string            { return proto.CompactTextString(m) }
func (*DealsReply) ProtoMessage()               {}
func (*DealsReply) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

func (m *DealsReply) GetDeal() []*Deal {
	if m != nil {
		return m.Deal
	}
	return nil
}

type Worker struct {
	ID     string                   `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Status WorkerRelationshipStatus `protobuf:"varint,2,opt,name=status,enum=sonm.WorkerRelationshipStatus" json:"status,omitempty"`
}

func (m *Worker) Reset()                    { *m = Worker{} }
func (m *Worker) String() string            { return proto.CompactTextString(m) }
func (*Worker) ProtoMessage()               {}
func (*Worker) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

func (m *Worker) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Worker) GetStatus() WorkerRelationshipStatus {
	if m != nil {
		return m.Status
	}
	return WorkerRelationshipStatus_RELATION_UNAPPROVED
}

type WorkerListReply struct {
	Workers []*Worker `protobuf:"bytes,1,rep,name=workers" json:"workers,omitempty"`
}

func (m *WorkerListReply) Reset()                    { *m = WorkerListReply{} }
func (m *WorkerListReply) String() string            { return proto.CompactTextString(m) }
func (*WorkerListReply) ProtoMessage()               {}
func (*WorkerListReply) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{3} }

func (m *WorkerListReply) GetWorkers() []*Worker {
	if m != nil {
		return m.Workers
	}
	return nil
}

func init() {
	proto.RegisterType((*JoinNetworkRequest)(nil), "sonm.JoinNetworkRequest")
	proto.RegisterType((*DealsReply)(nil), "sonm.DealsReply")
	proto.RegisterType((*Worker)(nil), "sonm.Worker")
	proto.RegisterType((*WorkerListReply)(nil), "sonm.WorkerListReply")
	proto.RegisterEnum("sonm.WorkerRelationshipStatus", WorkerRelationshipStatus_name, WorkerRelationshipStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TaskManagement service

type TaskManagementClient interface {
	// List produces a list of all tasks running on different SONM nodes
	List(ctx context.Context, in *EthAddress, opts ...grpc.CallOption) (*TaskListReply, error)
	// PushTask pushes image to Worker
	PushTask(ctx context.Context, opts ...grpc.CallOption) (TaskManagement_PushTaskClient, error)
	// Start starts a task on given resource
	Start(ctx context.Context, in *StartTaskRequest, opts ...grpc.CallOption) (*StartTaskReply, error)
	// JoinNetwork provides network specs to join specified task
	JoinNetwork(ctx context.Context, in *JoinNetworkRequest, opts ...grpc.CallOption) (*NetworkSpec, error)
	// Status produces a task status by their ID
	Status(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (*TaskStatusReply, error)
	// Logs retrieves a task log (stdin/stderr) from given task
	Logs(ctx context.Context, in *TaskLogsRequest, opts ...grpc.CallOption) (TaskManagement_LogsClient, error)
	// Stop stops a task by their ID
	Stop(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (*Empty, error)
	// PullTask pulls task image back
	PullTask(ctx context.Context, in *PullTaskRequest, opts ...grpc.CallOption) (TaskManagement_PullTaskClient, error)
}

type taskManagementClient struct {
	cc *grpc.ClientConn
}

func NewTaskManagementClient(cc *grpc.ClientConn) TaskManagementClient {
	return &taskManagementClient{cc}
}

func (c *taskManagementClient) List(ctx context.Context, in *EthAddress, opts ...grpc.CallOption) (*TaskListReply, error) {
	out := new(TaskListReply)
	err := grpc.Invoke(ctx, "/sonm.TaskManagement/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagementClient) PushTask(ctx context.Context, opts ...grpc.CallOption) (TaskManagement_PushTaskClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TaskManagement_serviceDesc.Streams[0], c.cc, "/sonm.TaskManagement/PushTask", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskManagementPushTaskClient{stream}
	return x, nil
}

type TaskManagement_PushTaskClient interface {
	Send(*Chunk) error
	Recv() (*Progress, error)
	grpc.ClientStream
}

type taskManagementPushTaskClient struct {
	grpc.ClientStream
}

func (x *taskManagementPushTaskClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *taskManagementPushTaskClient) Recv() (*Progress, error) {
	m := new(Progress)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *taskManagementClient) Start(ctx context.Context, in *StartTaskRequest, opts ...grpc.CallOption) (*StartTaskReply, error) {
	out := new(StartTaskReply)
	err := grpc.Invoke(ctx, "/sonm.TaskManagement/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagementClient) JoinNetwork(ctx context.Context, in *JoinNetworkRequest, opts ...grpc.CallOption) (*NetworkSpec, error) {
	out := new(NetworkSpec)
	err := grpc.Invoke(ctx, "/sonm.TaskManagement/JoinNetwork", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagementClient) Status(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (*TaskStatusReply, error) {
	out := new(TaskStatusReply)
	err := grpc.Invoke(ctx, "/sonm.TaskManagement/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagementClient) Logs(ctx context.Context, in *TaskLogsRequest, opts ...grpc.CallOption) (TaskManagement_LogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TaskManagement_serviceDesc.Streams[1], c.cc, "/sonm.TaskManagement/Logs", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskManagementLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TaskManagement_LogsClient interface {
	Recv() (*TaskLogsChunk, error)
	grpc.ClientStream
}

type taskManagementLogsClient struct {
	grpc.ClientStream
}

func (x *taskManagementLogsClient) Recv() (*TaskLogsChunk, error) {
	m := new(TaskLogsChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *taskManagementClient) Stop(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/sonm.TaskManagement/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagementClient) PullTask(ctx context.Context, in *PullTaskRequest, opts ...grpc.CallOption) (TaskManagement_PullTaskClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TaskManagement_serviceDesc.Streams[2], c.cc, "/sonm.TaskManagement/PullTask", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskManagementPullTaskClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TaskManagement_PullTaskClient interface {
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type taskManagementPullTaskClient struct {
	grpc.ClientStream
}

func (x *taskManagementPullTaskClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TaskManagement service

type TaskManagementServer interface {
	// List produces a list of all tasks running on different SONM nodes
	List(context.Context, *EthAddress) (*TaskListReply, error)
	// PushTask pushes image to Worker
	PushTask(TaskManagement_PushTaskServer) error
	// Start starts a task on given resource
	Start(context.Context, *StartTaskRequest) (*StartTaskReply, error)
	// JoinNetwork provides network specs to join specified task
	JoinNetwork(context.Context, *JoinNetworkRequest) (*NetworkSpec, error)
	// Status produces a task status by their ID
	Status(context.Context, *TaskID) (*TaskStatusReply, error)
	// Logs retrieves a task log (stdin/stderr) from given task
	Logs(*TaskLogsRequest, TaskManagement_LogsServer) error
	// Stop stops a task by their ID
	Stop(context.Context, *TaskID) (*Empty, error)
	// PullTask pulls task image back
	PullTask(*PullTaskRequest, TaskManagement_PullTaskServer) error
}

func RegisterTaskManagementServer(s *grpc.Server, srv TaskManagementServer) {
	s.RegisterService(&_TaskManagement_serviceDesc, srv)
}

func _TaskManagement_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EthAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagementServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.TaskManagement/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagementServer).List(ctx, req.(*EthAddress))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManagement_PushTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TaskManagementServer).PushTask(&taskManagementPushTaskServer{stream})
}

type TaskManagement_PushTaskServer interface {
	Send(*Progress) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type taskManagementPushTaskServer struct {
	grpc.ServerStream
}

func (x *taskManagementPushTaskServer) Send(m *Progress) error {
	return x.ServerStream.SendMsg(m)
}

func (x *taskManagementPushTaskServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TaskManagement_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagementServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.TaskManagement/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagementServer).Start(ctx, req.(*StartTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManagement_JoinNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagementServer).JoinNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.TaskManagement/JoinNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagementServer).JoinNetwork(ctx, req.(*JoinNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManagement_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagementServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.TaskManagement/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagementServer).Status(ctx, req.(*TaskID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManagement_Logs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TaskLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TaskManagementServer).Logs(m, &taskManagementLogsServer{stream})
}

type TaskManagement_LogsServer interface {
	Send(*TaskLogsChunk) error
	grpc.ServerStream
}

type taskManagementLogsServer struct {
	grpc.ServerStream
}

func (x *taskManagementLogsServer) Send(m *TaskLogsChunk) error {
	return x.ServerStream.SendMsg(m)
}

func _TaskManagement_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagementServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.TaskManagement/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagementServer).Stop(ctx, req.(*TaskID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManagement_PullTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullTaskRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TaskManagementServer).PullTask(m, &taskManagementPullTaskServer{stream})
}

type TaskManagement_PullTaskServer interface {
	Send(*Chunk) error
	grpc.ServerStream
}

type taskManagementPullTaskServer struct {
	grpc.ServerStream
}

func (x *taskManagementPullTaskServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

var _TaskManagement_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonm.TaskManagement",
	HandlerType: (*TaskManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _TaskManagement_List_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _TaskManagement_Start_Handler,
		},
		{
			MethodName: "JoinNetwork",
			Handler:    _TaskManagement_JoinNetwork_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _TaskManagement_Status_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _TaskManagement_Stop_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PushTask",
			Handler:       _TaskManagement_PushTask_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Logs",
			Handler:       _TaskManagement_Logs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PullTask",
			Handler:       _TaskManagement_PullTask_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "node.proto",
}

// Client API for DealManagement service

type DealManagementClient interface {
	// List produces a list of all deals made by client with given ID
	List(ctx context.Context, in *Count, opts ...grpc.CallOption) (*DealsReply, error)
	// Status produces a detailed info about deal with given ID.
	Status(ctx context.Context, in *ID, opts ...grpc.CallOption) (*DealInfoReply, error)
	// Finish finishes a deal with given ID
	Finish(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error)
}

type dealManagementClient struct {
	cc *grpc.ClientConn
}

func NewDealManagementClient(cc *grpc.ClientConn) DealManagementClient {
	return &dealManagementClient{cc}
}

func (c *dealManagementClient) List(ctx context.Context, in *Count, opts ...grpc.CallOption) (*DealsReply, error) {
	out := new(DealsReply)
	err := grpc.Invoke(ctx, "/sonm.DealManagement/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dealManagementClient) Status(ctx context.Context, in *ID, opts ...grpc.CallOption) (*DealInfoReply, error) {
	out := new(DealInfoReply)
	err := grpc.Invoke(ctx, "/sonm.DealManagement/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dealManagementClient) Finish(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/sonm.DealManagement/Finish", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DealManagement service

type DealManagementServer interface {
	// List produces a list of all deals made by client with given ID
	List(context.Context, *Count) (*DealsReply, error)
	// Status produces a detailed info about deal with given ID.
	Status(context.Context, *ID) (*DealInfoReply, error)
	// Finish finishes a deal with given ID
	Finish(context.Context, *ID) (*Empty, error)
}

func RegisterDealManagementServer(s *grpc.Server, srv DealManagementServer) {
	s.RegisterService(&_DealManagement_serviceDesc, srv)
}

func _DealManagement_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Count)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DealManagementServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.DealManagement/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DealManagementServer).List(ctx, req.(*Count))
	}
	return interceptor(ctx, in, info, handler)
}

func _DealManagement_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DealManagementServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.DealManagement/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DealManagementServer).Status(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DealManagement_Finish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DealManagementServer).Finish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.DealManagement/Finish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DealManagementServer).Finish(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

var _DealManagement_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonm.DealManagement",
	HandlerType: (*DealManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _DealManagement_List_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _DealManagement_Status_Handler,
		},
		{
			MethodName: "Finish",
			Handler:    _DealManagement_Finish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}

// Client API for MasterManagement service

type MasterManagementClient interface {
	// WorkersList returns worker's list for current master address.
	// List includes already registred workers and pending unapproved requests.
	WorkersList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*WorkerListReply, error)
	// WorkerConfirm (as master) confirms incoming request for given Worker address.
	WorkerConfirm(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error)
	// WorkerRemove (as master) unbinds given Worker address from Master address.
	WorkerRemove(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error)
}

type masterManagementClient struct {
	cc *grpc.ClientConn
}

func NewMasterManagementClient(cc *grpc.ClientConn) MasterManagementClient {
	return &masterManagementClient{cc}
}

func (c *masterManagementClient) WorkersList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*WorkerListReply, error) {
	out := new(WorkerListReply)
	err := grpc.Invoke(ctx, "/sonm.MasterManagement/WorkersList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterManagementClient) WorkerConfirm(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/sonm.MasterManagement/WorkerConfirm", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterManagementClient) WorkerRemove(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/sonm.MasterManagement/WorkerRemove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MasterManagement service

type MasterManagementServer interface {
	// WorkersList returns worker's list for current master address.
	// List includes already registred workers and pending unapproved requests.
	WorkersList(context.Context, *Empty) (*WorkerListReply, error)
	// WorkerConfirm (as master) confirms incoming request for given Worker address.
	WorkerConfirm(context.Context, *ID) (*Empty, error)
	// WorkerRemove (as master) unbinds given Worker address from Master address.
	WorkerRemove(context.Context, *ID) (*Empty, error)
}

func RegisterMasterManagementServer(s *grpc.Server, srv MasterManagementServer) {
	s.RegisterService(&_MasterManagement_serviceDesc, srv)
}

func _MasterManagement_WorkersList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterManagementServer).WorkersList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.MasterManagement/WorkersList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterManagementServer).WorkersList(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterManagement_WorkerConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterManagementServer).WorkerConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.MasterManagement/WorkerConfirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterManagementServer).WorkerConfirm(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterManagement_WorkerRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterManagementServer).WorkerRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.MasterManagement/WorkerRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterManagementServer).WorkerRemove(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

var _MasterManagement_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonm.MasterManagement",
	HandlerType: (*MasterManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WorkersList",
			Handler:    _MasterManagement_WorkersList_Handler,
		},
		{
			MethodName: "WorkerConfirm",
			Handler:    _MasterManagement_WorkerConfirm_Handler,
		},
		{
			MethodName: "WorkerRemove",
			Handler:    _MasterManagement_WorkerRemove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}

// Begin grpccmd
var _ = grpccmd.RunE

// TaskManagement
var _TaskManagementCmd = &cobra.Command{
	Use:   "taskManagement [method]",
	Short: "Subcommand for the TaskManagement service.",
}

var _TaskManagement_ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Make the List method call, input-type: sonm.EthAddress output-type: sonm.TaskListReply",
	RunE: grpccmd.RunE(
		"List",
		"sonm.EthAddress",
		func(c io.Closer) interface{} {
			cc := c.(*grpc.ClientConn)
			return NewTaskManagementClient(cc)
		},
	),
}

var _TaskManagement_ListCmd_gen = &cobra.Command{
	Use:   "list-gen",
	Short: "Generate JSON for method call of List (input-type: sonm.EthAddress)",
	RunE:  grpccmd.TypeToJson("sonm.EthAddress"),
}

var _TaskManagement_PushTaskCmd = &cobra.Command{
	Use:   "pushTask",
	Short: "Make the PushTask method call, input-type: sonm.Chunk output-type: sonm.Progress",
	RunE: grpccmd.RunE(
		"PushTask",
		"sonm.Chunk",
		func(c io.Closer) interface{} {
			cc := c.(*grpc.ClientConn)
			return NewTaskManagementClient(cc)
		},
	),
}

var _TaskManagement_PushTaskCmd_gen = &cobra.Command{
	Use:   "pushTask-gen",
	Short: "Generate JSON for method call of PushTask (input-type: sonm.Chunk)",
	RunE:  grpccmd.TypeToJson("sonm.Chunk"),
}

var _TaskManagement_StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Make the Start method call, input-type: sonm.StartTaskRequest output-type: sonm.StartTaskReply",
	RunE: grpccmd.RunE(
		"Start",
		"sonm.StartTaskRequest",
		func(c io.Closer) interface{} {
			cc := c.(*grpc.ClientConn)
			return NewTaskManagementClient(cc)
		},
	),
}

var _TaskManagement_StartCmd_gen = &cobra.Command{
	Use:   "start-gen",
	Short: "Generate JSON for method call of Start (input-type: sonm.StartTaskRequest)",
	RunE:  grpccmd.TypeToJson("sonm.StartTaskRequest"),
}

var _TaskManagement_JoinNetworkCmd = &cobra.Command{
	Use:   "joinNetwork",
	Short: "Make the JoinNetwork method call, input-type: sonm.JoinNetworkRequest output-type: sonm.NetworkSpec",
	RunE: grpccmd.RunE(
		"JoinNetwork",
		"sonm.JoinNetworkRequest",
		func(c io.Closer) interface{} {
			cc := c.(*grpc.ClientConn)
			return NewTaskManagementClient(cc)
		},
	),
}

var _TaskManagement_JoinNetworkCmd_gen = &cobra.Command{
	Use:   "joinNetwork-gen",
	Short: "Generate JSON for method call of JoinNetwork (input-type: sonm.JoinNetworkRequest)",
	RunE:  grpccmd.TypeToJson("sonm.JoinNetworkRequest"),
}

var _TaskManagement_StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Make the Status method call, input-type: sonm.TaskID output-type: sonm.TaskStatusReply",
	RunE: grpccmd.RunE(
		"Status",
		"sonm.TaskID",
		func(c io.Closer) interface{} {
			cc := c.(*grpc.ClientConn)
			return NewTaskManagementClient(cc)
		},
	),
}

var _TaskManagement_StatusCmd_gen = &cobra.Command{
	Use:   "status-gen",
	Short: "Generate JSON for method call of Status (input-type: sonm.TaskID)",
	RunE:  grpccmd.TypeToJson("sonm.TaskID"),
}

var _TaskManagement_LogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Make the Logs method call, input-type: sonm.TaskLogsRequest output-type: sonm.TaskLogsChunk",
	RunE: grpccmd.RunE(
		"Logs",
		"sonm.TaskLogsRequest",
		func(c io.Closer) interface{} {
			cc := c.(*grpc.ClientConn)
			return NewTaskManagementClient(cc)
		},
	),
}

var _TaskManagement_LogsCmd_gen = &cobra.Command{
	Use:   "logs-gen",
	Short: "Generate JSON for method call of Logs (input-type: sonm.TaskLogsRequest)",
	RunE:  grpccmd.TypeToJson("sonm.TaskLogsRequest"),
}

var _TaskManagement_StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Make the Stop method call, input-type: sonm.TaskID output-type: sonm.Empty",
	RunE: grpccmd.RunE(
		"Stop",
		"sonm.TaskID",
		func(c io.Closer) interface{} {
			cc := c.(*grpc.ClientConn)
			return NewTaskManagementClient(cc)
		},
	),
}

var _TaskManagement_StopCmd_gen = &cobra.Command{
	Use:   "stop-gen",
	Short: "Generate JSON for method call of Stop (input-type: sonm.TaskID)",
	RunE:  grpccmd.TypeToJson("sonm.TaskID"),
}

var _TaskManagement_PullTaskCmd = &cobra.Command{
	Use:   "pullTask",
	Short: "Make the PullTask method call, input-type: sonm.PullTaskRequest output-type: sonm.Chunk",
	RunE: grpccmd.RunE(
		"PullTask",
		"sonm.PullTaskRequest",
		func(c io.Closer) interface{} {
			cc := c.(*grpc.ClientConn)
			return NewTaskManagementClient(cc)
		},
	),
}

var _TaskManagement_PullTaskCmd_gen = &cobra.Command{
	Use:   "pullTask-gen",
	Short: "Generate JSON for method call of PullTask (input-type: sonm.PullTaskRequest)",
	RunE:  grpccmd.TypeToJson("sonm.PullTaskRequest"),
}

// Register commands with the root command and service command
func init() {
	grpccmd.RegisterServiceCmd(_TaskManagementCmd)
	_TaskManagementCmd.AddCommand(
		_TaskManagement_ListCmd,
		_TaskManagement_ListCmd_gen,
		_TaskManagement_PushTaskCmd,
		_TaskManagement_PushTaskCmd_gen,
		_TaskManagement_StartCmd,
		_TaskManagement_StartCmd_gen,
		_TaskManagement_JoinNetworkCmd,
		_TaskManagement_JoinNetworkCmd_gen,
		_TaskManagement_StatusCmd,
		_TaskManagement_StatusCmd_gen,
		_TaskManagement_LogsCmd,
		_TaskManagement_LogsCmd_gen,
		_TaskManagement_StopCmd,
		_TaskManagement_StopCmd_gen,
		_TaskManagement_PullTaskCmd,
		_TaskManagement_PullTaskCmd_gen,
	)
}

// DealManagement
var _DealManagementCmd = &cobra.Command{
	Use:   "dealManagement [method]",
	Short: "Subcommand for the DealManagement service.",
}

var _DealManagement_ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Make the List method call, input-type: sonm.Count output-type: sonm.DealsReply",
	RunE: grpccmd.RunE(
		"List",
		"sonm.Count",
		func(c io.Closer) interface{} {
			cc := c.(*grpc.ClientConn)
			return NewDealManagementClient(cc)
		},
	),
}

var _DealManagement_ListCmd_gen = &cobra.Command{
	Use:   "list-gen",
	Short: "Generate JSON for method call of List (input-type: sonm.Count)",
	RunE:  grpccmd.TypeToJson("sonm.Count"),
}

var _DealManagement_StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Make the Status method call, input-type: sonm.ID output-type: sonm.DealInfoReply",
	RunE: grpccmd.RunE(
		"Status",
		"sonm.ID",
		func(c io.Closer) interface{} {
			cc := c.(*grpc.ClientConn)
			return NewDealManagementClient(cc)
		},
	),
}

var _DealManagement_StatusCmd_gen = &cobra.Command{
	Use:   "status-gen",
	Short: "Generate JSON for method call of Status (input-type: sonm.ID)",
	RunE:  grpccmd.TypeToJson("sonm.ID"),
}

var _DealManagement_FinishCmd = &cobra.Command{
	Use:   "finish",
	Short: "Make the Finish method call, input-type: sonm.ID output-type: sonm.Empty",
	RunE: grpccmd.RunE(
		"Finish",
		"sonm.ID",
		func(c io.Closer) interface{} {
			cc := c.(*grpc.ClientConn)
			return NewDealManagementClient(cc)
		},
	),
}

var _DealManagement_FinishCmd_gen = &cobra.Command{
	Use:   "finish-gen",
	Short: "Generate JSON for method call of Finish (input-type: sonm.ID)",
	RunE:  grpccmd.TypeToJson("sonm.ID"),
}

// Register commands with the root command and service command
func init() {
	grpccmd.RegisterServiceCmd(_DealManagementCmd)
	_DealManagementCmd.AddCommand(
		_DealManagement_ListCmd,
		_DealManagement_ListCmd_gen,
		_DealManagement_StatusCmd,
		_DealManagement_StatusCmd_gen,
		_DealManagement_FinishCmd,
		_DealManagement_FinishCmd_gen,
	)
}

// MasterManagement
var _MasterManagementCmd = &cobra.Command{
	Use:   "masterManagement [method]",
	Short: "Subcommand for the MasterManagement service.",
}

var _MasterManagement_WorkersListCmd = &cobra.Command{
	Use:   "workersList",
	Short: "Make the WorkersList method call, input-type: sonm.Empty output-type: sonm.WorkerListReply",
	RunE: grpccmd.RunE(
		"WorkersList",
		"sonm.Empty",
		func(c io.Closer) interface{} {
			cc := c.(*grpc.ClientConn)
			return NewMasterManagementClient(cc)
		},
	),
}

var _MasterManagement_WorkersListCmd_gen = &cobra.Command{
	Use:   "workersList-gen",
	Short: "Generate JSON for method call of WorkersList (input-type: sonm.Empty)",
	RunE:  grpccmd.TypeToJson("sonm.Empty"),
}

var _MasterManagement_WorkerConfirmCmd = &cobra.Command{
	Use:   "workerConfirm",
	Short: "Make the WorkerConfirm method call, input-type: sonm.ID output-type: sonm.Empty",
	RunE: grpccmd.RunE(
		"WorkerConfirm",
		"sonm.ID",
		func(c io.Closer) interface{} {
			cc := c.(*grpc.ClientConn)
			return NewMasterManagementClient(cc)
		},
	),
}

var _MasterManagement_WorkerConfirmCmd_gen = &cobra.Command{
	Use:   "workerConfirm-gen",
	Short: "Generate JSON for method call of WorkerConfirm (input-type: sonm.ID)",
	RunE:  grpccmd.TypeToJson("sonm.ID"),
}

var _MasterManagement_WorkerRemoveCmd = &cobra.Command{
	Use:   "workerRemove",
	Short: "Make the WorkerRemove method call, input-type: sonm.ID output-type: sonm.Empty",
	RunE: grpccmd.RunE(
		"WorkerRemove",
		"sonm.ID",
		func(c io.Closer) interface{} {
			cc := c.(*grpc.ClientConn)
			return NewMasterManagementClient(cc)
		},
	),
}

var _MasterManagement_WorkerRemoveCmd_gen = &cobra.Command{
	Use:   "workerRemove-gen",
	Short: "Generate JSON for method call of WorkerRemove (input-type: sonm.ID)",
	RunE:  grpccmd.TypeToJson("sonm.ID"),
}

// Register commands with the root command and service command
func init() {
	grpccmd.RegisterServiceCmd(_MasterManagementCmd)
	_MasterManagementCmd.AddCommand(
		_MasterManagement_WorkersListCmd,
		_MasterManagement_WorkersListCmd_gen,
		_MasterManagement_WorkerConfirmCmd,
		_MasterManagement_WorkerConfirmCmd_gen,
		_MasterManagement_WorkerRemoveCmd,
		_MasterManagement_WorkerRemoveCmd_gen,
	)
}

// End grpccmd

func init() { proto.RegisterFile("node.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xb5, 0xdb, 0x7c, 0xfe, 0x9a, 0x49, 0x49, 0xd3, 0x2d, 0x85, 0xc8, 0x42, 0x55, 0x31, 0x88,
	0x06, 0x04, 0x69, 0xe5, 0x4a, 0x45, 0x48, 0xdc, 0x44, 0x75, 0x90, 0x5c, 0xa5, 0x69, 0xe4, 0x94,
	0x9f, 0x3b, 0xb4, 0x4d, 0xb6, 0xb1, 0x15, 0x7b, 0xd7, 0x78, 0xd7, 0x54, 0x3c, 0x01, 0xef, 0xc0,
	0x1b, 0xf0, 0x96, 0x68, 0xbd, 0xeb, 0xc4, 0x69, 0x95, 0x3b, 0xcf, 0x99, 0x33, 0x73, 0x66, 0x67,
	0x8e, 0x0c, 0x40, 0xd9, 0x94, 0x74, 0xd3, 0x8c, 0x09, 0x86, 0x6a, 0x9c, 0xd1, 0xc4, 0xde, 0x99,
	0x30, 0x2a, 0x70, 0x44, 0x49, 0xa6, 0x60, 0xbb, 0x1e, 0xe6, 0x37, 0xfa, 0x73, 0x27, 0xa2, 0x92,
	0x43, 0x23, 0xac, 0x81, 0xdd, 0x04, 0x67, 0x73, 0x22, 0xd2, 0x18, 0x4f, 0x74, 0x17, 0xe7, 0x1b,
	0xa0, 0x0b, 0x16, 0xd1, 0x21, 0x11, 0x77, 0x2c, 0x9b, 0x07, 0xe4, 0x47, 0x4e, 0xb8, 0x40, 0x2f,
	0xc1, 0x12, 0x98, 0xcf, 0x7d, 0xaf, 0x6d, 0x1e, 0x9a, 0x9d, 0x86, 0xbb, 0xdd, 0x95, 0x8d, 0xba,
	0xd7, 0x05, 0x16, 0xe8, 0x1c, 0x7a, 0x06, 0x75, 0x5d, 0xe7, 0x7b, 0xed, 0x8d, 0x43, 0xb3, 0x53,
	0x0f, 0x96, 0x80, 0xf3, 0x16, 0xc0, 0x23, 0x38, 0xe6, 0x01, 0x49, 0xe3, 0x5f, 0xe8, 0x00, 0x6a,
	0x53, 0x82, 0xe3, 0xb6, 0x79, 0xb8, 0xd9, 0x69, 0xb8, 0xa0, 0xfa, 0xc9, 0x7c, 0x50, 0xe0, 0xce,
	0x08, 0xac, 0xaf, 0x2c, 0x9b, 0x93, 0x0c, 0x35, 0x61, 0x43, 0xeb, 0xd6, 0x83, 0x0d, 0xdf, 0x43,
	0x67, 0x60, 0x71, 0x81, 0x45, 0xce, 0x0b, 0x89, 0xa6, 0x7b, 0xa0, 0x6a, 0x15, 0x3b, 0x20, 0x31,
	0x16, 0x11, 0xa3, 0x3c, 0x8c, 0xd2, 0x71, 0xc1, 0x0a, 0x34, 0xdb, 0xf9, 0x00, 0x3b, 0x8a, 0x33,
	0x88, 0xb8, 0x50, 0x43, 0xbc, 0x82, 0xff, 0xef, 0x0a, 0x88, 0xeb, 0x39, 0xb6, 0x57, 0x7a, 0x95,
	0xc9, 0x37, 0x17, 0xd0, 0x5e, 0xd7, 0x1e, 0x3d, 0x85, 0xbd, 0xa0, 0x3f, 0xe8, 0x5d, 0xfb, 0x57,
	0xc3, 0xef, 0x9f, 0x87, 0xbd, 0xd1, 0x28, 0xb8, 0xfa, 0xd2, 0xf7, 0x5a, 0x06, 0xda, 0x87, 0xdd,
	0x45, 0x62, 0x01, 0x9b, 0xee, 0xdf, 0x4d, 0x68, 0xca, 0xbd, 0x5d, 0x62, 0x8a, 0x67, 0x24, 0x21,
	0x54, 0xa0, 0x63, 0xa8, 0xc9, 0x99, 0x50, 0x4b, 0xa9, 0xf7, 0x45, 0xd8, 0x9b, 0x4e, 0x33, 0xc2,
	0xb9, 0xbd, 0xb7, 0xdc, 0xf3, 0x62, 0x6a, 0xc7, 0x40, 0xef, 0x60, 0x6b, 0x94, 0xf3, 0x50, 0xc2,
	0xa8, 0xa1, 0x28, 0xe7, 0x61, 0x4e, 0xe7, 0x76, 0x53, 0x05, 0xa3, 0x8c, 0xcd, 0x64, 0xbd, 0x63,
	0x74, 0xcc, 0x13, 0x13, 0xbd, 0x87, 0xff, 0xc6, 0x02, 0x67, 0x02, 0x3d, 0x51, 0xe9, 0x22, 0x90,
	0xc5, 0xfa, 0xbc, 0xf6, 0xe3, 0x07, 0xb8, 0xd2, 0xf9, 0x08, 0x8d, 0x8a, 0x19, 0x50, 0x5b, 0xd1,
	0x1e, 0xfa, 0xc3, 0xde, 0x55, 0x19, 0x8d, 0x8e, 0x53, 0x32, 0x71, 0x0c, 0x74, 0x0c, 0x96, 0xde,
	0xd1, 0x8a, 0x5d, 0xec, 0xfd, 0x65, 0xa4, 0x4f, 0xa4, 0xe5, 0xce, 0xa0, 0x36, 0x60, 0x33, 0x8e,
	0x2a, 0x04, 0x19, 0x97, 0x22, 0x7b, 0xab, 0x70, 0xf1, 0x62, 0xc7, 0x38, 0x31, 0xd1, 0x0b, 0xa8,
	0x8d, 0x05, 0x4b, 0xef, 0xc9, 0xe8, 0xc5, 0xf4, 0x93, 0x54, 0xc8, 0xe6, 0xae, 0xdc, 0x59, 0x1c,
	0x17, 0x3b, 0xd3, 0x02, 0x65, 0x5c, 0x0a, 0x54, 0x57, 0x29, 0x1b, 0xbb, 0xbf, 0x4d, 0x68, 0x4a,
	0x4f, 0x56, 0x6e, 0x75, 0xa4, 0x6f, 0x55, 0x72, 0x59, 0x4e, 0x85, 0xdd, 0x5a, 0xda, 0x77, 0xf1,
	0x98, 0xd7, 0x8b, 0xd7, 0x6f, 0xa9, 0xac, 0xef, 0x95, 0x2f, 0x90, 0x3c, 0x9f, 0xde, 0xb2, 0x92,
	0xfa, 0x1c, 0xac, 0x4f, 0x11, 0x8d, 0x78, 0x58, 0xa1, 0xae, 0x4e, 0xef, 0xfe, 0x31, 0xa1, 0x75,
	0x89, 0xb9, 0x20, 0x59, 0x65, 0x96, 0x53, 0x68, 0x28, 0x5b, 0xf2, 0xea, 0x48, 0x45, 0x49, 0xb9,
	0xe4, 0x7b, 0x8e, 0x77, 0x0c, 0xd4, 0x81, 0x47, 0x0a, 0x3c, 0x67, 0xf4, 0x36, 0xca, 0x92, 0xb5,
	0x9a, 0xe8, 0x08, 0xb6, 0x4b, 0xd7, 0x27, 0xec, 0x27, 0x59, 0x4b, 0xbc, 0xb1, 0x8a, 0x5f, 0xc7,
	0xe9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x56, 0x91, 0x5f, 0x7b, 0x8e, 0x04, 0x00, 0x00,
}
