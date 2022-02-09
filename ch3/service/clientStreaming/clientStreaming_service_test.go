package clientstreaming

import (
	"context"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/manofthelionarmy/learngRPC/ch3/clientStreaming"
	pb "github.com/manofthelionarmy/learngRPC/ch3/clientStreaming"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestPokemonClientStream(t *testing.T) {
	for scenario, f := range map[string]func(t *testing.T){
		"testClientStream": testClientStream,
	} {
		t.Run(scenario, f)
	}
}

func testClientStream(t *testing.T) {
	const bufSize = 1024

	listner := bufconn.Listen(bufSize)
	defer listner.Close()

	clientConnection, err := grpc.DialContext(
		context.Background(),
		listner.Addr().String(),
		grpc.WithDialer(func(s string, d time.Duration) (net.Conn, error) {
			return listner.Dial()
		}),
		grpc.WithInsecure(),
	)

	require.NoError(t, err)

	defer clientConnection.Close()

	pokemonService := PokemonService{
		Pokemon: make(map[string]*clientStreaming.Pokemon),
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPokemonServiceServer(grpcServer, pokemonService)

	pokemonGrpcClient := pb.NewPokemonServiceClient(clientConnection)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		grpcServer.Serve(listner)
	}()

	clientStream, err := pokemonGrpcClient.CatchPokemon(context.Background())
	require.NoError(t, err)

	err = clientStream.Send(&pb.Pokemon{Name: "Pikachu"})
	require.NoError(t, err)

	err = clientStream.Send(&pb.Pokemon{Name: "Eevee"})
	require.NoError(t, err)

	pokemonCaught, err := clientStream.CloseAndRecv()
	require.NoError(t, err)

	require.Equal(t, 2, int(pokemonCaught.Value))

	grpcServer.Stop()
	wg.Wait()
}
