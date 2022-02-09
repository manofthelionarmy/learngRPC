package simpleunary

import (
	"context"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/manofthelionarmy/learngRPC/ch3/simpleUnary"
	pb "github.com/manofthelionarmy/learngRPC/ch3/simpleUnary"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestSimpleUnaryService(t *testing.T) {
	for scenario, f := range map[string]func(t *testing.T){
		"testAnimeUnaryService": testAnimeUnaryService,
	} {
		t.Run(scenario, f)
	}
}

func testAnimeUnaryService(t *testing.T) {
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

	animeGrpcServer := &AnimeGrpcServer{
		AnimeMap: make(map[string]*simpleUnary.Anime),
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAnimeServiceServer(grpcServer, animeGrpcServer)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		grpcServer.Serve(listener)
	}()

	animeGrpcClient := pb.NewAnimeServiceClient(clientConnection)

	_, err = animeGrpcClient.GetAnime(context.Background(), &pb.AnimeId{Value: "Luffy"})
	require.Error(t, err)

	grpcServer.Stop()
	wg.Wait()
}
