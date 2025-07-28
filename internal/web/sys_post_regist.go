package web

import (
	utility "go_ruoyi_base/Utility"
	"go_ruoyi_base/internal/service"

	"github.com/gin-gonic/gin"
)

type SysPostHandler struct {
	svc *service.SysPostService
}

func NewSysPostHandler(svc *service.SysPostService) *SysPostHandler {
	return &SysPostHandler{
		svc: svc,
	}
}

func (h *SysPostHandler) RegistRoutes(server *gin.Engine) {
	group := server.Group(utility.BASE_API_PRE)
	{
		// 新增岗位
		group.POST("/system/post", h.AddPost)
		// 岗位列表
		group.GET("/system/post/list", h.QueryPostList)
		// 编辑岗位
		group.PUT("/system/post", h.UpdatePost)
		// 查询岗位详情
		group.GET("/system/post/:id", h.QueryPostDetail)
		// 删除岗位
		group.DELETE("/system/post/:id", h.DeletePost)
	}
}
