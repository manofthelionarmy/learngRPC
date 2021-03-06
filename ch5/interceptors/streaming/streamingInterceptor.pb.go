// Code generated by protoc-gen-go. DO NOT EDIT.
// source: streamingInterceptor.proto

package streamingInterceptor

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

type Pokemon struct {
	Id                   int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tag                  *PokemonTag `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Pokemon) Reset()         { *m = Pokemon{} }
func (m *Pokemon) String() string { return proto.CompactTextString(m) }
func (*Pokemon) ProtoMessage()    {}
func (*Pokemon) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab8084eb1a07bd8b, []int{0}
}

func (m *Pokemon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pokemon.Unmarshal(m, b)
}
func (m *Pokemon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pokemon.Marshal(b, m, deterministic)
}
func (m *Pokemon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pokemon.Merge(m, src)
}
func (m *Pokemon) XXX_Size() int {
	return xxx_messageInfo_Pokemon.Size(m)
}
func (m *Pokemon) XXX_DiscardUnknown() {
	xxx_messageInfo_Pokemon.DiscardUnknown(m)
}

var xxx_messageInfo_Pokemon proto.InternalMessageInfo

func (m *Pokemon) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Pokemon) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Pokemon) GetTag() *PokemonTag {
	if m != nil {
		return m.Tag
	}
	return nil
}

type PokemonId struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PokemonId) Reset()         { *m = PokemonId{} }
func (m *PokemonId) String() string { return proto.CompactTextString(m) }
func (*PokemonId) ProtoMessage()    {}
func (*PokemonId) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab8084eb1a07bd8b, []int{1}
}

func (m *PokemonId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PokemonId.Unmarshal(m, b)
}
func (m *PokemonId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PokemonId.Marshal(b, m, deterministic)
}
func (m *PokemonId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PokemonId.Merge(m, src)
}
func (m *PokemonId) XXX_Size() int {
	return xxx_messageInfo_PokemonId.Size(m)
}
func (m *PokemonId) XXX_DiscardUnknown() {
	xxx_messageInfo_PokemonId.DiscardUnknown(m)
}

var xxx_messageInfo_PokemonId proto.InternalMessageInfo

func (m *PokemonId) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type PokemonTag struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PokemonTag) Reset()         { *m = PokemonTag{} }
func (m *PokemonTag) String() string { return proto.CompactTextString(m) }
func (*PokemonTag) ProtoMessage()    {}
func (*PokemonTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab8084eb1a07bd8b, []int{2}
}

func (m *PokemonTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PokemonTag.Unmarshal(m, b)
}
func (m *PokemonTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PokemonTag.Marshal(b, m, deterministic)
}
func (m *PokemonTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PokemonTag.Merge(m, src)
}
func (m *PokemonTag) XXX_Size() int {
	return xxx_messageInfo_PokemonTag.Size(m)
}
func (m *PokemonTag) XXX_DiscardUnknown() {
	xxx_messageInfo_PokemonTag.DiscardUnknown(m)
}

var xxx_messageInfo_PokemonTag proto.InternalMessageInfo

func (m *PokemonTag) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Pokemon)(nil), "Pokemon")
	proto.RegisterType((*PokemonId)(nil), "PokemonId")
	proto.RegisterType((*PokemonTag)(nil), "PokemonTag")
}

func init() { proto.RegisterFile("streamingInterceptor.proto", fileDescriptor_ab8084eb1a07bd8b) }

var fileDescriptor_ab8084eb1a07bd8b = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x4f, 0xdd, 0x6a, 0x83, 0x30,
	0x14, 0x26, 0xba, 0x3f, 0x8f, 0xe0, 0x20, 0xec, 0x42, 0x1c, 0x03, 0xc9, 0x95, 0xec, 0x42, 0x86,
	0x7b, 0x83, 0xdd, 0x09, 0xbb, 0x18, 0xd9, 0x1e, 0x60, 0xa9, 0x39, 0x84, 0xd0, 0x9a, 0x48, 0x9a,
	0x0a, 0x7d, 0xfb, 0xa2, 0x8d, 0xb5, 0xf4, 0xee, 0x7c, 0x1f, 0xdf, 0xdf, 0x81, 0x62, 0xef, 0x1d,
	0x8a, 0x5e, 0x1b, 0xd5, 0x1a, 0x8f, 0xae, 0xc3, 0xc1, 0x5b, 0x57, 0x0f, 0xce, 0x7a, 0xcb, 0xbe,
	0xe1, 0xf1, 0xc7, 0x6e, 0xb1, 0xb7, 0x86, 0x66, 0x10, 0x69, 0x99, 0x93, 0x92, 0x54, 0x31, 0x8f,
	0xb4, 0xa4, 0x14, 0xee, 0x8c, 0xe8, 0x31, 0x8f, 0x4a, 0x52, 0x25, 0x7c, 0xbe, 0xe9, 0x1b, 0xc4,
	0x5e, 0xa8, 0x3c, 0x2e, 0x49, 0x95, 0x36, 0x69, 0x1d, 0xac, 0x7f, 0x42, 0xf1, 0x89, 0x67, 0xaf,
	0x90, 0x04, 0xaa, 0x95, 0xb7, 0x79, 0x8c, 0x01, 0xac, 0x7a, 0xfa, 0x02, 0xf7, 0xa3, 0xd8, 0x1d,
	0x70, 0x16, 0x24, 0xfc, 0x0c, 0x9a, 0x7f, 0xc8, 0x82, 0xe6, 0x17, 0xdd, 0xa8, 0x3b, 0xa4, 0x0c,
	0x40, 0x48, 0xb9, 0x6c, 0x7c, 0x5a, 0x2a, 0x0b, 0xa8, 0xd7, 0xa6, 0x77, 0x78, 0x56, 0xe8, 0x03,
	0xfe, 0x3a, 0x4e, 0xf1, 0xd7, 0xdb, 0x8a, 0x8b, 0xeb, 0x83, 0x6c, 0x1e, 0xe6, 0xbf, 0x3f, 0x4f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x9f, 0xe6, 0x61, 0x15, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PokemonServiceClient is the client API for PokemonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PokemonServiceClient interface {
	AddPokemon(ctx context.Context, in *Pokemon, opts ...grpc.CallOption) (*PokemonId, error)
	GetPokemonByTag(ctx context.Context, in *PokemonTag, opts ...grpc.CallOption) (PokemonService_GetPokemonByTagClient, error)
}

type pokemonServiceClient struct {
	cc *grpc.ClientConn
}

func NewPokemonServiceClient(cc *grpc.ClientConn) PokemonServiceClient {
	return &pokemonServiceClient{cc}
}

func (c *pokemonServiceClient) AddPokemon(ctx context.Context, in *Pokemon, opts ...grpc.CallOption) (*PokemonId, error) {
	out := new(PokemonId)
	err := c.cc.Invoke(ctx, "/PokemonService/addPokemon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonServiceClient) GetPokemonByTag(ctx context.Context, in *PokemonTag, opts ...grpc.CallOption) (PokemonService_GetPokemonByTagClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PokemonService_serviceDesc.Streams[0], "/PokemonService/getPokemonByTag", opts...)
	if err != nil {
		return nil, err
	}
	x := &pokemonServiceGetPokemonByTagClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PokemonService_GetPokemonByTagClient interface {
	Recv() (*Pokemon, error)
	grpc.ClientStream
}

type pokemonServiceGetPokemonByTagClient struct {
	grpc.ClientStream
}

func (x *pokemonServiceGetPokemonByTagClient) Recv() (*Pokemon, error) {
	m := new(Pokemon)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PokemonServiceServer is the server API for PokemonService service.
type PokemonServiceServer interface {
	AddPokemon(context.Context, *Pokemon) (*PokemonId, error)
	GetPokemonByTag(*PokemonTag, PokemonService_GetPokemonByTagServer) error
}

// UnimplementedPokemonServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPokemonServiceServer struct {
}

func (*UnimplementedPokemonServiceServer) AddPokemon(ctx context.Context, req *Pokemon) (*PokemonId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPokemon not implemented")
}
func (*UnimplementedPokemonServiceServer) GetPokemonByTag(req *PokemonTag, srv PokemonService_GetPokemonByTagServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPokemonByTag not implemented")
}

func RegisterPokemonServiceServer(s *grpc.Server, srv PokemonServiceServer) {
	s.RegisterService(&_PokemonService_serviceDesc, srv)
}

func _PokemonService_AddPokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pokemon)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServiceServer).AddPokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PokemonService/AddPokemon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServiceServer).AddPokemon(ctx, req.(*Pokemon))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokemonService_GetPokemonByTag_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PokemonTag)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PokemonServiceServer).GetPokemonByTag(m, &pokemonServiceGetPokemonByTagServer{stream})
}

type PokemonService_GetPokemonByTagServer interface {
	Send(*Pokemon) error
	grpc.ServerStream
}

type pokemonServiceGetPokemonByTagServer struct {
	grpc.ServerStream
}

func (x *pokemonServiceGetPokemonByTagServer) Send(m *Pokemon) error {
	return x.ServerStream.SendMsg(m)
}

var _PokemonService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PokemonService",
	HandlerType: (*PokemonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addPokemon",
			Handler:    _PokemonService_AddPokemon_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getPokemonByTag",
			Handler:       _PokemonService_GetPokemonByTag_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "streamingInterceptor.proto",
}
