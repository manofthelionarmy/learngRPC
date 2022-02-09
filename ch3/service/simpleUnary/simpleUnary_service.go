package simpleunary

import (
	"context"
	"fmt"

	pb "github.com/manofthelionarmy/learngRPC/ch3/simpleUnary"
)

// AnimeGrpcServer is my implementation of the simple/unary grpc server
type AnimeGrpcServer struct {
	AnimeMap map[string]*pb.Anime
}

var _ pb.AnimeServiceServer = (*AnimeGrpcServer)(nil)

// GetAnime invokes the remote function exposed by our grpc service
func (a AnimeGrpcServer) GetAnime(ctx context.Context, in *pb.AnimeId) (
	*pb.Anime,
	error,
) {
	anime, ok := a.AnimeMap[in.Value]
	if !ok {
		return nil, fmt.Errorf("anime not found")
	}
	return anime, nil
}
