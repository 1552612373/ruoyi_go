package service

import (
	"context"
	"go_ruoyi_base/ruoyi_go/domain"
	"go_ruoyi_base/ruoyi_go/repository"
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

func (svc *SysMenuService) Update(ctx context.Context, obj domain.SysMenu) error {
	return svc.repo.Update(ctx, obj)
}

func (svc *SysMenuService) QueryById(ctx context.Context, id int64) (domain.SysMenu, error) {
	return svc.repo.QueryById(ctx, id)
}

func (svc *SysMenuService) DeleteByDictId(ctx context.Context, id int64) error {
	return svc.repo.DeleteById(ctx, id)
}
