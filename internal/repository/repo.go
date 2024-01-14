package repository

import (
	"context"

	"github.com/alexandrhub/grpc-chat/internal/model"
)

type ChatRepository interface {
	Create(ctx context.Context, info *model.Chat) (int64, error)
	Send(ctx context.Context, info *model.Chat) error
	Delete(ctx context.Context, id int64) error
}
