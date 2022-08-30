package main

import (
	"context"
	"embed"
	"file-server/middleware"
	"file-server/routers"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func InitLogConfig() {
	log.SetPrefix("[file-server] ")
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
}

//go:embed templates
var template embed.FS

func main() {
	InitLogConfig()
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery(), middleware.Recorder(), middleware.Cors())
	routers.Init(r)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	//启动api服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("srv.ListenAndServe error:", err)
		}
	}()

	fads, err := fs.Sub(template, "templates")
	if err != nil {
		log.Println("fs.Sub error:", err)
		return
	}
	engine := gin.New()
	engine.Use(gin.Recovery(), middleware.Recorder(), middleware.Cors())
	engine.StaticFS("/", http.FS(fads))
	srv2 := &http.Server{
		Addr:    ":8081",
		Handler: engine,
	}
	//启动文件服务
	go func() {
		if err := srv2.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("srv2.ListenAndServe error:", err)
		}
	}()

	log.Println("before capture signal. the number of goroutines: ", runtime.NumGoroutine())
	C := make(chan os.Signal, 1)
	signal.Notify(C, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	<-C
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("srv.Shutdown error:", err)
	}
	log.Println("after capture signal. the remain number of goroutines: ", runtime.NumGoroutine())
}
