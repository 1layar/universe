package repo

import (
	"context"

	"github.com/1layar/universe/src/email_service/model"
	"github.com/uptrace/bun"
)

type EmailMessageRepository struct {
	db *bun.DB
}

func NewEmailMessageRepository(db *bun.DB) *EmailMessageRepository {
	return &EmailMessageRepository{db: db}
}

func (r *EmailMessageRepository) Create(ctx context.Context, message *model.EmailMessage) error {
	_, err := r.db.NewInsert().Model(message).Exec(ctx)
	return err
}

func (r *EmailMessageRepository) GetByID(ctx context.Context, id int) (*model.EmailMessage, error) {
	message := new(model.EmailMessage)
	err := r.db.NewSelect().Model(message).Where("id = ?", id).Scan(ctx)
	return message, err
}

func (r *EmailMessageRepository) Update(ctx context.Context, message *model.EmailMessage) error {
	_, err := r.db.NewUpdate().Model(message).Where("id = ?", message.ID).Exec(ctx)
	return err
}

func (r *EmailMessageRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.NewDelete().Model(&model.EmailMessage{}).Where("id = ?", id).Exec(ctx)
	return err
}
