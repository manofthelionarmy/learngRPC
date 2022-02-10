package unaryinterceptor

import (
	"context"
	"fmt"

	pb "github.com/manofthelionarmy/learngRPC/ch5/interceptors/unary"
	"google.golang.org/grpc"
)

var _ pb.PokemonServiceServer = (*PokemonServer)(nil)

// PokemonServer implements the remote functions exposed by my service definition
type PokemonServer struct {
	Pokemon map[string]*pb.Pokemon
}

// AddPokemon implements the remote function exposed by the service definition
func (p PokemonServer) AddPokemon(ctx context.Context, pokemon *pb.Pokemon) (*pb.PokemonId, error) {
	if _, ok := p.Pokemon[pokemon.Name]; !ok {
		pokemon.Id = uint32(len(p.Pokemon) + 1)
		p.Pokemon[pokemon.Name] = pokemon
	}
	return &pb.PokemonId{Id: p.Pokemon[pokemon.Name].Id}, nil
}

// We must create a unary server interceptor function that has the same signature as the UnaryInterceptor function,
// Similar to how we have to create a handler function with the signature matching HandlerFunc
func pokemonUnaryServerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	// Pre-Process request
	fmt.Printf("logging request %+v\n", req)
	// Invoke rpc
	msg, err := handler(ctx, req)

	// Post-process
	fmt.Printf("Post processing response message %+v\n", msg)
	resp = msg
	return
}
