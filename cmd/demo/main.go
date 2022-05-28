package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	app "github.com/reatang/gostart/application/demo"
	"github.com/reatang/gostart/pkg/database/mysql"
	"github.com/reatang/gostart/pkg/logger"
	"github.com/reatang/gostart/pkg/redis"
	"go.uber.org/zap"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// pprof
	go func() {
		_ = http.ListenAndServe(":6060", nil)
	}()
	
	flag.Parse()
	app.InitConfig()
	
	logger.InitLogger(app.GetConfig().Log)
	mysql.InitDatabase(app.GetConfig().Database)
	redis.InitRedis(app.GetConfig().Redis)

	gin.SetMode(gin.DebugMode)
	router := app.RegisterRouter()
	
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.GetConfig().Web.Port),
		Handler:      router,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	
	// 应用
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("http server startup err", zap.Error(err))
		}
	}()

	// 系统信号
	_ = waitForSignal()
	_ = server.Shutdown(context.Background())
	log.Printf("App closed.")
}

func waitForSignal() os.Signal {
	signalChan := make(chan os.Signal, 1)
	defer close(signalChan)
	signal.Notify(signalChan, os.Kill, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	s := <-signalChan
	log.Printf("Signal received: %s", s)
	signal.Stop(signalChan)
	return s
}

