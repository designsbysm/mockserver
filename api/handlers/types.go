package handlers

type Response struct {
	Success bool        `json:"success"`
	Results interface{} `json:"results,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
}

type ErrorInfo struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}
