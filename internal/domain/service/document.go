package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github/k-tsurumaki/quilldeck/internal/domain/models"
	"github/k-tsurumaki/quilldeck/internal/domain/repository"
	"github/k-tsurumaki/quilldeck/internal/pkg/errors"

	"github.com/google/uuid"
)

type DocumentService struct {
	docRepo     repository.DocumentRepository
	summaryRepo repository.SummaryRepository
	llmClinet   *LLMClient
}

type LLMClient struct {
	apiKey  string
	baseURL string
	model   string
	client  *http.Client
}

type LLMRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type LLMResponse struct {
	ID      string `json:"id"`
	Model   string `json:"model"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Choices []struct {
		Index        int     `json:"index"`
		FinishReason string  `json:"finish_reason"`
		Message      Message `json:"message"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

// NewLLMClient creates a new instance of LLMClient with configuration from the global Config.
func NewLLMClient(apiKey, baseURL, model string) *LLMClient {
	return &LLMClient{
		apiKey:  apiKey,
		baseURL: baseURL,
		model:   model,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func NewDocumentService(docRepo repository.DocumentRepository, summaryRepo repository.SummaryRepository, llmAPIKey string, llmBaseURL string, llmModel string) *DocumentService {
	llmClient := NewLLMClient(llmAPIKey, llmBaseURL, llmModel)
	return &DocumentService{
		docRepo:     docRepo,
		summaryRepo: summaryRepo,
		llmClinet:   llmClient,
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

func (s *DocumentService) GenerateSummary(ctx context.Context, documentID uuid.UUID) (*models.Summary, error) {
	document, err := s.docRepo.GetByID(ctx, documentID)
	if err != nil {
		return nil, errors.Wrap(err, errors.ErrCodeNotFound, "document not found")
	}

	// Generate summary using LLM API
	summaryContent, err := s.getSummaryFromLLM(ctx, document.Content)
	if err != nil {
		return nil, errors.Wrap(err, errors.ErrCodeInternal, "failed to get summary from LLM")
	}

	summary := models.NewSummary(documentID, summaryContent)

	if err := summary.Validate(); err != nil {
		return nil, errors.Wrap(err, errors.ErrCodeValidation, "invalid summary data")
	}

	if err := s.summaryRepo.Create(ctx, summary); err != nil {
		return nil, errors.Wrap(err, errors.ErrCodeInternal, "failed to create summary")
	}

	// Mark document as processed
	document.MarkProcessed()
	if err := s.docRepo.Update(ctx, document); err != nil {
		return nil, errors.Wrap(err, errors.ErrCodeInternal, "failed to update document")
	}

	return summary, nil
}

func (s *DocumentService) getSummaryFromLLM(ctx context.Context, content string) (string, error) {

	prompt := fmt.Sprintf("Please summarize the following content in Japanese business style within 200 characters:\n\n%s", content)
	log.Printf("LLM Request Prompt: %s", prompt)

	resBody := LLMRequest{
		Model: s.llmClinet.model,
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	jsonData, err := json.Marshal(resBody)
	if err != nil {
		return "", errors.Wrap(err, errors.ErrCodeInternal, "failed to marshal LLM request")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", s.llmClinet.baseURL+"/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", errors.Wrap(err, errors.ErrCodeInternal, "failed to create LLM request")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.llmClinet.apiKey)

	resp, err := s.llmClinet.client.Do(req)
	if err != nil {
		return "", errors.Wrap(err, errors.ErrCodeInternal, "failed to call LLM API")
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, errors.ErrCodeInternal, "failed to read LLM response body")
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("LLM API returned error: %s", string(bodyBytes))
	}

	var llmResponse LLMResponse
	if err := json.Unmarshal(bodyBytes, &llmResponse); err != nil {
		return "", errors.Wrap(err, errors.ErrCodeInternal, "failed to unmarshal LLM response")
	}

	if len(llmResponse.Choices) == 0 {
		return "", fmt.Errorf("LLM API returned no choices in response")
	}

	return llmResponse.Choices[0].Message.Content, nil
}
