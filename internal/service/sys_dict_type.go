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

func (svc *SysDictTypeService) QueryList(ctx context.Context, pageNum int, pageSize int) ([]domain.SysDictType, int, error) {
	return svc.repo.QueryList(ctx, pageNum, pageSize)
}

func (svc *SysDictTypeService) QueryByDictId(ctx context.Context, dictId int64) (domain.SysDictType, error) {
	return svc.repo.QueryByDictId(ctx, dictId)
}

func (svc *SysDictTypeService) DeleteByDictId(ctx context.Context, dictId int64) error {
	return svc.repo.DeleteByDictId(ctx, dictId)
}

func (svc *SysDictTypeService) Update(ctx context.Context, obj domain.SysDictType) error {
	return svc.repo.Update(ctx, obj)
}
