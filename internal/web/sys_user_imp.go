package web

import (
	utility "go_ruoyi_base/Utility"
	"go_ruoyi_base/internal/domain"
	rescode "go_ruoyi_base/resCode"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type resUserObj struct {
	UserId     int64  `json:"userId"`
	UserName   string `json:"userName"`
	NickName   string `json:"nickName"`
	Status     string `json:"status"`
	CreateBy   string `json:"createBy"`
	CreateTime string `json:"createTime"`
	UpdateBy   string `json:"updateBy"`
	UpdateTime string `json:"updateTime"`
	Remark     string `json:"remark"`
}

func toResUserObj(domainObj domain.SysUser) resUserObj {
	remark := ""
	if domainObj.Remark == nil {
		domainObj.Remark = &remark
	}
	return resUserObj{
		UserId:     domainObj.UserID,
		UserName:   domainObj.UserName,
		NickName:   domainObj.NickName,
		Status:     domainObj.Status,
		Remark:     remark,
		UpdateBy:   domainObj.UpdateBy,
		UpdateTime: utility.FormatTimestamp(utility.DefaultTimeFormat, domainObj.UpdateTime),
		CreateBy:   domainObj.CreateBy,
		CreateTime: utility.FormatTimestamp(utility.DefaultTimeFormat, domainObj.CreateTime),
	}
}

func (h *SysUserHandler) Signup(ctx *gin.Context) {
	type SignupReq struct {
		// 登录账号（唯一）✅ 必填
		UserName string `json:"userName" binding:"required" example:"admin"`

		// 用户昵称（显示用）❌ 非必填
		NickName string `json:"nickName" example:"管理员"`

		// 登录密码（加密存储）✅ 必填
		Password string `json:"password" binding:"required" example:"123456"`

		// 账户状态（0-启用，1-停用）✅ 必填
		Status string `json:"status" binding:"required" example:"0"`

		// 手机号码 ❌ 非必填
		Phonenumber string `json:"phonenumber" example:"13800000000"`

		// 邮箱地址 ❌ 非必填
		Email string `json:"email" example:"admin@example.com"`

		// 性别（男、女）❌ 非必填
		Sex string `json:"sex" example:"男"`

		// 所属部门ID ❌ 非必填
		DeptId int64 `json:"deptId,omitempty" example:"1"`

		// 用户拥有的角色ID列表 ❌ 非必填
		RoleIds []int64 `json:"roleIds,omitempty" example:"[1,2]"`

		// 备注信息 ❌ 非必填
		Remark string `json:"remark,omitempty" example:"系统管理员"`
	}
}

func (h *SysUserHandler) LoginJWT(ctx *gin.Context) {
	type LoginReq struct {
		UserName string `json:"userName" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var req LoginReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": rescode.ErrInvalidParam,
			"msg":  rescode.ErrInvalidParam.String(),
		})
		return
	}

	obj, err := h.svc.Login(ctx, req.UserName, req.Password)
	if err != nil {
		utility.ThrowSysErrowIfneeded(ctx, err)
		return
	}

	tokenStr := h.setJWTToken(ctx, obj.UserID, obj.UserName)

	ctx.JSON(http.StatusOK, gin.H{
		"code": rescode.Success,
		"msg":  rescode.Success.String(),
		// "data": map[string]interface{}{
		// 	"token": tokenStr,
		// },
		"token": tokenStr,
	})
}

func (h *SysUserHandler) setJWTToken(ctx *gin.Context, userId int64, userName string) string {
	ac := utility.UserClaims{
		UserId:    userId,
		UserName:  userName,
		UserAgent: ctx.GetHeader("User-Agent"),
		RegisteredClaims: jwt.RegisteredClaims{
			// x 分钟过期
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30000)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, ac)
	tokenStr, err := token.SignedString(utility.JWTKey)
	if err != nil {
		utility.ThrowSysErrowIfneeded(ctx, err)
	}
	ctx.Header("tlh-jwt-token", tokenStr)
	return tokenStr
}

func (h *SysUserHandler) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": rescode.Success,
		"msg":  rescode.Success.String(),
	})
}

func (h *SysUserHandler) GetInfo(ctx *gin.Context) {
	claimsObj, ok := ctx.MustGet(utility.ClaimsName).(utility.UserClaims)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": rescode.ErrUserUnauthorized,
			"msg":  rescode.ErrUserUnauthorized.String(),
		})
	}

	obj, err := h.svc.GetInfo(ctx, claimsObj.UserId)
	if err != nil {
		utility.ThrowSysErrowIfneeded(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": rescode.Success,
		"msg":  rescode.Success.String(),
		"user": map[string]interface{}{
			"userId":        obj.UserID,
			"deptId":        obj.DeptID,
			"userName":      obj.UserName,
			"nickName":      obj.NickName,
			"userType":      obj.UserType,
			"email":         obj.Email,
			"phonenumber":   obj.Phonenumber,
			"sex":           obj.Sex,
			"avatar":        obj.Avatar,
			"password":      obj.Password,
			"status":        obj.Status,
			"delFlag":       obj.DelFlag,
			"loginIp":       obj.LoginIP,
			"loginDate":     obj.LoginDate,
			"pwdUpdateDate": obj.PwdUpdateDate,
			"createBy":      obj.CreateBy,
			"createTime":    obj.CreateTime,
			"updateBy":      obj.UpdateBy,
			"updateTime":    obj.UpdateTime,
			"remark":        obj.Remark,
		},
		// 临时这样写
		"permissions": []string{
			"*:*:*",
		},
		"roles": []string{
			"admin",
		},
		"isDefaultModifyPwd": false,
		"isPasswordExpired":  false,
	})
}

func (h *SysUserHandler) GetRouters(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": 200,
		"data": []map[string]interface{}{
			{
				"name":       "System",
				"path":       "/system",
				"hidden":     false,
				"redirect":   "noRedirect",
				"component":  "Layout",
				"alwaysShow": true,
				"meta": map[string]interface{}{
					"title":   "系统管理",
					"icon":    "system",
					"noCache": false,
					"link":    nil,
				},
				"children": []map[string]interface{}{
					{
						"name":      "User",
						"path":      "user",
						"hidden":    false,
						"component": "system/user/index",
						"meta": map[string]interface{}{
							"title":   "用户管理",
							"icon":    "user",
							"noCache": false,
							"link":    nil,
						},
					},
					{
						"name":      "Role",
						"path":      "role",
						"hidden":    false,
						"component": "system/role/index",
						"meta": map[string]interface{}{
							"title":   "角色管理",
							"icon":    "peoples",
							"noCache": false,
							"link":    nil,
						},
					},
					{
						"name":      "Menu",
						"path":      "menu",
						"hidden":    false,
						"component": "system/menu/index",
						"meta": map[string]interface{}{
							"title":   "菜单管理",
							"icon":    "tree-table",
							"noCache": false,
							"link":    nil,
						},
					},
					{
						"name":      "Dept",
						"path":      "dept",
						"hidden":    false,
						"component": "system/dept/index",
						"meta": map[string]interface{}{
							"title":   "部门管理",
							"icon":    "tree",
							"noCache": false,
							"link":    nil,
						},
					},
					{
						"name":      "Post",
						"path":      "post",
						"hidden":    false,
						"component": "system/post/index",
						"meta": map[string]interface{}{
							"title":   "岗位管理",
							"icon":    "post",
							"noCache": false,
							"link":    nil,
						},
					},
					{
						"name":      "Dict",
						"path":      "dict",
						"hidden":    false,
						"component": "system/dict/index",
						"meta": map[string]interface{}{
							"title":   "字典管理",
							"icon":    "dict",
							"noCache": false,
							"link":    nil,
						},
					},
					{
						"name":      "Config",
						"path":      "config",
						"hidden":    false,
						"component": "system/config/index",
						"meta": map[string]interface{}{
							"title":   "参数设置",
							"icon":    "edit",
							"noCache": false,
							"link":    nil,
						},
					},
					{
						"name":      "Notice",
						"path":      "notice",
						"hidden":    false,
						"component": "system/notice/index",
						"meta": map[string]interface{}{
							"title":   "通知公告",
							"icon":    "message",
							"noCache": false,
							"link":    nil,
						},
					},
					{
						"name":       "Log",
						"path":       "log",
						"hidden":     false,
						"redirect":   "noRedirect",
						"component":  "ParentView",
						"alwaysShow": true,
						"meta": map[string]interface{}{
							"title":   "日志管理",
							"icon":    "log",
							"noCache": false,
							"link":    nil,
						},
						"children": []map[string]interface{}{
							{
								"name":      "Operlog",
								"path":      "operlog",
								"hidden":    false,
								"component": "monitor/operlog/index",
								"meta": map[string]interface{}{
									"title":   "操作日志",
									"icon":    "form",
									"noCache": false,
									"link":    nil,
								},
							},
							{
								"name":      "Logininfor",
								"path":      "logininfor",
								"hidden":    false,
								"component": "monitor/logininfor/index",
								"meta": map[string]interface{}{
									"title":   "登录日志",
									"icon":    "logininfor",
									"noCache": false,
									"link":    nil,
								},
							},
						},
					},
				},
			},
		},
	})
}

func (h *SysUserHandler) DefaultPassword(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "123456",
		"code": 200,
	})
}

// 查询用户列表
func (h *SysUserHandler) QueryUserList(ctx *gin.Context) {
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

	resList := []resUserObj{}
	for _, domainObj := range domainList {
		resList = append(resList, toResUserObj(domainObj))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  rescode.Success,
		"msg":   rescode.Success.String(),
		"total": total,
		"rows":  resList,
	})
}

// 查看通用系统用户：岗位post列表和角色role列表
func (h *SysUserHandler) GetSystemUserBase(ctx *gin.Context) {
	postObjList, roleObjList, err := h.svc.GetSystemUserBase(ctx)
	if err != nil {
		utility.ThrowSysErrowIfneeded(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  rescode.Success,
		"msg":   rescode.Success.String(),
		"posts": postObjList,
		"roles": roleObjList,
	})
}
