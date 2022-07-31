/*
 * @Description: 业务层 RBAC
 * @Author: Gavin
 * @Date: 2022-07-19 10:49:19
 * @LastEditTime: 2022-07-31 10:21:00
 * @LastEditors: Gavin
 */
package rbac_core

import (
	"Artemis-admin-web/model/RBAC/request"
	"Artemis-admin-web/model/global"
	"Artemis-admin-web/utils"

	"gorm.io/gorm"
)

type Role struct {
}

func (c Role) CreateItem(body request.SysRole) (*request.SysRole, error) {

	res, err := utils.InitSQL(func(db *gorm.DB) (any, error) {

		var sr request.SysRole
		if err := db.AutoMigrate(&sr); err != nil {
			return &body, err
		}
		err := db.Create(&body).Error

		return &body, err

	})

	return res.(*request.SysRole), err
}
func (r Role) GetRoleList(body *request.SysRole, info *global.PageInfo) ([]request.SysRole, error) {

	res, err := utils.InitSQL(func(db *gorm.DB) (any, error) {
		limit := info.PageSize
		offset := info.PageSize * (info.Page - 1)

		var temp = new(request.SysRole)
		if err := db.AutoMigrate(temp); err != nil {
			return nil, err
		}
		var result []request.SysRole
		err2 := db.Model(temp).Limit(limit).Offset(offset).Preload("SysUser").Find(&result).Error
		return result, err2
	})
	return res.([]request.SysRole), err

}

/**
 * @description: 返回一个RBACApi实例
 * @return {*}
 * @Date: 2022-07-20 11:40:52
 */
func GetRBACApi() Role {

	c := Role{}
	return c
}
