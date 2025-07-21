package service

import (
	"context"
	"go_ruoyi_base/internal/domain"
	"go_ruoyi_base/internal/repository"
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
