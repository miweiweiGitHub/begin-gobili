// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app/service/main/open/api/grpc/v1/api.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type AppIDReply struct {
	AppId                int64    `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppIDReply) Reset()         { *m = AppIDReply{} }
func (m *AppIDReply) String() string { return proto.CompactTextString(m) }
func (*AppIDReply) ProtoMessage()    {}
func (*AppIDReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_f5c7f1ae9578a514, []int{0}
}
func (m *AppIDReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AppIDReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AppIDReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *AppIDReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppIDReply.Merge(dst, src)
}
func (m *AppIDReply) XXX_Size() int {
	return m.Size()
}
func (m *AppIDReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AppIDReply.DiscardUnknown(m)
}

var xxx_messageInfo_AppIDReply proto.InternalMessageInfo

func (m *AppIDReply) GetAppId() int64 {
	if m != nil {
		return m.AppId
	}
	return 0
}

type AppIDReq struct {
	AppKey               string   `protobuf:"bytes,2,opt,name=app_key,json=appKey,proto3" json:"app_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppIDReq) Reset()         { *m = AppIDReq{} }
func (m *AppIDReq) String() string { return proto.CompactTextString(m) }
func (*AppIDReq) ProtoMessage()    {}
func (*AppIDReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_f5c7f1ae9578a514, []int{1}
}
func (m *AppIDReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AppIDReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AppIDReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *AppIDReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppIDReq.Merge(dst, src)
}
func (m *AppIDReq) XXX_Size() int {
	return m.Size()
}
func (m *AppIDReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AppIDReq.DiscardUnknown(m)
}

var xxx_messageInfo_AppIDReq proto.InternalMessageInfo

func (m *AppIDReq) GetAppKey() string {
	if m != nil {
		return m.AppKey
	}
	return ""
}

type CloseReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseReply) Reset()         { *m = CloseReply{} }
func (m *CloseReply) String() string { return proto.CompactTextString(m) }
func (*CloseReply) ProtoMessage()    {}
func (*CloseReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_f5c7f1ae9578a514, []int{2}
}
func (m *CloseReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CloseReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CloseReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *CloseReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseReply.Merge(dst, src)
}
func (m *CloseReply) XXX_Size() int {
	return m.Size()
}
func (m *CloseReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseReply.DiscardUnknown(m)
}

var xxx_messageInfo_CloseReply proto.InternalMessageInfo

type CloseReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseReq) Reset()         { *m = CloseReq{} }
func (m *CloseReq) String() string { return proto.CompactTextString(m) }
func (*CloseReq) ProtoMessage()    {}
func (*CloseReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_f5c7f1ae9578a514, []int{3}
}
func (m *CloseReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CloseReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CloseReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *CloseReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseReq.Merge(dst, src)
}
func (m *CloseReq) XXX_Size() int {
	return m.Size()
}
func (m *CloseReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseReq.DiscardUnknown(m)
}

var xxx_messageInfo_CloseReq proto.InternalMessageInfo

type PingReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingReply) Reset()         { *m = PingReply{} }
func (m *PingReply) String() string { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()    {}
func (*PingReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_f5c7f1ae9578a514, []int{4}
}
func (m *PingReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PingReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PingReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PingReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingReply.Merge(dst, src)
}
func (m *PingReply) XXX_Size() int {
	return m.Size()
}
func (m *PingReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PingReply.DiscardUnknown(m)
}

var xxx_messageInfo_PingReply proto.InternalMessageInfo

type PingReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingReq) Reset()         { *m = PingReq{} }
func (m *PingReq) String() string { return proto.CompactTextString(m) }
func (*PingReq) ProtoMessage()    {}
func (*PingReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_f5c7f1ae9578a514, []int{5}
}
func (m *PingReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PingReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PingReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PingReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingReq.Merge(dst, src)
}
func (m *PingReq) XXX_Size() int {
	return m.Size()
}
func (m *PingReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PingReq.DiscardUnknown(m)
}

var xxx_messageInfo_PingReq proto.InternalMessageInfo

type SecretReply struct {
	Res                  string   `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecretReply) Reset()         { *m = SecretReply{} }
func (m *SecretReply) String() string { return proto.CompactTextString(m) }
func (*SecretReply) ProtoMessage()    {}
func (*SecretReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_f5c7f1ae9578a514, []int{6}
}
func (m *SecretReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SecretReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SecretReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SecretReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecretReply.Merge(dst, src)
}
func (m *SecretReply) XXX_Size() int {
	return m.Size()
}
func (m *SecretReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SecretReply.DiscardUnknown(m)
}

var xxx_messageInfo_SecretReply proto.InternalMessageInfo

func (m *SecretReply) GetRes() string {
	if m != nil {
		return m.Res
	}
	return ""
}

type SecretReq struct {
	SappKey              string   `protobuf:"bytes,2,opt,name=sapp_key,json=sappKey,proto3" json:"sapp_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecretReq) Reset()         { *m = SecretReq{} }
func (m *SecretReq) String() string { return proto.CompactTextString(m) }
func (*SecretReq) ProtoMessage()    {}
func (*SecretReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_f5c7f1ae9578a514, []int{7}
}
func (m *SecretReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SecretReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SecretReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SecretReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecretReq.Merge(dst, src)
}
func (m *SecretReq) XXX_Size() int {
	return m.Size()
}
func (m *SecretReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SecretReq.DiscardUnknown(m)
}

var xxx_messageInfo_SecretReq proto.InternalMessageInfo

func (m *SecretReq) GetSappKey() string {
	if m != nil {
		return m.SappKey
	}
	return ""
}

func init() {
	proto.RegisterType((*AppIDReply)(nil), "manager.service.open.AppIDReply")
	proto.RegisterType((*AppIDReq)(nil), "manager.service.open.AppIDReq")
	proto.RegisterType((*CloseReply)(nil), "manager.service.open.CloseReply")
	proto.RegisterType((*CloseReq)(nil), "manager.service.open.CloseReq")
	proto.RegisterType((*PingReply)(nil), "manager.service.open.PingReply")
	proto.RegisterType((*PingReq)(nil), "manager.service.open.PingReq")
	proto.RegisterType((*SecretReply)(nil), "manager.service.open.SecretReply")
	proto.RegisterType((*SecretReq)(nil), "manager.service.open.SecretReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OpenClient is the client API for Open service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OpenClient interface {
	// Ping check dao health.
	Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingReply, error)
	// Close close all dao.
	Close(ctx context.Context, in *CloseReq, opts ...grpc.CallOption) (*CloseReply, error)
	// Secret .
	Secret(ctx context.Context, in *SecretReq, opts ...grpc.CallOption) (*SecretReply, error)
	// AppID .
	AppID(ctx context.Context, in *AppIDReq, opts ...grpc.CallOption) (*AppIDReply, error)
}

type openClient struct {
	cc *grpc.ClientConn
}

func NewOpenClient(cc *grpc.ClientConn) OpenClient {
	return &openClient{cc}
}

func (c *openClient) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/manager.service.open.Open/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openClient) Close(ctx context.Context, in *CloseReq, opts ...grpc.CallOption) (*CloseReply, error) {
	out := new(CloseReply)
	err := c.cc.Invoke(ctx, "/manager.service.open.Open/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openClient) Secret(ctx context.Context, in *SecretReq, opts ...grpc.CallOption) (*SecretReply, error) {
	out := new(SecretReply)
	err := c.cc.Invoke(ctx, "/manager.service.open.Open/Secret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openClient) AppID(ctx context.Context, in *AppIDReq, opts ...grpc.CallOption) (*AppIDReply, error) {
	out := new(AppIDReply)
	err := c.cc.Invoke(ctx, "/manager.service.open.Open/AppID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpenServer is the server API for Open service.
type OpenServer interface {
	// Ping check dao health.
	Ping(context.Context, *PingReq) (*PingReply, error)
	// Close close all dao.
	Close(context.Context, *CloseReq) (*CloseReply, error)
	// Secret .
	Secret(context.Context, *SecretReq) (*SecretReply, error)
	// AppID .
	AppID(context.Context, *AppIDReq) (*AppIDReply, error)
}

func RegisterOpenServer(s *grpc.Server, srv OpenServer) {
	s.RegisterService(&_Open_serviceDesc, srv)
}

func _Open_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.service.open.Open/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenServer).Ping(ctx, req.(*PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Open_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.service.open.Open/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenServer).Close(ctx, req.(*CloseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Open_Secret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecretReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenServer).Secret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.service.open.Open/Secret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenServer).Secret(ctx, req.(*SecretReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Open_AppID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenServer).AppID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.service.open.Open/AppID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenServer).AppID(ctx, req.(*AppIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Open_serviceDesc = grpc.ServiceDesc{
	ServiceName: "manager.service.open.Open",
	HandlerType: (*OpenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Open_Ping_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _Open_Close_Handler,
		},
		{
			MethodName: "Secret",
			Handler:    _Open_Secret_Handler,
		},
		{
			MethodName: "AppID",
			Handler:    _Open_AppID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/service/main/open/api/grpc/v1/api.proto",
}

func (m *AppIDReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AppIDReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.AppId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.AppId))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *AppIDReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AppIDReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.AppKey) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintApi(dAtA, i, uint64(len(m.AppKey)))
		i += copy(dAtA[i:], m.AppKey)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *CloseReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CloseReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *CloseReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CloseReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PingReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PingReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PingReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PingReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SecretReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SecretReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Res) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintApi(dAtA, i, uint64(len(m.Res)))
		i += copy(dAtA[i:], m.Res)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SecretReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SecretReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.SappKey) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintApi(dAtA, i, uint64(len(m.SappKey)))
		i += copy(dAtA[i:], m.SappKey)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintApi(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AppIDReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AppId != 0 {
		n += 1 + sovApi(uint64(m.AppId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AppIDReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AppKey)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CloseReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CloseReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PingReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PingReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SecretReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Res)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SecretReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SappKey)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovApi(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AppIDReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AppIDReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AppIDReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppId", wireType)
			}
			m.AppId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AppIDReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AppIDReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AppIDReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CloseReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CloseReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CloseReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CloseReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CloseReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CloseReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PingReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PingReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PingReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PingReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PingReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PingReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SecretReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SecretReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SecretReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Res", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Res = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SecretReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SecretReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SecretReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SappKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SappKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthApi
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowApi
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipApi(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthApi = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("app/service/main/open/api/grpc/v1/api.proto", fileDescriptor_api_f5c7f1ae9578a514)
}

var fileDescriptor_api_f5c7f1ae9578a514 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdb, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xd9, 0x9e, 0x33, 0xfd, 0x5f, 0x94, 0xe5, 0x2f, 0xd6, 0x82, 0x69, 0xdd, 0x82, 0x14,
	0x84, 0x2c, 0xd5, 0x27, 0xf0, 0x80, 0x50, 0x14, 0x94, 0x78, 0xe7, 0x8d, 0xac, 0xed, 0x50, 0x82,
	0x6d, 0x32, 0xd9, 0x94, 0x40, 0x5e, 0xcb, 0xa7, 0xf0, 0xd2, 0x47, 0x90, 0x3c, 0x89, 0x6c, 0x0e,
	0x8a, 0xd0, 0xe4, 0x6e, 0x87, 0xf9, 0x66, 0x7e, 0xdf, 0xc7, 0x2c, 0x9c, 0x29, 0x22, 0x19, 0xa1,
	0x8e, 0xbd, 0x25, 0xca, 0xad, 0xf2, 0x7c, 0x19, 0x10, 0xfa, 0x52, 0x91, 0x27, 0xd7, 0x9a, 0x96,
	0x32, 0x9e, 0x9b, 0xb7, 0x43, 0x3a, 0xd8, 0x05, 0xfc, 0xff, 0x56, 0xf9, 0x6a, 0x8d, 0xda, 0x29,
	0x06, 0x1c, 0xa3, 0x15, 0x53, 0x80, 0x4b, 0xa2, 0xc5, 0x8d, 0x8b, 0xb4, 0x49, 0xf8, 0x01, 0x74,
	0x14, 0xd1, 0x8b, 0xb7, 0x1a, 0xb2, 0x09, 0x9b, 0x35, 0xdd, 0xb6, 0x22, 0x5a, 0xac, 0xc4, 0x14,
	0x7a, 0x85, 0x28, 0xe4, 0x87, 0xd0, 0x35, 0x92, 0x37, 0x4c, 0x86, 0x8d, 0x09, 0x9b, 0x59, 0xae,
	0x99, 0xb8, 0xc3, 0x44, 0xfc, 0x03, 0xb8, 0xde, 0x04, 0x11, 0x66, 0x9b, 0x04, 0x40, 0xaf, 0xa8,
	0x42, 0xd1, 0x07, 0xeb, 0xd1, 0xf3, 0xd7, 0x79, 0xc3, 0x82, 0x6e, 0x5e, 0x84, 0x62, 0x0c, 0xfd,
	0x27, 0x5c, 0x6a, 0xdc, 0xe5, 0xf0, 0x01, 0x34, 0x35, 0x46, 0x19, 0xd9, 0x72, 0xcd, 0x53, 0x9c,
	0x82, 0x55, 0x0a, 0x42, 0x7e, 0x04, 0xbd, 0xe8, 0x2f, 0xb9, 0x1b, 0xe5, 0xe8, 0xf3, 0xf7, 0x06,
	0xb4, 0x1e, 0x08, 0x7d, 0x7e, 0x0b, 0x2d, 0xb3, 0x9c, 0x1f, 0x3b, 0xfb, 0xc2, 0x3a, 0x05, 0x78,
	0x34, 0xae, 0x6b, 0x1b, 0x2b, 0x0b, 0x68, 0x67, 0xee, 0xb9, 0xbd, 0x5f, 0x59, 0x46, 0x1b, 0x4d,
	0x6a, 0xfb, 0x66, 0xd5, 0x3d, 0x74, 0xf2, 0x0c, 0xbc, 0x82, 0xfa, 0x93, 0x70, 0x74, 0x52, 0x2f,
	0x28, 0x8c, 0x65, 0x97, 0xa8, 0x32, 0x56, 0x9e, 0xa9, 0xca, 0xd8, 0xef, 0xad, 0xaf, 0x06, 0x1f,
	0xa9, 0xcd, 0x3e, 0x53, 0x9b, 0x7d, 0xa5, 0x36, 0x7b, 0x6e, 0xc4, 0xf3, 0xd7, 0x4e, 0xf6, 0x51,
	0x2e, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x52, 0x90, 0x4e, 0x57, 0x02, 0x00, 0x00,
}
