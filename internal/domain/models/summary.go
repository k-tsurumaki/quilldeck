package models

import (
	"time"

	"github.com/google/uuid"
)

type Summary struct {
	ID         uuid.UUID `json:"id"`
	DocumentID uuid.UUID `json:"document_id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewSummary(documentID uuid.UUID, content string) *Summary {
	return &Summary{
		ID:         uuid.New(),
		DocumentID: documentID,
		Content:    content,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func (s *Summary) Validate() error {
	if s.DocumentID == uuid.Nil {
		return &ValidationError{Field: "document_id", Message: "document_id is required"}
	}
	if s.Content == "" {
		return &ValidationError{Field: "content", Message: "content is required"}
	}
	return nil
}

func (s *Summary) UpdateContent(content string) {
	s.Content = content
	s.UpdatedAt = time.Now()
}