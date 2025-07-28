package repository

import (
	"context"
	"go_ruoyi_base/internal/domain"
	"go_ruoyi_base/internal/repository/dao"
)

type SysRoleRepository struct {
	dao *dao.SysRoleDAO
}

func NewSysRoleRepository(dao *dao.SysRoleDAO) *SysRoleRepository {
	return &SysRoleRepository{
		dao: dao,
	}
}

func (repo *SysRoleRepository) Create(ctx context.Context, obj domain.SysRole, menuIds []int64, deptIds []int64) error {
	return repo.dao.Insert(ctx, repo.toDao(obj), menuIds, deptIds)
}

func (repo *SysRoleRepository) QueryList(ctx context.Context, pageNum int, pageSize int) ([]domain.SysRole, int, error) {
	daoList, total, err := repo.dao.QueryList(ctx, pageNum, pageSize)
	return repo.toDomainList(daoList), total, err
}

func (repo *SysRoleRepository) Update(ctx context.Context, obj domain.SysRole) error {
	return repo.dao.Update(ctx, repo.toDao(obj))
}

func (repo *SysRoleRepository) QueryById(ctx context.Context, id int64) (domain.SysRole, error) {
	daoObj, err := repo.dao.QueryById(ctx, id)
	return repo.toDomain(daoObj), err
}

func (repo *SysRoleRepository) DeleteById(ctx context.Context, id int64) error {
	err := repo.dao.DeleteById(ctx, id)
	return err
}

func (repo *SysRoleRepository) QueryRoleMenuTreeById(ctx context.Context, id int64) ([]*dao.MenuTreeNode, []int64, error) {
	tree, keys, err := repo.dao.QueryRoleMenuTreeById(ctx, id)
	return tree, keys, err
}

func (repo *SysRoleRepository) toDao(obj domain.SysRole) dao.SysRole {
	return dao.SysRole{
		RoleId:            obj.RoleId,
		RoleName:          obj.RoleName,
		RoleKey:           obj.RoleKey,
		RoleSort:          obj.RoleSort,
		DataScope:         obj.DataScope,
		MenuCheckStrictly: obj.MenuCheckStrictly,
		DeptCheckStrictly: obj.DeptCheckStrictly,
		Status:            obj.Status,
		DelFlag:           obj.DelFlag,
		CreateBy:          obj.CreateBy,
		CreateTime:        obj.CreateTime,
		UpdateBy:          obj.UpdateBy,
		UpdateTime:        obj.UpdateTime,
		Remark:            obj.Remark,
	}
}

func (repo *SysRoleRepository) toDomain(obj dao.SysRole) domain.SysRole {
	return domain.SysRole{
		RoleId:            obj.RoleId,
		RoleName:          obj.RoleName,
		RoleKey:           obj.RoleKey,
		RoleSort:          obj.RoleSort,
		DataScope:         obj.DataScope,
		MenuCheckStrictly: obj.MenuCheckStrictly,
		DeptCheckStrictly: obj.DeptCheckStrictly,
		Status:            obj.Status,
		DelFlag:           obj.DelFlag,
		CreateBy:          obj.CreateBy,
		CreateTime:        obj.CreateTime,
		UpdateBy:          obj.UpdateBy,
		UpdateTime:        obj.UpdateTime,
		Remark:            obj.Remark,
	}
}

func (repo *SysRoleRepository) toDomainList(daoList []dao.SysRole) []domain.SysRole {
	domainList := []domain.SysRole{}
	for _, daoObj := range daoList {
		domainObj := repo.toDomain(daoObj)
		domainList = append(domainList, domainObj)
	}
	return domainList
}
