package middleware

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"go_ruoyi_base/ruoyi_go/repository"
	"go_ruoyi_base/ruoyi_go/repository/dao"
	"go_ruoyi_base/ruoyi_go/service"
	"go_ruoyi_base/ruoyi_go/web"
)

type RuoYiMiddlewareBuilder struct {
	db *gorm.DB
}

// InitRuoYi 返回一个中间件，它会在第一次请求时注册所有路由
func (m *RuoYiMiddlewareBuilder) InitRuoYi(server *gin.Engine) {

	// 使用 sync.Once 如果你想确保只注册一次
	var initialized bool

	// 只在第一次请求时注册路由
	if !initialized {
		log.Println("🔧 正在初始化 RuoYi 模块路由...")

		m.db = initDB()

		// 创建各个 Handler 并注册路由
		sysUserHandler := m.initSysUser()
		sysUserHandler.RegistRoutes(server)

		sysDictTypeHandler := m.initSysDictType()
		sysDictTypeHandler.RegistRoutes(server)

		sysDictDataHandler := m.initSysDictData()
		sysDictDataHandler.RegistRoutes(server)

		sysDeptHandler := m.initSysDept()
		sysDeptHandler.RegistRoutes(server)

		SysPostHandler := m.initSysPost()
		SysPostHandler.RegistRoutes(server)

		SysMenuHandler := m.initSysMenu()
		SysMenuHandler.RegistRoutes(server)

		SysRoleHandler := m.initSysRole()
		SysRoleHandler.RegistRoutes(server)

		// 打印所有路由
		server := server
		for _, route := range server.Routes() {
			log.Printf("HTTP %s --> %s", route.Method, route.Path)
		}

		log.Println("✅ 所有模块路由注册完成")
		initialized = true
	}

}

// 下面是各个 init 方法，完全复制你 main.go 的逻辑

func (m *RuoYiMiddlewareBuilder) initSysUser() *web.SysUserHandler {
	postDao := dao.NewSysPostDAO(m.db)
	menuDao := dao.NewSysMenuDAO(m.db)
	roleDao := dao.NewSysRoleDAO(m.db, menuDao)
	deptDao := dao.NewSysDeptDAO(m.db)

	postRepo := repository.NewSysPostRepository(postDao)
	roleRepo := repository.NewSysRoleRepository(roleDao)
	deptRepo := repository.NewSysDeptRepository(deptDao)

	myDao := dao.NewSysUserDAO(m.db, postDao, roleDao, deptDao, menuDao)
	myRepo := repository.NewSysUserRepository(myDao, postRepo, roleRepo, deptRepo)
	mySvc := service.NewSysUserService(myRepo)
	return web.NewSysUserHandler(mySvc)
}

func (m *RuoYiMiddlewareBuilder) initSysDictType() *web.SysDictTypeHandler {
	myDao := dao.NewSysDictTypeDAO(m.db)
	myRepo := repository.NewSysDictTypeRepository(myDao)
	mySvc := service.NewSysDictTypeService(myRepo)
	return web.NewSysDictTypeHandler(mySvc)
}

func (m *RuoYiMiddlewareBuilder) initSysDictData() *web.SysDictDataHandler {
	myDao := dao.NewSysDictDataDAO(m.db)
	myRepo := repository.NewSysDictDataRepository(myDao)
	mySvc := service.NewSysDictDataService(myRepo)
	return web.NewSysDictDataHandler(mySvc)
}

func (m *RuoYiMiddlewareBuilder) initSysDept() *web.SysDeptHandler {
	myDao := dao.NewSysDeptDAO(m.db)
	myRepo := repository.NewSysDeptRepository(myDao)
	mySvc := service.NewSysDeptService(myRepo)
	return web.NewSysDeptHandler(mySvc)
}

func (m *RuoYiMiddlewareBuilder) initSysPost() *web.SysPostHandler {
	myDao := dao.NewSysPostDAO(m.db)
	myRepo := repository.NewSysPostRepository(myDao)
	mySvc := service.NewSysPostService(myRepo)
	return web.NewSysPostHandler(mySvc)
}

func (m *RuoYiMiddlewareBuilder) initSysMenu() *web.SysMenuHandler {
	myDao := dao.NewSysMenuDAO(m.db)
	myRepo := repository.NewSysMenuRepository(myDao)
	mySvc := service.NewSysMenuService(myRepo)
	return web.NewSysMenuHandler(mySvc)
}

func (m *RuoYiMiddlewareBuilder) initSysRole() *web.SysRoleHandler {
	menuDao := dao.NewSysMenuDAO(m.db)
	myDao := dao.NewSysRoleDAO(m.db, menuDao)
	myRepo := repository.NewSysRoleRepository(myDao)
	mySvc := service.NewSysRoleService(myRepo)
	return web.NewSysRoleHandler(mySvc)
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
