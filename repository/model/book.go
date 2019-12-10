package model

// Book :nodoc:
type Book struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Publisher  string `json:"publisher"`
	IsIssued   bool   `json:"is_issued"`
	IssueCount int64  `json:"issue_count"`
}

// Issue :nodoc:
type Issue struct {
	ID         string `json:"id"`
	IssuedBook string `json:"issued_book"`
	StartDate  string `json:"start_date"`
	FinishDate string `json:"finish_date"`
	IssuedBy   string `json:"issued_by"`
}
