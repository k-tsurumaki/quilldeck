package service

import (
	"context"
	"strings"

	"github/k-tsurumaki/quilldeck/internal/domain/models"
	"github/k-tsurumaki/quilldeck/internal/domain/repository"
	"github/k-tsurumaki/quilldeck/internal/pkg/errors"
	"github.com/google/uuid"
)

type DocumentService struct {
	docRepo     repository.DocumentRepository
	summaryRepo repository.SummaryRepository
}

func NewDocumentService(docRepo repository.DocumentRepository, summaryRepo repository.SummaryRepository) *DocumentService {
	return &DocumentService{
		docRepo:     docRepo,
		summaryRepo: summaryRepo,
	}
}

func (s *DocumentService) UploadDocument(ctx context.Context, userID uuid.UUID, title, content string, docType models.DocumentType) (*models.Document, error) {
	document := models.NewDocument(userID, title, content, docType, int64(len(content)))

	if err := document.Validate(); err != nil {
		return nil, errors.Wrap(err, errors.ErrCodeValidation, "invalid document data")
	}

	if err := s.docRepo.Create(ctx, document); err != nil {
		return nil, errors.Wrap(err, errors.ErrCodeInternal, "failed to create document")
	}

	return document, nil
}

func (s *DocumentService) GetDocument(ctx context.Context, documentID uuid.UUID) (*models.Document, error) {
	document, err := s.docRepo.GetByID(ctx, documentID)
	if err != nil {
		return nil, errors.Wrap(err, errors.ErrCodeNotFound, "document not found")
	}
	return document, nil
}

func (s *DocumentService) GetUserDocuments(ctx context.Context, userID uuid.UUID) ([]*models.Document, error) {
	documents, err := s.docRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, errors.ErrCodeInternal, "failed to get user documents")
	}
	return documents, nil
}

func (s *DocumentService) GenerateSummary(ctx context.Context, documentID uuid.UUID, length models.SummaryLength) (*models.Summary, error) {
	document, err := s.docRepo.GetByID(ctx, documentID)
	if err != nil {
		return nil, errors.Wrap(err, errors.ErrCodeNotFound, "document not found")
	}

	// 簡単な要約生成（実際のAI連携は後で実装）
	summaryContent := s.generateSimpleSummary(document.Content, length)
	keywords := s.extractKeywords(document.Content)

	summary := models.NewSummary(documentID, summaryContent, length, keywords)

	if err := summary.Validate(); err != nil {
		return nil, errors.Wrap(err, errors.ErrCodeValidation, "invalid summary data")
	}

	if err := s.summaryRepo.Create(ctx, summary); err != nil {
		return nil, errors.Wrap(err, errors.ErrCodeInternal, "failed to create summary")
	}

	// ドキュメントを処理済みとしてマーク
	document.MarkProcessed()
	if err := s.docRepo.Update(ctx, document); err != nil {
		return nil, errors.Wrap(err, errors.ErrCodeInternal, "failed to update document")
	}

	return summary, nil
}

func (s *DocumentService) generateSimpleSummary(content string, length models.SummaryLength) string {
	sentences := strings.Split(content, ".")
	var maxSentences int

	switch length {
	case models.SummaryLengthShort:
		maxSentences = 2
	case models.SummaryLengthMedium:
		maxSentences = 5
	case models.SummaryLengthLong:
		maxSentences = 10
	default:
		maxSentences = 5
	}

	if len(sentences) <= maxSentences {
		return content
	}

	return strings.Join(sentences[:maxSentences], ".") + "."
}

func (s *DocumentService) extractKeywords(content string) []string {
	words := strings.Fields(strings.ToLower(content))
	keywordMap := make(map[string]int)

	for _, word := range words {
		if len(word) > 3 {
			keywordMap[word]++
		}
	}

	var keywords []string
	for word, count := range keywordMap {
		if count > 1 {
			keywords = append(keywords, word)
		}
	}

	if len(keywords) > 5 {
		keywords = keywords[:5]
	}

	return keywords
}