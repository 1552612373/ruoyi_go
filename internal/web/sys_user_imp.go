package web

import (
	utility "go_ruoyi_base/Utility"
	rescode "go_ruoyi_base/resCode"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

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

	tokenStr := h.setJWTToken(ctx, obj.UserId, obj.UserName)

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
			"userId":        obj.UserId,
			"deptId":        obj.DeptId,
			"userName":      obj.UserName,
			"nickName":      obj.NickName,
			"email":         obj.Email,
			"avatar":        obj.Avatar,
			"phonenumber":   obj.Phonenumber,
			"sex":           obj.Sex,
			"password":      obj.Password,
			"status":        obj.Status,
			"delFlag":       obj.DelFlag,
			"loginIp":       obj.LoginIp,
			"loginDate":     obj.LoginDate,
			"pwdUpdateDate": obj.PwdUpdateDate,
			"createBy":      obj.CreateBy,
			"createTime":    obj.CreateTime,
			"updateBy":      obj.UpdateBy,
			"updateTime":    obj.UpdateTime,
			"remark":        obj.Remark,
		},
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
