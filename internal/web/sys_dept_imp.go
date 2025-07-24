package web

import (
	utility "go_ruoyi_base/Utility"
	"go_ruoyi_base/internal/domain"
	rescode "go_ruoyi_base/resCode"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 新增字典数据
func (h *SysDeptHandler) AddDept(ctx *gin.Context) {
	type deptReq struct {
		DeptName string `json:"deptName" binding:"required"`
		Email    string `json:"email"`
		Leader   string `json:"leader" `
		OrderNum int    `json:"orderNum" binding:"required"`
		ParentId int64  `json:"parentId" binding:"required"`
		Phone    string `json:"phone"`
		Status   string `json:"status"`
	}

	var req deptReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": rescode.ErrInvalidParam,
			"msg":  rescode.ErrInvalidParam.String(),
		})
		return
	}

	claimsObj, ok := ctx.MustGet(utility.ClaimsName).(utility.UserClaims)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": rescode.ErrUserUnauthorized,
			"msg":  rescode.ErrUserUnauthorized.String(),
		})
	}
	now := time.Now().UnixMilli()

	err := h.svc.Create(ctx, domain.SysDept{
		DeptName:   req.DeptName,
		Email:      &req.Email,
		Leader:     &req.Leader,
		OrderNum:   req.OrderNum,
		ParentID:   req.ParentId,
		Phone:      &req.Phone,
		Status:     req.Status,
		UpdateBy:   claimsObj.UserName,
		UpdateTime: now,
		CreateBy:   claimsObj.UserName,
		CreateTime: now,
	})
	if err != nil {
		utility.ThrowSysErrowIfneeded(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": rescode.Success,
		"msg":  rescode.Success.String(),
	})

}

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
