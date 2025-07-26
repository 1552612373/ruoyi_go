package web

import (
	utility "go_ruoyi_base/Utility"
	"go_ruoyi_base/internal/service"

	"github.com/gin-gonic/gin"
)

type SysMenuHandler struct {
	svc *service.SysMenuService
}

func NewSysMenuHandler(svc *service.SysMenuService) *SysMenuHandler {
	return &SysMenuHandler{
		svc: svc,
	}
}

func (h *SysMenuHandler) RegistRoutes(server *gin.Engine) {
	group := server.Group(utility.BASE_API_PRE)
	{
		// 新增菜单
		group.POST("/system/menu", h.AddMenu)
		// 查询菜单列表
		group.GET("/system/menu/list", h.QueryMenuList)
		// 编辑菜单
		group.PUT("/system/menu", h.UpdateMenu)
		// 查询菜单详情
		group.GET("/system/menu/:id", h.QueryMenuDetail)
		// 删除菜单
		group.DELETE("/system/menu/:id", h.DeleteMenu)
	}
}
