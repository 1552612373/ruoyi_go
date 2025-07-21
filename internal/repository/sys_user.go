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

func (repo *SysUserRepository) toDao(obj domain.SysUser) dao.SysUser {
	return dao.SysUser{
		UserId:        obj.UserId,
		DeptId:        obj.DeptId,
		UserName:      obj.UserName,
		NickName:      obj.NickName,
		Email:         obj.Email,
		Avatar:        obj.Avatar,
		Phonenumber:   obj.Phonenumber,
		Sex:           obj.Sex,
		Password:      obj.Password,
		Status:        obj.Status,
		DelFlag:       obj.DelFlag,
		LoginIp:       obj.LoginIp,
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
		UserId:        obj.UserId,
		DeptId:        obj.DeptId,
		UserName:      obj.UserName,
		NickName:      obj.NickName,
		Email:         obj.Email,
		Avatar:        obj.Avatar,
		Phonenumber:   obj.Phonenumber,
		Sex:           obj.Sex,
		Password:      obj.Password,
		Status:        obj.Status,
		DelFlag:       obj.DelFlag,
		LoginIp:       obj.LoginIp,
		LoginDate:     obj.LoginDate,
		PwdUpdateDate: obj.PwdUpdateDate,
		CreateBy:      obj.CreateBy,
		CreateTime:    obj.CreateTime,
		UpdateBy:      obj.UpdateBy,
		UpdateTime:    obj.UpdateTime,
		Remark:        obj.Remark,
	}
}
