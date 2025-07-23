package main

import (
	"fmt"
	"go_ruoyi_base/internal/repository"
	"go_ruoyi_base/internal/repository/dao"
	"go_ruoyi_base/internal/service"
	"go_ruoyi_base/internal/web"
	"go_ruoyi_base/internal/web/middleware"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	print("main")

	db := initDB()
	server := initWebServer()

	sysUserHandler := initSysUser(db)
	sysUserHandler.RegistRoutes(server)

	sysDictTypeHandler := initSysDictType(db)
	sysDictTypeHandler.RegistRoutes(server)

	sysDictDataHandler := initSysDictData(db)
	sysDictDataHandler.RegistRoutes(server)

	for _, route := range server.Routes() {
		log.Printf("HTTP %s --> %s\n", route.Method, route.Path)
	}

	server.Run(os.Getenv("SERVER_RUN_ADDRESS"))
}

func initDB() *gorm.DB {
	// 加载 .env 文件
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 获取配置
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	charset := os.Getenv("DB_CHARSET")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		user, password, host, port, dbName, charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// NoLowerCase: true, // 禁用字段名的小写转换
			SingularTable: true, // 禁用复数形式，即不再自动加 s
		},
	})

	if err != nil {
		panic(err)
	}

	errx := dao.InitTable(db)
	if errx != nil {
		panic(errx)
	}

	return db

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
		println("这是我的 Middleware")
	})
	// 定义下session
	// store := cookie.NewStore([]byte("secret")) // 基于cookie
	// store := memstore.NewStore([]byte("xxxxxx"), []byte("xxxxxx64wei")) // 基于内存
	// store, err := redis.NewStore(16, "tcp", "localhost:6379", "", "", []byte("k6CswdUm75WKcbM68UQUuxVsHSpTCwDD"), []byte("k6CswdUm75WKcbM68UQUuxVsHSpTCwEE"))
	// if err != nil {
	// 	panic(err)
	// }
	// server.Use(sessions.Sessions("ssid", store))

	// // 所有接口检验下
	// login := &middleware.LoginMiddlewareBuilder{}
	// server.Use(login.CheckLogin())

	login := &middleware.LoginJWTMiddlewareBuilder{}
	server.Use(login.CheckLogin())

	return server
}

func initSysUser(db *gorm.DB) *web.SysUserHandler {
	myDao := dao.NewSysUserDAO(db)
	myRepo := repository.NewSysUserRepository(myDao)
	mySvc := service.NewSysUserService(myRepo)
	myHandler := web.NewSysUserHandler(mySvc)
	return myHandler
}

func initSysDictType(db *gorm.DB) *web.SysDictTypeHandler {
	myDao := dao.NewSysDictTypeDAO(db)
	myRepo := repository.NewSysDictTypeRepository(myDao)
	mySvc := service.NewSysDictTypeService(myRepo)
	myHandler := web.NewSysDictTypeHandler(mySvc)
	return myHandler
}

func initSysDictData(db *gorm.DB) *web.SysDictDataHandler {
	myDao := dao.NewSysDictDataDAO(db)
	myRepo := repository.NewSysDictDataRepository(myDao)
	mySvc := service.NewSysDictDataService(myRepo)
	myHandler := web.NewSysDictDataHandler(mySvc)
	return myHandler
}
