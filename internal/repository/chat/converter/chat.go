package converter

import (
	"github.com/alexandrhub/grpc-chat/internal/model"
	modelRepo "github.com/alexandrhub/grpc-chat/internal/repository/chat/model"
)

func ToChatFromRepo(chat *modelRepo.Chat) *model.Chat {
	return &model.Chat{
		ID:   chat.ID,
		Name: chat.Name,
		Msg:  chat.Msg,
	}
}
