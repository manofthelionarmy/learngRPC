package serverstreaming

import (
	"context"
	"fmt"

	pb "github.com/manofthelionarmy/learngRPC/ch3/serverStreaming"
	"google.golang.org/grpc"
)

// Make sure that AnimeServer implements AnimeServiceServer
var _ pb.AnimeServiceServer = (*AnimeServer)(nil)

// AnimeServer is my grpc server implementation
type AnimeServer struct {
	AnimeByTag map[string][]pb.Anime
}

// GetAnimeStream invokes the remote funciton exposed by my service
func (a AnimeServer) GetAnimeStream(query *pb.AnimeId, streamServer pb.AnimeService_GetAnimeStreamServer) error {
	animeByTag, ok := a.AnimeByTag[query.Tag]
	if !ok {
		return fmt.Errorf("Tag %s hasn't been created yet", query.Tag)
	}

	for _, anime := range animeByTag {
		err := streamServer.Send(&anime)
		if err != nil {
			return err
		}
	}

	return nil
}

type wrapServerStream struct {
	grpc.ServerStream
}

func newWrappedServerStream(ss grpc.ServerStream) grpc.ServerStream {
	return &wrapServerStream{ss}
}

func (ws *wrapServerStream) RecvMsg(m interface{}) error {
	return ws.ServerStream.RecvMsg(m)
}

func (ws *wrapServerStream) SendMsg(m interface{}) error {
	return ws.ServerStream.SendMsg(m)
}

func serverStreamInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {

	err := handler(srv, newWrappedServerStream(ss))
	if err != nil {
		return fmt.Errorf("")
	}
	return nil
}

type wrapClientStream struct {
	grpc.ClientStream
}

func newWrapClientStream(cs grpc.ClientStream) grpc.ClientStream {
	return &wrapClientStream{ClientStream: cs}
}

func (cs wrapClientStream) RecvMsg(msg interface{}) error {
	return cs.ClientStream.RecvMsg(msg)
}

func (cs wrapClientStream) SendMsg(msg interface{}) error {
	return cs.ClientStream.SendMsg(msg)
}

func clientStreamInterceptor(
	ctx context.Context,
	desc *grpc.StreamDesc,
	cc *grpc.ClientConn,
	method string,
	streamer grpc.Streamer,
	opts ...grpc.CallOption,
) (grpc.ClientStream, error) {

	clientStream, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}

	return newWrapClientStream(clientStream), nil
}
