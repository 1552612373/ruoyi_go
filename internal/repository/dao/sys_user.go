package dao

import (
	"context"
	"errors"
	"fmt"
	"go_ruoyi_base/internal/domain"
	"sort"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type SysUser struct {
	// 用户ID
	// UserID int64 `gorm:"column:user_id;primaryKey;autoIncrement" json:"userId"`
	ID int64 `gorm:"column:user_id;primaryKey;autoIncrement" json:"userId"`

	// 部门ID
	DeptID *int64 `gorm:"column:dept_id" json:"deptId"`

	// 用户账号
	UserName string `gorm:"column:user_name" json:"userName"`

	// 用户昵称
	NickName string `gorm:"column:nick_name" json:"nickName"`

	// 用户类型（00系统用户）
	UserType string `gorm:"column:user_type" json:"userType"`

	// 用户邮箱
	Email string `gorm:"column:email" json:"email"`

	// 手机号码
	Phonenumber string `gorm:"column:phonenumber" json:"phonenumber"`

	// 用户性别（0男 1女 2未知）
	Sex string `gorm:"column:sex" json:"sex"`

	// 头像地址
	Avatar string `gorm:"column:avatar" json:"avatar"`

	// 密码
	Password string `gorm:"column:password" json:"password"`

	// 账号状态（0正常 1停用）
	Status string `gorm:"column:status" json:"status"`

	// 删除标志（0代表存在 2代表删除）
	DelFlag string `gorm:"column:del_flag" json:"delFlag"`

	// 最后登录IP
	LoginIP string `gorm:"column:login_ip" json:"loginIp"`

	// 最后登录时间（时间戳）
	LoginDate *time.Time `gorm:"column:login_date" json:"loginDate"`

	// 密码最后更新时间（时间戳）
	PwdUpdateDate *time.Time `gorm:"column:pwd_update_date" json:"pwdUpdateDate"`

	// 创建者
	CreateBy string `gorm:"column:create_by" json:"createBy"`

	// 创建时间（时间戳）
	CreateTime *time.Time `gorm:"column:create_time" json:"createTime"`

	// 更新者
	UpdateBy string `gorm:"column:update_by" json:"updateBy"`

	// 更新时间（时间戳）
	UpdateTime *time.Time `gorm:"column:update_time" json:"updateTime"`

	// 备注
	Remark *string `gorm:"column:remark" json:"remark"`
}

type SysUserPost struct {
	UserId int64 `gorm:"primaryKey"`
	PostId int64 `gorm:"primaryKey"`
}

type SysUserRole struct {
	UserId int64 `gorm:"primaryKey"`
	RoleId int64 `gorm:"primaryKey"`
}

type SysUserDAO struct {
	db      *gorm.DB
	postDao *SysPostDAO
	roleDao *SysRoleDAO
	deptDao *SysDeptDAO
	menuDao *SysMenuDAO
}

func NewSysUserDAO(db *gorm.DB, postDao *SysPostDAO, roleDao *SysRoleDAO, deptDao *SysDeptDAO, menuDao *SysMenuDAO) *SysUserDAO {
	return &SysUserDAO{
		db:      db,
		postDao: postDao,
		roleDao: roleDao,
		deptDao: deptDao,
		menuDao: menuDao,
	}
}

func (dao *SysUserDAO) Insert(ctx context.Context, obj SysUser, postIds []int64, roleIds []int64) error {
	// 开启事务
	tx := dao.db.WithContext(ctx).Begin()
	// “延迟执行 + panic 捕获” 机制，用于在发生 panic 时，自动回滚事务，防止数据不一致
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 先插入
	now := time.Now()
	obj.UpdateTime = &now
	obj.CreateTime = &now
	err := tx.Create(&obj).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			// 唯一键冲突
			tx.Rollback()
			return errors.New("ZT账户已存在")
		}
	}
	fmt.Printf("插入后 UserID: %d\n", obj.ID)
	// 如果有岗位，则插入关系表
	if len(postIds) > 0 {
		var userPosts []SysUserPost
		for _, postId := range postIds {
			userPosts = append(userPosts, SysUserPost{
				UserId: obj.ID,
				PostId: postId,
			})
		}
		if err := tx.Create(&userPosts).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	// 如果有角色，则插入关系表
	if len(roleIds) > 0 {
		var userRoles []SysUserRole
		for _, roleId := range roleIds {
			userRoles = append(userRoles, SysUserRole{
				UserId: obj.ID,
				RoleId: roleId,
			})
		}
		if err := tx.Create(&userRoles).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (dao *SysUserDAO) Update(ctx context.Context, obj SysUser, postIds []int64, roleIds []int64) error {
	// 开启事务
	tx := dao.db.WithContext(ctx).Begin()
	// “延迟执行 + panic 捕获” 机制，用于在发生 panic 时，自动回滚事务，防止数据不一致
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 检查 ID 是否有效
	if obj.ID == 0 {
		tx.Rollback()
		return errors.New("用户ID不能为空")
	}

	userUpdates := map[string]interface{}{
		"nick_name":   obj.NickName,
		"status":      obj.Status,
		"phonenumber": obj.Phonenumber,
		"email":       obj.Email,
		"sex":         obj.Sex,
		"dept_id":     *obj.DeptID, // 注意：obj.DeptId 是值，不是指针
		"remark":      *obj.Remark, // 注意：obj.Remark 是值，不是指针
	}

	if err := tx.Model(SysUser{}).Where("user_id = ?", obj.ID).Updates(userUpdates).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 处理岗位关联
	// 先删除该用户所有现有的岗位关联
	if err := tx.Where("user_id = ?", obj.ID).Delete(&SysUserPost{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 如果有新的岗位 ID，则插入新的关联
	if len(postIds) > 0 {
		var userPosts []SysUserPost
		for _, postId := range postIds {
			userPosts = append(userPosts, SysUserPost{
				UserId: obj.ID,
				PostId: postId,
			})
		}
		if err := tx.Create(&userPosts).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// 处理角色关联
	// 先删除该用户所有现有的角色关联
	if err := tx.Where("user_id = ?", obj.ID).Delete(&SysUserRole{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 如果有新的角色 ID，则插入新的关联
	if len(roleIds) > 0 {
		var userRoles []SysUserRole
		for _, roleId := range roleIds {
			userRoles = append(userRoles, SysUserRole{
				UserId: obj.ID,
				RoleId: roleId,
			})
		}
		if err := tx.Create(&userRoles).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// 提交事务
	return tx.Commit().Error
}

func (dao *SysUserDAO) ChangeStatus(ctx context.Context, userId int64, status string) error {

	// 检查 ID 是否有效
	if userId == 0 {
		return errors.New("用户ID不能为空")
	}

	userUpdates := map[string]interface{}{
		"status": status,
	}

	if err := dao.db.WithContext(ctx).Model(SysUser{}).Where("user_id = ?", userId).Updates(userUpdates).Error; err != nil {
		return err
	}

	return nil
}

func (dao *SysUserDAO) ResetPwd(ctx context.Context, userId int64, password string) error {

	// 检查 ID 是否有效
	if userId == 0 {
		return errors.New("用户ID不能为空")
	}

	userUpdates := map[string]interface{}{
		"password": password,
	}

	if err := dao.db.WithContext(ctx).Model(SysUser{}).Where("user_id = ?", userId).Updates(userUpdates).Error; err != nil {
		return err
	}

	return nil
}

func (dao *SysUserDAO) FindByAccount(ctx context.Context, account string) (SysUser, error) {
	sysUser := SysUser{}
	err := dao.db.WithContext(ctx).Where("user_name = ?", account).First(&sysUser).Error
	return sysUser, err
}

func (dao *SysUserDAO) QueryList(ctx context.Context, req domain.UserListReq) ([]SysUser, int, error) {
	objList := []SysUser{}
	db := dao.db.WithContext(ctx).Model(&SysUser{})

	var total int64

	// --- 构建查询条件 ---
	// 注意：使用指针或零值检查来判断参数是否提供
	if req.UserName != "" {
		// 模糊查询，使用 LIKE
		db = db.Where("user_name LIKE ?", "%"+req.UserName+"%")
		// 或者精确查询: db = db.Where("user_name = ?", req.UserName)
	}
	if req.Phonenumber != "" {
		// 通常手机号是精确匹配
		db = db.Where("phonenumber = ?", req.Phonenumber)
	}
	if req.Status != "" {
		// 状态通常是精确匹配
		db = db.Where("status = ?", req.Status)
	}

	// 处理时间范围查询
	// 假设 req.Params.BeginTime 和 req.Params.EndTime 是 string 类型
	if req.Params.BeginTime != "" {
		// 将字符串解析为 time.Time 进行比较更安全
		// 这里简化处理，直接拼接字符串 (注意 SQL 注入风险极低，因为是日期格式)
		// 更好的做法：解析成 time.Time 然后比较
		// parsedTime, err := time.Parse("2006-01-02", req.Params.BeginTime)
		// if err == nil {
		//     db = db.Where("create_time >= ?", parsedTime)
		// }
		db = db.Where("create_time >= ?", req.Params.BeginTime+" 00:00:00")
	}
	if req.Params.EndTime != "" {
		// 注意：EndTime 通常包含当天的 23:59:59
		db = db.Where("create_time <= ?", req.Params.EndTime+" 23:59:59")
	}
	// --- 条件构建结束 ---

	// 查询总数 (Count 会忽略 Limit 和 Offset，但会应用前面的 Where 条件)
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err // 如果 Count 出错，直接返回
	}

	// 分页处理
	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 执行分页查询
	// 注意：Find 会应用前面的 Where, Offset, Limit
	err := db.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Find(&objList).Error

	return objList, int(total), err
}

// 查看通用系统用户：岗位post列表和角色role列表
func (dao *SysUserDAO) GetSystemUserBase(ctx context.Context) ([]SysPost, []SysRole, error) {
	postObjList, _, err := dao.postDao.QueryList(ctx, 1, 99)
	if err != nil {
		return []SysPost{}, []SysRole{}, err
	}
	roleObjList, _, err := dao.roleDao.QueryList(ctx, 1, 99)
	if err != nil {
		return []SysPost{}, []SysRole{}, err
	}
	return postObjList, roleObjList, nil
}

// 需要带权限、角色等
func (dao *SysUserDAO) FindById(ctx context.Context, id int64) (SysUser, SysDept, []string, []string, error) {
	// 查询用户详情
	sysUser := SysUser{}
	err := dao.db.WithContext(ctx).Where("user_id = ?", id).First(&sysUser).Error
	if err != nil {
		return SysUser{}, SysDept{}, []string{}, []string{}, err
	}
	sysDept, errx := dao.deptDao.QueryByDeptId(ctx, *sysUser.DeptID)
	if errx != nil {
		return SysUser{}, SysDept{}, []string{}, []string{}, errx
	}
	permissions, _ := dao.menuDao.GetPermissionsByUserID(ctx, id)

	roles, _ := dao.menuDao.GetRoleKeysByUserID(ctx, id)

	return sysUser, sysDept, permissions, roles, nil
}

// 需要带post的id和详情列表，role的id和详情列表
func (dao *SysUserDAO) QueryById(ctx context.Context, id int64) (SysUser, SysDept, []int64, []SysPost, []int64, []SysRole, error) {
	// 查询用户详情
	sysUser := SysUser{}
	err := dao.db.WithContext(ctx).Where("user_id = ?", id).First(&sysUser).Error
	if err != nil {
		return SysUser{}, SysDept{}, []int64{}, []SysPost{}, []int64{}, []SysRole{}, err
	}
	sysDept, errx := dao.deptDao.QueryByDeptId(ctx, *sysUser.DeptID)
	if errx != nil {
		return SysUser{}, SysDept{}, []int64{}, []SysPost{}, []int64{}, []SysRole{}, errx
	}

	// 查该用户岗位
	var postRelations []SysUserPost
	erry := dao.db.Where("user_id = ?", id).Find(&postRelations).Error
	if erry != nil {
		return SysUser{}, SysDept{}, []int64{}, []SysPost{}, []int64{}, []SysRole{}, erry
	}
	// 提取出该用户对应的id
	var postIds []int64
	for _, rel := range postRelations {
		postIds = append(postIds, rel.PostId)
	}

	// 查该用户角色
	var roleRelations []SysUserRole
	errz := dao.db.Where("user_id = ?", id).Find(&roleRelations).Error
	if errz != nil {
		return SysUser{}, SysDept{}, []int64{}, []SysPost{}, []int64{}, []SysRole{}, errz
	}
	// 提取出该用户对应的id
	var roleIds []int64
	for _, rel := range roleRelations {
		roleIds = append(roleIds, rel.RoleId)
	}

	// 查询所有岗位列表
	daoPosts, _, _ := dao.postDao.QueryList(ctx, 1, 99)
	// 查询所有角色列表
	daoRoles, _, _ := dao.roleDao.QueryList(ctx, 1, 99)

	return sysUser, sysDept, postIds, daoPosts, roleIds, daoRoles, nil
}

func (dao *SysUserDAO) DeleteById(ctx context.Context, id int64) error {
	// 开启事务
	tx := dao.db.WithContext(ctx).Begin()
	// “延迟执行 + panic 捕获” 机制，用于在发生 panic 时，自动回滚事务，防止数据不一致
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 先删角色
	err := dao.db.WithContext(ctx).Where("user_id = ?", id).Delete(&SysUser{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 再删关系表
	if err := tx.Where("user_id = ?", id).Delete(&SysUserPost{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("user_id = ?", id).Delete(&SysUserRole{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	return tx.Commit().Error
}

func (dao *SysUserDAO) GetRoutersById(ctx context.Context, userId int64) ([]map[string]interface{}, error) {
	var menus []SysMenu

	// ====== 第一步：检查用户是否是超级管理员 ======
	var superAdmin bool
	err := dao.db.WithContext(ctx).
		Table("sys_user_role").
		Where("user_id = ? AND role_id = ?", userId, 1).
		Select("1").
		Take(&superAdmin).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("检查超级管理员角色失败: %w", err)
	}

	// 如果是超级管理员，返回所有启用的 M/C 菜单
	if superAdmin {
		err = dao.db.WithContext(ctx).
			Model(&domain.SysMenu{}).
			Where("menu_type IN ? AND status = ?", []string{"M", "C"}, "0"). // status=0 正常
			Order("parent_id ASC, order_num ASC, menu_id ASC").
			Find(&menus).Error

		if err != nil {
			return nil, fmt.Errorf("查询所有菜单失败: %w", err)
		}
	} else {
		// 普通用户：查询其角色拥有的菜单
		err = dao.db.WithContext(ctx).
			Table("sys_menu m").
			Joins("JOIN sys_role_menu rm ON m.menu_id = rm.menu_id").
			Joins("JOIN sys_user_role ur ON rm.role_id = ur.role_id").
			Where("ur.user_id = ? AND m.menu_type IN ?", userId, []string{"M", "C"}).
			Order("m.parent_id ASC, m.order_num ASC, m.menu_id ASC").
			Find(&menus).Error

		if err != nil {
			return nil, fmt.Errorf("查询用户菜单失败: %w", err)
		}
	}

	if len(menus) == 0 {
		return []map[string]interface{}{}, nil
	}

	// ====== 新增：将 menu_id 为 1,2,3,4 的菜单移到最后 ======
	specialIDs := map[int64]bool{1: true, 2: true, 3: true, 4: true}
	var normalMenus, specialMenus []SysMenu

	for _, menu := range menus {
		if specialIDs[menu.MenuID] { // 请根据你的结构体字段名调整：MenuID / MenuId / ID
			specialMenus = append(specialMenus, menu)
		} else {
			normalMenus = append(normalMenus, menu)
		}
	}

	// 对 specialMenus 按 menu_id 升序排序
	sort.Slice(specialMenus, func(i, j int) bool {
		return specialMenus[i].MenuID < specialMenus[j].MenuID
	})

	// 合并：正常菜单 + 特殊菜单
	menus = append(normalMenus, specialMenus...)
	// =======================================

	// 构建路由树（使用之前修正的 buildMenuTree）
	return buildMenuTree(menus), nil
}

func buildMenuTree(menus []SysMenu) []map[string]interface{} {
	menuMap := make(map[int64]map[string]interface{})
	var rootMenus []map[string]interface{}

	// 第一步：构建所有节点
	for _, m := range menus {
		meta := map[string]interface{}{
			"title":   m.MenuName,
			"icon":    m.Icon,
			"noCache": m.IsCache == 1,
			"link":    interface{}(nil),
		}

		if m.IsFrame == 0 { // 外链
			meta["link"] = m.Path
		}

		// ====== 路径处理：根节点加 /，子节点不加 / ======
		path := m.Path
		if m.ParentID == 0 {
			// 根节点：确保以 / 开头，且不以 / 结尾（除非是 "/")
			if !strings.HasPrefix(path, "/") {
				path = "/" + path
			}
			// 避免 // 或 /xxx/
			path = strings.TrimSuffix(path, "/")
			if path != "/" {
				path = strings.TrimSuffix(path, "/")
			}
		} else {
			// 子节点：确保不以 / 开头
			path = strings.TrimPrefix(path, "/")
		}

		item := map[string]interface{}{
			"name":      m.RouteName,
			"path":      path, // 使用处理后的 path
			"component": m.Component,
			"hidden":    m.Visible == "1",
			"meta":      meta,
			"children":  []map[string]interface{}{},
		}

		// 检查是否有子菜单
		hasChildren := false
		for _, child := range menus {
			if child.ParentID == m.MenuID {
				hasChildren = true
				break
			}
		}

		if hasChildren {
			item["redirect"] = "noRedirect"

			// 设置 component
			if m.MenuType == "M" {
				if m.ParentID == 0 {
					if m.Component == "" || m.Component == "Layout" {
						item["component"] = "Layout"
					}
				} else {
					if m.Component == "" || m.Component == "ParentView" {
						item["component"] = "ParentView"
					}
				}
			}

			item["alwaysShow"] = m.MenuType == "M"
		} else {
			item["redirect"] = nil
			item["alwaysShow"] = false
		}

		menuMap[m.MenuID] = item

		if m.ParentID == 0 {
			rootMenus = append(rootMenus, item)
		}
	}

	// 第二步：挂载 children
	for _, m := range menus {
		if m.ParentID != 0 {
			if parent, exists := menuMap[m.ParentID]; exists {
				children := parent["children"].([]map[string]interface{})
				parent["children"] = append(children, menuMap[m.MenuID])
			}
		}
	}

	return rootMenus
}
