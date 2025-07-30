package repository

import (
	"context"
	"go_ruoyi_base/ruoyi_go/domain"
	"go_ruoyi_base/ruoyi_go/repository/dao"
)

type SysMenuRepository struct {
	dao *dao.SysMenuDAO
}

func NewSysMenuRepository(dao *dao.SysMenuDAO) *SysMenuRepository {
	return &SysMenuRepository{
		dao: dao,
	}
}

func (repo *SysMenuRepository) Create(ctx context.Context, obj domain.SysMenu) error {
	return repo.dao.Insert(ctx, repo.toDao(obj))
}

func (repo *SysMenuRepository) QueryList(ctx context.Context, pageNum int, pageSize int) ([]domain.SysMenu, int, error) {
	daoList, total, err := repo.dao.QueryList(ctx, pageNum, pageSize)
	return repo.toDomainList(daoList), total, err
}

func (repo *SysMenuRepository) Update(ctx context.Context, obj domain.SysMenu) error {
	return repo.dao.Update(ctx, repo.toDao(obj))
}

func (repo *SysMenuRepository) QueryById(ctx context.Context, id int64) (domain.SysMenu, error) {
	daoObj, err := repo.dao.QueryById(ctx, id)
	return repo.toDomain(daoObj), err
}

func (repo *SysMenuRepository) DeleteById(ctx context.Context, id int64) error {
	err := repo.dao.DeleteById(ctx, id)
	return err
}

func (repo *SysMenuRepository) toDao(obj domain.SysMenu) dao.SysMenu {
	return dao.SysMenu{
		MenuID:     obj.MenuID,
		MenuName:   obj.MenuName,
		ParentID:   obj.ParentID,
		OrderNum:   obj.OrderNum,
		Path:       obj.Path,
		Component:  obj.Component,
		Query:      obj.Query,
		RouteName:  obj.RouteName,
		IsFrame:    obj.IsFrame,
		IsCache:    obj.IsCache,
		MenuType:   obj.MenuType,
		Visible:    obj.Visible,
		Status:     obj.Status,
		Perms:      obj.Perms,
		Icon:       obj.Icon,
		CreateBy:   obj.CreateBy,
		CreateTime: obj.CreateTime,
		UpdateBy:   obj.UpdateBy,
		UpdateTime: obj.UpdateTime,
		Remark:     obj.Remark,
	}
}

func (repo *SysMenuRepository) toDomain(obj dao.SysMenu) domain.SysMenu {
	return domain.SysMenu{
		MenuID:     obj.MenuID,
		MenuName:   obj.MenuName,
		ParentID:   obj.ParentID,
		OrderNum:   obj.OrderNum,
		Path:       obj.Path,
		Component:  obj.Component,
		Query:      obj.Query,
		RouteName:  obj.RouteName,
		IsFrame:    obj.IsFrame,
		IsCache:    obj.IsCache,
		MenuType:   obj.MenuType,
		Visible:    obj.Visible,
		Status:     obj.Status,
		Perms:      obj.Perms,
		Icon:       obj.Icon,
		CreateBy:   obj.CreateBy,
		CreateTime: obj.CreateTime,
		UpdateBy:   obj.UpdateBy,
		UpdateTime: obj.UpdateTime,
		Remark:     obj.Remark,
	}
}

func (repo *SysMenuRepository) toDomainList(daoList []dao.SysMenu) []domain.SysMenu {
	domainList := []domain.SysMenu{}
	for _, daoObj := range daoList {
		domainObj := repo.toDomain(daoObj)
		domainList = append(domainList, domainObj)
	}
	return domainList
}
