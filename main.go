package main

import (
	"context"
	"embed"
	"file-server/middleware"
	"file-server/router"
	"file-server/utils"
	"flag"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {

}

func InitLogConfig() {
	log.SetPrefix("[file-server] ")
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
}

//go:embed templates
var tmpl embed.FS

func main() {
	InitLogConfig()
	var uiPort, srvPort int
	flag.IntVar(&uiPort, "up", 8080, "ui port")
	flag.IntVar(&srvPort, "sp", 8081, "server port")
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery(), middleware.Recorder(), middleware.Cors())
	router.Init(r)
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(srvPort),
		Handler: r,
	}
	//启动api服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("srv.ListenAndServe error:", err)
		}
	}()

	fads, err := fs.Sub(tmpl, "templates")
	if err != nil {
		log.Println("fs.Sub error:", err)
		return
	}
	engine := gin.New()
	engine.Use(gin.Recovery(), middleware.Recorder(), middleware.Cors())
	engine.StaticFS("/", http.FS(fads))
	srv2 := &http.Server{
		Addr:    ":" + strconv.Itoa(uiPort),
		Handler: engine,
	}
	usingIP := utils.GetLocalIP()
	log.Printf("ui run(net):\thttp:\\\\%s:%d\n", usingIP, uiPort)
	log.Printf("server run(net):\thttp:\\\\%s:%d\n", usingIP, srvPort)
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
