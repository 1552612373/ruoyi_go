package service

import (
	"context"
	"go_ruoyi_base/internal/domain"
	"go_ruoyi_base/internal/repository"
)

type SysPostService struct {
	repo *repository.SysPostRepository
}

func NewSysPostService(repo *repository.SysPostRepository) *SysPostService {
	return &SysPostService{
		repo: repo,
	}
}

func (svc *SysPostService) Create(ctx context.Context, obj domain.SysPost) error {
	return svc.repo.Create(ctx, obj)
}

func (svc *SysPostService) QueryList(ctx context.Context, pageNum int, pageSize int) ([]domain.SysPost, int, error) {
	return svc.repo.QueryList(ctx, pageNum, pageSize)
}

func (svc *SysPostService) Update(ctx context.Context, obj domain.SysPost) error {
	return svc.repo.Update(ctx, obj)
}

func (svc *SysPostService) QueryById(ctx context.Context, id int64) (domain.SysPost, error) {
	return svc.repo.QueryById(ctx, id)
}

func (svc *SysPostService) DeleteByDictId(ctx context.Context, id int64) error {
	return svc.repo.DeleteById(ctx, id)
}
