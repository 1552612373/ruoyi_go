package main

import (
	"go_ruoyi_base/ruoyi_go/web/middleware"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	print("main")

	server := initWebServer()

	for _, route := range server.Routes() {
		log.Printf("HTTP %s --> %s\n", route.Method, route.Path)
	}

	server.Run(os.Getenv("SERVER_RUN_ADDRESS"))
}

func initWebServer() *gin.Engine {
	server := gin.Default()
	server.Use(func(ctx *gin.Context) {
		println("这是一个middleware，作用于这个server")

	})

	server.Use(cors.New(cors.Config{
		//AllowAllOrigins: true,
		//AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,

		AllowHeaders: []string{"Content-Type", "Authorization"},
		// 这个是允许前端访问你的后端响应中带的头部
		ExposeHeaders: []string{"zt-jwt-token"},
		//AllowHeaders: []string{"content-type"},
		//AllowMethods: []string{"POST"},
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") || strings.HasPrefix(origin, "http://127.0.0.1") {
				//if strings.Contains(origin, "localhost") {
				return true
			}
			return strings.Contains(origin, "120.55.44.222:7123")
		},
		MaxAge: 1222 * time.Hour,
	}), func(ctx *gin.Context) {
		// println("这是我的 Middleware")
	})

	login := &middleware.LoginJWTMiddlewareBuilder{}
	server.Use(login.CheckLogin())

	ry := &middleware.RuoYiMiddlewareBuilder{}
	ry.InitRuoYi(server)

	return server
}
