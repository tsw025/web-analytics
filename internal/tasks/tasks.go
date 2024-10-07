package tasks

import "encoding/json"

// Task Types
const (
	TypeAnalyzeWebsite = "analyze:website"
)

// AnalyzeWebsitePayload defines the payload for the AnalyzeWebsite task
type AnalyzeWebsitePayload struct {
	URL         string `json:"url"`
	AnalyticsID uint   `json:"analytics_id"`
}

// Marshal to JSON
func (p *AnalyzeWebsitePayload) Marshal() ([]byte, error) {
	return json.Marshal(p)
}

// Unmarshal from JSON
func (p *AnalyzeWebsitePayload) Unmarshal(b []byte) error {
	return json.Unmarshal(b, p)
}
