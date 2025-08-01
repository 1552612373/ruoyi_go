package service

import (
	"context"
	"go_ruoyi_base/ruoyi_go/domain"
	"go_ruoyi_base/ruoyi_go/repository"
	"go_ruoyi_base/ruoyi_go/repository/dao"
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

func (svc *SysDeptService) QueryListExclude(ctx context.Context, excludeDeptId int64) ([]domain.SysDept, int, error) {
	return svc.repo.QueryListExclude(ctx, excludeDeptId)
}

func (svc *SysDeptService) QueryByDeptId(ctx context.Context, deptId int64) (domain.SysDept, error) {
	return svc.repo.QueryByDeptId(ctx, deptId)
}

func (svc *SysDeptService) DeleteByDictId(ctx context.Context, id int64) error {
	return svc.repo.DeleteById(ctx, id)
}

func (svc *SysDeptService) Update(ctx context.Context, obj domain.SysDept) error {
	return svc.repo.Update(ctx, obj)
}

func (svc *SysDeptService) GetDeptTree(ctx context.Context) ([]*dao.DeptTreeNode, error) {
	tree, err := svc.repo.GetDeptTree(ctx)
	return tree, err
}
