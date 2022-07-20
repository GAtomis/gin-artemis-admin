/*
 * @Description: 路由层 RBAC
 * @Author: Gavin
 * @Date: 2022-07-19 16:30:11
 * @LastEditTime: 2022-07-20 17:19:36
 * @LastEditors: Gavin
 */
package gateway

import (
	"Artemis-admin-web/api/RBAC"
	"Artemis-admin-web/router/interceptor"
	"fmt"

	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

func (r *RoleRouter) InitRoleRouter(g *gin.RouterGroup) {
	roleGateway := g.Group("role").Use(interceptor.MiddleCommon(func(c *gin.Context) {

		fmt.Println("请求前")
	}, func(c *gin.Context) {

		fmt.Println("请求后")

	}))
	roleGateway.POST("role", RBAC.CreateRole)
}
