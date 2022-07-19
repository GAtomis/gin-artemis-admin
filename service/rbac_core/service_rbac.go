/*
 * @Description: 业务层 RBAC
 * @Author: Gavin
 * @Date: 2022-07-19 10:49:19
 * @LastEditTime: 2022-07-19 19:16:16
 * @LastEditors: Gavin
 */
package rbac_core

import (
	"Artemis-admin-web/model/RBAC/request"
	"errors"
)

type RBACApi struct {
}

func (c RBACApi) QueryItemById(id string) (error, any) {
	return errors.New("错误"), "sdasd"
}
func (c RBACApi) CreateItem(body request.SysUser) (error, any) {
	return errors.New("sda"), ""
}

// func (c RBACApi) UpdateItemByPut() (error, any) {

// }
// func (c RBACApi) ChangeItembyDelete() (error, any) {

// }

func GetApi() RBACApi {

	c := RBACApi{}
	return c
}
