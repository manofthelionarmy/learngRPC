package serverstreaminterceptor

import (
	"context"
	"fmt"
	"io"
	"net"
	"sync"
	"testing"

	pb "github.com/manofthelionarmy/learngRPC/ch5/interceptors/streaming"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestServerStreamInterceptor(t *testing.T) {
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
	)

	require.NoError(t, err)

	defer clientConnection.Close()

	pokemonServer := PokemonStreamServer{
		Pokemon: make(map[string][]*pb.Pokemon),
	}

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(
			pokmeonStreamInterceptor,
		),
	)

	pb.RegisterPokemonServiceServer(grpcServer, pokemonServer)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		grpcServer.Serve(listener)
	}()

	pokemonClient := pb.NewPokemonServiceClient(clientConnection)

	_, err = pokemonClient.AddPokemon(
		context.Background(),
		&pb.Pokemon{
			Name: "Pikachu",
			Tag: &pb.PokemonTag{
				Value: "favorites",
			},
		},
	)

	require.NoError(t, err)

	_, err = pokemonClient.AddPokemon(
		context.Background(),
		&pb.Pokemon{
			Name: "Zubat",
			Tag: &pb.PokemonTag{
				Value: "favorites",
			},
		},
	)

	require.NoError(t, err)

	pokemonServerStream, err := pokemonClient.GetPokemonByTag(
		context.Background(),
		&pb.PokemonTag{
			Value: "favorites",
		},
	)

	require.NoError(t, err)
	for {
		pokemon, err := pokemonServerStream.Recv()
		if err == io.EOF {
			break
		}

		fmt.Printf("Got pokemon from PC: %+v\n", pokemon)
	}

	// Tell it to stop receiving messages from client
	pokemonServerStream.CloseSend()

	grpcServer.GracefulStop()
	result()
}
