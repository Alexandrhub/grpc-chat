package chat

import (
	"context"

	"github.com/alexandrhub/grpc-chat/internal/model"
)

func (s *serv) Create(ctx context.Context, info *model.Chat) (int64, error) {
	var id int64
	err := s.txManager.ReadCommitted(
		ctx, func(ctx context.Context) error {
			var errTx error
			id, errTx = s.chatRepository.Create(ctx, info)
			return errTx
		},
	)

	if err != nil {
		return 0, err
	}

	return id, nil
}
