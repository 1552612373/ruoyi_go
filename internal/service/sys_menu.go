package service

import (
	"context"
	"go_ruoyi_base/internal/domain"
	"go_ruoyi_base/internal/repository"
)

type SysMenuService struct {
	repo *repository.SysMenuRepository
}

func NewSysMenuService(repo *repository.SysMenuRepository) *SysMenuService {
	return &SysMenuService{
		repo: repo,
	}
}

func (svc *SysMenuService) Create(ctx context.Context, obj domain.SysMenu) error {
	return svc.repo.Create(ctx, obj)
}

func (svc *SysMenuService) QueryList(ctx context.Context, pageNum int, pageSize int) ([]domain.SysMenu, int, error) {
	return svc.repo.QueryList(ctx, pageNum, pageSize)
}
