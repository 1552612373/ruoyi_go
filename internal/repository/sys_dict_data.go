package repository

import (
	"context"
	"go_ruoyi_base/internal/domain"
	"go_ruoyi_base/internal/repository/dao"
)

type SysDictDataRepository struct {
	dao *dao.SysDictDataDAO
}

func NewSysDictDataRepository(dao *dao.SysDictDataDAO) *SysDictDataRepository {
	return &SysDictDataRepository{
		dao: dao,
	}
}

func (repo *SysDictDataRepository) Create(ctx context.Context, obj domain.SysDictData) error {
	return repo.dao.Insert(ctx, repo.toDao(obj))
}

func (repo *SysDictDataRepository) QueryList(ctx context.Context, pageNum int, pageSize int, dictType string) ([]domain.SysDictData, int, error) {
	daoList, total, err := repo.dao.QueryList(ctx, pageNum, pageSize, dictType)
	return repo.toDomainList(daoList), total, err
}

func (repo *SysDictDataRepository) toDao(obj domain.SysDictData) dao.SysDictData {
	return dao.SysDictData{
		DictCode:   obj.DictCode,
		DictSort:   obj.DictSort,
		DictLabel:  obj.DictLabel,
		DictValue:  obj.DictValue,
		DictType:   obj.DictType,
		Status:     obj.Status,
		CreateBy:   obj.CreateBy,
		CreateTime: obj.CreateTime,
		UpdateBy:   obj.UpdateBy,
		UpdateTime: obj.UpdateTime,
		Remark:     obj.Remark,
	}
}

func (repo *SysDictDataRepository) toDomain(obj dao.SysDictData) domain.SysDictData {
	return domain.SysDictData{
		DictCode:   obj.DictCode,
		DictSort:   obj.DictSort,
		DictLabel:  obj.DictLabel,
		DictValue:  obj.DictValue,
		DictType:   obj.DictType,
		Status:     obj.Status,
		CreateBy:   obj.CreateBy,
		CreateTime: obj.CreateTime,
		UpdateBy:   obj.UpdateBy,
		UpdateTime: obj.UpdateTime,
		Remark:     obj.Remark,
	}
}

func (repo *SysDictDataRepository) toDomainList(daoList []dao.SysDictData) []domain.SysDictData {
	domainList := []domain.SysDictData{}
	for _, daoObj := range daoList {
		domainObj := repo.toDomain(daoObj)
		domainList = append(domainList, domainObj)
	}
	return domainList
}
