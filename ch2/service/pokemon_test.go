package service

import (
	"context"
	"fmt"
	"net"
	"sync"
	"testing"

	pb "github.com/manofthelionarmy/learngRPC/ch2/pokemon"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestPokemonGrpcService(t *testing.T) {
	for scenario, f := range map[string]func(t *testing.T){
		"testPokemonGrpcService":  testPokemonGrpcService,
		"testPokemGrpcGetPokemon": testPokemGrpcGetPokemon,
		"testGrpcDialCtx":         testGrpcDialCtx,
	} {
		t.Run(scenario, f)
	}
}

func testPokemonGrpcService(t *testing.T) {
	// I think the server starts, but how do we test it?
	// t.Skip()
	// TODO: lern how to use grpc.DialContext
	server := &PokemonGrpcServer{
		Pokemon: make(map[string]*pb.Pokemon),
	}

	lis, err := net.Listen("tcp", ":8080")
	require.NoError(t, err)

	//  This is to test our client connection to our server
	clientOptions := []grpc.DialOption{grpc.WithInsecure()}
	clientConnection, err := grpc.Dial(lis.Addr().String(), clientOptions...)
	require.NoError(t, err)

	server.GetPokemon(context.Background(),
		&pb.PokemonID{},
	)

	s := grpc.NewServer()
	// In the book, it accepts a pointer, idk why it accecpts a nil pointer now
	pb.RegisterPokemonInfoServer(s, *server)
	wg := sync.WaitGroup{}
	wg.Add(1)
	// Why a go function? I guess we schedule this to happen in the background/ or in another goroutine
	// so it doesn't block the main go routine
	// I added it as part of the waitgroup to ensure we don't have a gorotuine/memory leak
	go func() {
		defer wg.Done()
		err := s.Serve(lis)
		// When we call stop, the error must be grpc.ErrServerStopped
		require.Error(t, err)
		require.EqualError(t, err, grpc.ErrServerStopped.Error())
	}()
	defer wg.Wait()

	s.Stop()
	clientConnection.Close()
	lis.Close()
}

func testPokemGrpcGetPokemon(t *testing.T) {
	lis, err := net.Listen("tcp", ":8000")
	require.NoError(t, err)

	// defer lis.Close()

	server := PokemonGrpcServer{
		Pokemon: make(map[string]*pb.Pokemon),
	}

	s := grpc.NewServer()
	pb.RegisterPokemonInfoServer(s, server)
	// wg := sync.WaitGroup{}
	// wg.Add(1)
	// done := make(chan bool)
	g := errgroup.Group{}
	g.Go(func() error {
		return s.Serve(lis)
	})
	// go func() {
	// 	defer wg.Done()
	// 	// create a semaphore to give enough time for the server to be tore down
	// 	<-done
	// 	require.Error(t, err)
	// 	require.EqualError(t, err, grpc.ErrServerStopped.Error())
	// }()

	// Recall defer atcs like a stack of function calls
	// defer wg.Wait() // last to be called on exit/return of evenloping function
	// defer lis.Close() // second to be called
	// defer s.Stop()    // first to be called, must be called first, we can't close listener before stoping server

	pokemon, err := server.GetPokemon(context.Background(), &pb.PokemonID{})
	require.NoError(t, err)

	require.Equal(t, "zubat", pokemon.Name)
	fmt.Printf("%+v\n", pokemon)
	s.Stop()
	lis.Close()
	// done <- true
	err = g.Wait()
	require.NoError(t, err)
	// TODO: figure out a way to terminate the server
	// require.Error(t, err)
	// require.EqualError(t, err, grpc.ErrServerStopped.Error())
}

// This allows me to test but without using a port :)
// I needed to create a listener that uses a buffer conneciton made possible Dialers
func testGrpcDialCtx(t *testing.T) {
	// TODO: Familiarize myself with Dialers, or buffered connections
	const bufSize = 1024

	// Listen returns a Listener that can only be contacted by its own Dialers and
	// creates buffered connections between the two.
	lis := bufconn.Listen(bufSize)
	// The target name syntax is defined in
	// https://github.com/grpc/grpc/blob/master/doc/naming.md.
	// target name, is this the one provided by bufconn? It's the not
	// the same as the "network" we call in net.Dial. It's not the same as http/scheme.
	// Dude, it's the target address provided by listener
	clientConnection, err := grpc.DialContext(
		context.Background(),
		lis.Addr().String(),
		grpc.WithContextDialer(
			func(ctx context.Context, addr string) (net.Conn, error) {
				return lis.Dial()
			},
		),
		grpc.WithInsecure(),
	)

	require.NoError(t, err)
	clientConnection.Close()
	lis.Close()
}
