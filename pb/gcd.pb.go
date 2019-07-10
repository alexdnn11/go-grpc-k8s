// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gcd.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type GenerateRequest struct {
	Attributes           []byte   `protobuf:"bytes,1,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateRequest) Reset()         { *m = GenerateRequest{} }
func (m *GenerateRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateRequest) ProtoMessage()    {}
func (*GenerateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_178e0e204cde370a, []int{0}
}

func (m *GenerateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateRequest.Unmarshal(m, b)
}
func (m *GenerateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateRequest.Marshal(b, m, deterministic)
}
func (m *GenerateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateRequest.Merge(m, src)
}
func (m *GenerateRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateRequest.Size(m)
}
func (m *GenerateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateRequest proto.InternalMessageInfo

func (m *GenerateRequest) GetAttributes() []byte {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type GenerateResponse struct {
	Result               []byte   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateResponse) Reset()         { *m = GenerateResponse{} }
func (m *GenerateResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateResponse) ProtoMessage()    {}
func (*GenerateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_178e0e204cde370a, []int{1}
}

func (m *GenerateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateResponse.Unmarshal(m, b)
}
func (m *GenerateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateResponse.Marshal(b, m, deterministic)
}
func (m *GenerateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateResponse.Merge(m, src)
}
func (m *GenerateResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateResponse.Size(m)
}
func (m *GenerateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateResponse proto.InternalMessageInfo

func (m *GenerateResponse) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

type VerifyRequest struct {
	Proof                []byte   `protobuf:"bytes,1,opt,name=proof,proto3" json:"proof,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyRequest) Reset()         { *m = VerifyRequest{} }
func (m *VerifyRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyRequest) ProtoMessage()    {}
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_178e0e204cde370a, []int{2}
}

func (m *VerifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyRequest.Unmarshal(m, b)
}
func (m *VerifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyRequest.Marshal(b, m, deterministic)
}
func (m *VerifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyRequest.Merge(m, src)
}
func (m *VerifyRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyRequest.Size(m)
}
func (m *VerifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyRequest proto.InternalMessageInfo

func (m *VerifyRequest) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

type VerifyResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyResponse) Reset()         { *m = VerifyResponse{} }
func (m *VerifyResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyResponse) ProtoMessage()    {}
func (*VerifyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_178e0e204cde370a, []int{3}
}

func (m *VerifyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyResponse.Unmarshal(m, b)
}
func (m *VerifyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyResponse.Marshal(b, m, deterministic)
}
func (m *VerifyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyResponse.Merge(m, src)
}
func (m *VerifyResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyResponse.Size(m)
}
func (m *VerifyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyResponse proto.InternalMessageInfo

func (m *VerifyResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func init() {
	proto.RegisterType((*GenerateRequest)(nil), "pb.GenerateRequest")
	proto.RegisterType((*GenerateResponse)(nil), "pb.GenerateResponse")
	proto.RegisterType((*VerifyRequest)(nil), "pb.VerifyRequest")
	proto.RegisterType((*VerifyResponse)(nil), "pb.VerifyResponse")
}

func init() { proto.RegisterFile("gcd.proto", fileDescriptor_178e0e204cde370a) }

var fileDescriptor_178e0e204cde370a = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x4f, 0x4e, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x32, 0xe4, 0xe2, 0x77, 0x4f, 0xcd,
	0x4b, 0x2d, 0x4a, 0x2c, 0x49, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe3, 0xe2,
	0x4a, 0x2c, 0x29, 0x29, 0xca, 0x4c, 0x2a, 0x2d, 0x49, 0x2d, 0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x09, 0x42, 0x12, 0x51, 0xd2, 0xe2, 0x12, 0x40, 0x68, 0x29, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15,
	0x12, 0xe3, 0x62, 0x2b, 0x4a, 0x2d, 0x2e, 0xcd, 0x29, 0x81, 0xaa, 0x87, 0xf2, 0x94, 0x54, 0xb9,
	0x78, 0xc3, 0x52, 0x8b, 0x32, 0xd3, 0x2a, 0x61, 0x86, 0x8b, 0x70, 0xb1, 0x16, 0x14, 0xe5, 0xe7,
	0xa7, 0x41, 0xd5, 0x41, 0x38, 0x4a, 0x1a, 0x5c, 0x7c, 0x30, 0x65, 0x58, 0x0d, 0xe4, 0x80, 0x19,
	0x68, 0x54, 0xca, 0xc5, 0x1e, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x64, 0xce, 0xc5, 0x01,
	0x73, 0x87, 0x90, 0xb0, 0x5e, 0x41, 0x92, 0x1e, 0x9a, 0x47, 0xa4, 0x44, 0x50, 0x05, 0x21, 0x26,
	0x2b, 0x31, 0x08, 0x19, 0x72, 0xb1, 0x41, 0x6c, 0x13, 0x12, 0x04, 0xa9, 0x40, 0x71, 0xa0, 0x94,
	0x10, 0xb2, 0x10, 0x4c, 0x4b, 0x12, 0x1b, 0x38, 0xc4, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xe4, 0xa3, 0x14, 0x5a, 0x3e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error)
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error) {
	out := new(GenerateResponse)
	err := c.cc.Invoke(ctx, "/pb.Service/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error) {
	out := new(VerifyResponse)
	err := c.cc.Invoke(ctx, "/pb.Service/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	Generate(context.Context, *GenerateRequest) (*GenerateResponse, error)
	Verify(context.Context, *VerifyRequest) (*VerifyResponse, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Generate(ctx context.Context, req *GenerateRequest) (*GenerateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (*UnimplementedServiceServer) Verify(ctx context.Context, req *VerifyRequest) (*VerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Service/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Generate(ctx, req.(*GenerateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Service/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _Service_Generate_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _Service_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gcd.proto",
}
