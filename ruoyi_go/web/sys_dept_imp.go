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

// 新增部门
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
	now := time.Now()

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

// 查询部门列表
func (h *SysDeptHandler) QueryDeptList(ctx *gin.Context) {

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

// 查询部门列表exclude （排除节点）
func (h *SysDeptHandler) QueryDeptListExclude(ctx *gin.Context) {

	// 获取路径参数 id
	deptIdStr := ctx.Param("deptId")
	deptId, err := strconv.ParseInt(deptIdStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": rescode.ErrInvalidParam,
			"msg":  "无效的字典类型ID",
		})
		return
	}

	domainList, _, err := h.svc.QueryListExclude(ctx, deptId)
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

// 查询部门详情
func (h *SysDeptHandler) QueryDeptDetail(ctx *gin.Context) {
	// 获取路径参数 id
	deptIdStr := ctx.Param("deptId")
	deptId, err := strconv.ParseInt(deptIdStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": rescode.ErrInvalidParam,
			"msg":  "无效的字典类型ID",
		})
		return
	}

	domainObj, err := h.svc.QueryByDeptId(ctx, deptId)

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

// 更新部门
func (h *SysDeptHandler) UpdateDept(ctx *gin.Context) {

	type deptReq struct {
		Ancestors  string    `json:"ancestors"`
		CreateBy   string    `json:"createBy"`
		CreateTime time.Time `json:"createTime"`
		DelFlag    string    `json:"delFlag" `
		DeptId     int64     `json:"deptId"`
		DeptName   string    `json:"deptName"`
		Email      string    `json:"email"`
		Leader     string    `json:"leader"`
		OrderNum   int       `json:"orderNum"`
		ParentId   int64     `json:"parentId"`
		Phone      string    `json:"phone"`
		Status     string    `json:"status"`
		UpdateBy   string    `json:"updateBy"`
		UpdateTime time.Time `json:"updateTime"`
	}

	var req deptReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": rescode.ErrInvalidParam,
			"msg":  rescode.ErrInvalidParam.String(),
		})
		return
	}

	err := h.svc.Update(ctx, domain.SysDept{
		Ancestors:  req.Ancestors,
		CreateBy:   req.CreateBy,
		CreateTime: req.CreateTime,
		DelFlag:    req.DelFlag,
		DeptID:     req.DeptId,
		DeptName:   req.DeptName,
		Email:      &req.Email,
		Leader:     &req.Leader,
		OrderNum:   req.OrderNum,
		ParentID:   req.ParentId,
		Phone:      &req.Phone,
		Status:     req.Status,
		UpdateBy:   req.UpdateBy,
		UpdateTime: req.UpdateTime,
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

// 查询部门数
func (h *SysDeptHandler) QueryDeptTree(ctx *gin.Context) {

	treeObj, err := h.svc.GetDeptTree(ctx)

	if err != nil {
		utility.ThrowSysErrowIfneeded(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": rescode.Success,
		"msg":  rescode.Success.String(),
		"data": treeObj,
	})
}

// 删除部门
func (h *SysDeptHandler) DeleteDept(ctx *gin.Context) {
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
