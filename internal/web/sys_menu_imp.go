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

type resMenuObj struct {
	MenuID     int64  `json:"menuId"`
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

func toResMenuObj(domainObj domain.SysMenu) resMenuObj {
	return resMenuObj{
		MenuID:    domainObj.MenuID,
		Component: domainObj.Component,
		Icon:      domainObj.Icon,
		IsCache:   domainObj.IsCache,
		IsFrame:   domainObj.IsFrame,
		MenuName:  domainObj.MenuName,
		MenuType:  domainObj.MenuType,
		OrderNum:  domainObj.OrderNum,
		ParentId:  domainObj.ParentID,
		// ParentName: domainObj.ParentName,
		Path:       domainObj.Path,
		Perms:      domainObj.Perms,
		Query:      domainObj.Query,
		Remark:     domainObj.Remark,
		RouteName:  domainObj.RouteName,
		Status:     domainObj.Status,
		CreateBy:   domainObj.CreateBy,
		CreateTime: utility.FormatTimestamp(utility.DefaultTimeFormat, domainObj.CreateTime),
		UpdateBy:   domainObj.UpdateBy,
		UpdateTime: utility.FormatTimestamp(utility.DefaultTimeFormat, domainObj.UpdateTime),
		Visible:    domainObj.Visible,
	}
}

// 新增菜单类型
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
		Path:       req.Path,
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

// 查询菜单列表
func (h *SysMenuHandler) QueryMenuList(ctx *gin.Context) {
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

	resList := []resMenuObj{}
	for _, domainObj := range domainList {
		resList = append(resList, toResMenuObj(domainObj))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  rescode.Success,
		"msg":   rescode.Success.String(),
		"total": total,
		"data":  resList,
	})
}

// 编辑菜单
func (h *SysMenuHandler) UpdateMenu(ctx *gin.Context) {
	type addReq struct {
		MenuID     int64  `json:"menuId"`
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

	err := h.svc.Update(ctx, domain.SysMenu{
		MenuID:    req.MenuID,
		Component: req.Component,
		Icon:      req.Icon,
		IsCache:   req.IsCache,
		IsFrame:   req.IsFrame,
		MenuName:  req.MenuName,
		MenuType:  req.MenuType,
		OrderNum:  req.OrderNum,
		ParentID:  req.ParentId,
		// ParentName: req.ParentName,
		Path:       req.Path,
		Perms:      req.Perms,
		Query:      req.Query,
		Remark:     req.Remark,
		RouteName:  req.RouteName,
		Status:     req.Status,
		UpdateBy:   claimsObj.UserName,
		UpdateTime: now,
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

// 查询菜单详情
func (h *SysMenuHandler) QueryMenuDetail(ctx *gin.Context) {
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
		"data": toResMenuObj(domainObj),
	})
}

// 删除菜单
func (h *SysMenuHandler) DeleteMenu(ctx *gin.Context) {
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
