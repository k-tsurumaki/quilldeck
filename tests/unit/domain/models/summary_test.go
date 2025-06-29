package models

import (
	"testing"

	"github/k-tsurumaki/quilldeck/internal/domain/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewSummary(t *testing.T) {
	documentID := uuid.New()
	content := "This is a summary"

	summary := models.NewSummary(documentID, content)

	assert.NotNil(t, summary)
	assert.Equal(t, documentID, summary.DocumentID)
	assert.Equal(t, content, summary.Content)
	assert.NotEqual(t, "", summary.ID.String())
	assert.False(t, summary.CreatedAt.IsZero())
	assert.False(t, summary.UpdatedAt.IsZero())
}

func TestSummary_Validate(t *testing.T) {
	documentID := uuid.New()

	tests := []struct {
		name    string
		summary *models.Summary
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid summary",
			summary: models.NewSummary(documentID, "Content"),
			wantErr: false,
		},
		{
			name: "empty document_id",
			summary: &models.Summary{
				DocumentID: uuid.Nil,
				Content:    "Content",
			},
			wantErr: true,
			errMsg:  "document_id is required",
		},
		{
			name: "empty content",
			summary: &models.Summary{
				DocumentID: documentID,
				Content:    "",
			},
			wantErr: true,
			errMsg:  "content is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.summary.Validate()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSummary_UpdateContent(t *testing.T) {
	summary := models.NewSummary(uuid.New(), "Original content")
	originalUpdatedAt := summary.UpdatedAt

	newContent := "Updated content"
	summary.UpdateContent(newContent)

	assert.Equal(t, newContent, summary.Content)
	assert.True(t, summary.UpdatedAt.After(originalUpdatedAt))
}