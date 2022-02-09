package client

import (
	"context"
	"fmt"
	"net"
	"sync"
	"testing"

	pb "github.com/manofthelionarmy/learngRPC/ch2/pokemon"
	"github.com/manofthelionarmy/learngRPC/ch2/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestGrpcClient(t *testing.T) {
	for scenario, f := range map[string]func(t *testing.T){
		"testGrpcClient": testGrpcClient,
	} {
		t.Run(scenario, f)
	}
}

func testGrpcClient(t *testing.T) {

	const bufSize = 1024 // MB
	lis := bufconn.Listen(bufSize)
	defer lis.Close()

	pokemonServer := &service.PokemonGrpcServer{
		Pokemon: make(map[string]pb.Pokemon),
	}

	server := grpc.NewServer()
	pb.RegisterPokemonInfoServer(server, pokemonServer)

	// Establish connection with a buffered connection
	clientConnection, err := grpc.DialContext(
		context.Background(),
		lis.Addr().String(),
		grpc.WithContextDialer(
			func(c context.Context, s string) (net.Conn, error) {
				return lis.Dial()
			},
		),
		grpc.WithInsecure(),
	)

	require.NoError(t, err)
	defer clientConnection.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)

	done := make(chan bool)
	go func() {
		defer wg.Done()
		// why doesn't this throw an error?
		err := server.Serve(lis)
		<-done
		require.NoError(t, err)
		// require.Error(t, err)
		// require.EqualError(t, err, grpc.ErrServerStopped.Error())
	}()

	require.NoError(t, err)

	pokemonGrpcClient, err := NewPokemonInfoClient(clientConnection)
	require.NoError(t, err)

	zubat, err := pokemonGrpcClient.GetPokemonDetail(context.Background(), &pb.PokemonID{})
	require.Equal(t, zubat.Name, "zubat")

	fmt.Println(zubat.String())

	server.Stop()
	done <- true
	wg.Wait()
}
