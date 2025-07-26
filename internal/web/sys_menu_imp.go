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
func (h *SysMenuHandler) AddMenu(ctx *gin.Context) {
	type addReq struct {
		Component  string `json:"component"`
		CreateBy   string `json:"createBy"`
		CreateTime string `json:"createTime"`
		Icon       string `json:"icon"`
		IsCache    int    `json:"isCache"`
		IsFrame    int    `json:"isFrame"`
		MenuName   string `json:"menuName"`
		MenuType   string `json:"menuType"`
		OrderNum   int    `json:"orderNum"`
		ParentId   int64  `json:"parentId"`
		ParentName string `json:"parentName"`
		Path       string `json:"path"`
		Perms      string `json:"perms"`
		Query      string `json:"query"`
		Remark     string `json:"remark"`
		RouteName  string `json:"routeName"`
		Status     string `json:"status"`
		UpdateBy   string `json:"updateBy"`
		UpdateTime string `json:"updateTime"`
		Visible    string `json:"visible"`
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
	now := time.Now().UnixMilli()

	err := h.svc.Create(ctx, domain.SysMenu{
		Component: req.Component,
		Icon:      req.Icon,
		IsCache:   req.IsCache,
		IsFrame:   req.IsFrame,
		MenuName:  req.MenuName,
		MenuType:  req.MenuType,
		OrderNum:  req.OrderNum,
		ParentID:  req.ParentId,
		// ParentName: req.ParentName,
		Perms:      req.Perms,
		Query:      req.Query,
		Remark:     req.Remark,
		RouteName:  req.RouteName,
		Status:     req.Status,
		CreateBy:   claimsObj.UserName,
		CreateTime: now,
		Visible:    req.Visible,
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
