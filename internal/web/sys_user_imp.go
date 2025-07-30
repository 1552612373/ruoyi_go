package web

import (
	utility "go_ruoyi_base/Utility"
	"go_ruoyi_base/internal/domain"
	rescode "go_ruoyi_base/resCode"
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type resUserObj struct {
	// Admin      bool   `json:"admin"`
	// dept字典;password;postIds;pwdUpdateDate;roleId;roleIds;roles

	Avatar      string         `json:"avatar"`
	DelFlag     string         `json:"delFlag"`
	DeptId      int64          `json:"deptId"`
	Dept        domain.SysDept `json:"dept"`
	Email       string         `json:"email"`
	LoginDate   string         `json:"loginDate"`
	LoginIp     string         `json:"loginIp"`
	Phonenumber string         `json:"phonenumber"`
	Sex         string         `json:"sex"`
	UserId      int64          `json:"userId"`
	UserName    string         `json:"userName"`
	NickName    string         `json:"nickName"`
	Status      string         `json:"status"`
	CreateBy    string         `json:"createBy"`
	CreateTime  string         `json:"createTime"`
	UpdateBy    string         `json:"updateBy"`
	UpdateTime  string         `json:"updateTime"`
	Remark      string         `json:"remark"`
}

func toResUserObj(domainObj domain.SysUser) resUserObj {
	remark := ""
	if domainObj.Remark == nil {
		domainObj.Remark = &remark
	}
	return resUserObj{
		Avatar:    domainObj.Avatar,
		DelFlag:   domainObj.DelFlag,
		DeptId:    *domainObj.DeptID,
		Dept:      domainObj.Dept,
		Email:     domainObj.Email,
		LoginDate: utility.FormatTimePtr(utility.DefaultTimeFormat, domainObj.LoginDate),
		// LoginIp:     domainObj.LoginIp,
		Phonenumber: domainObj.Phonenumber,
		Sex:         domainObj.Sex,
		UserId:      domainObj.ID,
		UserName:    domainObj.UserName,
		NickName:    domainObj.NickName,
		Status:      domainObj.Status,
		Remark:      *domainObj.Remark,
		UpdateBy:    domainObj.UpdateBy,
		UpdateTime:  utility.FormatTimePtr(utility.DefaultTimeFormat, domainObj.UpdateTime),
		CreateBy:    domainObj.CreateBy,
		CreateTime:  utility.FormatTimePtr(utility.DefaultTimeFormat, domainObj.CreateTime),
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

		PostIds []int64 `json:"postIds,omitempty" example:"[1,2]"`

		// 备注信息 ❌ 非必填
		Remark string `json:"remark,omitempty" example:"系统管理员"`
	}

	var req SignupReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": rescode.ErrInvalidParam,
			"msg":  rescode.ErrInvalidParam.String(),
		})
		return
	}

	err := h.svc.Create(ctx, domain.SysUser{
		UserName:    req.UserName,
		NickName:    req.NickName,
		Password:    req.Password,
		Status:      req.Status,
		Phonenumber: req.Phonenumber,
		Email:       req.Email,
		Sex:         req.Sex,
		DeptID:      &req.DeptId,
		Remark:      &req.Remark,
	}, req.PostIds, req.RoleIds)

	if err != nil {
		utility.ThrowSysErrowIfneeded(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": rescode.Success,
		"msg":  rescode.Success.String(),
	})
}

func (h *SysUserHandler) Update(ctx *gin.Context) {
	type UpdateReq struct {
		ID          int64   `json:"userId"`
		NickName    string  `json:"nickName"`
		Status      string  `json:"status" binding:"required"`
		Phonenumber string  `json:"phonenumber"`
		Email       string  `json:"email"`
		Sex         string  `json:"sex"`
		DeptId      int64   `json:"deptId,omitempty"`
		RoleIds     []int64 `json:"roleIds,omitempty"`
		PostIds     []int64 `json:"postIds,omitempty"`
		Remark      string  `json:"remark,omitempty"`
	}

	var req UpdateReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": rescode.ErrInvalidParam,
			"msg":  rescode.ErrInvalidParam.String(),
		})
		return
	}

	err := h.svc.Update(ctx, domain.SysUser{
		ID:          req.ID,
		NickName:    req.NickName,
		Status:      req.Status,
		Phonenumber: req.Phonenumber,
		Email:       req.Email,
		Sex:         req.Sex,
		DeptID:      &req.DeptId,
		Remark:      &req.Remark,
	}, req.PostIds, req.RoleIds)

	if err != nil {
		utility.ThrowSysErrowIfneeded(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": rescode.Success,
		"msg":  rescode.Success.String(),
	})
}

func (h *SysUserHandler) ChangeStatus(ctx *gin.Context) {
	type UpdateReq struct {
		ID     int64  `json:"userId" binding:"required"`
		Status string `json:"status" binding:"required"`
	}

	var req UpdateReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": rescode.ErrInvalidParam,
			"msg":  rescode.ErrInvalidParam.String(),
		})
		return
	}

	err := h.svc.ChangeStatus(ctx, req.ID, req.Status)

	if err != nil {
		utility.ThrowSysErrowIfneeded(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": rescode.Success,
		"msg":  rescode.Success.String(),
	})
}

func (h *SysUserHandler) ResetPwd(ctx *gin.Context) {
	type UpdateReq struct {
		ID       int64  `json:"userId" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var req UpdateReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": rescode.ErrInvalidParam,
			"msg":  rescode.ErrInvalidParam.String(),
		})
		return
	}

	err := h.svc.ResetPwd(ctx, req.ID, req.Password)

	if err != nil {
		utility.ThrowSysErrowIfneeded(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": rescode.Success,
		"msg":  rescode.Success.String(),
	})
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

	tokenStr := h.setJWTToken(ctx, obj.ID, obj.UserName)

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

// 查询用户详情
func (h *SysUserHandler) QueryUserDetail(ctx *gin.Context) {
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

	domainSysUser, postIds, domainPosts, roleIds, domainRoles, err := h.svc.QueryById(ctx, id)
	if err != nil {
		utility.ThrowSysErrowIfneeded(ctx, err)
		return
	}

	resObj := toResUserObj(domainSysUser)

	resPostList := []resPostObj{}
	for _, domainObj := range domainPosts {
		resPostList = append(resPostList, ToResPostObj(domainObj))
	}

	resRoleList := []resRoleObj{}
	for _, domainObj := range domainRoles {
		resRoleList = append(resRoleList, ToResRoleObj(domainObj))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    rescode.Success,
		"msg":     rescode.Success.String(),
		"data":    resObj,
		"postIds": postIds,
		"posts":   resPostList,
		"roleIds": roleIds,
		"roles":   resRoleList,
	})
}

// 查询当前用户详情，和上面用户详情返回字段和格式有所不同
func (h *SysUserHandler) GetInfo(ctx *gin.Context) {
	claimsObj, ok := ctx.MustGet(utility.ClaimsName).(utility.UserClaims)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": rescode.ErrUserUnauthorized,
			"msg":  rescode.ErrUserUnauthorized.String(),
		})
	}

	obj, permissions, roles, err := h.svc.GetInfo(ctx, claimsObj.UserId)
	if err != nil {
		utility.ThrowSysErrowIfneeded(ctx, err)
		return
	}

	resObj := toResUserObj(obj)

	// 如果roles有admin，给permissions新赋值
	for _, roleKey := range roles {
		if roleKey == "admin" {
			permissions = []string{"*:*:*"}
			break
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":               rescode.Success,
		"msg":                rescode.Success.String(),
		"user":               resObj,
		"permissions":        permissions,
		"roles":              roles,
		"isDefaultModifyPwd": false,
		"isPasswordExpired":  false,
	})
}

func (h *SysUserHandler) GetRouters(ctx *gin.Context) {
	claimsObj, ok := ctx.MustGet(utility.ClaimsName).(utility.UserClaims)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": rescode.ErrUserUnauthorized,
			"msg":  rescode.ErrUserUnauthorized.String(),
		})
	}

	menuMap, err := h.svc.GetRoutersById(ctx, claimsObj.UserId)
	if err != nil {
		utility.ThrowSysErrowIfneeded(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": 200,
		"data": menuMap,
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

	var req domain.UserListReq

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": rescode.ErrInvalidParam,
			"msg":  rescode.ErrInvalidParam.String(),
		})
		return
	}

	domainList, total, err := h.svc.QueryList(ctx, req)
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

// 删除用户
func (h *SysUserHandler) DeleteUser(ctx *gin.Context) {
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

	errx := h.svc.DeleteById(ctx, id)
	if errx != nil {
		utility.ThrowSysErrowIfneeded(ctx, errx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": rescode.Success,
		"msg":  rescode.Success.String(),
	})

}

// 认证角色列表
func (h *SysUserHandler) QueryAuthRoleList(ctx *gin.Context) {
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

	domainRoleList, errx := h.svc.QueryAuthRoleListById(ctx, id)
	if errx != nil {
		utility.ThrowSysErrowIfneeded(ctx, errx)
		return
	}

	domainSysUser, _, _, roleIds, _, err := h.svc.QueryById(ctx, id)
	if err != nil {
		utility.ThrowSysErrowIfneeded(ctx, err)
		return
	}
	// println(roleIds)
	resObj := toResUserObj(domainSysUser)

	resList := []resRoleObj{}
	for _, domainObj := range domainRoleList {
		resRoleObj := ToResRoleObj(domainObj)
		if slices.Contains(roleIds, domainObj.RoleId) {
			// 该用户roleId包含于该角色内
			resRoleObj.Flag = 1
		}
		resList = append(resList, resRoleObj)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  rescode.Success,
		"msg":   rescode.Success.String(),
		"roles": resList,
		"user":  resObj,
	})

}
