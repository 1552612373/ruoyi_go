package service

import (
	"context"
	"go_ruoyi_base/internal/domain"
	"go_ruoyi_base/internal/repository"
)

type SysDictDataService struct {
	repo *repository.SysDictDataRepository
}

func NewSysDictDataService(repo *repository.SysDictDataRepository) *SysDictDataService {
	return &SysDictDataService{
		repo: repo,
	}
}

func (svc *SysDictDataService) Create(ctx context.Context, obj domain.SysDictData) error {
	return svc.repo.Create(ctx, obj)
}

func (svc *SysDictDataService) QueryByDictCode(ctx context.Context, dictCode int64) (domain.SysDictData, error) {
	return svc.repo.QueryByDictCode(ctx, dictCode)
}

func (svc *SysDictDataService) QueryList(ctx context.Context, pageNum int, pageSize int, dictType string) ([]domain.SysDictData, int, error) {
	return svc.repo.QueryList(ctx, pageNum, pageSize, dictType)
}

func (svc *SysDictDataService) DeleteByDictCode(ctx context.Context, dictCode int64) error {
	return svc.repo.DeleteByDictCode(ctx, dictCode)
}

func (svc *SysDictDataService) Update(ctx context.Context, obj domain.SysDictData) error {
	return svc.repo.Update(ctx, obj)
}
