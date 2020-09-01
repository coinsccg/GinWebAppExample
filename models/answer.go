package models

import "time"

type Answer struct {
	AnswerID    uint64    `json:"answer_id,string" db:"answer_id"`
	AuthorID    uint64    `json:"author_id,string" db:"author_id"`
	QuestionID  uint64    `json:"question_id,string" db:"question_id"`
	Content     string    `json:"content" db:"content"`
	VoteUpCount int       `json:"vote_up_count" db:"vote_up_count"`
	Status      int8      `json:"status" db:"status"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
	UpdateTime  time.Time `json:"update_time" db:"update_time"`
}

type ApiAnswer struct {
	Answer
	AuthorName string `json:"author_name" db:"author_name"`
}

type ApiAnswerList struct {
	AnswerList []*ApiAnswer `json:"answer_list"`
	TotalCount int          `json:"total_count"`
}
