/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-26 08:50:58
 * @LastEditTime: 2022-07-27 22:06:36
 * @LastEditors: Gavin
 */
package rbac_core

type Permission struct {
}

// func (p Permission) CreateItem(body request.SysPermission) {
// 	res, err := utils.InitSQL(func(db *gorm.DB) (any, error) {

// 		var sr request.SysRole
// 		if err := db.AutoMigrate(&sr); err != nil {
// 			return &body, err
// 		}
// 		err := db.Create(&body).Error

// 		return &body, err

// 	})

// 	return res.(*request.SysRole), err

// }
