package chat

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/alexandrhub/grpc-chat/internal/converter"
	pb "github.com/alexandrhub/grpc-chat/pkg/chat_v1"
)

func (i *Implementation) Send(ctx context.Context, req *pb.SendMessageRequest) (*empty.Empty, error) {
	err := i.chatService.Send(ctx, converter.ToChatFromDescSend(req))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}
