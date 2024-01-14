package chat

import (
	"github.com/alexandrhub/grpc-chat/internal/db"
	"github.com/alexandrhub/grpc-chat/internal/repository"
	"github.com/alexandrhub/grpc-chat/internal/service"
)

type serv struct {
	chatRepository repository.ChatRepository
	txManager      db.TxManager
}

func NewService(chatRepository repository.ChatRepository, txManager db.TxManager) service.ChatService {
	return &serv{
		chatRepository: chatRepository,
		txManager:      txManager,
	}
}
