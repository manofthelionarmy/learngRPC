package service

import (
	"context"
	"fmt"

	pb "github.com/manofthelionarmy/learngRPC/ch2/pokemon"
)

var _ pb.PokemonInfoServer = (*PokemonGrpcServer)(nil)

// PokemonGrpcServer is the implementation of my grpc server
type PokemonGrpcServer struct {
	pb.UnimplementedPokemonInfoServer
	Pokemon map[string]*pb.Pokemon
}

// GetPokemon is the implementation of the function my service exposes
func (p PokemonGrpcServer) GetPokemon(
	ctx context.Context,
	in *pb.PokemonID,
) (
	*pb.Pokemon,
	error,
) {

	pokemon := &pb.Pokemon{
		Name:      "zubat",
		Attack:    35,
		Defense:   45,
		SpAttack:  40,
		SpDefense: 45,
		Speed:     65,
	}
	return pokemon, nil
}

// AddPokemon implements the AddPokemon stub
func (p PokemonGrpcServer) AddPokemon(ctx context.Context, in *pb.Pokemon) (
	*pb.PokemonID,
	error,
) {
	p.Pokemon[in.Name] = in

	pokemonID := pb.PokemonID{
		Value: in.Name,
	}
	return &pokemonID, nil
}
