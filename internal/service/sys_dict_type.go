package service

import (
	"context"
	"go_ruoyi_base/internal/domain"
	"go_ruoyi_base/internal/repository"
)

type SysDictTypeService struct {
	repo *repository.SysDictTypeRepository
}

func NewSysDictTypeService(repo *repository.SysDictTypeRepository) *SysDictTypeService {
	return &SysDictTypeService{
		repo: repo,
	}
}

func (svc *SysDictTypeService) Create(ctx context.Context, obj domain.SysDictType) error {
	return svc.repo.Create(ctx, obj)
}
