package repository

import (
	"context"

	"github/k-tsurumaki/quilldeck/internal/domain/models"
	"github.com/google/uuid"
)

type DocumentRepository interface {
	Create(ctx context.Context, document *models.Document) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Document, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*models.Document, error)
	Update(ctx context.Context, document *models.Document) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type SummaryRepository interface {
	Create(ctx context.Context, summary *models.Summary) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Summary, error)
	GetByDocumentID(ctx context.Context, documentID uuid.UUID) ([]*models.Summary, error)
	Update(ctx context.Context, summary *models.Summary) error
	Delete(ctx context.Context, id uuid.UUID) error
}