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

type resRoleObj struct {
	RoleId            int64   `json:"roleId"`
	RoleName          string  `json:"roleName"`
	RoleKey           string  `json:"roleKey"`
	RoleSort          int     `json:"roleSort"`
	DataScope         string  `json:"dataScope"`
	MenuCheckStrictly int     `json:"menuCheckStrictly"`
	DeptCheckStrictly int     `json:"deptCheckStrictly"`
	Flag              int     `json:"flag"`
	Admin             int     `json:"admin"`
	Remark            string  `json:"remark"`
	Status            string  `json:"status"`
	DelFlag           string  `json:"delFlag"`
	MenuIds           []int64 `json:"menuIds"`
	DeptIds           []int64 `json:"deptIds"`
	Permissions       string  `json:"permissions"`
	UpdateBy          string  `json:"updateBy"`
	UpdateTime        string  `json:"updateTime"`
	CreateBy          string  `json:"createBy"`
	CreateTime        string  `json:"createTime"`
}

func toResRoleObj(domainObj domain.SysRole) resRoleObj {
	return resRoleObj{
		RoleId:            domainObj.RoleId,
		RoleName:          domainObj.RoleName,
		RoleKey:           domainObj.RoleKey,
		RoleSort:          domainObj.RoleSort,
		DataScope:         domainObj.DataScope,
		MenuCheckStrictly: domainObj.MenuCheckStrictly,
		DeptCheckStrictly: domainObj.DeptCheckStrictly,
		// Flag:              domainObj.Flag,
		// Admin:             domainObj.Admin,
		Remark:  domainObj.Remark,
		Status:  domainObj.Status,
		DelFlag: domainObj.DelFlag,
		// MenuIds:           domainObj.MenuIds,
		// DeptIds:           domainObj.DeptIds,
		// Permissions:       domainObj.Permissions,
		UpdateBy:   domainObj.UpdateBy,
		UpdateTime: utility.FormatTimestamp(utility.DefaultTimeFormat, domainObj.UpdateTime),
		CreateBy:   domainObj.CreateBy,
		CreateTime: utility.FormatTimestamp(utility.DefaultTimeFormat, domainObj.CreateTime),
	}
}

// 新增角色
func (h *SysRoleHandler) AddRole(ctx *gin.Context) {
	type addReq struct {
		RoleName          string  `json:"roleName"`
		RoleKey           string  `json:"roleKey"`
		RoleSort          int     `json:"roleSort"`
		Remark            string  `json:"remark"`
		Status            string  `json:"status"`
		MenuCheckStrictly int     `json:"menuCheckStrictly"`
		DeptCheckStrictly int     `json:"deptCheckStrictly"`
		MenuIds           []int64 `json:"menuIds"`
		DeptIds           []int64 `json:"deptIds"`
	}

	var req addReq

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
	now := time.Now()

	err := h.svc.Create(ctx, domain.SysRole{
		RoleName:          req.RoleName,
		RoleKey:           req.RoleKey,
		RoleSort:          req.RoleSort,
		Remark:            req.Remark,
		Status:            req.Status,
		MenuCheckStrictly: req.MenuCheckStrictly,
		DeptCheckStrictly: req.DeptCheckStrictly,
		CreateBy:          claimsObj.UserName,
		CreateTime:        now,
	}, req.MenuIds, req.DeptIds)
	if err != nil {
		utility.ThrowSysErrowIfneeded(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": rescode.Success,
		"msg":  rescode.Success.String(),
	})

}

// // 编辑角色
// func (h *SysRoleHandler) UpdatePost(ctx *gin.Context) {
// 	type addReq struct {
// 		PostId     int64  `json:"postId" binding:"required"`
// 		CreateBy   string `json:"createBy"`
// 		CreateTime string `json:"createTime"`
// 		Flag       bool   `json:"flag"`
// 		PostCode   string `json:"postCode"`
// 		PostName   string `json:"postName"`
// 		PostSort   int32  `json:"postSort"`
// 		Remark     string `json:"remark"`
// 		Status     string `json:"status"`
// 		UpdateBy   string `json:"updateBy"`
// 		UpdateTime string `json:"updateTime"`
// 	}

// 	var req addReq

// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"code": rescode.ErrInvalidParam,
// 			"msg":  rescode.ErrInvalidParam.String(),
// 		})
// 		return
// 	}

// 	claimsObj, ok := ctx.MustGet(utility.ClaimsName).(utility.UserClaims)
// 	if !ok {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"code": rescode.ErrUserUnauthorized,
// 			"msg":  rescode.ErrUserUnauthorized.String(),
// 		})
// 	}
// 	now := time.Now()

// 	err := h.svc.Update(ctx, domain.SysRole{
// 		PostID: req.PostId,
// 		// Flag:       req.Flag,
// 		PostCode:   req.PostCode,
// 		PostName:   req.PostName,
// 		PostSort:   req.PostSort,
// 		Remark:     &req.Remark,
// 		Status:     req.Status,
// 		UpdateBy:   claimsObj.UserName,
// 		UpdateTime: now,
// 		CreateBy:   claimsObj.UserName,
// 		CreateTime: now,
// 	})
// 	if err != nil {
// 		utility.ThrowSysErrowIfneeded(ctx, err)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"code": rescode.Success,
// 		"msg":  rescode.Success.String(),
// 	})

// }

// 查询角色列表
func (h *SysRoleHandler) QueryRoleList(ctx *gin.Context) {
	type typeReq struct {
		PageNum  int `json:"pageNum" form:"pageNum"`
		PageSize int `json:"pageSize" form:"pageSize"`
	}

	var req typeReq

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": rescode.ErrInvalidParam,
			"msg":  rescode.ErrInvalidParam.String(),
		})
		return
	}

	domainList, total, err := h.svc.QueryList(ctx, req.PageNum, req.PageSize)
	if err != nil {
		utility.ThrowSysErrowIfneeded(ctx, err)
		return
	}

	resList := []resRoleObj{}
	for _, domainObj := range domainList {
		resList = append(resList, toResRoleObj(domainObj))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  rescode.Success,
		"msg":   rescode.Success.String(),
		"total": total,
		"rows":  resList,
	})
}

// 查询角色详情
func (h *SysRoleHandler) QueryRoleDetail(ctx *gin.Context) {
	// 获取路径参数 id
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": rescode.ErrInvalidParam,
			"msg":  "无效的字典类型ID",
		})
		return
	}

	domainObj, err := h.svc.QueryById(ctx, id)

	if err != nil {
		utility.ThrowSysErrowIfneeded(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": rescode.Success,
		"msg":  rescode.Success.String(),
		"data": toResRoleObj(domainObj),
	})
}

// 删除角色
func (h *SysRoleHandler) DeleteRole(ctx *gin.Context) {
	// 获取路径参数 id
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": rescode.ErrInvalidParam,
			"msg":  "无效的字典类型ID",
		})
		return
	}

	errx := h.svc.DeleteByDictId(ctx, id)
	if errx != nil {
		utility.ThrowSysErrowIfneeded(ctx, errx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": rescode.Success,
		"msg":  rescode.Success.String(),
	})

}
