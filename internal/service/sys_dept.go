package service

import (
	"context"
	"go_ruoyi_base/internal/domain"
	"go_ruoyi_base/internal/repository"
)

type SysDeptService struct {
	repo *repository.SysDeptRepository
}

func NewSysDeptService(repo *repository.SysDeptRepository) *SysDeptService {
	return &SysDeptService{
		repo: repo,
	}
}

func (svc *SysDeptService) Create(ctx context.Context, obj domain.SysDept) error {
	return svc.repo.Create(ctx, obj)
}

func (svc *SysDeptService) QueryList(ctx context.Context) ([]domain.SysDept, int, error) {
	return svc.repo.QueryList(ctx)
}
