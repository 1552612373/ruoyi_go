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

func (svc *SysUserService) GetInfo(ctx context.Context, id int64) (domain.SysUser, error) {
	domainSysUser, err := svc.repo.FindById(ctx, id)
	if err != nil {
		return domain.SysUser{}, errors.New("ZT查询失败")
	}
	return domainSysUser, err
}

func (svc *SysUserService) QueryList(ctx context.Context, pageNum int, pageSize int) ([]domain.SysUser, int, error) {
	return svc.repo.QueryList(ctx, pageNum, pageSize)
}

// 查看通用系统用户：岗位post列表和角色role列表
func (svc *SysUserService) GetSystemUserBase(ctx context.Context) ([]dao.SysPost, []dao.SysRole, error) {
	return svc.repo.GetSystemUserBase(ctx)
}

func (svc *SysUserService) QueryById(ctx context.Context, id int64) (domain.SysUser, error) {
	return svc.repo.QueryById(ctx, id)
}

func (svc *SysUserService) DeleteById(ctx context.Context, id int64) error {
	return svc.repo.DeleteById(ctx, id)
}
