package web

import (
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
	group := server.Group("/api/user")
	group.POST("/signup", h.Signup)
	group.POST("/login", h.LoginJWT)

}
