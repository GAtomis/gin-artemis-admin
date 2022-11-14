/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-22 16:05:43
 * @LastEditTime: 2022-11-14 17:13:29
 * @LastEditors: Gavin 850680822@qq.com
 */
package rbac_core

import (
	"Artemis-admin-web/model/RBAC/request"
	"Artemis-admin-web/model/global"
	"Artemis-admin-web/utils"
)

type User struct {
}

func (c User) CreateItem(body request.SysUser) (*request.SysUser, error) {
	db := utils.GAA_SQL.GetDB()
	err := db.Create(&body).Error
	return &body, err
}
func (c User) UpdateItem(body request.SysUser) (*request.SysUser, error) {
	db := utils.GAA_SQL.GetDB()
	err := db.Model(&body).Updates(body).Error
	return &body, err
}

// 查
func (c User) GetItem(body *global.Primarykey) (request.SysUser, error) {
	db := utils.GAA_SQL.GetDB()
	sr := request.SysUser{
		DBModel: global.DBModel{
			ID: body.ID,
		},
	}
	err2 := db.First(&sr).Error
	return sr, err2
}
