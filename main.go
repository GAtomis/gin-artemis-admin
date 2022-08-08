/*
 * @Description: 初始化
 * @Author: Gavin
 * @Date: 2022-07-18 15:11:28
 * @LastEditTime: 2022-08-08 19:42:52
 * @LastEditors: Gavin
 */
package main

import (
	"Artemis-admin-web/router"
	"Artemis-admin-web/utils"
	"sync"
)

/**
 * @description: 入口函数....
 * @return {*}
 * @Date: 2022-07-18 15:12:06
 */
//声明一个wg类型
var wg sync.WaitGroup

func main() {

	db, _ := utils.GAA_SQL.StartSQL()
	if db != nil {
		sqlDB, _ := db.DB()
		defer sqlDB.Close()
	}

	router.Start()

}
