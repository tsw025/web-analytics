package services

import (
	"encoding/json"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/tsw025/web_analytics/internal/repositories"
	"io"
	"net/http"
	"strings"
	"time"
)

type WorkerService interface {
	PerformAnalysis(urlStr string) (string, error)
}

type workerService struct {
	analyticsRepo repositories.AnalyticsRepository
}

func NewWorkerService(analyticsRepo repositories.AnalyticsRepository) *workerService {
	return &workerService{
		analyticsRepo: analyticsRepo,
	}
}

func (s *workerService) PerformAnalysis(urlStr string) (string, error) {
	resp, err := http.Get(urlStr)
	if err != nil {
		return "", errors.New("failed to fetch the webpage content")
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return "", errors.New("webpage returned a non-success status code")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("failed to read webpage content")
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		return "", errors.New("failed to parse HTML content")
	}

	htmlVersion := extractHTMLVersion(doc)
	pageTitle := extractPageTitle(doc)
	headings := countHeadings(doc)
	linksInfo, err := analyzeLinks(doc, urlStr)
	if err != nil {
		return "", err
	}
	hasLoginForm := detectLoginForm(doc)

	result := map[string]interface{}{
		"html_version":          htmlVersion,
		"page_title":            pageTitle,
		"headings_count":        headings,
		"links":                 linksInfo,
		"contains_login_form":   hasLoginForm,
		"analysis_completed_at": time.Now().Format(time.RFC3339),
	}

	resultJSON, err := json.Marshal(result)
	if err != nil {
		return "", errors.New("failed to marshal analysis result")
	}

	return string(resultJSON), nil
}
