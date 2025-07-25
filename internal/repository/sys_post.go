package repository

import (
	"context"
	"go_ruoyi_base/internal/domain"
	"go_ruoyi_base/internal/repository/dao"
)

type SysPostRepository struct {
	dao *dao.SysPostDAO
}

func NewSysPostRepository(dao *dao.SysPostDAO) *SysPostRepository {
	return &SysPostRepository{
		dao: dao,
	}
}

func (repo *SysPostRepository) Create(ctx context.Context, obj domain.SysPost) error {
	return repo.dao.Insert(ctx, repo.toDao(obj))
}

func (repo *SysPostRepository) QueryList(ctx context.Context, pageNum int, pageSize int) ([]domain.SysPost, int, error) {
	daoList, total, err := repo.dao.QueryList(ctx, pageNum, pageSize)
	return repo.toDomainList(daoList), total, err
}

func (repo *SysPostRepository) toDao(obj domain.SysPost) dao.SysPost {
	return dao.SysPost{
		PostID:     obj.PostID,
		PostCode:   obj.PostCode,
		PostName:   obj.PostName,
		PostSort:   obj.PostSort,
		Status:     obj.Status,
		CreateBy:   obj.CreateBy,
		CreateTime: obj.CreateTime,
		UpdateBy:   obj.UpdateBy,
		UpdateTime: obj.UpdateTime,
		Remark:     obj.Remark,
	}
}

func (repo *SysPostRepository) toDomain(obj dao.SysPost) domain.SysPost {
	return domain.SysPost{
		PostID:     obj.PostID,
		PostCode:   obj.PostCode,
		PostName:   obj.PostName,
		PostSort:   obj.PostSort,
		Status:     obj.Status,
		CreateBy:   obj.CreateBy,
		CreateTime: obj.CreateTime,
		UpdateBy:   obj.UpdateBy,
		UpdateTime: obj.UpdateTime,
		Remark:     obj.Remark,
	}
}

func (repo *SysPostRepository) toDomainList(daoList []dao.SysPost) []domain.SysPost {
	domainList := []domain.SysPost{}
	for _, daoObj := range daoList {
		domainObj := repo.toDomain(daoObj)
		domainList = append(domainList, domainObj)
	}
	return domainList
}
