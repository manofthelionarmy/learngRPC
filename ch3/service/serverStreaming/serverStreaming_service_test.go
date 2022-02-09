package serverstreaming

import (
	"context"
	"fmt"
	"io"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/manofthelionarmy/learngRPC/ch3/serverStreaming"
	pb "github.com/manofthelionarmy/learngRPC/ch3/serverStreaming"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestServerStreamingPattern(t *testing.T) {
	for scenario, f := range map[string]func(t *testing.T){
		"testServerStreaming": testServerStreaming,
	} {
		t.Run(scenario, f)
	}
}

func testServerStreaming(t *testing.T) {
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

	animeServiceServer := AnimeServer{
		AnimeByTag: make(map[string][]serverStreaming.Anime),
	}

	favorites := make([]pb.Anime, 0)
	favorites = append(favorites,
		[]pb.Anime{
			{
				Title: "Black Clover",
			},
			{
				Title: "One Piece",
			},
			{
				Title: "Dragonball Z",
			},
		}...,
	)

	animeServiceServer.AnimeByTag["favorites"] = favorites

	grpcServer := grpc.NewServer()

	pb.RegisterAnimeServiceServer(grpcServer, animeServiceServer)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		grpcServer.Serve(listener)
	}()

	animeServiceClient := pb.NewAnimeServiceClient(clientConnection)

	animeStream, err := animeServiceClient.GetAnimeStream(context.Background(), &pb.AnimeId{Tag: "favorites"})
	require.NoError(t, err)

	for {
		anime, err := animeStream.Recv()
		// When we return nil, from our implementation of GetAnimeStream for AnimeServiceServer, it'll return io.EOF
		if err == io.EOF {
			break
		}
		require.NoError(t, err)
		fmt.Println(anime.Title)
	}

	grpcServer.Stop()
	wg.Wait()
}
