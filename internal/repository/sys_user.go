package repository

import (
	"context"
	"go_ruoyi_base/internal/domain"
	"go_ruoyi_base/internal/repository/dao"
)

type SysUserRepository struct {
	dao      *dao.SysUserDAO
	postRepo *SysPostRepository
	roleRepo *SysRoleRepository
	deptRepo *SysDeptRepository
}

func NewSysUserRepository(dao *dao.SysUserDAO, postRepo *SysPostRepository, roleRepo *SysRoleRepository, deptRepo *SysDeptRepository) *SysUserRepository {
	return &SysUserRepository{
		dao:      dao,
		postRepo: postRepo,
		roleRepo: roleRepo,
		deptRepo: deptRepo,
	}
}

func (repo *SysUserRepository) Create(ctx context.Context, obj domain.SysUser, postIds []int64, roleIds []int64) error {
	return repo.dao.Insert(ctx, repo.toDao(obj), postIds, roleIds)
}

func (repo *SysUserRepository) Update(ctx context.Context, obj domain.SysUser, postIds []int64, roleIds []int64) error {
	return repo.dao.Update(ctx, repo.toDao(obj), postIds, roleIds)
}

func (repo *SysUserRepository) ChangeStatus(ctx context.Context, userId int64, status string) error {
	return repo.dao.ChangeStatus(ctx, userId, status)
}

func (repo *SysUserRepository) ResetPwd(ctx context.Context, userId int64, password string) error {
	return repo.dao.ResetPwd(ctx, userId, password)
}

func (repo *SysUserRepository) FindByAccount(ctx context.Context, account string) (domain.SysUser, error) {
	daoSysUser, err := repo.dao.FindByAccount(ctx, account)
	if err != nil {
		return domain.SysUser{}, err
	}
	domainSysUser := repo.toDomain(ctx, daoSysUser)
	return domainSysUser, err
}

func (repo *SysUserRepository) FindById(ctx context.Context, id int64) (domain.SysUser, []string, []string, error) {
	daoSysUser, daoSysDept, permissions, roles, err := repo.dao.FindById(ctx, id)
	if err != nil {
		return domain.SysUser{}, []string{}, []string{}, err
	}
	domainSysUser := repo.toDomain(ctx, daoSysUser)
	domainSysDept := repo.deptRepo.toDomain(daoSysDept)
	domainSysUser.Dept = domainSysDept
	return domainSysUser, permissions, roles, err
}

func (repo *SysUserRepository) QueryList(ctx context.Context, req domain.UserListReq) ([]domain.SysUser, int, error) {
	daoList, total, err := repo.dao.QueryList(ctx, req)
	return repo.toDomainList(ctx, daoList), total, err
}

func (repo *SysUserRepository) QueryById(ctx context.Context, id int64) (domain.SysUser, []int64, []domain.SysPost, []int64, []domain.SysRole, error) {
	daoSysUser, daoSysDept, postIds, daoPosts, roleIds, daoRoles, err := repo.dao.QueryById(ctx, id)
	domainSysUser := repo.toDomain(ctx, daoSysUser)
	domainSysDept := repo.deptRepo.toDomain(daoSysDept)
	domainSysUser.Dept = domainSysDept
	domainPosts := repo.postRepo.toDomainList(daoPosts)
	domainRoles := repo.roleRepo.toDomainList(daoRoles)
	return domainSysUser, postIds, domainPosts, roleIds, domainRoles, err
}

func (repo *SysUserRepository) DeleteById(ctx context.Context, id int64) error {
	err := repo.dao.DeleteById(ctx, id)
	return err
}

func (repo *SysUserRepository) GetRoutersById(ctx context.Context, userId int64) ([]map[string]interface{}, error) {
	menusMap, err := repo.dao.GetRoutersById(ctx, userId)
	if err != nil {
		return []map[string]interface{}{}, err
	}
	return menusMap, nil
}

func (repo *SysUserRepository) QueryAuthRoleListById(ctx context.Context, id int64) ([]domain.SysRole, error) {
	daoRoleList, err := repo.dao.QueryAuthRoleListById(ctx, id)
	domainRoleList := repo.roleRepo.toDomainList(daoRoleList)
	return domainRoleList, err
}

// 查看通用系统用户：岗位post列表和角色role列表
func (repo *SysUserRepository) GetSystemUserBase(ctx context.Context) ([]dao.SysPost, []dao.SysRole, error) {
	return repo.dao.GetSystemUserBase(ctx)
}

func (repo *SysUserRepository) toDao(obj domain.SysUser) dao.SysUser {
	return dao.SysUser{
		ID:            obj.ID,
		DeptID:        obj.DeptID,
		UserName:      obj.UserName,
		NickName:      obj.NickName,
		UserType:      obj.UserType,
		Email:         obj.Email,
		Phonenumber:   obj.Phonenumber,
		Sex:           obj.Sex,
		Avatar:        obj.Avatar,
		Password:      obj.Password,
		Status:        obj.Status,
		DelFlag:       obj.DelFlag,
		LoginIP:       obj.LoginIP,
		LoginDate:     obj.LoginDate,
		PwdUpdateDate: obj.PwdUpdateDate,
		CreateBy:      obj.CreateBy,
		CreateTime:    obj.CreateTime,
		UpdateBy:      obj.UpdateBy,
		UpdateTime:    obj.UpdateTime,
		Remark:        obj.Remark,
	}
}

func (repo *SysUserRepository) toDomain(ctx context.Context, obj dao.SysUser) domain.SysUser {
	dept, err := repo.deptRepo.dao.QueryByDeptId(ctx, *obj.DeptID)
	deptVar := domain.SysDept{}
	if err == nil {
		deptVar = repo.deptRepo.toDomain(dept)
	}
	return domain.SysUser{
		ID:            obj.ID,
		Dept:          deptVar,
		DeptID:        obj.DeptID,
		UserName:      obj.UserName,
		NickName:      obj.NickName,
		UserType:      obj.UserType,
		Email:         obj.Email,
		Phonenumber:   obj.Phonenumber,
		Sex:           obj.Sex,
		Avatar:        obj.Avatar,
		Password:      obj.Password,
		Status:        obj.Status,
		DelFlag:       obj.DelFlag,
		LoginIP:       obj.LoginIP,
		LoginDate:     obj.LoginDate,
		PwdUpdateDate: obj.PwdUpdateDate,
		CreateBy:      obj.CreateBy,
		CreateTime:    obj.CreateTime,
		UpdateBy:      obj.UpdateBy,
		UpdateTime:    obj.UpdateTime,
		Remark:        obj.Remark,
	}
}

func (repo *SysUserRepository) toDomainList(ctx context.Context, daoList []dao.SysUser) []domain.SysUser {
	domainList := []domain.SysUser{}
	for _, daoObj := range daoList {
		domainObj := repo.toDomain(ctx, daoObj)
		domainList = append(domainList, domainObj)
	}
	return domainList
}
