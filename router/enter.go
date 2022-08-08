/*
 * @Description: gin启动服务
 * @Author: Gavin
 * @Date: 2022-07-20 16:10:20
 * @LastEditTime: 2022-08-03 11:01:36
 * @LastEditors: Gavin
 */
package router

import (
	"Artemis-admin-web/router/gateway"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	rg := r.Group("api")
	r2 := new(gateway.Router)
	r2.InitCaptcha(rg)
	r2.InitRoleRouter(rg)
	r2.InitUserRouter(rg)
	r2.InitPermissionRouter(rg)

	r.Run(":8888")
}
