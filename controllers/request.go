package controllers

import (
	"errors"
	"strconv"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

var (
	ErrorUserNotLogin = errors.New("当前用户未登录")
)

func getCurrentUserID(c *gin.Context) (userID uint64, err error) {
	_userID, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = _userID.(uint64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

func getOffsetAndLimit(c *gin.Context) (offset, limit int) {
	offsetStr := c.Query("page")
	limitStr := c.Query("limit")

	var err error
	offset, err = strconv.Atoi(offsetStr)
	if err != nil {
		zap.L().Warn("invalid offset", zap.String("offsetStr", offsetStr), zap.Error(err))
	}
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		zap.L().Warn("invalid limit", zap.String("limitStr", limitStr), zap.Error(err))
		limit = viper.GetInt("app.default_page_size")
	}
	return
}

func getIDFromQuery(c *gin.Context, queryKey string) (uint64, error) {
	idStr := c.Query(queryKey)
	return strconv.ParseUint(idStr, 10, 64)
}
