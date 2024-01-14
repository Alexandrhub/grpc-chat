package chat

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/alexandrhub/grpc-chat/internal/converter"
	pb "github.com/alexandrhub/grpc-chat/pkg/chat_v1"
)

func (i *Implementation) Delete(ctx context.Context, req *pb.DeleteChatRequest) (*empty.Empty, error) {
	err := i.chatService.Delete(ctx, converter.ToChatFromDescDelete(req))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
