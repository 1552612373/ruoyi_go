package web

import (
	"github.com/gin-gonic/gin"
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
	// var req LoginReq

	// if err := ctx.ShouldBindJSON(&req); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"code": rescode.ErrInvalidParam,
	// 		"msg":  rescode.ErrInvalidParam.String(),
	// 	})
	// 	return
	// }

	// ad, err := h.svc.Login(ctx, req.Account, req.Password)
	// if err != nil {
	// 	utility.ThrowSysErrowIfneeded(ctx, err)
	// 	return
	// }

	// tokenStr := h.setJWTToken(ctx, ad.Id, ad.Account)

	// fmt.Println(ad)

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"code": rescode.Success,
	// 	"msg":  rescode.Success.String(),
	// 	"data": map[string]interface{}{
	// 		"id":         ad.Id,
	// 		"token":      tokenStr,
	// 		"account":    ad.Account,
	// 		"adminLevel": ad.AdminLevel,
	// 		"type":       ad.Type,
	// 	},
	// })
}
