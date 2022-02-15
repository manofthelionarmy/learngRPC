package clientstreaminterceptor

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/manofthelionarmy/learngRPC/ch5/interceptors/clientStreamInterceptor"
	"google.golang.org/grpc"
)

var _ pb.PokemonServiceServer = (*PokemonServer)(nil)

// PokemonServer implements the grpc service defined in our service definition
type PokemonServer struct {
	Pokemon map[string][]*pb.Pokemon
}

// AddPokemon implements the remote funciton exposed by the service
func (p PokemonServer) AddPokemon(
	ctx context.Context,
	pokemon *pb.Pokemon,
) (*pb.PokemonId, error) {
	if _, ok := p.Pokemon[pokemon.GetTag().GetValue()]; !ok {
		p.Pokemon[pokemon.GetTag().GetValue()] = make([]*pb.Pokemon, 0)
	}

	p.Pokemon[pokemon.GetTag().GetValue()] = append(
		p.Pokemon[pokemon.GetTag().GetValue()],
		pokemon,
	)

	total := 0
	for _, pokemons := range p.Pokemon {
		total += len(pokemons)
	}
	return &pb.PokemonId{Id: int64(total)}, nil
}

// GetPokemonByTag implements the remote function exposed by the service
func (p PokemonServer) GetPokemonByTag(
	tagStream pb.PokemonService_GetPokemonByTagServer,
) error {
	for {
		_, err := tagStream.Recv()
		if err == io.EOF {
			pokemons := make([]*pb.Pokemon, 0)
			for _, p := range p.Pokemon {
				pokemons = append(pokemons, p...)
			}
			return tagStream.SendAndClose(&pb.Pokemons{Pokemons: pokemons})
		}

		if err != nil {
			return err
		}
	}
}

type clientStreamWrapper struct {
	grpc.ClientStream
}

func newWrappedStream(clientStream grpc.ClientStream) grpc.ClientStream {
	return &clientStreamWrapper{ClientStream: clientStream}
}

func (cs clientStreamWrapper) RecvMsg(msg interface{}) error {
	log.Printf(
		"=== [Client Stream Interceptor] ==="+
			"Received a msg (Type: %T) at %+v\n",
		msg,
		time.Now().Format(time.RFC3339),
	)
	return cs.ClientStream.RecvMsg(msg)
}

func (cs clientStreamWrapper) SendMsg(msg interface{}) error {
	log.Printf(
		"=== [Client Stream Interceptor] ==="+
			"Sent a msg (Type: %T) at %+v\n",
		msg,
		time.Now().Format(time.RFC3339),
	)
	return cs.ClientStream.SendMsg(msg)
}

// pokemonClientStreamInterceptor is a function that is of type StreamClientInterceptor
func pokemonClientStreamInterceptor(
	ctx context.Context,
	desc *grpc.StreamDesc,
	cc *grpc.ClientConn,
	method string,
	streamer grpc.Streamer, // Streamer is called by StreamClientInterceptor to create a ClientStream.
	opts ...grpc.CallOption,
) (grpc.ClientStream, error) {
	log.Println("=== [Client Stream Interceptor] ===", method)
	clientStreamer, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	return newWrappedStream(clientStreamer), nil
}
