package repo

import (
	"context"

	"github.com/1layar/universe/pkg/email_service/model"

	"github.com/uptrace/bun"
)

type EmailEventRepository struct {
	db *bun.DB
}

func NewEmailEventRepository(db *bun.DB) *EmailEventRepository {
	return &EmailEventRepository{db: db}
}

func (r *EmailEventRepository) Create(ctx context.Context, event *model.EmailEvent) error {
	_, err := r.db.NewInsert().Model(event).Exec(ctx)
	return err
}

func (r *EmailEventRepository) GetByID(ctx context.Context, id int) (*model.EmailEvent, error) {
	event := new(model.EmailEvent)
	err := r.db.NewSelect().
		Model(event).
		Where("id = ?", id).
		Relation("Message").
		Relation("Template").
		Scan(ctx)
	return event, err
}

func (r *EmailEventRepository) Update(ctx context.Context, event *model.EmailEvent) error {
	_, err := r.db.NewUpdate().Model(event).Where("id = ?", event.ID).Exec(ctx)
	return err
}

func (r *EmailEventRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.NewDelete().Model(&model.EmailEvent{}).Where("id = ?", id).Exec(ctx)
	return err
}
