package converter

import (
	"github.com/alexandrhub/grpc-chat/internal/model"
	pb "github.com/alexandrhub/grpc-chat/pkg/chat_v1"
)

func ToChatFromDescCreate(chat *pb.CreateChatRequest) *model.Chat {
	return &model.Chat{
		ID:   0,
		Name: chat.Username,
		Msg:  chat.Msg,
	}
}

func ToChatFromDescDelete(chat *pb.DeleteChatRequest) int64 {
	return chat.Id
}

func ToChatFromDescSend(chat *pb.SendMessageRequest) *model.Chat {
	return &model.Chat{
		ID:   0,
		Name: []string{chat.From},
		Msg:  chat.Text,
	}
}
