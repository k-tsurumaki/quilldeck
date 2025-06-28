package models

import (
	"time"

	"github.com/google/uuid"
)

type SummaryLength string

const (
	SummaryLengthShort  SummaryLength = "short"
	SummaryLengthMedium SummaryLength = "medium"
	SummaryLengthLong   SummaryLength = "long"
)

type Summary struct {
	ID         uuid.UUID     `json:"id"`
	DocumentID uuid.UUID     `json:"document_id"`
	Content    string        `json:"content"`
	Length     SummaryLength `json:"length"`
	Keywords   []string      `json:"keywords"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at"`
}

func NewSummary(documentID uuid.UUID, content string, length SummaryLength, keywords []string) *Summary {
	return &Summary{
		ID:         uuid.New(),
		DocumentID: documentID,
		Content:    content,
		Length:     length,
		Keywords:   keywords,
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
	if s.Length != SummaryLengthShort && s.Length != SummaryLengthMedium && s.Length != SummaryLengthLong {
		return &ValidationError{Field: "length", Message: "invalid summary length"}
	}
	return nil
}

func (s *Summary) UpdateContent(content string) {
	s.Content = content
	s.UpdatedAt = time.Now()
}