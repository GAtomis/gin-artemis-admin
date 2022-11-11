/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-26 08:50:58
 * @LastEditTime: 2022-08-13 10:53:00
 * @LastEditors: Gavin
 */
package rbac_core

import (
	"Artemis-admin-web/model/RBAC/request"
	"Artemis-admin-web/model/global"
	"Artemis-admin-web/utils"
	"fmt"
)

type Permission struct {
}

// 查列表通过关联RoleId
func (p *Permission) GetItemByRoleId(ID string) ([]request.SysPermission, error) {
	db := utils.GAA_SQL.GetDB()
	fmt.Printf("ID: %v\n", ID)
	sr := request.SysRole{
		DBModel: global.DBModel{
			ID: ID,
		},
	}
	var sp []request.SysPermission

	err2 := db.Model(&sr).Association("SysPermissions").Find(&sp)
	return sp, err2
}

// 查询列表
func (r *Permission) GetList(body *request.SysPermission, info *global.PageInfo) (map[string]any, error) {

	db := utils.GAA_SQL.GetDB()
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	var temp = new(request.SysPermission)
	var result []request.SysPermission
	var total int64 = 0
	if err1 := db.Model(temp).Count(&total).Error; err1 != nil {
		return nil, err1
	}
	if err2 := db.Model(temp).Limit(limit).Offset(offset).Find(&result).Error; err2 != nil {
		return map[string]any{"item": nil, "total": total}, err2
	}
	return map[string]any{"item": result, "total": total}, nil

}

func (p *Permission) CreateItem(body *request.SysPermission) (*request.SysPermission, error) {

	db := utils.GAA_SQL.GetDB()
	err := db.Create(body).Error

	return body, err

}
func (p *Permission) UpdateItem(body *request.SysPermission) (*request.SysPermission, error) {
	db := utils.GAA_SQL.GetDB()
	err := db.Model(body).Updates(body).Error
	return body, err
}

func (p *Permission) DelItem(body *request.SysPermission) (*request.SysPermission, error) {
	db := utils.GAA_SQL.GetDB()
	err := db.Delete(body).Error
	return body, err
}
func (p *Permission) GetItem(body *global.Primarykey) (*request.SysPermission, error) {
	db := utils.GAA_SQL.GetDB()
	sp := request.SysPermission{}
	err := db.First(&sp, body.ID).Error
	return &sp, err
}
