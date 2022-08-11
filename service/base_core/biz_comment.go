/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-08-10 13:40:27
 * @LastEditTime: 2022-08-10 16:58:40
 * @LastEditors: Gavin
 */
package base_core

import (
	"Artemis-admin-web/model/business/request"
	"Artemis-admin-web/model/global"
	"Artemis-admin-web/utils"
)

type Comment struct {
}

func (c Comment) GetCommentsByUid(body *global.Primarykey) (*[]request.BizComment, error) {
	db := utils.GAA_SQL.GetDB()
	var bc []request.BizComment
	err := db.Where("user_id= ? ", body.ID).Find(&bc).Error
	return &bc, err
}
func (c Comment) CreateItem(body *request.BizComment) (*request.BizComment, error) {
	db := utils.GAA_SQL.GetDB()

	// if err := db.AutoMigrate(*body); err != nil {
	// 	return nil, err
	// }
	err := db.Create(body).Error

	return body, err
}
