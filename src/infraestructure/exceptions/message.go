package exceptions

type Message struct {
	StatusCode int    `json:"statusCode"`
	Timestamp  string `json:"timestamp"`
	Path       string `json:"path"`
	Message    string `json:"message"`
}
