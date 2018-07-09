// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ocr/ocr.proto

package ocr

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

type OcrResult struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OcrResult) Reset()         { *m = OcrResult{} }
func (m *OcrResult) String() string { return proto.CompactTextString(m) }
func (*OcrResult) ProtoMessage()    {}
func (*OcrResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ocr_d5501d46cbae3d5f, []int{0}
}
func (m *OcrResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OcrResult.Unmarshal(m, b)
}
func (m *OcrResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OcrResult.Marshal(b, m, deterministic)
}
func (dst *OcrResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OcrResult.Merge(dst, src)
}
func (m *OcrResult) XXX_Size() int {
	return xxx_messageInfo_OcrResult.Size(m)
}
func (m *OcrResult) XXX_DiscardUnknown() {
	xxx_messageInfo_OcrResult.DiscardUnknown(m)
}

var xxx_messageInfo_OcrResult proto.InternalMessageInfo

func (m *OcrResult) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *OcrResult) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *OcrResult) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type File struct {
	Bytes                []byte   `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_ocr_d5501d46cbae3d5f, []int{1}
}
func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (dst *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(dst, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

func init() {
	proto.RegisterType((*OcrResult)(nil), "ocr.OcrResult")
	proto.RegisterType((*File)(nil), "ocr.File")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OcrServiceClient is the client API for OcrService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OcrServiceClient interface {
	GetResult(ctx context.Context, in *File, opts ...grpc.CallOption) (*OcrResult, error)
}

type ocrServiceClient struct {
	cc *grpc.ClientConn
}

func NewOcrServiceClient(cc *grpc.ClientConn) OcrServiceClient {
	return &ocrServiceClient{cc}
}

func (c *ocrServiceClient) GetResult(ctx context.Context, in *File, opts ...grpc.CallOption) (*OcrResult, error) {
	out := new(OcrResult)
	err := c.cc.Invoke(ctx, "/ocr.OcrService/GetResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OcrServiceServer is the server API for OcrService service.
type OcrServiceServer interface {
	GetResult(context.Context, *File) (*OcrResult, error)
}

func RegisterOcrServiceServer(s *grpc.Server, srv OcrServiceServer) {
	s.RegisterService(&_OcrService_serviceDesc, srv)
}

func _OcrService_GetResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(File)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcrServiceServer).GetResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocr.OcrService/GetResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcrServiceServer).GetResult(ctx, req.(*File))
	}
	return interceptor(ctx, in, info, handler)
}

var _OcrService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ocr.OcrService",
	HandlerType: (*OcrServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetResult",
			Handler:    _OcrService_GetResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ocr/ocr.proto",
}

func init() { proto.RegisterFile("ocr/ocr.proto", fileDescriptor_ocr_d5501d46cbae3d5f) }

var fileDescriptor_ocr_d5501d46cbae3d5f = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4f, 0x2e, 0xd2,
	0xcf, 0x4f, 0x2e, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x4f, 0x2e, 0x52, 0xf2,
	0xe5, 0xe2, 0xf4, 0x4f, 0x2e, 0x0a, 0x4a, 0x2d, 0x2e, 0xcd, 0x29, 0x11, 0x12, 0xe2, 0x62, 0x49,
	0xce, 0x4f, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x02, 0xb3, 0x85, 0x24, 0xb8, 0xd8,
	0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c,
	0x90, 0xea, 0x94, 0xc4, 0x92, 0x44, 0x09, 0x66, 0xb0, 0x30, 0x98, 0xad, 0x24, 0xc3, 0xc5, 0xe2,
	0x96, 0x99, 0x93, 0x2a, 0x24, 0xc2, 0xc5, 0x9a, 0x54, 0x59, 0x92, 0x5a, 0x0c, 0x36, 0x8a, 0x27,
	0x08, 0xc2, 0x31, 0x32, 0xe3, 0xe2, 0xf2, 0x4f, 0x2e, 0x0a, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e,
	0x15, 0xd2, 0xe0, 0xe2, 0x74, 0x4f, 0x2d, 0x81, 0x5a, 0xcd, 0xa9, 0x07, 0x72, 0x18, 0x48, 0xaf,
	0x14, 0x1f, 0x98, 0x09, 0x77, 0x95, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0xc1, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xc2, 0xea, 0x19, 0x5e, 0xc1, 0x00, 0x00, 0x00,
}