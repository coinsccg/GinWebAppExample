package controllers

import (
	"tanjunchen.io.webapp/dao/mysql"
	"tanjunchen.io.webapp/models"
	"tanjunchen.io.webapp/pkg/myJWT"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	// 1. 提取用户提交的注册信息
	// 2. 并进行数据校验
	var sd models.SignUpForm
	if err := c.ShouldBindJSON(&sd); err != nil {
		// 返回参数错误的响应
		zap.L().Error("invalid params", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	// 3. 保存到数据库
	u := &models.User{
		UserName: sd.UserName,
		Password: sd.Password,
	}
	if err := mysql.Register(u); err != nil {
		zap.L().Error("sign up user failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 4. 返回响应
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	// 1. 获取登录请求携带的用户名、密码数据 -> 获取请求的数据
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	// 2. 去数据库校验用户名和密码是否正确
	if err := mysql.Login(&u); err != nil {
		ResponseError(c, CodeInvalidPassword)
		return
	}
	// 3. 生成JWT Token
	tokenString, err := myJWT.GenToken(u.UserID)
	if err != nil {
		zap.L().Error("myJWT.GenToken failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 如果是基于Cookie/Session的模式
	// 生成Cookie值，并生成后端的session数据
	// Cookie值返回给客户端种到浏览器的Cookie中
	// 4. 返回响应
	ResponseSuccess(c, gin.H{
		"token":    tokenString,
		"userID":   u.UserID,
		"username": u.UserName,
	})
}
