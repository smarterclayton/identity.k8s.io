// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/apis/workload/workload.proto

/*
Package pkg_apis_workload is a generated protocol buffer package.

It is generated from these files:
	pkg/apis/workload/workload.proto

It has these top-level messages:
	GetTokenRequest
	GetTokenResponse
	ValidateTokenRequest
	ValidateTokenResponse
*/
package pkg_apis_workload

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetTokenRequest struct {
	Audience []string `protobuf:"bytes,1,rep,name=audience" json:"audience,omitempty"`
}

func (m *GetTokenRequest) Reset()                    { *m = GetTokenRequest{} }
func (m *GetTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*GetTokenRequest) ProtoMessage()               {}
func (*GetTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetTokenRequest) GetAudience() []string {
	if m != nil {
		return m.Audience
	}
	return nil
}

type GetTokenResponse struct {
	Token      []byte `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Expiration int64  `protobuf:"varint,2,opt,name=expiration" json:"expiration,omitempty"`
}

func (m *GetTokenResponse) Reset()                    { *m = GetTokenResponse{} }
func (m *GetTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*GetTokenResponse) ProtoMessage()               {}
func (*GetTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetTokenResponse) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *GetTokenResponse) GetExpiration() int64 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

type ValidateTokenRequest struct {
	Audience []string `protobuf:"bytes,1,rep,name=audience" json:"audience,omitempty"`
}

func (m *ValidateTokenRequest) Reset()                    { *m = ValidateTokenRequest{} }
func (m *ValidateTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*ValidateTokenRequest) ProtoMessage()               {}
func (*ValidateTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ValidateTokenRequest) GetAudience() []string {
	if m != nil {
		return m.Audience
	}
	return nil
}

type ValidateTokenResponse struct {
	Valid bool `protobuf:"varint,1,opt,name=valid" json:"valid,omitempty"`
}

func (m *ValidateTokenResponse) Reset()                    { *m = ValidateTokenResponse{} }
func (m *ValidateTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*ValidateTokenResponse) ProtoMessage()               {}
func (*ValidateTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ValidateTokenResponse) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func init() {
	proto.RegisterType((*GetTokenRequest)(nil), "pkg.apis.workload.GetTokenRequest")
	proto.RegisterType((*GetTokenResponse)(nil), "pkg.apis.workload.GetTokenResponse")
	proto.RegisterType((*ValidateTokenRequest)(nil), "pkg.apis.workload.ValidateTokenRequest")
	proto.RegisterType((*ValidateTokenResponse)(nil), "pkg.apis.workload.ValidateTokenResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Workload service

type WorkloadClient interface {
	GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error)
	ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error)
}

type workloadClient struct {
	cc *grpc.ClientConn
}

func NewWorkloadClient(cc *grpc.ClientConn) WorkloadClient {
	return &workloadClient{cc}
}

func (c *workloadClient) GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error) {
	out := new(GetTokenResponse)
	err := grpc.Invoke(ctx, "/pkg.apis.workload.Workload/GetToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workloadClient) ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error) {
	out := new(ValidateTokenResponse)
	err := grpc.Invoke(ctx, "/pkg.apis.workload.Workload/ValidateToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Workload service

type WorkloadServer interface {
	GetToken(context.Context, *GetTokenRequest) (*GetTokenResponse, error)
	ValidateToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error)
}

func RegisterWorkloadServer(s *grpc.Server, srv WorkloadServer) {
	s.RegisterService(&_Workload_serviceDesc, srv)
}

func _Workload_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkloadServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.apis.workload.Workload/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkloadServer).GetToken(ctx, req.(*GetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workload_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkloadServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.apis.workload.Workload/ValidateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkloadServer).ValidateToken(ctx, req.(*ValidateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Workload_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pkg.apis.workload.Workload",
	HandlerType: (*WorkloadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetToken",
			Handler:    _Workload_GetToken_Handler,
		},
		{
			MethodName: "ValidateToken",
			Handler:    _Workload_ValidateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/apis/workload/workload.proto",
}

func init() { proto.RegisterFile("pkg/apis/workload/workload.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0xc8, 0x4e, 0xd7,
	0x4f, 0x2c, 0xc8, 0x2c, 0xd6, 0x2f, 0xcf, 0x2f, 0xca, 0xce, 0xc9, 0x4f, 0x4c, 0x81, 0x33, 0xf4,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0x0b, 0xb2, 0xd3, 0xf5, 0x40, 0x2a, 0xf4, 0x60, 0x12,
	0x4a, 0xba, 0x5c, 0xfc, 0xee, 0xa9, 0x25, 0x21, 0xf9, 0xd9, 0xa9, 0x79, 0x41, 0xa9, 0x85, 0xa5,
	0xa9, 0xc5, 0x25, 0x42, 0x52, 0x5c, 0x1c, 0x89, 0xa5, 0x29, 0x99, 0xa9, 0x79, 0xc9, 0xa9, 0x12,
	0x8c, 0x0a, 0xcc, 0x1a, 0x9c, 0x41, 0x70, 0xbe, 0x92, 0x07, 0x97, 0x00, 0x42, 0x79, 0x71, 0x41,
	0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x08, 0x17, 0x6b, 0x09, 0x48, 0x40, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x27, 0x08, 0xc2, 0x11, 0x92, 0xe3, 0xe2, 0x4a, 0xad, 0x28, 0xc8, 0x2c, 0x4a, 0x2c, 0xc9, 0xcc,
	0xcf, 0x93, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x0e, 0x42, 0x12, 0x51, 0x32, 0xe2, 0x12, 0x09, 0x4b,
	0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0x25, 0xda, 0x76, 0x5d, 0x2e, 0x51, 0x34, 0x3d, 0x08, 0x27,
	0x94, 0x81, 0x24, 0xc0, 0x4e, 0xe0, 0x08, 0x82, 0x70, 0x8c, 0x8e, 0x33, 0x72, 0x71, 0x84, 0x43,
	0x3d, 0x2a, 0x14, 0xca, 0xc5, 0x01, 0x73, 0xb9, 0x90, 0x92, 0x1e, 0x46, 0x40, 0xe8, 0xa1, 0x85,
	0x82, 0x94, 0x32, 0x5e, 0x35, 0x10, 0x7b, 0x95, 0x18, 0x84, 0x52, 0xb8, 0x78, 0x51, 0x9c, 0x24,
	0xa4, 0x8e, 0x45, 0x1f, 0x36, 0x8f, 0x4a, 0x69, 0x10, 0x56, 0x08, 0xb3, 0x25, 0x89, 0x0d, 0x1c,
	0x7f, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x26, 0xa5, 0x0f, 0xe3, 0x01, 0x00, 0x00,
}