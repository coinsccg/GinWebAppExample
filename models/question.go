package models

import "time"

type Question struct {
	QuestionID uint64    `json:"question_id,string" db:"question_id"`
	AuthorID   uint64    `json:"author_id,string" db:"author_id"`
	CategoryID int32     `json:"category_id" db:"category_id"`
	Status     int8      `json:"status" db:"status"`
	Caption    string    `json:"caption" db:"caption"`
	Content    string    `json:"content" db:"content"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}

type ApiQuestion struct {
	Question
	AuthorName string `json:"author_name"`
}

type ApiQuestionDetail struct {
	Question
	AuthorName   string `json:"author_name"`
	CategoryName string `json:"category_name"`
}

type ApiQuestionList struct {
	QuestionList []*ApiQuestion `json:"question_list"`
	TotalCount   int            `json:"total_count"`
}
