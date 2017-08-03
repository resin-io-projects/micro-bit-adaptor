// Code generated by protoc-gen-go. DO NOT EDIT.
// source: update.proto

/*
Package update is a generated protocol buffer package.

It is generated from these files:
	update.proto

It has these top-level messages:
	StartRequest
	StartResponse
	StatusRequest
	StatusResponse
*/
package update

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

type StatusResponse_State int32

const (
	StatusResponse_CREATED          StatusResponse_State = 0
	StatusResponse_QUEUED           StatusResponse_State = 1
	StatusResponse_STARTED          StatusResponse_State = 2
	StatusResponse_EXTRACTING       StatusResponse_State = 3
	StatusResponse_CONNECTING       StatusResponse_State = 4
	StatusResponse_FLASHING         StatusResponse_State = 5
	StatusResponse_VALIDATING       StatusResponse_State = 6
	StatusResponse_COMPLETED        StatusResponse_State = 7
	StatusResponse_CANCELLED        StatusResponse_State = 8
	StatusResponse_EXTRACT_FAILURE  StatusResponse_State = 9
	StatusResponse_CONNECT_FAILURE  StatusResponse_State = 10
	StatusResponse_FLASH_FAILURE    StatusResponse_State = 11
	StatusResponse_VALIDATE_FAILURE StatusResponse_State = 12
)

var StatusResponse_State_name = map[int32]string{
	0:  "CREATED",
	1:  "QUEUED",
	2:  "STARTED",
	3:  "EXTRACTING",
	4:  "CONNECTING",
	5:  "FLASHING",
	6:  "VALIDATING",
	7:  "COMPLETED",
	8:  "CANCELLED",
	9:  "EXTRACT_FAILURE",
	10: "CONNECT_FAILURE",
	11: "FLASH_FAILURE",
	12: "VALIDATE_FAILURE",
}
var StatusResponse_State_value = map[string]int32{
	"CREATED":          0,
	"QUEUED":           1,
	"STARTED":          2,
	"EXTRACTING":       3,
	"CONNECTING":       4,
	"FLASHING":         5,
	"VALIDATING":       6,
	"COMPLETED":        7,
	"CANCELLED":        8,
	"EXTRACT_FAILURE":  9,
	"CONNECT_FAILURE":  10,
	"FLASH_FAILURE":    11,
	"VALIDATE_FAILURE": 12,
}

func (x StatusResponse_State) String() string {
	return proto.EnumName(StatusResponse_State_name, int32(x))
}
func (StatusResponse_State) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type StartRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Payload string `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
	Timeout int64  `protobuf:"varint,3,opt,name=timeout" json:"timeout,omitempty"`
	Device  string `protobuf:"bytes,4,opt,name=device" json:"device,omitempty"`
}

func (m *StartRequest) Reset()                    { *m = StartRequest{} }
func (m *StartRequest) String() string            { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()               {}
func (*StartRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StartRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *StartRequest) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *StartRequest) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *StartRequest) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

type StartResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *StartResponse) Reset()                    { *m = StartResponse{} }
func (m *StartResponse) String() string            { return proto.CompactTextString(m) }
func (*StartResponse) ProtoMessage()               {}
func (*StartResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StartResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type StatusRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *StatusRequest) Reset()                    { *m = StatusRequest{} }
func (m *StatusRequest) String() string            { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()               {}
func (*StatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StatusRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type StatusResponse struct {
	Id           string               `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	StartRequest *StartRequest        `protobuf:"bytes,2,opt,name=startRequest" json:"startRequest,omitempty"`
	State        StatusResponse_State `protobuf:"varint,3,opt,name=state,enum=update.StatusResponse_State" json:"state,omitempty"`
	Progress     int32                `protobuf:"varint,4,opt,name=progress" json:"progress,omitempty"`
	Message      string               `protobuf:"bytes,5,opt,name=message" json:"message,omitempty"`
	Started      int64                `protobuf:"varint,6,opt,name=started" json:"started,omitempty"`
	Completed    int64                `protobuf:"varint,7,opt,name=completed" json:"completed,omitempty"`
	Duration     int64                `protobuf:"varint,8,opt,name=duration" json:"duration,omitempty"`
}

func (m *StatusResponse) Reset()                    { *m = StatusResponse{} }
func (m *StatusResponse) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()               {}
func (*StatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StatusResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StatusResponse) GetStartRequest() *StartRequest {
	if m != nil {
		return m.StartRequest
	}
	return nil
}

func (m *StatusResponse) GetState() StatusResponse_State {
	if m != nil {
		return m.State
	}
	return StatusResponse_CREATED
}

func (m *StatusResponse) GetProgress() int32 {
	if m != nil {
		return m.Progress
	}
	return 0
}

func (m *StatusResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *StatusResponse) GetStarted() int64 {
	if m != nil {
		return m.Started
	}
	return 0
}

func (m *StatusResponse) GetCompleted() int64 {
	if m != nil {
		return m.Completed
	}
	return 0
}

func (m *StatusResponse) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func init() {
	proto.RegisterType((*StartRequest)(nil), "update.StartRequest")
	proto.RegisterType((*StartResponse)(nil), "update.StartResponse")
	proto.RegisterType((*StatusRequest)(nil), "update.StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "update.StatusResponse")
	proto.RegisterEnum("update.StatusResponse_State", StatusResponse_State_name, StatusResponse_State_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Update service

type UpdateClient interface {
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	Cancel(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type updateClient struct {
	cc *grpc.ClientConn
}

func NewUpdateClient(cc *grpc.ClientConn) UpdateClient {
	return &updateClient{cc}
}

func (c *updateClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := grpc.Invoke(ctx, "/update.Update/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *updateClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := grpc.Invoke(ctx, "/update.Update/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *updateClient) Cancel(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := grpc.Invoke(ctx, "/update.Update/Cancel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Update service

type UpdateServer interface {
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	Cancel(context.Context, *StatusRequest) (*StatusResponse, error)
}

func RegisterUpdateServer(s *grpc.Server, srv UpdateServer) {
	s.RegisterService(&_Update_serviceDesc, srv)
}

func _Update_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/update.Update/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Update_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/update.Update/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Update_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/update.Update/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateServer).Cancel(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Update_serviceDesc = grpc.ServiceDesc{
	ServiceName: "update.Update",
	HandlerType: (*UpdateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Update_Start_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Update_Status_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _Update_Cancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "update.proto",
}

func init() { proto.RegisterFile("update.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x8e, 0xd2, 0x50,
	0x14, 0xc6, 0x2d, 0x0c, 0x05, 0xce, 0x14, 0xac, 0x47, 0x66, 0xd2, 0x10, 0x12, 0x27, 0x5d, 0x4d,
	0x66, 0x31, 0x24, 0xb8, 0x31, 0xee, 0x9a, 0x72, 0xc7, 0x21, 0xe9, 0x30, 0x5a, 0xc0, 0xb8, 0x33,
	0x57, 0x7a, 0x43, 0x9a, 0x00, 0xb7, 0xf6, 0x5e, 0x4c, 0x8c, 0x71, 0xe3, 0x2b, 0xf8, 0x68, 0x6e,
	0x5d, 0xfa, 0x00, 0xee, 0xdc, 0x9a, 0xfb, 0x07, 0x64, 0x08, 0x2b, 0x97, 0xbf, 0xf3, 0x9d, 0x7e,
	0xdf, 0x3d, 0xe7, 0x00, 0x78, 0x9b, 0x22, 0xa3, 0x92, 0x5d, 0x17, 0x25, 0x97, 0x1c, 0x5d, 0x43,
	0xdd, 0xde, 0x82, 0xf3, 0xc5, 0x92, 0xf5, 0x69, 0x91, 0xf7, 0xe9, 0x7a, 0xcd, 0x25, 0x95, 0x39,
	0x5f, 0x0b, 0xd3, 0x15, 0x4a, 0xf0, 0x26, 0x92, 0x96, 0x32, 0x65, 0x1f, 0x37, 0x4c, 0x48, 0x0c,
	0xa0, 0x4e, 0xb3, 0xac, 0x64, 0x42, 0x04, 0xce, 0x85, 0x73, 0xd9, 0x4c, 0xb7, 0xa8, 0x94, 0x82,
	0x7e, 0x5e, 0x72, 0x9a, 0x05, 0x15, 0xa3, 0x58, 0x54, 0x8a, 0xcc, 0x57, 0x8c, 0x6f, 0x64, 0x50,
	0xbd, 0x70, 0x2e, 0xab, 0xe9, 0x16, 0xf1, 0x1c, 0xdc, 0x8c, 0x7d, 0xca, 0xe7, 0x2c, 0x38, 0xd1,
	0x9f, 0x58, 0x0a, 0x9f, 0x41, 0xcb, 0xa6, 0x8a, 0x82, 0xaf, 0x05, 0xc3, 0x36, 0x54, 0xf2, 0xcc,
	0x26, 0x56, 0xf2, 0xcc, 0x36, 0xc8, 0x8d, 0xd8, 0xbe, 0xeb, 0xb0, 0xe1, 0x4f, 0x15, 0xda, 0xdb,
	0x8e, 0xe3, 0x1e, 0xf8, 0x02, 0x3c, 0xb1, 0x37, 0x9a, 0x7e, 0xf5, 0xe9, 0xa0, 0x73, 0x6d, 0xb7,
	0xb4, 0x3f, 0x76, 0xfa, 0xa0, 0x13, 0x07, 0x50, 0x13, 0x92, 0x4a, 0xa6, 0xc7, 0x69, 0x0f, 0x7a,
	0x7b, 0x9f, 0xec, 0x05, 0x6a, 0x64, 0xa9, 0x69, 0xc5, 0x2e, 0x34, 0x8a, 0x92, 0x2f, 0xf4, 0xe6,
	0xd4, 0xb0, 0xb5, 0x74, 0xc7, 0x6a, 0x41, 0x2b, 0x26, 0x04, 0x5d, 0xb0, 0xa0, 0x66, 0x56, 0x67,
	0x51, 0x29, 0x3a, 0x99, 0x65, 0x81, 0x6b, 0x56, 0x67, 0x11, 0x7b, 0xd0, 0x9c, 0xf3, 0x55, 0xb1,
	0x64, 0x4a, 0xab, 0x6b, 0xed, 0x5f, 0x41, 0xa5, 0x65, 0x9b, 0x52, 0x5f, 0x32, 0x68, 0x68, 0x71,
	0xc7, 0xe1, 0x4f, 0x07, 0x6a, 0xfa, 0x69, 0x78, 0x0a, 0xf5, 0x38, 0x25, 0xd1, 0x94, 0x0c, 0xfd,
	0x47, 0x08, 0xe0, 0xbe, 0x99, 0x91, 0x19, 0x19, 0xfa, 0x8e, 0x12, 0x26, 0xd3, 0x28, 0x55, 0x42,
	0x05, 0xdb, 0x00, 0xe4, 0xdd, 0x34, 0x8d, 0xe2, 0xe9, 0x68, 0xfc, 0xca, 0xaf, 0x2a, 0x8e, 0xef,
	0xc7, 0x63, 0x62, 0xf8, 0x04, 0x3d, 0x68, 0xdc, 0x24, 0xd1, 0xe4, 0x56, 0x51, 0x4d, 0xa9, 0x6f,
	0xa3, 0x64, 0x34, 0x8c, 0xb4, 0xea, 0x62, 0x0b, 0x9a, 0xf1, 0xfd, 0xdd, 0xeb, 0x84, 0x28, 0xb3,
	0xba, 0xc6, 0x68, 0x1c, 0x93, 0x24, 0x21, 0x43, 0xbf, 0x81, 0x4f, 0xe1, 0xb1, 0xf5, 0x7e, 0x7f,
	0x13, 0x8d, 0x92, 0x59, 0x4a, 0xfc, 0xa6, 0x2a, 0xda, 0x80, 0x5d, 0x11, 0xf0, 0x09, 0xb4, 0x74,
	0xca, 0xae, 0x74, 0x8a, 0x1d, 0xf0, 0x6d, 0x14, 0xd9, 0x55, 0xbd, 0xc1, 0x6f, 0x07, 0xdc, 0x99,
	0xbe, 0x07, 0xde, 0xea, 0x41, 0x4b, 0x89, 0x47, 0x8f, 0xda, 0x3d, 0x3b, 0xa8, 0x9a, 0xb3, 0x85,
	0xf8, 0xed, 0xc7, 0xaf, 0xef, 0x15, 0x2f, 0xac, 0xf7, 0x8d, 0xfc, 0xd2, 0xb9, 0xc2, 0x3b, 0x70,
	0xcd, 0x71, 0xf1, 0xec, 0xf0, 0xd8, 0xc6, 0xeb, 0xfc, 0xf8, 0x6f, 0x20, 0xec, 0x68, 0xb3, 0x36,
	0x7a, 0xd6, 0xac, 0xff, 0x25, 0xcf, 0xbe, 0x2a, 0xbb, 0x98, 0xae, 0xe7, 0x6c, 0xf9, 0x9f, 0x76,
	0x57, 0x0f, 0xec, 0x3e, 0xb8, 0xfa, 0xbf, 0xfa, 0xfc, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5d,
	0x5d, 0x6f, 0xef, 0xe1, 0x03, 0x00, 0x00,
}
