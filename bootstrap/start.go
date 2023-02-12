package bootstrap

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"project/global"
	"project/middleware"
	"project/routers"
	"syscall"
	"time"
)

/*
StartServer 启动服务
*/
func StartServer() {

	r := setUpRouter()

	srv := &http.Server{
		Addr:    ":" + global.App.Config.App.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

func setUpRouter() *gin.Engine {
	r := gin.Default()
	// 注册中间件
	r.Use(middleware.InitEnv)
	r.Use(middleware.WriteActionLog)
	r.Use(middleware.CatchError)
	// 注册路由
	routers.SetRouter(r)
	return r
}
