package controllers

import (
	"tanjunchen.io.webapp/logic"
	"tanjunchen.io.webapp/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AnswerListHandler(c *gin.Context) {
	questionID, err := getIDFromQuery(c, "question_id")
	if err != nil {
		zap.L().Error("invalid query param", zap.Uint64("questionID", questionID), zap.Error(err))
	}
	offset, limit := getOffsetAndLimit(c)
	zap.L().Debug("get answer list param success",
		zap.Uint64("questionID", questionID),
		zap.Int("offset", offset),
		zap.Int("limit", limit),
	)

	data, err := logic.GetAnswerList(questionID, offset, limit)
	if err != nil {
		zap.L().Error("logic.GetAnswerList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func AnswerCommitHandler(c *gin.Context) {

	var answer = new(models.Answer)
	err := c.BindJSON(answer)
	if err != nil {
		zap.L().Error("bind json failed", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	if answer.QuestionID == 0 || len(answer.Content) == 0 {
		ResponseError(c, CodeInvalidParams)
		return
	}

	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Warn("user not login", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}

	answer.AuthorID = userID

	if err := logic.CreateAnswer(answer); err != nil {
		zap.L().Error("logic.CreateAnswer()", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
