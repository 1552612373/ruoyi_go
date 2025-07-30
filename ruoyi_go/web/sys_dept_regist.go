package web

import (
	utility "go_ruoyi_base/ruoyi_go/Utility"
	"go_ruoyi_base/ruoyi_go/service"

	"github.com/gin-gonic/gin"
)

type SysDeptHandler struct {
	svc *service.SysDeptService
}

func NewSysDeptHandler(svc *service.SysDeptService) *SysDeptHandler {
	return &SysDeptHandler{
		svc: svc,
	}
}

func (h *SysDeptHandler) RegistRoutes(server *gin.Engine) {
	group := server.Group(utility.BASE_API_PRE)
	{
		// 新增部门
		group.POST("/system/dept", h.AddDept)
		// 更新部门
		group.PUT("/system/dept", h.UpdateDept)
		// 查询部门列表
		group.GET("/system/dept/list", h.QueryDeptList)
		// 查询部门列表exclude （排除节点）
		group.GET("/system/dept/list/exclude/:deptId", h.QueryDeptListExclude)
		// 删除部门
		group.DELETE("/system/dept/:id", h.DeleteDept)
		// 查询部门详情
		group.GET("/system/dept/:deptId", h.QueryDeptDetail)
		// 查询部门数
		group.GET("/system/user/deptTree", h.QueryDeptTree)
	}
}
