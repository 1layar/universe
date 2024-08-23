package repo

import (
	"context"

	"github.com/1layar/universe/pkg/email_service/model"

	"github.com/uptrace/bun"
)

type EmailTemplateRepository struct {
	db *bun.DB
}

func NewEmailTemplateRepository(db *bun.DB) *EmailTemplateRepository {
	return &EmailTemplateRepository{db: db}
}

func (r *EmailTemplateRepository) Create(ctx context.Context, template *model.EmailTemplate) error {
	_, err := r.db.NewInsert().Model(template).Exec(ctx)
	return err
}

func (r *EmailTemplateRepository) GetByID(ctx context.Context, id int) (*model.EmailTemplate, error) {
	template := new(model.EmailTemplate)
	err := r.db.NewSelect().Model(template).Where("id = ?", id).Scan(ctx)
	return template, err
}

// get email by code
func (r *EmailTemplateRepository) GetByCode(ctx context.Context, code string) (*model.EmailTemplate, error) {
	template := new(model.EmailTemplate)
	err := r.db.NewSelect().Model(template).Where("name = ?", code).Scan(ctx)
	return template, err
}

func (r *EmailTemplateRepository) Update(ctx context.Context, template *model.EmailTemplate) error {
	_, err := r.db.NewUpdate().Model(template).Where("id = ?", template.ID).Exec(ctx)
	return err
}

func (r *EmailTemplateRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.NewDelete().Model(&model.EmailTemplate{}).Where("id = ?", id).Exec(ctx)
	return err
}
