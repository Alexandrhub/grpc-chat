package chat

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/alexandrhub/grpc-chat/internal/converter"
	pb "github.com/alexandrhub/grpc-chat/pkg/chat_v1"
)

func (i *Implementation) Create(ctx context.Context, req *pb.CreateChatRequest) (*pb.CreateChatResponse, error) {
	id, err := i.chatService.Create(ctx, converter.ToChatFromDescCreate(req))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	log.Println("created chat", id)

	return &pb.CreateChatResponse{
		Id: id,
	}, nil
}
