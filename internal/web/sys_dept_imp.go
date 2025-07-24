package web

import (
	utility "go_ruoyi_base/Utility"
	"go_ruoyi_base/internal/domain"
	rescode "go_ruoyi_base/resCode"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 查询字典数据列表
func (h *SysDeptHandler) QueryDeptList(ctx *gin.Context) {
	// type typeReq struct {
	// 	PageNum  int    `form:"pageNum" json:"pageNum"`   // 添加 form 标签
	// 	PageSize int    `form:"pageSize" json:"pageSize"` // 添加 form 标签
	// }

	// var req typeReq

	// if err := ctx.ShouldBindQuery(&req); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"code": rescode.ErrInvalidParam,
	// 		"msg":  rescode.ErrInvalidParam.String(),
	// 	})
	// 	return
	// }

	domainList, _, err := h.svc.QueryList(ctx)
	if err != nil {
		utility.ThrowSysErrowIfneeded(ctx, err)
		return
	}

	resList := []domain.SysDept{}
	for _, domainObj := range domainList {
		// resList = append(resList, toResDictDataObj(domainObj))
		resList = append(resList, domainObj)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": rescode.Success,
		"msg":  rescode.Success.String(),
		"data": resList,
	})
}
