/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-22 16:05:43
 * @LastEditTime: 2022-07-27 22:50:15
 * @LastEditors: Gavin
 */
package rbac_core

import (
	"Artemis-admin-web/model/RBAC/request"
	"Artemis-admin-web/utils"

	"gorm.io/gorm"
)

type User struct {
}

func (c User) CreateItem(body request.SysUser) (*request.SysUser, error) {

	res, err := utils.InitSQL(func(db *gorm.DB) (any, error) {

		var sr request.SysUser
		if err := db.AutoMigrate(&sr); err != nil {
			return &body, err
		}
		err := db.Create(&body).Error
		return &body, err

	})

	return res.(*request.SysUser), err
}
func (c User) UpdateItem(body request.SysUser) (*request.SysUser, error) {

	res, err := utils.InitSQL(func(db *gorm.DB) (any, error) {
		var sr request.SysUser
		if err := db.AutoMigrate(&sr); err != nil {
			return &body, err
		}
		err := db.Model(&body).Updates(body).Error
		return &body, err
	})

	return res.(*request.SysUser), err
}
