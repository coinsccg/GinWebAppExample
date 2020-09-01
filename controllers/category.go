package controllers

import (
	"tanjunchen.io.webapp/dao/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CategoryListHandler(c *gin.Context) {
	data, err := mysql.GetCategoryList()
	if err != nil {
		zap.L().Error("mysql.GetCategoryList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	zap.L().Debug("get category list success", zap.Any("data", data))
	ResponseSuccess(c, data)
}
