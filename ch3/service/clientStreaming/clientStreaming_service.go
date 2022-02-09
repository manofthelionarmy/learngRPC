package clientstreaming

import (
	"io"

	pb "github.com/manofthelionarmy/learngRPC/ch3/clientStreaming"
)

var _ pb.PokemonServiceServer = (*PokemonService)(nil)

// PokemonService is my implementation of the the PokemonService
type PokemonService struct {
	Pokemon map[string]*pb.Pokemon
}

// CatchPokemon is the implementation of the remote function exposed by my service interface/definition
func (p PokemonService) CatchPokemon(clientStream pb.PokemonService_CatchPokemonServer) error {
	numCaught := 0
	for {
		pokemon, err := clientStream.Recv()
		if err == io.EOF {
			// finish reading the client stream of pokemon caught
			return clientStream.SendAndClose(&pb.CaughtPokemon{Value: uint64(numCaught)})
		}

		if _, ok := p.Pokemon[pokemon.Name]; !ok {
			p.Pokemon[pokemon.Name] = pokemon
		}
		numCaught++
	}
}
