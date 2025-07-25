package web

import (
	utility "go_ruoyi_base/Utility"
	"go_ruoyi_base/internal/service"

	"github.com/gin-gonic/gin"
)

type SysUserHandler struct {
	svc *service.SysUserService
}

func NewSysUserHandler(svc *service.SysUserService) *SysUserHandler {
	return &SysUserHandler{
		svc: svc,
	}
}

func (h *SysUserHandler) RegistRoutes(server *gin.Engine) {
	group := server.Group(utility.BASE_API_PRE)
	{
		group.POST("/signup", h.Signup)
		group.POST("/login", h.LoginJWT)
		group.POST("/logout", h.Logout)
		group.GET("/getInfo", h.GetInfo)

		// 查询用户列表
		group.GET("/system/user/list", h.QueryUserList)

		// 临时
		group.GET("/getRouters", h.GetRouters)
		group.GET("/system/config/configKey/sys.user.initPassword", h.VoidResponse)

	}

}
