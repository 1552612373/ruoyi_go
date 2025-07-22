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
		// 新增字典类型
		group.POST("/system/dict/data", h.AddDictData)
		// 查询字典列表
		// group.GET("/system/dict/type/list", h.QueryTypeList)
	}
}
