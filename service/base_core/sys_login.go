/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-21 18:01:53
 * @LastEditTime: 2022-11-15 12:32:24
 * @LastEditors: Gavin 850680822@qq.com
 */
package base_core

import (
	reqRBAC "Artemis-admin-web/model/RBAC/request"
	reqBase "Artemis-admin-web/model/base/request"
	"Artemis-admin-web/utils"
	"errors"
)

type Base struct {
}

// func (b Base) OldLogin(body reqBase.Login) (*respBase.Token, error) {
// 	res, err := utils.InitSQL(func(db *gorm.DB) (any, error) {
// 		var sr reqRBAC.SysUserLogin
// 		if err := db.AutoMigrate(&sr); err != nil {
// 			return &body, err
// 		}
// 		err := db.Where("username = ? AND password = ?", body.Username, body.Password).First(&sr).Error
// 		fmt.Printf("sr: %v\n", sr)
// 		return &sr, err

//		})
//		if err != nil {
//			return &respBase.Token{Token: "获取失败"}, err
//		}
//		if isMatch(&body, res.(*reqRBAC.SysUserLogin)) {
//			jwt := new(utils.JWT)
//			s, err2 := jwt.InitJWT(*res.(*reqRBAC.SysUserLogin))
//			if err2 == nil {
//				return &respBase.Token{
//					Token: s,
//				}, err
//			}
//		}
//		return nil, err
//	}
func (b Base) Login(body reqBase.Login) (*reqRBAC.SysUserLogin, error) {

	db := utils.GAA_SQL.GetDB()

	var sr reqRBAC.SysUserLogin
	// if err := db.AutoMigrate(&sr); err != nil {
	// 	return nil, err
	// }

	if err := db.Where("username = ?", body.Username).First(&sr).Error; err != nil {
		return nil, err
	}
	if isMatch(&body, &sr) {

		return &sr, nil
		// jwt := new(utils.JWT)
		// s, err2 := jwt.InitJWT(sr)
		// if err2 == nil {
		// 	return &respBase.Token{
		// 		Token: s,
		// 	}, err
		// }
	}
	return nil, errors.New("Username or password incorrect")
}

func isMatch(target *reqBase.Login, source *reqRBAC.SysUserLogin) bool {
	return target.Password == source.Password
}
