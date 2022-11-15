/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-22 16:05:43
 * @LastEditTime: 2022-11-15 11:37:51
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

func (c User) CreateItem(body request.SysUserInfo) (*request.SysUserInfo, error) {
	db := utils.GAA_SQL.GetDB()
	err := db.Create(&body).Error
	return &body, err
}
func (c User) UpdateItem(body request.SysUserInfo) (*request.SysUserInfo, error) {
	db := utils.GAA_SQL.GetDB()
	err := db.Model(&body).Updates(body).Error
	return &body, err
}

// 查
func (c User) GetItem(body *global.Primarykey) (request.SysUserInfo, error) {
	db := utils.GAA_SQL.GetDB()
	sr := request.SysUserInfo{
		DBModel: global.DBModel{
			ID: body.ID,
		},
	}
	err2 := db.First(&sr).Error
	return sr, err2
}
