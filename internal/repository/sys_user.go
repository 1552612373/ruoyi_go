package repository

import (
	"context"
	"go_ruoyi_base/internal/domain"
	"go_ruoyi_base/internal/repository/dao"
)

type SysUserRepository struct {
	dao *dao.SysUserDAO
}

func NewSysUserRepository(dao *dao.SysUserDAO) *SysUserRepository {
	return &SysUserRepository{
		dao: dao,
	}
}

func (repo *SysUserRepository) Create(ctx context.Context, obj domain.SysUser) error {
	return repo.dao.Insert(ctx, repo.toDao(obj))
}

func (repo *SysUserRepository) FindByAccount(ctx context.Context, account string) (domain.SysUser, error) {
	daoSysUser, err := repo.dao.FindByAccount(ctx, account)
	domainSysUser := repo.toDomain(daoSysUser)
	return domainSysUser, err
}

func (repo *SysUserRepository) FindById(ctx context.Context, id int64) (domain.SysUser, error) {
	daoSysUser, err := repo.dao.FindById(ctx, id)
	domainSysUser := repo.toDomain(daoSysUser)
	return domainSysUser, err
}

func (repo *SysUserRepository) QueryList(ctx context.Context, pageNum int, pageSize int) ([]domain.SysUser, int, error) {
	daoList, total, err := repo.dao.QueryList(ctx, pageNum, pageSize)
	return repo.toDomainList(daoList), total, err
}

// 查看通用系统用户：岗位post列表和角色role列表
func (repo *SysUserRepository) GetSystemUserBase(ctx context.Context) ([]dao.SysPost, []dao.SysRole, error) {
	return repo.dao.GetSystemUserBase(ctx)
}

func (repo *SysUserRepository) toDao(obj domain.SysUser) dao.SysUser {
	return dao.SysUser{
		UserID:        obj.UserID,
		DeptID:        obj.DeptID,
		UserName:      obj.UserName,
		NickName:      obj.NickName,
		UserType:      obj.UserType,
		Email:         obj.Email,
		Phonenumber:   obj.Phonenumber,
		Sex:           obj.Sex,
		Avatar:        obj.Avatar,
		Password:      obj.Password,
		Status:        obj.Status,
		DelFlag:       obj.DelFlag,
		LoginIP:       obj.LoginIP,
		LoginDate:     obj.LoginDate,
		PwdUpdateDate: obj.PwdUpdateDate,
		CreateBy:      obj.CreateBy,
		CreateTime:    obj.CreateTime,
		UpdateBy:      obj.UpdateBy,
		UpdateTime:    obj.UpdateTime,
		Remark:        obj.Remark,
	}
}

func (repo *SysUserRepository) toDomain(obj dao.SysUser) domain.SysUser {
	return domain.SysUser{
		UserID:        obj.UserID,
		DeptID:        obj.DeptID,
		UserName:      obj.UserName,
		NickName:      obj.NickName,
		UserType:      obj.UserType,
		Email:         obj.Email,
		Phonenumber:   obj.Phonenumber,
		Sex:           obj.Sex,
		Avatar:        obj.Avatar,
		Password:      obj.Password,
		Status:        obj.Status,
		DelFlag:       obj.DelFlag,
		LoginIP:       obj.LoginIP,
		LoginDate:     obj.LoginDate,
		PwdUpdateDate: obj.PwdUpdateDate,
		CreateBy:      obj.CreateBy,
		CreateTime:    obj.CreateTime,
		UpdateBy:      obj.UpdateBy,
		UpdateTime:    obj.UpdateTime,
		Remark:        obj.Remark,
	}
}

func (repo *SysUserRepository) toDomainList(daoList []dao.SysUser) []domain.SysUser {
	domainList := []domain.SysUser{}
	for _, daoObj := range daoList {
		domainObj := repo.toDomain(daoObj)
		domainList = append(domainList, domainObj)
	}
	return domainList
}
