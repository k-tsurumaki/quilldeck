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
	length := models.SummaryLengthMedium
	keywords := []string{"test", "summary"}

	summary := models.NewSummary(documentID, content, length, keywords)

	assert.NotNil(t, summary)
	assert.Equal(t, documentID, summary.DocumentID)
	assert.Equal(t, content, summary.Content)
	assert.Equal(t, length, summary.Length)
	assert.Equal(t, keywords, summary.Keywords)
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
			name: "valid summary",
			summary: models.NewSummary(documentID, "Content", models.SummaryLengthMedium, []string{"test"}),
			wantErr: false,
		},
		{
			name: "empty document_id",
			summary: &models.Summary{
				DocumentID: uuid.Nil,
				Content:    "Content",
				Length:     models.SummaryLengthMedium,
			},
			wantErr: true,
			errMsg:  "document_id is required",
		},
		{
			name: "empty content",
			summary: &models.Summary{
				DocumentID: documentID,
				Content:    "",
				Length:     models.SummaryLengthMedium,
			},
			wantErr: true,
			errMsg:  "content is required",
		},
		{
			name: "invalid length",
			summary: &models.Summary{
				DocumentID: documentID,
				Content:    "Content",
				Length:     "invalid",
			},
			wantErr: true,
			errMsg:  "invalid summary length",
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
	summary := models.NewSummary(uuid.New(), "Original content", models.SummaryLengthMedium, []string{"test"})
	originalUpdatedAt := summary.UpdatedAt

	newContent := "Updated content"
	summary.UpdateContent(newContent)

	assert.Equal(t, newContent, summary.Content)
	assert.True(t, summary.UpdatedAt.After(originalUpdatedAt))
}