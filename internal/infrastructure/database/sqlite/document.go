package sqlite

import (
	"context"
	"database/sql"
	"strings"

	"github/k-tsurumaki/quilldeck/internal/domain/models"
	"github.com/google/uuid"
)

type DocumentRepository struct {
	db *DB
}

func NewDocumentRepository(db *DB) *DocumentRepository {
	return &DocumentRepository{db: db}
}

func (r *DocumentRepository) Create(ctx context.Context, document *models.Document) error {
	query := `
		INSERT INTO documents (id, user_id, title, content, type, size, uploaded_at, processed_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query,
		document.ID.String(),
		document.UserID.String(),
		document.Title,
		document.Content,
		string(document.Type),
		document.Size,
		document.UploadedAt,
		document.ProcessedAt,
	)
	return err
}

func (r *DocumentRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Document, error) {
	query := `SELECT id, user_id, title, content, type, size, uploaded_at, processed_at FROM documents WHERE id = ?`

	var document models.Document
	var idStr, userIDStr, typeStr string
	err := r.db.QueryRowContext(ctx, query, id.String()).Scan(
		&idStr,
		&userIDStr,
		&document.Title,
		&document.Content,
		&typeStr,
		&document.Size,
		&document.UploadedAt,
		&document.ProcessedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	document.ID = uuid.MustParse(idStr)
	document.UserID = uuid.MustParse(userIDStr)
	document.Type = models.DocumentType(typeStr)
	return &document, nil
}

func (r *DocumentRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*models.Document, error) {
	query := `SELECT id, user_id, title, content, type, size, uploaded_at, processed_at FROM documents WHERE user_id = ?`

	rows, err := r.db.QueryContext(ctx, query, userID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var documents []*models.Document
	for rows.Next() {
		var document models.Document
		var idStr, userIDStr, typeStr string
		err := rows.Scan(
			&idStr,
			&userIDStr,
			&document.Title,
			&document.Content,
			&typeStr,
			&document.Size,
			&document.UploadedAt,
			&document.ProcessedAt,
		)
		if err != nil {
			return nil, err
		}

		document.ID = uuid.MustParse(idStr)
		document.UserID = uuid.MustParse(userIDStr)
		document.Type = models.DocumentType(typeStr)
		documents = append(documents, &document)
	}

	return documents, nil
}

func (r *DocumentRepository) Update(ctx context.Context, document *models.Document) error {
	query := `UPDATE documents SET title = ?, content = ?, processed_at = ? WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query,
		document.Title,
		document.Content,
		document.ProcessedAt,
		document.ID.String(),
	)
	return err
}

func (r *DocumentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM documents WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id.String())
	return err
}

type SummaryRepository struct {
	db *DB
}

func NewSummaryRepository(db *DB) *SummaryRepository {
	return &SummaryRepository{db: db}
}

func (r *SummaryRepository) Create(ctx context.Context, summary *models.Summary) error {
	keywords := strings.Join(summary.Keywords, ",")
	query := `
		INSERT INTO summaries (id, document_id, content, length, keywords, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query,
		summary.ID.String(),
		summary.DocumentID.String(),
		summary.Content,
		string(summary.Length),
		keywords,
		summary.CreatedAt,
		summary.UpdatedAt,
	)
	return err
}

func (r *SummaryRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Summary, error) {
	query := `SELECT id, document_id, content, length, keywords, created_at, updated_at FROM summaries WHERE id = ?`

	var summary models.Summary
	var idStr, docIDStr, lengthStr, keywordsStr string
	err := r.db.QueryRowContext(ctx, query, id.String()).Scan(
		&idStr,
		&docIDStr,
		&summary.Content,
		&lengthStr,
		&keywordsStr,
		&summary.CreatedAt,
		&summary.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	summary.ID = uuid.MustParse(idStr)
	summary.DocumentID = uuid.MustParse(docIDStr)
	summary.Length = models.SummaryLength(lengthStr)
	if keywordsStr != "" {
		summary.Keywords = strings.Split(keywordsStr, ",")
	}
	return &summary, nil
}

func (r *SummaryRepository) GetByDocumentID(ctx context.Context, documentID uuid.UUID) ([]*models.Summary, error) {
	query := `SELECT id, document_id, content, length, keywords, created_at, updated_at FROM summaries WHERE document_id = ?`

	rows, err := r.db.QueryContext(ctx, query, documentID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var summaries []*models.Summary
	for rows.Next() {
		var summary models.Summary
		var idStr, docIDStr, lengthStr, keywordsStr string
		err := rows.Scan(
			&idStr,
			&docIDStr,
			&summary.Content,
			&lengthStr,
			&keywordsStr,
			&summary.CreatedAt,
			&summary.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		summary.ID = uuid.MustParse(idStr)
		summary.DocumentID = uuid.MustParse(docIDStr)
		summary.Length = models.SummaryLength(lengthStr)
		if keywordsStr != "" {
			summary.Keywords = strings.Split(keywordsStr, ",")
		}
		summaries = append(summaries, &summary)
	}

	return summaries, nil
}

func (r *SummaryRepository) Update(ctx context.Context, summary *models.Summary) error {
	keywords := strings.Join(summary.Keywords, ",")
	query := `UPDATE summaries SET content = ?, keywords = ?, updated_at = ? WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query,
		summary.Content,
		keywords,
		summary.UpdatedAt,
		summary.ID.String(),
	)
	return err
}

func (r *SummaryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM summaries WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id.String())
	return err
}