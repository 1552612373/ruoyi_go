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

type resDictTypeObj struct {
	DictId     int64  `json:"dictId"`
	DictName   string `json:"dictName"`
	DictType   string `json:"dictType"`
	Status     string `json:"status"`
	CreateBy   string `json:"createBy"`
	CreateTime string `json:"createTime"`
	UpdateBy   string `json:"updateBy"`
	UpdateTime string `json:"updateTime"`
	Remark     string `json:"remark"`
}

func toResDictTypeObj(domainObj domain.SysDictType) resDictTypeObj {
	return resDictTypeObj{
		DictId:     domainObj.DictId,
		DictName:   domainObj.DictName,
		DictType:   domainObj.DictType,
		Status:     domainObj.Status,
		Remark:     domainObj.Remark,
		UpdateBy:   domainObj.UpdateBy,
		UpdateTime: utility.FormatTimestamp(utility.DefaultTimeFormat, domainObj.UpdateTime),
		CreateBy:   domainObj.CreateBy,
		CreateTime: utility.FormatTimestamp(utility.DefaultTimeFormat, domainObj.CreateTime),
	}
}

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

// 更新字典类型
func (h *SysDictTypeHandler) UpdateType(ctx *gin.Context) {

	type typeReq struct {
		DictId     int64  `json:"dictId"`
		DictName   string `json:"dictName"`
		DictType   string `json:"dictType"`
		Status     string `json:"status"`
		CreateBy   string `json:"createBy"`
		CreateTime string `json:"createTime"`
		UpdateBy   string `json:"updateBy"`
		UpdateTime string `json:"updateTime"`
		Remark     string `json:"remark"`
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

	err := h.svc.Update(ctx, domain.SysDictType{
		DictId:     req.DictId,
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

// 查询字典详情
func (h *SysDictTypeHandler) QueryTypeDetail(ctx *gin.Context) {
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

	domainObj, err := h.svc.QueryByDictId(ctx, id)

	ctx.JSON(http.StatusOK, gin.H{
		"code": rescode.Success,
		"msg":  rescode.Success.String(),
		"data": toResDictTypeObj(domainObj),
	})
}

// 查询字典列表
func (h *SysDictTypeHandler) QueryTypeList(ctx *gin.Context) {
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

	var resList []resDictTypeObj
	for _, domainObj := range domainList {
		resList = append(resList, toResDictTypeObj(domainObj))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  rescode.Success,
		"msg":   rescode.Success.String(),
		"total": total,
		"rows":  resList,
	})
}

// 查询字典下拉列表
func (h *SysDictTypeHandler) QueryOptionselect(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "操作成功",
		"data": []interface{}{
			gin.H{
				"createBy":   "admin",
				"createTime": "2025-07-14 16:15:27",
				"updateBy":   nil,
				"updateTime": nil,
				"remark":     "用户性别列表",
				"dictId":     1,
				"dictName":   "用户性别",
				"dictType":   "sys_user_sex",
				"status":     "0",
			},
			gin.H{
				"createBy":   "admin",
				"createTime": "2025-07-14 16:15:27",
				"updateBy":   nil,
				"updateTime": nil,
				"remark":     "菜单状态列表",
				"dictId":     2,
				"dictName":   "菜单状态",
				"dictType":   "sys_show_hide",
				"status":     "0",
			},
			gin.H{
				"createBy":   "admin",
				"createTime": "2025-07-14 16:15:27",
				"updateBy":   nil,
				"updateTime": nil,
				"remark":     "系统开关列表",
				"dictId":     3,
				"dictName":   "系统开关",
				"dictType":   "sys_normal_disable",
				"status":     "0",
			},
			gin.H{
				"createBy":   "admin",
				"createTime": "2025-07-14 16:15:27",
				"updateBy":   nil,
				"updateTime": nil,
				"remark":     "任务状态列表",
				"dictId":     4,
				"dictName":   "任务状态",
				"dictType":   "sys_job_status",
				"status":     "0",
			},
			gin.H{
				"createBy":   "admin",
				"createTime": "2025-07-14 16:15:27",
				"updateBy":   nil,
				"updateTime": nil,
				"remark":     "任务分组列表",
				"dictId":     5,
				"dictName":   "任务分组",
				"dictType":   "sys_job_group",
				"status":     "0",
			},
			gin.H{
				"createBy":   "admin",
				"createTime": "2025-07-14 16:15:27",
				"updateBy":   nil,
				"updateTime": nil,
				"remark":     "系统是否列表",
				"dictId":     6,
				"dictName":   "系统是否",
				"dictType":   "sys_yes_no",
				"status":     "0",
			},
			gin.H{
				"createBy":   "admin",
				"createTime": "2025-07-14 16:15:27",
				"updateBy":   nil,
				"updateTime": nil,
				"remark":     "通知类型列表",
				"dictId":     7,
				"dictName":   "通知类型",
				"dictType":   "sys_notice_type",
				"status":     "0",
			},
			gin.H{
				"createBy":   "admin",
				"createTime": "2025-07-14 16:15:27",
				"updateBy":   nil,
				"updateTime": nil,
				"remark":     "通知状态列表",
				"dictId":     8,
				"dictName":   "通知状态",
				"dictType":   "sys_notice_status",
				"status":     "0",
			},
			gin.H{
				"createBy":   "admin",
				"createTime": "2025-07-14 16:15:27",
				"updateBy":   nil,
				"updateTime": nil,
				"remark":     "操作类型列表",
				"dictId":     9,
				"dictName":   "操作类型",
				"dictType":   "sys_oper_type",
				"status":     "0",
			},
			gin.H{
				"createBy":   "admin",
				"createTime": "2025-07-14 16:15:27",
				"updateBy":   nil,
				"updateTime": nil,
				"remark":     "登录状态列表",
				"dictId":     10,
				"dictName":   "系统状态",
				"dictType":   "sys_common_status",
				"status":     "0",
			},
			gin.H{
				"createBy":   "admin",
				"createTime": "2025-07-22 13:37:48",
				"updateBy":   nil,
				"updateTime": nil,
				"remark":     "remark!",
				"dictId":     100,
				"dictName":   "测试sys_test",
				"dictType":   "sys_test",
				"status":     "0",
			},
			gin.H{
				"createBy":   "admin",
				"createTime": "2025-07-22 13:39:32",
				"updateBy":   nil,
				"updateTime": nil,
				"remark":     "re",
				"dictId":     101,
				"dictName":   "测试2",
				"dictType":   "sys_test2",
				"status":     "0",
			},
		},
	})
}
