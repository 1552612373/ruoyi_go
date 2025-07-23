package web

import (
	utility "go_ruoyi_base/Utility"
	"go_ruoyi_base/internal/service"

	"github.com/gin-gonic/gin"
)

type SysDictDataHandler struct {
	svc *service.SysDictDataService
}

func NewSysDictDataHandler(svc *service.SysDictDataService) *SysDictDataHandler {
	return &SysDictDataHandler{
		svc: svc,
	}
}

func (h *SysDictDataHandler) RegistRoutes(server *gin.Engine) {
	group := server.Group(utility.BASE_API_PRE)
	{
		// 新增字典数据
		group.POST("/system/dict/data", h.AddDictData)
		// 更新字典数据
		// group.PUT("/system/dict/data", h.UpdateDictData)
		// 查询字典数据列表
		group.GET("/system/dict/data/list", h.QueryDictDataList)
		// 查询字典数据列表
		group.GET("/system/dict/data/type/:type", h.QueryDictDataType)
		// 查询字典数据详情
		group.GET("/system/dict/data/:dictCode", h.QueryDataDetail)
	}
}
