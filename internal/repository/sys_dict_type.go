package repository

import (
	"context"
	"go_ruoyi_base/internal/domain"
	"go_ruoyi_base/internal/repository/dao"
)

type SysDictTypeRepository struct {
	dao *dao.SysDictTypeDAO
}

func NewSysDictTypeRepository(dao *dao.SysDictTypeDAO) *SysDictTypeRepository {
	return &SysDictTypeRepository{
		dao: dao,
	}
}

func (repo *SysDictTypeRepository) Create(ctx context.Context, obj domain.SysDictType) error {
	return repo.dao.Insert(ctx, repo.toDao(obj))
}

func (repo *SysDictTypeRepository) toDao(obj domain.SysDictType) dao.SysDictType {
	return dao.SysDictType{
		DictId:     obj.DictId,
		DictName:   obj.DictName,
		DictType:   obj.DictType,
		Status:     obj.Status,
		CreateBy:   obj.CreateBy,
		CreateTime: obj.CreateTime,
		UpdateBy:   obj.UpdateBy,
		UpdateTime: obj.UpdateTime,
		Remark:     obj.Remark,
	}
}
