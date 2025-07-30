package web

import (
	utility "go_ruoyi_base/ruoyi_go/Utility"
	"go_ruoyi_base/ruoyi_go/domain"
	rescode "go_ruoyi_base/ruoyi_go/resCode"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type resPostObj struct {
	PostId     int64  `json:"postId" binding:"required"`
	CreateBy   string `json:"createBy"`
	CreateTime string `json:"createTime"`
	Flag       bool   `json:"flag"`
	PostCode   string `json:"postCode"`
	PostName   string `json:"postName"`
	PostSort   int32  `json:"postSort"`
	Remark     string `json:"remark"`
	Status     string `json:"status"`
	UpdateBy   string `json:"updateBy"`
	UpdateTime string `json:"updateTime"`
}

func ToResPostObj(domainObj domain.SysPost) resPostObj {
	return resPostObj{
		PostId:     domainObj.PostID,
		CreateBy:   domainObj.CreateBy,
		CreateTime: utility.FormatTimestamp(utility.DefaultTimeFormat, domainObj.CreateTime),
		PostCode:   domainObj.PostCode,
		PostName:   domainObj.PostName,
		PostSort:   domainObj.PostSort,
		Remark:     *domainObj.Remark,
		Status:     domainObj.Status,
		UpdateBy:   domainObj.UpdateBy,
		UpdateTime: utility.FormatTimestamp(utility.DefaultTimeFormat, domainObj.UpdateTime),
	}
}

// 新增岗位
func (h *SysPostHandler) AddPost(ctx *gin.Context) {
	type addReq struct {
		PostCode string `json:"postCode"`
		PostName string `json:"postName"`
		PostSort int32  `json:"postSort"`
		Remark   string `json:"remark"`
		Status   string `json:"status"`
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

	err := h.svc.Create(ctx, domain.SysPost{
		PostCode:   req.PostCode,
		PostName:   req.PostName,
		PostSort:   req.PostSort,
		Remark:     &req.Remark,
		Status:     req.Status,
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

// 编辑岗位
func (h *SysPostHandler) UpdatePost(ctx *gin.Context) {
	type addReq struct {
		PostId     int64  `json:"postId" binding:"required"`
		CreateBy   string `json:"createBy"`
		CreateTime string `json:"createTime"`
		Flag       bool   `json:"flag"`
		PostCode   string `json:"postCode"`
		PostName   string `json:"postName"`
		PostSort   int32  `json:"postSort"`
		Remark     string `json:"remark"`
		Status     string `json:"status"`
		UpdateBy   string `json:"updateBy"`
		UpdateTime string `json:"updateTime"`
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

	err := h.svc.Update(ctx, domain.SysPost{
		PostID: req.PostId,
		// Flag:       req.Flag,
		PostCode:   req.PostCode,
		PostName:   req.PostName,
		PostSort:   req.PostSort,
		Remark:     &req.Remark,
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

// 查询岗位列表
func (h *SysPostHandler) QueryPostList(ctx *gin.Context) {
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

	resList := []resPostObj{}
	for _, domainObj := range domainList {
		resList = append(resList, ToResPostObj(domainObj))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  rescode.Success,
		"msg":   rescode.Success.String(),
		"total": total,
		"rows":  resList,
	})
}

// 查询岗位详情
func (h *SysPostHandler) QueryPostDetail(ctx *gin.Context) {
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
		"data": ToResPostObj(domainObj),
	})
}

// 删除岗位
func (h *SysPostHandler) DeletePost(ctx *gin.Context) {
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
