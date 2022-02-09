package bidirectionalstreaming

import (
	"io"

	pb "github.com/manofthelionarmy/learngRPC/ch3/bidirectionalStreaming"
)

var _ pb.TextServiceServer = (*TextMessageServer)(nil)

// TextMessageServer is the implementation of the Text Service
type TextMessageServer struct {
	Conversation map[string][]*pb.TextMessage
}

// ProduceMessage implements the remote function exposed by the service defintion
func (t TextMessageServer) ProduceMessage(bidirectionalStream pb.TextService_ProduceMessageServer) error {
	completeConversation := &pb.CompleteConversation{}
	for {
		textMsg, err := bidirectionalStream.Recv()
		if err == io.EOF {
			return bidirectionalStream.Send(completeConversation)
		}

		if err != nil {
			return err
		}

		if _, ok := t.Conversation[textMsg.Day]; !ok {
			t.Conversation[textMsg.Day] = make([]*pb.TextMessage, 0)
		}

		conversation, _ := t.Conversation[textMsg.Day]
		conversation = append(conversation, textMsg)
		completeConversation.TextMessages = append(completeConversation.TextMessages, []*pb.TextMessage{textMsg}...)

		bidirectionalStream.Send(&pb.CompleteConversation{TextMessages: conversation})
	}
}
