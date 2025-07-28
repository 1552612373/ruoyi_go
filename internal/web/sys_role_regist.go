package web

import (
	utility "go_ruoyi_base/Utility"
	"go_ruoyi_base/internal/service"

	"github.com/gin-gonic/gin"
)

type SysRoleHandler struct {
	svc *service.SysRoleService
}

func NewSysRoleHandler(svc *service.SysRoleService) *SysRoleHandler {
	return &SysRoleHandler{
		svc: svc,
	}
}

func (h *SysRoleHandler) RegistRoutes(server *gin.Engine) {
	group := server.Group(utility.BASE_API_PRE)
	{
		// 新增角色
		group.POST("/system/role", h.AddRole)
		// 角色列表
		group.GET("/system/role/list", h.QueryRoleList)
		// // 编辑角色
		// group.PUT("/system/role", h.UpdateRole)
		// 查询角色详情
		group.GET("/system/role/:id", h.QueryRoleDetail)
		// 删除角色
		group.DELETE("/system/role/:id", h.DeleteRole)
		// 查询角色的所有菜单树和权限
		group.GET("/system/menu/roleMenuTreeselect/:id", h.QueryRoleMenuTree)
	}
}
