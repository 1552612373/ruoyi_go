package service

import (
	"context"
	"go_ruoyi_base/ruoyi_go/domain"
	"go_ruoyi_base/ruoyi_go/repository"
	"go_ruoyi_base/ruoyi_go/repository/dao"
)

type SysRoleService struct {
	repo *repository.SysRoleRepository
}

func NewSysRoleService(repo *repository.SysRoleRepository) *SysRoleService {
	return &SysRoleService{
		repo: repo,
	}
}

func (svc *SysRoleService) Create(ctx context.Context, obj domain.SysRole, menuIds []int64, deptIds []int64) error {
	return svc.repo.Create(ctx, obj, menuIds, deptIds)
}

func (svc *SysRoleService) QueryList(ctx context.Context, pageNum int, pageSize int) ([]domain.SysRole, int, error) {
	return svc.repo.QueryList(ctx, pageNum, pageSize)
}

func (svc *SysRoleService) Update(ctx context.Context, obj domain.SysRole, menuIds []int64) error {
	return svc.repo.Update(ctx, obj, menuIds)
}

func (svc *SysRoleService) QueryById(ctx context.Context, id int64) (domain.SysRole, error) {
	return svc.repo.QueryById(ctx, id)
}

func (svc *SysRoleService) DeleteByDictId(ctx context.Context, id int64) error {
	return svc.repo.DeleteById(ctx, id)
}

func (svc *SysRoleService) QueryRoleMenuTreeById(ctx context.Context, id int64) ([]*dao.MenuTreeNode, []int64, error) {
	tree, keys, err := svc.repo.QueryRoleMenuTreeById(ctx, id)
	return tree, keys, err
}
