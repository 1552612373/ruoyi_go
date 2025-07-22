package web

import (
	utility "go_ruoyi_base/Utility"
	"go_ruoyi_base/internal/service"

	"github.com/gin-gonic/gin"
)

type SysDictTypeHandler struct {
	svc *service.SysDictTypeService
}

func NewSysDictTypeHandler(svc *service.SysDictTypeService) *SysDictTypeHandler {
	return &SysDictTypeHandler{
		svc: svc,
	}
}

func (h *SysDictTypeHandler) RegistRoutes(server *gin.Engine) {
	group := server.Group(utility.BASE_API_PRE)
	{
		// 新增字典类型
		group.POST("/system/dict/type", h.AddDictType)
		// 查询字典列表
		group.GET("/system/dict/type/list", h.QueryTypeList)
		// 查询字典下拉列表
		group.GET("/system/dict/type/optionselect", h.QueryOptionselect)
		// 查询字典详情
		group.GET("/system/dict/type/:id", h.QueryTypeDetail)
		// 更新字典类型
		group.PUT("/system/dict/type", h.UpdateType)
	}
}
