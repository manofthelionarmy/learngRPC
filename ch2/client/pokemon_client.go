package client

import (
	"context"
	"fmt"

	pb "github.com/manofthelionarmy/learngRPC/ch2/pokemon"
	"google.golang.org/grpc"
)

// PokemonGrpcClient wraps my grpcStub/client
type PokemonGrpcClient struct {
	pb.PokemonInfoClient
}

// NewPokemonInfoClient creates a new pokemon info grpc client
func NewPokemonInfoClient(conn *grpc.ClientConn) (*PokemonGrpcClient, error) {
	if conn == nil {
		return nil, fmt.Errorf("Must past a grpc client connection")
	}
	return &PokemonGrpcClient{
		PokemonInfoClient: pb.NewPokemonInfoClient(conn),
	}, nil
}

// GetPokemonDetail is a wrapper function to use our stub to call/invoke functions exposed by our service interface
func (c *PokemonGrpcClient) GetPokemonDetail(ctx context.Context, pokemonID *pb.PokemonID) (*pb.Pokemon, error) {
	return c.PokemonInfoClient.GetPokemon(ctx, pokemonID)
}
