package web

import (
	utility "go_ruoyi_base/Utility"
	"go_ruoyi_base/internal/domain"
	rescode "go_ruoyi_base/resCode"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// type resDictDataObj struct {
// 	DictCode   int64  `json:"dictCode"`
// 	DictSort   int    `json:"dictSort"`
// 	DictLabel  string `json:"dictLabel"`
// 	DictValue  string `json:"dictValue"`
// 	DictType   string `json:"dictType"`
// 	Status     string `json:"status"`
// 	CreateBy   string `json:"createBy"`
// 	CreateTime string `json:"createTime"`
// 	UpdateBy   string `json:"updateBy"`
// 	UpdateTime string `json:"updateTime"`
// 	Remark     string `json:"remark"`
// }

// func toResDictDataObj(domainObj domain.SysDictData) resDictDataObj {
// 	return resDictDataObj{
// 		DictCode:   domainObj.DictCode,
// 		DictSort:   domainObj.DictSort,
// 		DictLabel:  domainObj.DictLabel,
// 		DictValue:  domainObj.DictValue,
// 		DictType:   domainObj.DictType,
// 		Status:     domainObj.Status,
// 		Remark:     domainObj.Remark,
// 		UpdateBy:   domainObj.UpdateBy,
// 		UpdateTime: utility.FormatTimestamp(utility.DefaultTimeFormat, domainObj.UpdateTime),
// 		CreateBy:   domainObj.CreateBy,
// 		CreateTime: utility.FormatTimestamp(utility.DefaultTimeFormat, domainObj.CreateTime),
// 	}
// }

// 新增字典数据
func (h *SysDictDataHandler) AddDictData(ctx *gin.Context) {
	type typeReq struct {
		DictLabel string `json:"dictLabel" binding:"required"`
		DictValue string `json:"dictValue" binding:"required"`
		ListClass string `json:"listClass" `
		DictSort  int    `json:"dictSort" binding:"required"`
		Status    string `json:"status" binding:"required"`
		DictType  string `json:"dictType" binding:"required"`
		CssClass  string `json:"cssClass"`
		Remark    string `json:"remark"`
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

	err := h.svc.Create(ctx, domain.SysDictData{
		DictLabel:  req.DictLabel,
		DictValue:  req.DictValue,
		ListClass:  &req.ListClass,
		DictSort:   req.DictSort,
		Status:     req.Status,
		DictType:   req.DictType,
		CssClass:   &req.CssClass,
		Remark:     &req.Remark,
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
func (h *SysDictDataHandler) QueryDictDataList(ctx *gin.Context) {
	type typeReq struct {
		PageNum  int    `form:"pageNum" json:"pageNum"`   // 添加 form 标签
		PageSize int    `form:"pageSize" json:"pageSize"` // 添加 form 标签
		DictType string `form:"dictType" json:"dictType" binding:"required"`
	}

	var req typeReq

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": rescode.ErrInvalidParam,
			"msg":  rescode.ErrInvalidParam.String(),
		})
		return
	}

	domainList, total, err := h.svc.QueryList(ctx, req.PageNum, req.PageSize, req.DictType)
	if err != nil {
		utility.ThrowSysErrowIfneeded(ctx, err)
		return
	}

	var resList []domain.SysDictData
	for _, domainObj := range domainList {
		// resList = append(resList, toResDictDataObj(domainObj))
		resList = append(resList, domainObj)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  rescode.Success,
		"msg":   rescode.Success.String(),
		"total": total,
		"rows":  resList,
	})
}

// 查询字典数据详情
func (h *SysDictDataHandler) QueryDictDataType(ctx *gin.Context) {
	// 获取路径参数 id
	typeStr := ctx.Param("type")

	domainList, total, err := h.svc.QueryList(ctx, 1, 20, typeStr)
	if err != nil {
		utility.ThrowSysErrowIfneeded(ctx, err)
		return
	}

	var resList []domain.SysDictData
	for _, domainObj := range domainList {
		resList = append(resList, domainObj)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  rescode.Success,
		"msg":   rescode.Success.String(),
		"total": total,
		"rows":  resList,
	})
}

// 查询字典数据详情
func (h *SysDictDataHandler) QueryDataDetail(ctx *gin.Context) {
	// 获取路径参数 id
	dictCodeStr := ctx.Param("dictCode")
	dictCode, err := strconv.ParseInt(dictCodeStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": rescode.ErrInvalidParam,
			"msg":  "无效的字典类型ID",
		})
		return
	}

	domainObj, err := h.svc.QueryByDictCode(ctx, dictCode)

	if err != nil {
		utility.ThrowSysErrowIfneeded(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": rescode.Success,
		"msg":  rescode.Success.String(),
		"data": domainObj,
	})
}

// 删除字典数据
func (h *SysDictDataHandler) DeleteDictData(ctx *gin.Context) {
	// 获取路径参数 id
	idStr := ctx.Param("dictCode")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": rescode.ErrInvalidParam,
			"msg":  "无效的字典类型ID",
		})
		return
	}

	errx := h.svc.DeleteByDictCode(ctx, id)
	if errx != nil {
		utility.ThrowSysErrowIfneeded(ctx, errx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": rescode.Success,
		"msg":  rescode.Success.String(),
	})

}

// 更新字典类型
func (h *SysDictDataHandler) UpdateDictData(ctx *gin.Context) {

	type dictReq struct {
		DictCode   int64  `json:"dictCode"`
		DictSort   int    `json:"dictSort"`
		DictLabel  string `json:"dictLabel"`
		DictValue  string `json:"dictValue"`
		DictType   string `json:"dictType"`
		CssClass   string `json:"cssClass"`
		ListClass  string `json:"listClass"`
		IsDefault  string `json:"isDefault"`
		Status     string `json:"status"`
		Default    bool   `json:"default"`
		CreateBy   string `json:"createBy"`
		CreateTime string `json:"createTime"`
		UpdateBy   string `json:"updateBy"`
		UpdateTime string `json:"updateTime"`
		Remark     string `json:"remark"`
	}

	var req dictReq

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

	err := h.svc.Update(ctx, domain.SysDictData{
		DictCode:  req.DictCode,
		DictSort:  req.DictSort,
		DictLabel: req.DictLabel,
		DictValue: req.DictValue,
		DictType:  req.DictType,
		CssClass:  &req.CssClass,
		ListClass: &req.ListClass,
		// IsDefault:  req.IsDefault,
		Status: req.Status,
		// Default:    req.Default,
		Remark:     &req.Remark,
		CreateBy:   req.CreateBy,
		CreateTime: utility.ParseToTimestamp(utility.DefaultTimeFormat, req.CreateTime),
		UpdateBy:   claimsObj.UserName,
		UpdateTime: now,
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
