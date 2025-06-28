package models

import (
	"testing"

	"github/k-tsurumaki/quilldeck/internal/domain/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewDocument(t *testing.T) {
	userID := uuid.New()
	title := "Test Document"
	content := "This is test content"
	docType := models.DocumentTypeTXT
	size := int64(len(content))

	doc := models.NewDocument(userID, title, content, docType, size)

	assert.NotNil(t, doc)
	assert.Equal(t, userID, doc.UserID)
	assert.Equal(t, title, doc.Title)
	assert.Equal(t, content, doc.Content)
	assert.Equal(t, docType, doc.Type)
	assert.Equal(t, size, doc.Size)
	assert.NotEqual(t, "", doc.ID.String())
	assert.False(t, doc.UploadedAt.IsZero())
	assert.Nil(t, doc.ProcessedAt)
}

func TestDocument_Validate(t *testing.T) {
	userID := uuid.New()

	tests := []struct {
		name    string
		doc     *models.Document
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid document",
			doc:  models.NewDocument(userID, "Test", "Content", models.DocumentTypeTXT, 7),
			wantErr: false,
		},
		{
			name: "empty user_id",
			doc: &models.Document{
				UserID:  uuid.Nil,
				Title:   "Test",
				Content: "Content",
				Type:    models.DocumentTypeTXT,
			},
			wantErr: true,
			errMsg:  "user_id is required",
		},
		{
			name: "empty title",
			doc: &models.Document{
				UserID:  userID,
				Title:   "",
				Content: "Content",
				Type:    models.DocumentTypeTXT,
			},
			wantErr: true,
			errMsg:  "title is required",
		},
		{
			name: "empty content",
			doc: &models.Document{
				UserID:  userID,
				Title:   "Test",
				Content: "",
				Type:    models.DocumentTypeTXT,
			},
			wantErr: true,
			errMsg:  "content is required",
		},
		{
			name: "invalid type",
			doc: &models.Document{
				UserID:  userID,
				Title:   "Test",
				Content: "Content",
				Type:    "invalid",
			},
			wantErr: true,
			errMsg:  "invalid document type",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.doc.Validate()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDocument_MarkProcessed(t *testing.T) {
	doc := models.NewDocument(uuid.New(), "Test", "Content", models.DocumentTypeTXT, 7)
	
	assert.False(t, doc.IsProcessed())
	assert.Nil(t, doc.ProcessedAt)

	doc.MarkProcessed()

	assert.True(t, doc.IsProcessed())
	assert.NotNil(t, doc.ProcessedAt)
	assert.False(t, doc.ProcessedAt.IsZero())
}