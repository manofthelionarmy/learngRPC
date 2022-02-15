package clientstreaminterceptor

import (
	"context"
	"fmt"
	"net"
	"sync"
	"testing"

	pb "github.com/manofthelionarmy/learngRPC/ch5/interceptors/clientStreamInterceptor"
	"github.com/stretchr/testify/require"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestPokemonClientStreamInterceptor(t *testing.T) {
	for scenario, f := range map[string]func(t *testing.T){
		"testClientStream": testClientStream,
	} {
		t.Run(scenario, f)
	}
}

func testClientStream(t *testing.T) {
	const bufSize = 1024

	listener := bufconn.Listen(bufSize)
	defer listener.Close()
	clientConnection, err := grpc.DialContext(
		context.Background(),
		listener.Addr().String(),
		grpc.WithContextDialer(func(c context.Context, s string) (net.Conn, error) {
			return listener.Dial()
		}),
		grpc.WithInsecure(),
		grpc.WithStreamInterceptor(
			pokemonClientStreamInterceptor,
		),
	)

	require.NoError(t, err)

	defer clientConnection.Close()

	pokemonServer := PokemonServer{
		Pokemon: make(map[string][]*pb.Pokemon),
	}

	grpcServer := grpc.NewServer()

	pb.RegisterPokemonServiceServer(grpcServer, pokemonServer)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		grpcServer.Serve(listener)
	}()

	pokemonGrpcClient := pb.NewPokemonServiceClient(clientConnection)

	pokemonGrpcClient.AddPokemon(
		context.Background(),
		&pb.Pokemon{
			Name: "Zubat",
			Tag: &pb.PokemonTag{
				Value: "favorites",
			},
		},
	)

	pokemonGrpcClient.AddPokemon(
		context.Background(),
		&pb.Pokemon{
			Name: "Riolu",
			Tag: &pb.PokemonTag{
				Value: "favorites",
			},
		},
	)

	pokemonGrpcClient.AddPokemon(
		context.Background(),
		&pb.Pokemon{
			Name: "Pikachu",
			Tag: &pb.PokemonTag{
				Value: "team",
			},
		},
	)

	pokemonClientStream, err := pokemonGrpcClient.GetPokemonByTag(
		context.Background(),
	)

	require.NoError(t, err)

	err = pokemonClientStream.Send(&pb.PokemonTag{
		Value: "favorites",
	})

	require.NoError(t, err)

	err = pokemonClientStream.Send(&pb.PokemonTag{
		Value: "team",
	})

	require.NoError(t, err)

	pokemons, err := pokemonClientStream.CloseAndRecv()

	require.NoError(t, err)

	fmt.Printf("%+v\n", pokemons)

	require.NoError(t, err)

	grpcServer.GracefulStop()
	wg.Wait()
}
