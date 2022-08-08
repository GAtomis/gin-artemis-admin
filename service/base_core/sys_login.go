/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-21 18:01:53
 * @LastEditTime: 2022-08-08 22:12:13
 * @LastEditors: Gavin
 */
package base_core

import (
	reqRBAC "Artemis-admin-web/model/RBAC/request"
	reqBase "Artemis-admin-web/model/base/request"
	respBase "Artemis-admin-web/model/base/response"
	"Artemis-admin-web/utils"
	"fmt"

	"gorm.io/gorm"
)

type Base struct {
}

func (b Base) OldLogin(body reqBase.Login) (*respBase.Token, error) {
	res, err := utils.InitSQL(func(db *gorm.DB) (any, error) {
		var sr reqRBAC.SysUser
		if err := db.AutoMigrate(&sr); err != nil {
			return &body, err
		}
		err := db.Where("username = ? AND password = ?", body.Username, body.Password).First(&sr).Error
		fmt.Printf("sr: %v\n", sr)
		return &sr, err

	})
	if err != nil {
		return &respBase.Token{Token: "获取失败"}, err
	}
	if isMatch(&body, res.(*reqRBAC.SysUser)) {
		jwt := new(utils.JWT)
		s, err2 := jwt.InitJWT(*res.(*reqRBAC.SysUser))
		if err2 == nil {
			return &respBase.Token{
				Token: s,
			}, err
		}
	}
	return nil, err
}
func (b Base) Login(body reqBase.Login) (*respBase.Token, error) {

	db := utils.GAA_SQL.GetDB()

	var sr reqRBAC.SysUser
	// if err := db.AutoMigrate(&sr); err != nil {
	// 	return nil, err
	// }
	err := db.Where("username = ? AND password = ?", body.Username, body.Password).First(&sr).Error
	fmt.Printf("sr: %v\n", sr)

	if err != nil {
		return &respBase.Token{Token: "获取失败"}, err
	}
	if isMatch(&body, &sr) {
		jwt := new(utils.JWT)
		s, err2 := jwt.InitJWT(sr)
		if err2 == nil {
			return &respBase.Token{
				Token: s,
			}, err
		}
	}
	return nil, err
}

func isMatch(target *reqBase.Login, source *reqRBAC.SysUser) bool {
	return target.Password == source.Password
}
