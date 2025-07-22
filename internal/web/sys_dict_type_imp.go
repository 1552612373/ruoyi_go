package web

import (
	utility "go_ruoyi_base/Utility"
	"go_ruoyi_base/internal/domain"
	rescode "go_ruoyi_base/resCode"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 新增字典类型
func (h *SysDictTypeHandler) AddDictType(ctx *gin.Context) {
	type typeReq struct {
		// 字典名称（如：用户性别）
		DictName string `json:"dictName" gorm:"column:dict_name" binding:"required"`

		// 字典类型编码（如：sys_user_sex）
		DictType string `json:"dictType" gorm:"column:dict_type;uniqueIndex;type:varchar(255)" binding:"required"`

		// 状态（0正常 1停用）
		Status string `json:"status" gorm:"column:status" binding:"required"`

		// 备注
		Remark string `json:"remark" gorm:"column:remark"`
	}

	var req typeReq

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
	err := h.svc.Create(ctx, domain.SysDictType{
		DictName:   req.DictName,
		DictType:   req.DictType,
		Status:     req.Status,
		Remark:     req.Remark,
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

// 查询字典列表
func (h *SysDictTypeHandler) QueryTypeList(ctx *gin.Context) {

}
