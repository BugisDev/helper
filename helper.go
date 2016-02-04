package helper

// ErrorMessage struct
type ErrorMessage struct {
	Code    int          `json:"code"`
	Source  SourceErrors `json:"source,omitempty"`
	Title   string       `json:"title"`
	Details string       `json:"details,omitempty"`
}

// SourceErrors Struct
type SourceErrors struct {
	Pointer string `json:"pointer,omitempty"`
}

// Response Struct Response
type Response struct {
	Data     Data     `json:"data"`
	Included Included `json:"included"`
}

// Data Struct Response
// Part from Response
type Data struct {
	ID            int         `json:"id"`
	Type          string      `json:"type"`
	Attributes    interface{} `json:"attributes"`
	Relationships interface{} `json:"relationships"`
}

// Included Struct Response
// Part from Response
type Included struct {
	ID         int         `json:"id"`
	Type       string      `json:"type"`
	Attributes interface{} `json:"attributes"`
}
