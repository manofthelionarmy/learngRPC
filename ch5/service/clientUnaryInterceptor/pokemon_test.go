package clientunaryinterceptor

import (
	"context"
	"net"
	"sync"
	"testing"

	pb "github.com/manofthelionarmy/learngRPC/ch5/interceptors/clientUnaryInterceptor"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestUnaryClientSideInterceptor(t *testing.T) {
	const bufSize = 1024
	listener := bufconn.Listen(bufSize)

	clientConnection, err := grpc.DialContext(
		context.Background(),
		listener.Addr().String(),
		grpc.WithContextDialer(func(c context.Context, s string) (net.Conn, error) {
			return listener.Dial()
		}),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			pokemonUnaryClientInterceptor, // Similar to how we passed the interceptor as an option to creating the grpcServer
		),
	)

	require.NoError(t, err)

	defer clientConnection.Close()

	pokemonServer := PokemonServer{
		Pokemon: make(map[string]*pb.Pokemon),
	}

	grpcServer := grpc.NewServer()

	pb.RegisterPokemonServiceServer(grpcServer, pokemonServer)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		grpcServer.Serve(listener)
	}()

	pokemonClient := pb.NewPokemonServiceClient(clientConnection)

	pokemonClient.AddPokemon(
		context.Background(),
		&pb.Pokemon{
			Name: "Zubat",
		},
	)

	grpcServer.GracefulStop()
}
