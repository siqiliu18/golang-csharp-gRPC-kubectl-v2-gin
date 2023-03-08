// Code generated by protoc-gen-go. DO NOT EDIT.
// source: merge-proto/template.proto

package templates

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

// The request message
type TemplateRequest struct {
	Templatename         string   `protobuf:"bytes,1,opt,name=templatename,proto3" json:"templatename,omitempty"`
	Templatetype         string   `protobuf:"bytes,2,opt,name=templatetype,proto3" json:"templatetype,omitempty"`
	Cbedna               string   `protobuf:"bytes,3,opt,name=cbedna,proto3" json:"cbedna,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TemplateRequest) Reset()         { *m = TemplateRequest{} }
func (m *TemplateRequest) String() string { return proto.CompactTextString(m) }
func (*TemplateRequest) ProtoMessage()    {}
func (*TemplateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b49f1d056affe7f6, []int{0}
}

func (m *TemplateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemplateRequest.Unmarshal(m, b)
}
func (m *TemplateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemplateRequest.Marshal(b, m, deterministic)
}
func (m *TemplateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemplateRequest.Merge(m, src)
}
func (m *TemplateRequest) XXX_Size() int {
	return xxx_messageInfo_TemplateRequest.Size(m)
}
func (m *TemplateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TemplateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TemplateRequest proto.InternalMessageInfo

func (m *TemplateRequest) GetTemplatename() string {
	if m != nil {
		return m.Templatename
	}
	return ""
}

func (m *TemplateRequest) GetTemplatetype() string {
	if m != nil {
		return m.Templatetype
	}
	return ""
}

func (m *TemplateRequest) GetCbedna() string {
	if m != nil {
		return m.Cbedna
	}
	return ""
}

// The response message
type TemplateResponse struct {
	Documenturi          string   `protobuf:"bytes,1,opt,name=documenturi,proto3" json:"documenturi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TemplateResponse) Reset()         { *m = TemplateResponse{} }
func (m *TemplateResponse) String() string { return proto.CompactTextString(m) }
func (*TemplateResponse) ProtoMessage()    {}
func (*TemplateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b49f1d056affe7f6, []int{1}
}

func (m *TemplateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemplateResponse.Unmarshal(m, b)
}
func (m *TemplateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemplateResponse.Marshal(b, m, deterministic)
}
func (m *TemplateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemplateResponse.Merge(m, src)
}
func (m *TemplateResponse) XXX_Size() int {
	return xxx_messageInfo_TemplateResponse.Size(m)
}
func (m *TemplateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TemplateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TemplateResponse proto.InternalMessageInfo

func (m *TemplateResponse) GetDocumenturi() string {
	if m != nil {
		return m.Documenturi
	}
	return ""
}

func init() {
	proto.RegisterType((*TemplateRequest)(nil), "templates.TemplateRequest")
	proto.RegisterType((*TemplateResponse)(nil), "templates.TemplateResponse")
}

func init() {
	proto.RegisterFile("merge-proto/template.proto", fileDescriptor_b49f1d056affe7f6)
}

var fileDescriptor_b49f1d056affe7f6 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xca, 0x4d, 0x2d, 0x4a,
	0x4f, 0xd5, 0x2d, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x49, 0xcd, 0x2d, 0xc8, 0x49, 0x2c, 0x49,
	0xd5, 0x03, 0x73, 0x85, 0x38, 0x61, 0xfc, 0x62, 0xa5, 0x42, 0x2e, 0xfe, 0x10, 0x28, 0x27, 0x28,
	0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x89, 0x8b, 0x07, 0x26, 0x9f, 0x97, 0x98, 0x9b, 0x2a,
	0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x22, 0x86, 0xac, 0xa6, 0xa4, 0xb2, 0x20, 0x55, 0x82,
	0x09, 0x55, 0x0d, 0x48, 0x4c, 0x48, 0x8c, 0x8b, 0x2d, 0x39, 0x29, 0x35, 0x25, 0x2f, 0x51, 0x82,
	0x19, 0x2c, 0x0b, 0xe5, 0x29, 0x99, 0x70, 0x09, 0x20, 0xac, 0x2c, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e,
	0x15, 0x52, 0xe0, 0xe2, 0x4e, 0xc9, 0x4f, 0x2e, 0xcd, 0x4d, 0xcd, 0x2b, 0x29, 0x2d, 0xca, 0x84,
	0x5a, 0x89, 0x2c, 0x64, 0x14, 0xc3, 0xc5, 0x0b, 0xd3, 0xe5, 0x9b, 0x98, 0x9d, 0x5a, 0x24, 0xe4,
	0xcd, 0x25, 0x10, 0x94, 0x5a, 0x52, 0x94, 0x99, 0x5a, 0x96, 0x0a, 0x93, 0x10, 0x92, 0xd2, 0x83,
	0xfb, 0x4c, 0x0f, 0xcd, 0x5b, 0x52, 0xd2, 0x58, 0xe5, 0x20, 0xf6, 0x27, 0xb1, 0x81, 0x03, 0xc6,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xfd, 0xc2, 0x08, 0x36, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TemplateMakerClient is the client API for TemplateMaker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TemplateMakerClient interface {
	RetrieveTemplate(ctx context.Context, in *TemplateRequest, opts ...grpc.CallOption) (*TemplateResponse, error)
}

type templateMakerClient struct {
	cc grpc.ClientConnInterface
}

func NewTemplateMakerClient(cc grpc.ClientConnInterface) TemplateMakerClient {
	return &templateMakerClient{cc}
}

func (c *templateMakerClient) RetrieveTemplate(ctx context.Context, in *TemplateRequest, opts ...grpc.CallOption) (*TemplateResponse, error) {
	out := new(TemplateResponse)
	err := c.cc.Invoke(ctx, "/templates.TemplateMaker/RetrieveTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemplateMakerServer is the server API for TemplateMaker service.
type TemplateMakerServer interface {
	RetrieveTemplate(context.Context, *TemplateRequest) (*TemplateResponse, error)
}

// UnimplementedTemplateMakerServer can be embedded to have forward compatible implementations.
type UnimplementedTemplateMakerServer struct {
}

func (*UnimplementedTemplateMakerServer) RetrieveTemplate(ctx context.Context, req *TemplateRequest) (*TemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveTemplate not implemented")
}

func RegisterTemplateMakerServer(s *grpc.Server, srv TemplateMakerServer) {
	s.RegisterService(&_TemplateMaker_serviceDesc, srv)
}

func _TemplateMaker_RetrieveTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateMakerServer).RetrieveTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/templates.TemplateMaker/RetrieveTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateMakerServer).RetrieveTemplate(ctx, req.(*TemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TemplateMaker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "templates.TemplateMaker",
	HandlerType: (*TemplateMakerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RetrieveTemplate",
			Handler:    _TemplateMaker_RetrieveTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "merge-proto/template.proto",
}
