/*
 * @Description: 业务层 RBAC
 * @Author: Gavin
 * @Date: 2022-07-19 10:49:19
 * @LastEditTime: 2022-07-20 16:09:05
 * @LastEditors: Gavin
 */
package rbac_core

import (
	"Artemis-admin-web/model/RBAC/request"
	"Artemis-admin-web/utils"

	"gorm.io/gorm"
)

type RBACApi struct {
}

func (c RBACApi) CreateItem(body request.SysRole) (*request.SysRole, error) {

	res, err := utils.InitSQL(func(db *gorm.DB) (any, error) {

		var sr request.SysRole = request.SysRole{}
		if err := db.AutoMigrate(&sr); err != nil {
			return &body, err
		}
		err := db.Create(&body).Error

		return &body, err

	})

	return res.(*request.SysRole), err
}

// func (c RBACApi) UpdateItemByPut() (error, any) {

// }
// func (c RBACApi) ChangeItembyDelete() (error, any) {

// }

/**
 * @description: 返回一个RBACApi实例
 * @return {*}
 * @Date: 2022-07-20 11:40:52
 */
func GetRBACApi() RBACApi {

	c := RBACApi{}
	return c
}
