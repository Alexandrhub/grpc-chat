package chat

import (
	"github.com/alexandrhub/grpc-chat/internal/service"
	pb "github.com/alexandrhub/grpc-chat/pkg/chat_v1"
)

type Implementation struct {
	pb.UnimplementedChatV1Server
	chatService service.ChatService
}

func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}
