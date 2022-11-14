/*
 * @Description: 业务层 RBAC
 * @Author: Gavin
 * @Date: 2022-07-19 10:49:19
 * @LastEditTime: 2022-11-14 23:37:47
 * @LastEditors: Gavin 850680822@qq.com
 */
package rbac_core

import (
	"Artemis-admin-web/model/RBAC/request"
	"Artemis-admin-web/model/global"
	"Artemis-admin-web/utils"
	"fmt"
)

type Role struct {
}

// 增
func (c Role) CreateItem(body request.SysRole) (*request.SysRole, error) {
	db := utils.GAA_SQL.GetDB()
	err := db.Create(&body).Error
	return &body, err
}

// 查询列表
func (r Role) GetRoleList(body *request.SysRole, info *global.PageInfo) (map[string]any, error) {
	db := utils.GAA_SQL.GetDB()
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	var temp = new(request.SysRole)
	var result []request.SysRole
	var total int64 = 0
	if err1 := db.Model(temp).Count(&total).Error; err1 != nil {
		return nil, err1
	}
	fmt.Printf("当前用户等级%v", body.Level)
	if err2 := db.Model(temp).Where("level >= ?", body.Level).Limit(limit).Offset(offset).Preload("SysPermissions").Find(&result).Error; err2 != nil {
		return map[string]any{"item": nil, "total": total}, err2
	}
	return map[string]any{"item": result, "total": total}, nil
}

// 查
func (c Role) GetItem(body *global.Primarykey) (request.SysRole, error) {
	db := utils.GAA_SQL.GetDB()
	sr := request.SysRole{
		DBModel: global.DBModel{
			ID: body.ID,
		},
	}
	err2 := db.Preload("SysPermissions").Preload("SysUsers").First(&sr).Error
	return sr, err2
}

func (c Role) UpdateItem(body *request.SysRole) (request.SysRole, error) {
	db := utils.GAA_SQL.GetDB()

	err := db.Model(body).Omit("SysPermissions", "SysUsers").Updates(body).Error

	// err = db.Model(body).Association("SysPermissions").Replace(spl)
	return *body, err
}
func (c Role) UpdateRoleOfPermission(body *request.SysRole) (request.SysRole, error) {
	db := utils.GAA_SQL.GetDB()
	spl := body.SysPermissions

	var err error
	// if len(spl) > 0 {
	// } else {
	// 		err = db.Model(body).Association("SysPermissions").Clear()
	// }
	err = db.Model(body).Association("SysPermissions").Replace(spl)
	return *body, err
}
