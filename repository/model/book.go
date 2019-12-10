package model

import "time"

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
	ID         string    `json:"id"`
	IssuedBook int64     `json:"issued_book"`
	StartTime  time.Time `json:"start_time"`
	FinishTime time.Time `json:"finish_time"`
	IssuedBy   int64     `json:"issued_by"`
}
