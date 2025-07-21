package service

import (
	"context"
	"errors"
	"go_ruoyi_base/internal/domain"
	"go_ruoyi_base/internal/repository"

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

func (svc *SysUserService) Create(ctx context.Context, obj domain.SysUser) error {
	return svc.repo.Create(ctx, obj)
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
	ad, err := svc.repo.FindById(ctx, id)
	if err != nil {
		return domain.SysUser{}, errors.New("ZT查询失败")
	}
	return ad, nil
}
