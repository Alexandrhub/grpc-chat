package service

import (
	"context"

	"github.com/alexandrhub/grpc-chat/internal/model"
)

type ChatService interface {
	Create(ctx context.Context, info *model.Chat) (int64, error)
	Delete(ctx context.Context, id int64) error
	Send(ctx context.Context, info *model.Chat) error
}
