package note

type TextNote struct {
	Note      string `json:"note"`
	Timestamp string `json:"timestamp"`
	Bucket    string `json:"bucket"`
}
