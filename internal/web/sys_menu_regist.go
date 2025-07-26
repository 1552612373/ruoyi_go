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

	}
}
