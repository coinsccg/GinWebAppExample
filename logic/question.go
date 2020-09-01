package logic

import (
	"strconv"

	"tanjunchen.io.webapp/dao/mysql"
	"tanjunchen.io.webapp/models"
	"tanjunchen.io.webapp/pkg/snowflake"

	"go.uber.org/zap"
)

// CreateQuestion 创建问题
func CreateQuestion(question *models.Question) (err error) {
	//- 防 XSS 攻击  --> `html/template`内容检测和转义
	//- 敏感词过滤  --> 涉黄、涉恐等等屏蔽词
	qid, err := snowflake.GenID()
	if err != nil {
		zap.L().Error("generate question id failed", zap.Error(err))
		return
	}

	question.QuestionID = qid

	err = mysql.CreateQuestion(question)
	if err != nil {
		zap.L().Error("mysql.CreateQuestion failed", zap.Error(err))
		return
	}
	return
}

func GetQuestionList(offset, limit int) (apiQuestionList *models.ApiQuestionList, err error) {
	questionList, err := mysql.GetQuestionList(offset, limit)
	if err != nil {
		zap.L().Error("mysql.GetQuestionList failed", zap.Error(err))
		return
	}
	if len(questionList) == 0 {
		zap.L().Info("question list is null")
		return
	}

	var userIDList []uint64
	for _, v := range questionList {
		userIDList = append(userIDList, v.AuthorID)
	}
	userInfoList, err := mysql.GetUserInfoList(userIDList)
	if err != nil {
		zap.L().Error("mysql.GetUserInfoList failed",
			zap.Any("userIDList", userIDList),
			zap.Error(err),
		)
		return
	}
	zap.L().Debug("GetQuestionList", zap.Any("questionList", questionList), zap.Any("userInfoList", userInfoList))
	// 涉及到大量数据的 slice，如果知道数据的数量，应当使用 make 预先初始化
	// 比 append() 扩容效率更高
	apiQuestionList = &models.ApiQuestionList{
		// QuestionList: make([]*models.ApiQuestion, len(questionList)),
		QuestionList: make([]*models.ApiQuestion, 0, len(questionList)),
	}

	for _, v := range questionList {
		apiQuestion := &models.ApiQuestion{
			Question: *v,
		}
		for _, user := range userInfoList {
			if user.UserID == v.AuthorID {
				apiQuestion.AuthorName = user.UserName
				break
			}
		}
		apiQuestionList.QuestionList = append(apiQuestionList.QuestionList, apiQuestion)
	}
	count, err := mysql.GetQuestionCount()
	if err != nil {
		zap.L().Error("mysql.GetQuestionCount failed",
			zap.Error(err),
		)
		return
	}

	apiQuestionList.TotalCount = count
	return
}

// GetQuestionDetail 获取问题详情
func GetQuestionDetail(questionID uint64) (data *models.ApiQuestionDetail, err error) {
	var question *models.Question
	question, err = mysql.GetQuestion(questionID)
	if err != nil {
		zap.L().Error("get question failed", zap.Error(err))
		return
	}

	categoryMap, err := mysql.MGetCategory([]int32{question.CategoryID})
	if err != nil {
		zap.L().Error("get category failed", zap.Error(err))
		return
	}

	category, ok := categoryMap[question.CategoryID]
	if !ok {
		zap.L().Error("get category failed", zap.Error(err))
		return
	}

	user, err := mysql.GetUserByID(strconv.Itoa(int(question.AuthorID)))
	if err != nil {
		zap.L().Error("get user info failed", zap.Error(err))
		return
	}

	apiQuestionDetail := &models.ApiQuestionDetail{}
	apiQuestionDetail.Question = *question
	apiQuestionDetail.AuthorName = user.UserName
	apiQuestionDetail.CategoryName = category.CategoryName
	return apiQuestionDetail, nil
}
