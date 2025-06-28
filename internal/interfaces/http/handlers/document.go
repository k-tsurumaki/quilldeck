package handlers

import (
	"io"
	"net/http"
	"strings"

	"github.com/k-tsurumaki/fuselage"
	"github/k-tsurumaki/quilldeck/internal/domain/models"
	"github/k-tsurumaki/quilldeck/internal/domain/service"
	"github.com/google/uuid"
)

type DocumentHandler struct {
	docService *service.DocumentService
}

func NewDocumentHandler(docService *service.DocumentService) *DocumentHandler {
	return &DocumentHandler{docService: docService}
}

type UploadResponse struct {
	Message    string `json:"message"`
	DocumentID string `json:"document_id"`
}

type SummaryRequest struct {
	DocumentID string                `json:"document_id"`
	Length     models.SummaryLength `json:"length"`
}

type SummaryResponse struct {
	Message   string   `json:"message"`
	SummaryID string   `json:"summary_id"`
	Content   string   `json:"content"`
	Keywords  []string `json:"keywords"`
}

func (h *DocumentHandler) Upload(c *fuselage.Context) error {
	// TODO: 実際の認証実装後にユーザーIDを取得
	userID := uuid.New()

	err := c.Request.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to parse form"})
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "No file uploaded"})
	}
	defer file.Close()

	// ファイル拡張子チェック
	filename := header.Filename
	var docType models.DocumentType
	if strings.HasSuffix(strings.ToLower(filename), ".txt") {
		docType = models.DocumentTypeTXT
	} else if strings.HasSuffix(strings.ToLower(filename), ".md") {
		docType = models.DocumentTypeMD
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Only .txt and .md files are supported"})
	}

	// ファイル内容読み取り
	content, err := io.ReadAll(file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to read file"})
	}

	// ドキュメント作成
	document, err := h.docService.UploadDocument(c.Request.Context(), userID, filename, string(content), docType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, UploadResponse{
		Message:    "File uploaded successfully",
		DocumentID: document.ID.String(),
	})
}

func (h *DocumentHandler) GenerateSummary(c *fuselage.Context) error {
	var req SummaryRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	documentID, err := uuid.Parse(req.DocumentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid document ID"})
	}

	summary, err := h.docService.GenerateSummary(c.Request.Context(), documentID, req.Length)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, SummaryResponse{
		Message:   "Summary generated successfully",
		SummaryID: summary.ID.String(),
		Content:   summary.Content,
		Keywords:  summary.Keywords,
	})
}