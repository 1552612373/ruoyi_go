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
