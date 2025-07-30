package repository

import (
	"context"
	"go_ruoyi_base/ruoyi_go/domain"
	"go_ruoyi_base/ruoyi_go/repository/dao"
)

type SysDeptRepository struct {
	dao *dao.SysDeptDAO
}

func NewSysDeptRepository(dao *dao.SysDeptDAO) *SysDeptRepository {
	return &SysDeptRepository{
		dao: dao,
	}
}

func (repo *SysDeptRepository) Create(ctx context.Context, obj domain.SysDept) error {
	return repo.dao.Insert(ctx, repo.toDao(obj))
}

func (repo *SysDeptRepository) QueryList(ctx context.Context) ([]domain.SysDept, int, error) {
	daoList, total, err := repo.dao.QueryList(ctx)
	return repo.toDomainList(daoList), total, err
}

func (repo *SysDeptRepository) QueryListExclude(ctx context.Context, excludeDeptId int64) ([]domain.SysDept, int, error) {
	daoList, total, err := repo.dao.QueryListExclude(ctx, excludeDeptId)
	return repo.toDomainList(daoList), total, err
}

func (repo *SysDeptRepository) QueryByDeptId(ctx context.Context, deptId int64) (domain.SysDept, error) {
	daoObj, err := repo.dao.QueryByDeptId(ctx, deptId)
	return repo.toDomain(daoObj), err
}

func (repo *SysDeptRepository) DeleteById(ctx context.Context, id int64) error {
	err := repo.dao.DeleteById(ctx, id)
	return err
}

func (repo *SysDeptRepository) Update(ctx context.Context, obj domain.SysDept) error {
	return repo.dao.Update(ctx, repo.toDao(obj))
}

func (repo *SysDeptRepository) GetDeptTree(ctx context.Context) ([]*dao.DeptTreeNode, error) {
	tree, err := repo.dao.GetDeptTree(ctx)
	return tree, err
}

func (repo *SysDeptRepository) toDao(obj domain.SysDept) dao.SysDept {
	return dao.SysDept{
		DeptID:     obj.DeptID,
		ParentID:   obj.ParentID,
		Ancestors:  obj.Ancestors,
		DeptName:   obj.DeptName,
		OrderNum:   obj.OrderNum,
		Leader:     obj.Leader,
		Phone:      obj.Phone,
		Email:      obj.Email,
		Status:     obj.Status,
		DelFlag:    obj.DelFlag,
		CreateBy:   obj.CreateBy,
		CreateTime: obj.CreateTime,
		UpdateBy:   obj.UpdateBy,
		UpdateTime: obj.UpdateTime,
	}
}

func (repo *SysDeptRepository) toDomain(obj dao.SysDept) domain.SysDept {
	return domain.SysDept{
		DeptID:     obj.DeptID,
		ParentID:   obj.ParentID,
		Ancestors:  obj.Ancestors,
		DeptName:   obj.DeptName,
		OrderNum:   obj.OrderNum,
		Leader:     obj.Leader,
		Phone:      obj.Phone,
		Email:      obj.Email,
		Status:     obj.Status,
		DelFlag:    obj.DelFlag,
		CreateBy:   obj.CreateBy,
		CreateTime: obj.CreateTime,
		UpdateBy:   obj.UpdateBy,
		UpdateTime: obj.UpdateTime,
	}
}

func (repo *SysDeptRepository) toDomainList(daoList []dao.SysDept) []domain.SysDept {
	domainList := []domain.SysDept{}
	for _, daoObj := range daoList {
		domainObj := repo.toDomain(daoObj)
		domainList = append(domainList, domainObj)
	}
	return domainList
}
