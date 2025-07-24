package web

import (
	utility "go_ruoyi_base/Utility"
	"go_ruoyi_base/internal/service"

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
		// 新增字典数据
		group.POST("/system/dept", h.AddDept)
		// // 更新字典数据
		// group.PUT("/system/dict/data", h.UpdateDictData)
		// // 删除字典数据
		// group.DELETE("/system/dict/data/:dictCode", h.DeleteDictData)
		// 查询部门列表
		group.GET("/system/dept/list", h.QueryDeptList)
		// // 查询字典数据列表
		// group.GET("/system/dict/data/type/:type", h.QueryDictDataType)
		// // 查询字典数据详情
		// group.GET("/system/dict/data/:dictCode", h.QueryDataDetail)
	}
}
