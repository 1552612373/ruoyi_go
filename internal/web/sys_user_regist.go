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

		// 新增用户
		group.POST("/system/user", h.Signup)
		// 更新用户
		group.PUT("/system/user", h.Update)
		// 更新用户状态
		group.PUT("/system/user/changeStatus", h.ChangeStatus)
		// 更改用户密码
		group.PUT("/system/user/resetPwd", h.ResetPwd)
		// 查询用户列表
		group.GET("/system/user/list", h.QueryUserList)
		// 查看通用系统用户：岗位post列表和角色role列表
		group.GET("/system/user/", h.GetSystemUserBase)

		// 查询用户详情
		group.GET("/system/user/:id", h.QueryUserDetail)
		// 用户角色
		group.DELETE("/system/user/:id", h.DeleteUser)

		group.GET("/system/config/configKey/sys.user.initPassword", h.DefaultPassword)

		// 获取首页菜单项
		group.GET("/getRouters", h.GetRouters)

	}

}
