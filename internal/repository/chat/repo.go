package chat

import (
	"context"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"

	"github.com/alexandrhub/grpc-chat/internal/db"
	"github.com/alexandrhub/grpc-chat/internal/model"
	"github.com/alexandrhub/grpc-chat/internal/repository"
)

const (
	tableName = "chats"
	idColumn  = "id"
	username  = "username"
	msg       = "msg"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, info *model.Chat) (int64, error) {
	query := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(username, msg).
		Values(info.Name, info.Msg).
		Suffix("RETURNING id")

	queryText, args, err := query.ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed to build query: %w", err)
	}

	q := db.Query{
		Name:     "chat_repo_create",
		QueryRaw: queryText,
	}

	var chatServerID int64
	if err := r.db.DB().QueryRowContext(ctx, q, args...).Scan(&chatServerID); err != nil {
		return 0, fmt.Errorf("failed to insert chat_server: %w", err)
	}

	log.Printf("inserted chat_server with id: %d", chatServerID)

	return chatServerID, nil
}

func (r *repo) Delete(ctx context.Context, deleteID int64) error {
	query := sq.Delete(tableName).
		Where(sq.Eq{idColumn: deleteID}).
		PlaceholderFormat(sq.Dollar)

	queryText, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	q := db.Query{
		Name:     "chat_repo_delete",
		QueryRaw: queryText,
	}

	res, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("failed to delete chat_server: %v with error: %w", res, err)
	}

	log.Printf("deleted chat_server with id: %d", deleteID)

	return nil
}

func (r *repo) Send(ctx context.Context, info *model.Chat) error {
	return nil
}
