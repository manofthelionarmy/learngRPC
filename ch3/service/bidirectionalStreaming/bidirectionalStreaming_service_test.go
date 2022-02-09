package bidirectionalstreaming

import (
	"context"
	"fmt"
	"io"
	"net"
	"sync"
	"testing"
	"time"

	pb "github.com/manofthelionarmy/learngRPC/ch3/bidirectionalStreaming"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestBidirectionalStreaming(t *testing.T) {
	for scenario, f := range map[string]func(t *testing.T){
		"testBidirecitonalStreaming": testBidirecitonalStreaming,
	} {
		t.Run(scenario, f)
	}
}

func testBidirecitonalStreaming(t *testing.T) {
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

	textMsgServer := TextMessageServer{
		Conversation: make(map[string][]*pb.TextMessage),
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTextServiceServer(grpcServer, textMsgServer)

	textMsgClient := pb.NewTextServiceClient(clientConnection)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		grpcServer.Serve(listener)
	}()

	textMsgClientStream, err := textMsgClient.ProduceMessage(context.Background())
	require.NoError(t, err)

	err = textMsgClientStream.Send(&pb.TextMessage{Day: "Monday", Value: "Monday's suck"})
	require.NoError(t, err)

	wg.Add(1)
	// Concurrently receive messages bidirectionally from server as we are sending messages via the client
	go func() {
		defer wg.Done()
		for {
			// receive form the server, maybe a better approach is pass this to a channel
			msgFromServer, err := textMsgClientStream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				fmt.Println(err.Error())
				break // break breaks out of the if statement
			}
			require.NoError(t, err)

			fmt.Printf("The server said: %+v\n", msgFromServer.TextMessages)
		}
	}()
	// I guess it's not enough time for the messages to be recieved and sent back
	err = textMsgClientStream.CloseSend()
	require.NoError(t, err)

	// Use graceful stop?
	// GracefulStop stops the gRPC server gracefully. It stops the server from
	// accepting new connections and RPCs and blocks until all the pending RPCs are
	// finished.
	grpcServer.GracefulStop() // gracefully stops and let's all pending rpcs finish
	wg.Wait()
}
