package unaryinterceptor

import (
	"context"
	"net"
	"sync"
	"testing"
	"time"

	pb "github.com/manofthelionarmy/learngRPC/ch5/interceptors/unary"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestPokemonUnaryInterceptor(t *testing.T) {
	for scenario, f := range map[string]func(t *testing.T){
		"testUnaryInterceptor": testUnaryInterceptor,
	} {
		t.Run(scenario, f)
	}
}

func testUnaryInterceptor(t *testing.T) {
	const bufSize = 1024
	listener := bufconn.Listen(bufSize)
	defer listener.Close()

	clientConnection, err := grpc.DialContext(
		context.Background(),
		listener.Addr().String(),
		grpc.WithDialer(func(s string, d time.Duration) (net.Conn, error) {
			return listener.Dial()
		}),
		grpc.WithInsecure(),
	)

	require.NoError(t, err)

	defer clientConnection.Close()

	pokemonServer := PokemonServer{
		Pokemon: make(map[string]*pb.Pokemon),
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(pokemonUnaryServerInterceptor),
	)

	pb.RegisterPokemonServiceServer(grpcServer, pokemonServer)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		grpcServer.Serve(listener)
	}()

	pokemonClient := pb.NewPokemonServiceClient(clientConnection)
	pokemonID, err := pokemonClient.AddPokemon(context.Background(), &pb.Pokemon{Name: "Zubat"})

	require.Equal(t, uint32(1), pokemonID.Id)

	grpcServer.GracefulStop()
}
