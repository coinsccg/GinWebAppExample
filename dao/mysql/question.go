package mysql

import (
	"tanjunchen.io.webapp/models"

	"go.uber.org/zap"
)

func CreateQuestion(question *models.Question) (err error) {
	sqlStr := `insert into question(question_id,  caption, content, author_id, category_id) values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, question.QuestionID, question.Caption,
		question.Content, question.AuthorID, question.CategoryID)
	if err != nil {
		zap.L().Error("create question failed", zap.Any("question", question), zap.Error(err))
		return
	}
	return
}

func GetQuestionList(offset, limit int) (questionList []*models.Question, err error) {
	sqlStr := `select 
						question_id, caption, content, author_id, category_id, create_time
					from 
						question order by id desc
						limit ?, ?`
	err = db.Select(&questionList, sqlStr, offset, limit)
	if err != nil {
		zap.L().Error("get question list failed", zap.Error(err))
		return
	}
	return
}

func GetQuestionCount() (answerCount int, err error) {

	sqlStr := `select 
							count(question_id)
						from 
							question`
	err = db.Get(&answerCount, sqlStr)
	if err != nil {
		zap.L().Error("GetQuestionCount failed", zap.Error(err))
		return
	}
	return
}

func CheckQuestionIDExist(questionID uint64) bool {
	sqlStr := `select count(id) from question where question_id=?`
	count := 0
	db.Get(&count, sqlStr, questionID)
	return count > 0
}

func GetQuestion(questionID uint64) (question *models.Question, err error) {
	question = &models.Question{}
	sqlStr := `select 
							question_id, caption, content, author_id, category_id, create_time
						from 
							question
						where question_id=?`
	err = db.Get(question, sqlStr, questionID)
	if err != nil {
		zap.L().Error("get question failed", zap.Error(err))
		return
	}
	return
}
