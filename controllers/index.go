package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func IndexHandler(c *gin.Context) {
	zap.L().Error("this is a error")
	zap.L().Debug("this is index handler")
	zap.L().Info("this is a test log")
	c.String(http.StatusOK, viper.GetString("app.ver"))
}

func NeedLoginHandler(c *gin.Context) {
	// 因为之前的JWT中间件中在 上下文c 中保存了 当前登录的userID
	// 我在这个函数中就可以通过 上下文c 取获取当前登录的用户
	v, ok := c.Get(ContextUserIDKey)
	if !ok {
		// 没有取到userID， ---> /login
		ResponseError(c, CodeNotLogin)
		return
	}
	userID := v.(uint64) // 因为c中取出来的值是空接口类型，需要做类型断言
	ResponseSuccess(c, userID)
}
