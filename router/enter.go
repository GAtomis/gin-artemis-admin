/*
 * @Description: gin启动服务
 * @Author: Gavin
 * @Date: 2022-07-20 16:10:20
 * @LastEditTime: 2022-07-20 16:23:01
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
	new(gateway.RoleRouter).InitRoleRouter(rg)

	r.Run(":9527")
}
