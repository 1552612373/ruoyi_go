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
