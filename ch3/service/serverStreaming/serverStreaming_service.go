package serverstreaming

import (
	"fmt"

	pb "github.com/manofthelionarmy/learngRPC/ch3/serverStreaming"
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
