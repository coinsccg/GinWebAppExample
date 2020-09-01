package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"

	"github.com/jmoiron/sqlx"
	"tanjunchen.io.webapp/models"
	"tanjunchen.io.webapp/pkg/snowflake"

	"go.uber.org/zap"
)

// 存放数据库相关的操作，增删改查

var secret = "夏天夏天悄悄过去"

// encryptPassword 加密
func encryptPassword(data []byte) (result string) {
	h := md5.New()
	h.Write([]byte(secret))                // 先加盐
	return hex.EncodeToString(h.Sum(data)) // 再md5并转为十六进制字符串
}

// Register 将指定用户注册
func Register(user *models.User) (err error) {
	// 1. 先判断当前用户名是否已经被注册过
	if CheckUserIsExist(user.UserName) {
		return ErrorUserExit
	}
	// 2. 明文的密码要加盐加密处理处理才能入库
	password := encryptPassword([]byte(user.Password))
	// user_id
	userID, err := snowflake.GenID()
	if err != nil {
		zap.L().Error("snowflake.GenID failed", zap.Error(err))
		return err
	}
	// 3. 入库
	sqlStr := `insert into user (user_id, username, password) values (?, ?,?)`
	_, err = db.Exec(sqlStr, userID, user.UserName, password)
	return
	// 4. 返回
}

func CheckUserIsExist(username string) bool {
	sqlStr := "select count(user_id) from user where username = ?"
	var count int64
	err := db.Get(&count, sqlStr, username)
	if err != nil && err != sql.ErrNoRows {
		zap.L().Error("query user exist failed", zap.Error(err))
		return true
	}
	return count > 0
}



func Login(user *models.User) (err error) {
	originPassword := user.Password // 记录一下原始密码
	sqlStr := "select user_id, username, password from user where username = ?"
	err = db.Get(user, sqlStr, user.UserName)
	if err != nil && err != sql.ErrNoRows {
		// 查询数据库出错
		return
	}
	if err == sql.ErrNoRows {
		// 用户不存在
		return ErrorUserNotExit
	}
	// 生成加密密码与查询到的密码比较
	password := encryptPassword([]byte(originPassword))
	if user.Password != password {
		return ErrorPasswordWrong
	}
	return
}


func GetUserInfoList(idList []uint64) (userList []*models.User, err error) {
	userList = make([]*models.User, 0, len(idList))
	sqlStr := `select user_id, username from user where user_id in (?)`
	inSql, args, err := sqlx.In(sqlStr, idList)
	if err != nil {
		zap.L().Error("build sqlx.In failed", zap.Error(err))
		return
	}
	queryStr := db.Rebind(inSql)
	err = db.Select(&userList, queryStr, args...)
	if err != nil {
		zap.L().Error("db.Select failed", zap.Error(err))
		return
	}
	return
}

func GetUserByID(idStr string) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id, username from user where user_id = ?`
	err = db.Get(user, sqlStr, idStr)
	return
}
