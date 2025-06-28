package models

import (
	"time"

	"github.com/google/uuid"
)

type DocumentType string

const (
	DocumentTypeTXT DocumentType = "txt"
	DocumentTypeMD  DocumentType = "md"
)

type Document struct {
	ID          uuid.UUID    `json:"id"`
	UserID      uuid.UUID    `json:"user_id"`
	Title       string       `json:"title"`
	Content     string       `json:"content"`
	Type        DocumentType `json:"type"`
	Size        int64        `json:"size"`
	UploadedAt  time.Time    `json:"uploaded_at"`
	ProcessedAt *time.Time   `json:"processed_at,omitempty"`
}

func NewDocument(userID uuid.UUID, title, content string, docType DocumentType, size int64) *Document {
	return &Document{
		ID:         uuid.New(),
		UserID:     userID,
		Title:      title,
		Content:    content,
		Type:       docType,
		Size:       size,
		UploadedAt: time.Now(),
	}
}

func (d *Document) Validate() error {
	if d.UserID == uuid.Nil {
		return &ValidationError{Field: "user_id", Message: "user_id is required"}
	}
	if d.Title == "" {
		return &ValidationError{Field: "title", Message: "title is required"}
	}
	if d.Content == "" {
		return &ValidationError{Field: "content", Message: "content is required"}
	}
	if d.Type != DocumentTypeTXT && d.Type != DocumentTypeMD {
		return &ValidationError{Field: "type", Message: "invalid document type"}
	}
	return nil
}

func (d *Document) MarkProcessed() {
	now := time.Now()
	d.ProcessedAt = &now
}

func (d *Document) IsProcessed() bool {
	return d.ProcessedAt != nil
}