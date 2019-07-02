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

type GCDRequest struct {
	Attributes           []byte   `protobuf:"bytes,1,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GCDRequest) Reset()         { *m = GCDRequest{} }
func (m *GCDRequest) String() string { return proto.CompactTextString(m) }
func (*GCDRequest) ProtoMessage()    {}
func (*GCDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_178e0e204cde370a, []int{0}
}

func (m *GCDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GCDRequest.Unmarshal(m, b)
}
func (m *GCDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GCDRequest.Marshal(b, m, deterministic)
}
func (m *GCDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GCDRequest.Merge(m, src)
}
func (m *GCDRequest) XXX_Size() int {
	return xxx_messageInfo_GCDRequest.Size(m)
}
func (m *GCDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GCDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GCDRequest proto.InternalMessageInfo

func (m *GCDRequest) GetAttributes() []byte {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type GCDResponse struct {
	Result               []byte   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GCDResponse) Reset()         { *m = GCDResponse{} }
func (m *GCDResponse) String() string { return proto.CompactTextString(m) }
func (*GCDResponse) ProtoMessage()    {}
func (*GCDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_178e0e204cde370a, []int{1}
}

func (m *GCDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GCDResponse.Unmarshal(m, b)
}
func (m *GCDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GCDResponse.Marshal(b, m, deterministic)
}
func (m *GCDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GCDResponse.Merge(m, src)
}
func (m *GCDResponse) XXX_Size() int {
	return xxx_messageInfo_GCDResponse.Size(m)
}
func (m *GCDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GCDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GCDResponse proto.InternalMessageInfo

func (m *GCDResponse) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*GCDRequest)(nil), "pb.GCDRequest")
	proto.RegisterType((*GCDResponse)(nil), "pb.GCDResponse")
}

func init() { proto.RegisterFile("gcd.proto", fileDescriptor_178e0e204cde370a) }

var fileDescriptor_178e0e204cde370a = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x4f, 0x4e, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe1, 0xe2, 0x72, 0x77, 0x76,
	0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe3, 0xe2, 0x4a, 0x2c, 0x29, 0x29, 0xca,
	0x4c, 0x2a, 0x2d, 0x49, 0x2d, 0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x42, 0x12, 0x51, 0x52,
	0xe5, 0xe2, 0x06, 0xab, 0x2e, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12, 0xe3, 0x62, 0x2b, 0x4a,
	0x2d, 0x2e, 0xcd, 0x29, 0x81, 0x2a, 0x85, 0xf2, 0x8c, 0xac, 0xc0, 0x86, 0x06, 0xa7, 0x16, 0x95,
	0x65, 0x26, 0xa7, 0x0a, 0xe9, 0x70, 0xb1, 0x3b, 0xe7, 0xe7, 0x16, 0x94, 0x96, 0xa4, 0x0a, 0xf1,
	0xe9, 0x15, 0x24, 0xe9, 0x21, 0xec, 0x93, 0xe2, 0x87, 0xf3, 0x21, 0x26, 0x2a, 0x31, 0x24, 0xb1,
	0x81, 0xdd, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x12, 0x62, 0xbd, 0x3c, 0xa8, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GCDServiceClient is the client API for GCDService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GCDServiceClient interface {
	Compute(ctx context.Context, in *GCDRequest, opts ...grpc.CallOption) (*GCDResponse, error)
}

type gCDServiceClient struct {
	cc *grpc.ClientConn
}

func NewGCDServiceClient(cc *grpc.ClientConn) GCDServiceClient {
	return &gCDServiceClient{cc}
}

func (c *gCDServiceClient) Compute(ctx context.Context, in *GCDRequest, opts ...grpc.CallOption) (*GCDResponse, error) {
	out := new(GCDResponse)
	err := c.cc.Invoke(ctx, "/pb.GCDService/Compute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GCDServiceServer is the server API for GCDService service.
type GCDServiceServer interface {
	Compute(context.Context, *GCDRequest) (*GCDResponse, error)
}

// UnimplementedGCDServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGCDServiceServer struct {
}

func (*UnimplementedGCDServiceServer) Compute(ctx context.Context, req *GCDRequest) (*GCDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Compute not implemented")
}

func RegisterGCDServiceServer(s *grpc.Server, srv GCDServiceServer) {
	s.RegisterService(&_GCDService_serviceDesc, srv)
}

func _GCDService_Compute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GCDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GCDServiceServer).Compute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GCDService/Compute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GCDServiceServer).Compute(ctx, req.(*GCDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GCDService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.GCDService",
	HandlerType: (*GCDServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Compute",
			Handler:    _GCDService_Compute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gcd.proto",
}
