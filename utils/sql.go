/*
 * @Description: sql工具
 * @Author: Gavin
 * @Date: 2022-07-20 12:48:34
 * @LastEditTime: 2022-07-20 16:02:21
 * @LastEditors: Gavin
 */
package utils

import (
	"Artemis-admin-web/config"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
 * @description: 方法说明....
 * @param {*} cb 内容cb
 * @return {*}
 * @Date: 2022-07-20 13:53:33
 */
func InitSQL(cb func(*gorm.DB) (any, error)) (any, error) {

	dsn := config.GetDsn()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err != nil {
		return "", err
	}
	res, err2 := cb(db)
	return res, err2

}
