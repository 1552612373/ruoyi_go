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

func (repo *SysDictTypeRepository) QueryByDictId(ctx context.Context, dictId int64) (domain.SysDictType, error) {
	daoObj, err := repo.dao.QueryByDictId(ctx, dictId)
	return repo.toDomain(daoObj), err
}

func (repo *SysDictTypeRepository) QueryList(ctx context.Context, pageNum int, pageSize int) ([]domain.SysDictType, int, error) {
	daoList, total, err := repo.dao.QueryList(ctx, pageNum, pageSize)
	return repo.toDomainList(daoList), total, err
}

func (repo *SysDictTypeRepository) Update(ctx context.Context, obj domain.SysDictType) error {
	return repo.dao.Update(ctx, repo.toDao(obj))
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

func (repo *SysDictTypeRepository) toDomain(obj dao.SysDictType) domain.SysDictType {
	return domain.SysDictType{
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

func (repo *SysDictTypeRepository) toDomainList(daoList []dao.SysDictType) []domain.SysDictType {
	domainList := []domain.SysDictType{}
	for _, daoObj := range daoList {
		domainObj := repo.toDomain(daoObj)
		domainList = append(domainList, domainObj)
	}
	return domainList
}
