package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"tanjunchen.io.webapp/dao/mysql"
	"tanjunchen.io.webapp/logger"
	"tanjunchen.io.webapp/pkg/snowflake"
	"tanjunchen.io.webapp/routers"
	"tanjunchen.io.webapp/setting"

	"go.uber.org/zap"

	"github.com/spf13/viper"
)

func main() {
	// 加载配置信息
	if err := setting.Init(); err != nil {
		fmt.Printf("load settings failed, err:%v\n", err)
		return
	}
	fmt.Println("load settings success")
	// 初始化日志库
	logger.Init()
	zap.L().Info("init logger success")
	// 初始化MySQL连接
	if err := mysql.Init(); err != nil {
		zap.L().Error("init mysql failed",
			zap.Error(err),
			zap.String("mysql", "xxx"),
			zap.Int("port", 3306))
		return
	}
	defer mysql.Close()
	zap.L().Info("init mysql success")
	// 初始化Redis连接
	//if err := redis.Init(); err != nil {
	//	zap.L().Error("init redis failed", zap.Error(err))
	//	return
	//}
	//zap.L().Info("init redis success")
	// kafka.Init
	// etcd.Init()
	// 加载路由信息
	//r := routers.SetupRouters()
	//r.Run(fmt.Sprintf(":%d", viper.GetInt("app.port")))

	// 初始化 ID 生成器
	if err := snowflake.Init(uint16(viper.GetInt("app.machine_id"))); err != nil {
		zap.L().Error("init snowflake failed", zap.Error(err))
	}

	// kafka.Init
	// etcd.Init()
	// 加载路由信息
	 r := routers.SetupRouters()
	//r.Run(fmt.Sprintf(":%d", viper.GetInt("app.port")))

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	// 开启一个goroutine启动服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	// 相当于告诉程序我给你5秒钟的时间你把没完成的请求处理一下，之后我们就要关机啦
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}
	zap.L().Info("Server exiting...")
}
