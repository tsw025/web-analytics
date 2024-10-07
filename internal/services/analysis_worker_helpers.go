package services

import (
	"errors"
	"golang.org/x/net/html"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func extractHTMLVersion(doc *goquery.Document) string {
	// Convert goquery document to html.Node
	node := doc.Nodes[0]

	// Traverse the nodes to find the DOCTYPE
	for n := node; n != nil; n = n.NextSibling {
		if n.Type == html.DoctypeNode {
			doctype := strings.ToLower(n.Data)
			if strings.Contains(doctype, "html") {
				return "HTML5"
			} else if strings.Contains(doctype, "xhtml") {
				return "XHTML"
			} else if strings.Contains(doctype, "transitional") {
				return "HTML 4.01 Transitional"
			} else if strings.Contains(doctype, "strict") {
				return "HTML 4.01 Strict"
			} else if strings.Contains(doctype, "frameset") {
				return "HTML 4.01 Frameset"
			}
		}
	}
	return "HTML5"
}

func extractPageTitle(doc *goquery.Document) string {
	title := doc.Find("title").First().Text()
	return strings.TrimSpace(title)
}

func countHeadings(doc *goquery.Document) map[string]int {
	headings := make(map[string]int)
	for i := 1; i <= 6; i++ {
		tag := "h" + string('0'+i) // nolint:all
		count := doc.Find(tag).Length()
		headings[tag] = count
	}
	return headings
}

type LinksInfo struct {
	InternalLinks     int `json:"internal_links"`
	ExternalLinks     int `json:"external_links"`
	InaccessibleLinks int `json:"inaccessible_links"`
}

func analyzeLinks(doc *goquery.Document, baseURL string) (LinksInfo, error) {
	var info LinksInfo
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Parse base URL to determine internal vs external links
	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return info, errors.New("invalid base URL")
	}
	baseHost := parsedBaseURL.Host

	// Find all <a> tags with href attribute and analyze links
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return
		}

		// Resolve relative URLs
		link, err := parsedBaseURL.Parse(strings.TrimSpace(href))
		if err != nil {
			// Skip malformed URLs
			return
		}

		// Determine if the link is internal or external
		if link.Host == baseHost || link.Host == "" {
			mu.Lock()
			info.InternalLinks++
			mu.Unlock()
		} else {
			mu.Lock()
			info.ExternalLinks++
			mu.Unlock()
		}

		// Check accessibility of the link
		wg.Add(1)
		go func(linkStr string) {
			defer wg.Done()
			accessible := checkLinkAccessibility(linkStr)
			if !accessible {
				mu.Lock()
				info.InaccessibleLinks++
				mu.Unlock()
			}
		}(link.String())
	})

	wg.Wait()

	return info, nil
}

func checkLinkAccessibility(link string) bool {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Head(link)
	if err != nil {
		return false
	}
	_ = resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return true
	}
	return false
}

func detectLoginForm(doc *goquery.Document) bool {
	// Look for forms containing password input fields
	hasPassword := false
	doc.Find("form").EachWithBreak(func(i int, s *goquery.Selection) bool {
		if s.Find("input[type='password']").Length() > 0 {
			hasPassword = true
			return false
		}
		return true
	})
	return hasPassword
}
