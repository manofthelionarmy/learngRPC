// Code generated by protoc-gen-go. DO NOT EDIT.
// source: clientStreaming.proto

package clientStreaming

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

type CaughtPokemon struct {
	Value                uint64   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CaughtPokemon) Reset()         { *m = CaughtPokemon{} }
func (m *CaughtPokemon) String() string { return proto.CompactTextString(m) }
func (*CaughtPokemon) ProtoMessage()    {}
func (*CaughtPokemon) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2ccfb4be0af4408, []int{0}
}

func (m *CaughtPokemon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CaughtPokemon.Unmarshal(m, b)
}
func (m *CaughtPokemon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CaughtPokemon.Marshal(b, m, deterministic)
}
func (m *CaughtPokemon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CaughtPokemon.Merge(m, src)
}
func (m *CaughtPokemon) XXX_Size() int {
	return xxx_messageInfo_CaughtPokemon.Size(m)
}
func (m *CaughtPokemon) XXX_DiscardUnknown() {
	xxx_messageInfo_CaughtPokemon.DiscardUnknown(m)
}

var xxx_messageInfo_CaughtPokemon proto.InternalMessageInfo

func (m *CaughtPokemon) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Pokemon struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pokemon) Reset()         { *m = Pokemon{} }
func (m *Pokemon) String() string { return proto.CompactTextString(m) }
func (*Pokemon) ProtoMessage()    {}
func (*Pokemon) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2ccfb4be0af4408, []int{1}
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

func (m *Pokemon) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*CaughtPokemon)(nil), "CaughtPokemon")
	proto.RegisterType((*Pokemon)(nil), "Pokemon")
}

func init() { proto.RegisterFile("clientStreaming.proto", fileDescriptor_e2ccfb4be0af4408) }

var fileDescriptor_e2ccfb4be0af4408 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0x09, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0xcd, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x57, 0x52, 0xe5, 0xe2, 0x75, 0x4e, 0x2c, 0x4d, 0xcf, 0x28, 0x09, 0xc8, 0xcf, 0x4e, 0xcd,
	0xcd, 0xcf, 0x13, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4,
	0x60, 0x09, 0x82, 0x70, 0x94, 0x64, 0xb9, 0xd8, 0x61, 0x0a, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73,
	0x21, 0xf2, 0x9c, 0x41, 0x60, 0xb6, 0x91, 0x0d, 0x17, 0x1f, 0x54, 0x3a, 0x38, 0xb5, 0xa8, 0x2c,
	0x33, 0x39, 0x55, 0x48, 0x8b, 0x8b, 0xc7, 0x39, 0xb1, 0x24, 0x39, 0x03, 0xa6, 0x8b, 0x43, 0x0f,
	0xca, 0x92, 0xe2, 0xd3, 0x43, 0xb1, 0x50, 0x83, 0x31, 0x89, 0x0d, 0xec, 0x14, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xc5, 0xdc, 0x62, 0x72, 0xa3, 0x00, 0x00, 0x00,
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
	CatchPokemon(ctx context.Context, opts ...grpc.CallOption) (PokemonService_CatchPokemonClient, error)
}

type pokemonServiceClient struct {
	cc *grpc.ClientConn
}

func NewPokemonServiceClient(cc *grpc.ClientConn) PokemonServiceClient {
	return &pokemonServiceClient{cc}
}

func (c *pokemonServiceClient) CatchPokemon(ctx context.Context, opts ...grpc.CallOption) (PokemonService_CatchPokemonClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PokemonService_serviceDesc.Streams[0], "/PokemonService/CatchPokemon", opts...)
	if err != nil {
		return nil, err
	}
	x := &pokemonServiceCatchPokemonClient{stream}
	return x, nil
}

type PokemonService_CatchPokemonClient interface {
	Send(*Pokemon) error
	CloseAndRecv() (*CaughtPokemon, error)
	grpc.ClientStream
}

type pokemonServiceCatchPokemonClient struct {
	grpc.ClientStream
}

func (x *pokemonServiceCatchPokemonClient) Send(m *Pokemon) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pokemonServiceCatchPokemonClient) CloseAndRecv() (*CaughtPokemon, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CaughtPokemon)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PokemonServiceServer is the server API for PokemonService service.
type PokemonServiceServer interface {
	CatchPokemon(PokemonService_CatchPokemonServer) error
}

// UnimplementedPokemonServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPokemonServiceServer struct {
}

func (*UnimplementedPokemonServiceServer) CatchPokemon(srv PokemonService_CatchPokemonServer) error {
	return status.Errorf(codes.Unimplemented, "method CatchPokemon not implemented")
}

func RegisterPokemonServiceServer(s *grpc.Server, srv PokemonServiceServer) {
	s.RegisterService(&_PokemonService_serviceDesc, srv)
}

func _PokemonService_CatchPokemon_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PokemonServiceServer).CatchPokemon(&pokemonServiceCatchPokemonServer{stream})
}

type PokemonService_CatchPokemonServer interface {
	SendAndClose(*CaughtPokemon) error
	Recv() (*Pokemon, error)
	grpc.ServerStream
}

type pokemonServiceCatchPokemonServer struct {
	grpc.ServerStream
}

func (x *pokemonServiceCatchPokemonServer) SendAndClose(m *CaughtPokemon) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pokemonServiceCatchPokemonServer) Recv() (*Pokemon, error) {
	m := new(Pokemon)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PokemonService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PokemonService",
	HandlerType: (*PokemonServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CatchPokemon",
			Handler:       _PokemonService_CatchPokemon_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "clientStreaming.proto",
}
