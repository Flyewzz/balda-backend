// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dict.proto

package go_micro_srv_consignment

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Availability struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Availability) Reset()         { *m = Availability{} }
func (m *Availability) String() string { return proto.CompactTextString(m) }
func (*Availability) ProtoMessage()    {}
func (*Availability) Descriptor() ([]byte, []int) {
	return fileDescriptor_67812e90854f6714, []int{0}
}

func (m *Availability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Availability.Unmarshal(m, b)
}
func (m *Availability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Availability.Marshal(b, m, deterministic)
}
func (m *Availability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Availability.Merge(m, src)
}
func (m *Availability) XXX_Size() int {
	return xxx_messageInfo_Availability.Size(m)
}
func (m *Availability) XXX_DiscardUnknown() {
	xxx_messageInfo_Availability.DiscardUnknown(m)
}

var xxx_messageInfo_Availability proto.InternalMessageInfo

func (m *Availability) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type Word struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Word) Reset()         { *m = Word{} }
func (m *Word) String() string { return proto.CompactTextString(m) }
func (*Word) ProtoMessage()    {}
func (*Word) Descriptor() ([]byte, []int) {
	return fileDescriptor_67812e90854f6714, []int{1}
}

func (m *Word) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Word.Unmarshal(m, b)
}
func (m *Word) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Word.Marshal(b, m, deterministic)
}
func (m *Word) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Word.Merge(m, src)
}
func (m *Word) XXX_Size() int {
	return xxx_messageInfo_Word.Size(m)
}
func (m *Word) XXX_DiscardUnknown() {
	xxx_messageInfo_Word.DiscardUnknown(m)
}

var xxx_messageInfo_Word proto.InternalMessageInfo

func (m *Word) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type Nothing struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nothing) Reset()         { *m = Nothing{} }
func (m *Nothing) String() string { return proto.CompactTextString(m) }
func (*Nothing) ProtoMessage()    {}
func (*Nothing) Descriptor() ([]byte, []int) {
	return fileDescriptor_67812e90854f6714, []int{2}
}

func (m *Nothing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nothing.Unmarshal(m, b)
}
func (m *Nothing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nothing.Marshal(b, m, deterministic)
}
func (m *Nothing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nothing.Merge(m, src)
}
func (m *Nothing) XXX_Size() int {
	return xxx_messageInfo_Nothing.Size(m)
}
func (m *Nothing) XXX_DiscardUnknown() {
	xxx_messageInfo_Nothing.DiscardUnknown(m)
}

var xxx_messageInfo_Nothing proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Availability)(nil), "go.micro.srv.consignment.Availability")
	proto.RegisterType((*Word)(nil), "go.micro.srv.consignment.Word")
	proto.RegisterType((*Nothing)(nil), "go.micro.srv.consignment.Nothing")
}

func init() { proto.RegisterFile("dict.proto", fileDescriptor_67812e90854f6714) }

var fileDescriptor_67812e90854f6714 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xc9, 0x4c, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x48, 0xcf, 0xd7, 0xcb, 0xcd, 0x4c, 0x2e, 0xca,
	0xd7, 0x2b, 0x2e, 0x2a, 0xd3, 0x4b, 0xce, 0xcf, 0x2b, 0xce, 0x4c, 0xcf, 0xcb, 0x4d, 0xcd, 0x2b,
	0x51, 0x52, 0xe3, 0xe2, 0x71, 0x2c, 0x4b, 0xcc, 0xcc, 0x49, 0x4c, 0xca, 0xcc, 0xc9, 0x2c, 0xa9,
	0x14, 0x12, 0xe3, 0x62, 0x2b, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x08, 0x82, 0xf2, 0x94, 0x64, 0xb8, 0x58, 0xc2, 0xf3, 0x8b, 0x52, 0x84, 0x44, 0xb8, 0x58, 0x4b,
	0x32, 0x4b, 0x72, 0x52, 0xc1, 0xd2, 0x9c, 0x41, 0x10, 0x8e, 0x12, 0x27, 0x17, 0xbb, 0x5f, 0x7e,
	0x49, 0x46, 0x66, 0x5e, 0xba, 0x51, 0x03, 0x13, 0x17, 0x97, 0x4b, 0x66, 0x72, 0x49, 0x66, 0x7e,
	0x5e, 0x62, 0x51, 0xa5, 0x50, 0x30, 0x17, 0xa7, 0x73, 0x46, 0x6a, 0x72, 0x36, 0x58, 0xb3, 0x9c,
	0x1e, 0x2e, 0x77, 0xe8, 0x81, 0xe4, 0xa5, 0xd4, 0x70, 0xcb, 0xa3, 0x38, 0xd2, 0x87, 0x8b, 0xdd,
	0x31, 0x25, 0x85, 0x28, 0x23, 0x15, 0x71, 0xcb, 0x43, 0x5d, 0x2c, 0xe4, 0xcf, 0xc5, 0x15, 0x94,
	0x9a, 0x9b, 0x5f, 0x96, 0x4a, 0x25, 0x03, 0x93, 0xd8, 0xc0, 0x81, 0x6e, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0xc7, 0x49, 0x48, 0x23, 0x82, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DictionaryClient is the client API for Dictionary service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DictionaryClient interface {
	CheckWord(ctx context.Context, in *Word, opts ...grpc.CallOption) (*Availability, error)
	AddWord(ctx context.Context, in *Word, opts ...grpc.CallOption) (*Nothing, error)
	RemoveWord(ctx context.Context, in *Word, opts ...grpc.CallOption) (*Nothing, error)
}

type dictionaryClient struct {
	cc *grpc.ClientConn
}

func NewDictionaryClient(cc *grpc.ClientConn) DictionaryClient {
	return &dictionaryClient{cc}
}

func (c *dictionaryClient) CheckWord(ctx context.Context, in *Word, opts ...grpc.CallOption) (*Availability, error) {
	out := new(Availability)
	err := c.cc.Invoke(ctx, "/go.micro.srv.consignment.Dictionary/CheckWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionaryClient) AddWord(ctx context.Context, in *Word, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/go.micro.srv.consignment.Dictionary/AddWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionaryClient) RemoveWord(ctx context.Context, in *Word, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/go.micro.srv.consignment.Dictionary/RemoveWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DictionaryServer is the server API for Dictionary service.
type DictionaryServer interface {
	CheckWord(context.Context, *Word) (*Availability, error)
	AddWord(context.Context, *Word) (*Nothing, error)
	RemoveWord(context.Context, *Word) (*Nothing, error)
}

func RegisterDictionaryServer(s *grpc.Server, srv DictionaryServer) {
	s.RegisterService(&_Dictionary_serviceDesc, srv)
}

func _Dictionary_CheckWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Word)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictionaryServer).CheckWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.consignment.Dictionary/CheckWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictionaryServer).CheckWord(ctx, req.(*Word))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dictionary_AddWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Word)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictionaryServer).AddWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.consignment.Dictionary/AddWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictionaryServer).AddWord(ctx, req.(*Word))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dictionary_RemoveWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Word)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictionaryServer).RemoveWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.consignment.Dictionary/RemoveWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictionaryServer).RemoveWord(ctx, req.(*Word))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dictionary_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.srv.consignment.Dictionary",
	HandlerType: (*DictionaryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckWord",
			Handler:    _Dictionary_CheckWord_Handler,
		},
		{
			MethodName: "AddWord",
			Handler:    _Dictionary_AddWord_Handler,
		},
		{
			MethodName: "RemoveWord",
			Handler:    _Dictionary_RemoveWord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dict.proto",
}
