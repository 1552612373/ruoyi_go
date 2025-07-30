package service

import (
	"context"
	"errors"
	"go_ruoyi_base/internal/domain"
	"go_ruoyi_base/internal/repository"
	"go_ruoyi_base/internal/repository/dao"

	"gorm.io/gorm"
)

type SysUserService struct {
	repo *repository.SysUserRepository
}

func NewSysUserService(repo *repository.SysUserRepository) *SysUserService {
	return &SysUserService{
		repo: repo,
	}
}

func (svc *SysUserService) Create(ctx context.Context, obj domain.SysUser, postIds []int64, roleIds []int64) error {
	return svc.repo.Create(ctx, obj, postIds, roleIds)
}

func (svc *SysUserService) Update(ctx context.Context, obj domain.SysUser, postIds []int64, roleIds []int64) error {
	return svc.repo.Update(ctx, obj, postIds, roleIds)
}

func (svc *SysUserService) ChangeStatus(ctx context.Context, userId int64, status string) error {
	return svc.repo.ChangeStatus(ctx, userId, status)
}

func (svc *SysUserService) ResetPwd(ctx context.Context, userId int64, password string) error {
	return svc.repo.ResetPwd(ctx, userId, password)
}

func (svc *SysUserService) Login(ctx context.Context, account string, password string) (domain.SysUser, error) {
	domainSysUser, err := svc.repo.FindByAccount(ctx, account)
	if err == gorm.ErrRecordNotFound {
		return domain.SysUser{}, errors.New("ZT账号或密码不正确")
	}
	if err != nil {
		return domain.SysUser{}, err
	}
	// 对比密码,先简单不加密
	if domainSysUser.Password != password {
		return domain.SysUser{}, errors.New("ZT账号或密码不正确")
	}
	return domainSysUser, err
}

func (svc *SysUserService) GetInfo(ctx context.Context, id int64) (domain.SysUser, []string, []string, error) {
	domainSysUser, permissions, roles, err := svc.repo.FindById(ctx, id)
	if err != nil {
		return domain.SysUser{}, []string{}, []string{}, errors.New("ZT查询失败")
	}
	return domainSysUser, permissions, roles, nil
}

func (svc *SysUserService) QueryList(ctx context.Context, req domain.UserListReq) ([]domain.SysUser, int, error) {
	return svc.repo.QueryList(ctx, req)
}

// 查看通用系统用户：岗位post列表和角色role列表
func (svc *SysUserService) GetSystemUserBase(ctx context.Context) ([]dao.SysPost, []dao.SysRole, error) {
	return svc.repo.GetSystemUserBase(ctx)
}

func (svc *SysUserService) QueryById(ctx context.Context, id int64) (domain.SysUser, []int64, []domain.SysPost, []int64, []domain.SysRole, error) {
	return svc.repo.QueryById(ctx, id)
}

func (svc *SysUserService) DeleteById(ctx context.Context, id int64) error {
	return svc.repo.DeleteById(ctx, id)
}

func (svc *SysUserService) GetRoutersById(ctx context.Context, userId int64) ([]map[string]interface{}, error) {
	menusMap, err := svc.repo.GetRoutersById(ctx, userId)
	if err != nil {
		return []map[string]interface{}{}, err
	}
	return menusMap, nil
}

func (svc *SysUserService) QueryAuthRoleListById(ctx context.Context, id int64) ([]domain.SysRole, error) {
	domainRoleList, err := svc.repo.QueryAuthRoleListById(ctx, id)
	return domainRoleList, err
}
