/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-27 22:17:49
 * @LastEditTime: 2022-07-31 09:55:46
 * @LastEditors: Gavin
 */
package global

import (
	"time"

	"gorm.io/gorm"
)

type DBModel struct {
	ID        uint `json:"id" gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type PageInfo struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}
