package clientunaryinterceptor

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "github.com/manofthelionarmy/learngRPC/ch5/interceptors/clientUnaryInterceptor"
)

// PokemonServer implements the grpc service defined in our service definition
type PokemonServer struct {
	Pokemon map[string]*pb.Pokemon
}

// AddPokemon implements the remote function exposed by our grpc service
func (p PokemonServer) AddPokemon(
	ctx context.Context,
	pokemon *pb.Pokemon,
) (*pb.PokemonId, error) {
	if _, ok := p.Pokemon[pokemon.Name]; !ok {
		p.Pokemon[pokemon.Name] = pokemon
	}

	id := len(p.Pokemon)
	return &pb.PokemonId{Id: int32(id)}, nil
}

func pokemonUnaryClientInterceptor(
	ctx context.Context, method string,
	req, reply interface{}, cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
) error {
	// Pre-process

	log.Println("=== [Client Side Unary Interceptor] ===", method, req)
	// Invoke
	err := invoker(ctx, method, req, reply, cc, opts...)

	// Post-process
	log.Println("=== [Client Side Unary Interceptor: Reply] ===", reply)
	return err
}
