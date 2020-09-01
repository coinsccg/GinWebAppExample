package controllers

import (
	"strings"

	"tanjunchen.io.webapp/pkg/myJWT"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	ContextUserIDKey = "userID"
)

// 基于JWT实现的登录认证中间件
// 对于需要登陆才能访问的API来说
// 该中间件需要从请求头中获取JWT Token
// 如果没有Token --> /login
// 如果Token过期 --> /login
// 从JWT中解析我们需要的UserID字段  --> 根据userID我们就能从数据库查询到当前请求的用户是谁

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的中 Authorization: Bearer token_string
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			ResponseErrorWithMsg(c, CodeInvalidToken, "请求头缺少Auth Token")
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ResponseErrorWithMsg(c, CodeInvalidToken, "请求头中auth格式有误")
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := myJWT.ParseToken(parts[1])
			if err != nil {
			ResponseError(c, CodeInvalidToken)
			zap.L().Warn("invalid JWT token", zap.Error(err))
			c.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set(ContextUserIDKey, mc.UserID)
		c.Next()
		// 后续的处理函数可以用过c.Get("userID")来获取当前请求的用户信息
		// 返回响应的时候可以做Token/Cookie续期
	}
}

// 基于Cookie和Session认证的中间件
// 对于需要登陆才能访问的API来说
// 该中间件需要从请求中获取Cookie值
// 如果没有Cookie --> /login
// 拿到Cookie值取session数据中找对应的数据，找不到(1.session过期了2.无效的cookie值) --> /login
// session值 也可以通过 c.Set() 直接赋值到 上下文c 上
