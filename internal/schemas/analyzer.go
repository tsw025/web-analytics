package schemas

// AnalyserRequest is the request schema for the analyze endpoint
type AnalyserRequest struct {
	URL string `json:"url" validate:"required,url"`
}
