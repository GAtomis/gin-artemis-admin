/*
 * @Description: sql工具
 * @Author: Gavin
 * @Date: 2022-07-20 12:48:34
 * @LastEditTime: 2022-11-15 13:01:10
 * @LastEditors: Gavin 850680822@qq.com
 */
package utils

import (
	"Artemis-admin-web/config"
	"Artemis-admin-web/model/RBAC/request"
	reqBiz "Artemis-admin-web/model/business/request"
	"fmt"

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

type SqlType struct {
	db *gorm.DB
}

var GAA_SQL = new(SqlType)

// 初始化cb调用方法
func InitSQL(cb func(*gorm.DB) (any, error)) (any, error) {

	dsn := config.GetDsn()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return "", err
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	res, err2 := cb(db)
	return res, err2

}

func (sql *SqlType) StartSQL() (db *gorm.DB, err error) {

	dsn := config.GetDsn()

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("数据库连接异常,%s", err.Error())
		return
	}

	sp := request.SysPermission{}
	sr := request.SysRole{}
	sui := request.SysUserInfo{}
	sul := request.SysUserLogin{}
	bc := reqBiz.BizComment{}
	db.AutoMigrate(&sr, &sul, &sui, &sp, &bc)
	sql.db = db
	return db, err
}

func (sql *SqlType) GetDB() (db *gorm.DB) {

	return sql.db

}
