// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pokemon

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PokemonInfoClient is the client API for PokemonInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PokemonInfoClient interface {
	GetPokemon(ctx context.Context, in *PokemonID, opts ...grpc.CallOption) (*Pokemon, error)
	AddPokemon(ctx context.Context, in *Pokemon, opts ...grpc.CallOption) (*PokemonID, error)
}

type pokemonInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewPokemonInfoClient(cc grpc.ClientConnInterface) PokemonInfoClient {
	return &pokemonInfoClient{cc}
}

func (c *pokemonInfoClient) GetPokemon(ctx context.Context, in *PokemonID, opts ...grpc.CallOption) (*Pokemon, error) {
	out := new(Pokemon)
	err := c.cc.Invoke(ctx, "/pokemon.PokemonInfo/getPokemon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonInfoClient) AddPokemon(ctx context.Context, in *Pokemon, opts ...grpc.CallOption) (*PokemonID, error) {
	out := new(PokemonID)
	err := c.cc.Invoke(ctx, "/pokemon.PokemonInfo/addPokemon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PokemonInfoServer is the server API for PokemonInfo service.
// All implementations must embed UnimplementedPokemonInfoServer
// for forward compatibility
type PokemonInfoServer interface {
	GetPokemon(context.Context, *PokemonID) (*Pokemon, error)
	AddPokemon(context.Context, *Pokemon) (*PokemonID, error)
	mustEmbedUnimplementedPokemonInfoServer()
}

// UnimplementedPokemonInfoServer must be embedded to have forward compatible implementations.
type UnimplementedPokemonInfoServer struct {
}

func (UnimplementedPokemonInfoServer) GetPokemon(context.Context, *PokemonID) (*Pokemon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPokemon not implemented")
}
func (UnimplementedPokemonInfoServer) AddPokemon(context.Context, *Pokemon) (*PokemonID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPokemon not implemented")
}
func (UnimplementedPokemonInfoServer) mustEmbedUnimplementedPokemonInfoServer() {}

// UnsafePokemonInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PokemonInfoServer will
// result in compilation errors.
type UnsafePokemonInfoServer interface {
	mustEmbedUnimplementedPokemonInfoServer()
}

func RegisterPokemonInfoServer(s grpc.ServiceRegistrar, srv PokemonInfoServer) {
	s.RegisterService(&PokemonInfo_ServiceDesc, srv)
}

func _PokemonInfo_GetPokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PokemonID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonInfoServer).GetPokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.PokemonInfo/getPokemon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonInfoServer).GetPokemon(ctx, req.(*PokemonID))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokemonInfo_AddPokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pokemon)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonInfoServer).AddPokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.PokemonInfo/addPokemon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonInfoServer).AddPokemon(ctx, req.(*Pokemon))
	}
	return interceptor(ctx, in, info, handler)
}

// PokemonInfo_ServiceDesc is the grpc.ServiceDesc for PokemonInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PokemonInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pokemon.PokemonInfo",
	HandlerType: (*PokemonInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getPokemon",
			Handler:    _PokemonInfo_GetPokemon_Handler,
		},
		{
			MethodName: "addPokemon",
			Handler:    _PokemonInfo_AddPokemon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ch2/pokemon/pokemon.proto",
}
