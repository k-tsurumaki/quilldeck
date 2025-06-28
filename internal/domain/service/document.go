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
	// TODO: AI連携による実際の要約生成を実装
	// 現在は入力の一文目を返すサンプル実装
	content = strings.TrimSpace(content)
	
	// 日本語の句点で分割
	if strings.Contains(content, "。") {
		sentences := strings.Split(content, "。")
		if len(sentences) > 0 && strings.TrimSpace(sentences[0]) != "" {
			return strings.TrimSpace(sentences[0]) + "。"
		}
	}
	
	// 英語のピリオドで分割（改行を除去してから処理）
	if strings.Contains(content, ".") {
		cleanContent := strings.ReplaceAll(content, "\n", " ")
		sentences := strings.Split(cleanContent, ".")
		if len(sentences) > 0 && strings.TrimSpace(sentences[0]) != "" {
			return strings.TrimSpace(sentences[0]) + "."
		}
	}
	
	// 改行で分割して最初の行を返す
	lines := strings.Split(content, "\n")
	if len(lines) > 0 && strings.TrimSpace(lines[0]) != "" {
		return strings.TrimSpace(lines[0])
	}
	
	return "No content to summarize."
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