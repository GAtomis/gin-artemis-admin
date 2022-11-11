/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-27 22:17:49
 * @LastEditTime: 2022-08-03 05:44:29
 * @LastEditors: Gavin
 */
package global

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type DBModel struct {
	ID        string         `json:"id" gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `json:"createdAt" `
	UpdatedAt time.Time      `json:"updatedAt" `
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (u *DBModel) BeforeCreate(tx *gorm.DB) (err error) {
	if !(len(u.ID) > 0) {
		u.ID = uuid.NewV4().String()
	}
	return
}

type PageInfo struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}

// GetAuthorityId Get role by id structure
type Primarykey struct {
	ID string `json:"id" form:"id"` // 主键
}
