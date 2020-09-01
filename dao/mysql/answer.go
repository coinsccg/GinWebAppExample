package mysql

import (
	"tanjunchen.io.webapp/models"

	"go.uber.org/zap"
)

const (
	createAnswer = `insert into answer(answer_id, content, author_id, question_id) values(?,?,?,?)`
)

func CreateAnswer(answer *models.Answer) (err error) {

	_, err = db.Exec(createAnswer, answer.AnswerID, answer.Content, answer.AuthorID, answer.QuestionID)
	if err != nil {
		zap.L().Error("create answer failed", zap.Error(err))
		return
	}
	return
}

func GetAnswerList(questionID uint64, offset, limit int) (answerList []*models.Answer, err error) {

	sqlStr := `select 
					answer_id, content,
					vote_up_count, author_id, status,
					create_time, update_time
				 from 
				 	answer
				where question_id=? order by id desc
				limit ?, ?`

	err = db.Select(&answerList, sqlStr, questionID, offset, limit)
	if err != nil {
		zap.L().Error("GetAnswerList failed", zap.Error(err))
		return
	}
	return
}

func GetAnswerCount(questionID uint64) (answerCount int, err error) {

	sqlStr := `select 
							count(answer_id)
						from 
							answer
						where question_id=?`

	err = db.Get(&answerCount, sqlStr, questionID)
	if err != nil {
		zap.L().Error("GetAnswerCount failed", zap.Error(err))
		return
	}
	return
}
