package mysql

import (
	"fmt"

	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// Init 初始化MySQL连接
func Init() (err error) {
	// "user:password@tcp(host:port)/dbname"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.db"),
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	// 设置最大连接数
	db.SetMaxOpenConns(200)
	// 设置最大空闲连接数
	db.SetMaxIdleConns(20)

	return
}

// Close 程序退出时释放MySQL连接
// 不直接对外暴露db变量，而是对外暴露一个Close方法
func Close() {
	err := db.Close()
	fmt.Println(err)
}
