/*
 * @Description: 路由层 RBAC
 * @Author: Gavin
 * @Date: 2022-07-19 16:30:11
 * @LastEditTime: 2022-07-20 11:01:55
 * @LastEditors: Gavin
 */
package gateway

import (
	"Artemis-admin-web/api/RBAC"
	"Artemis-admin-web/router/interceptor"

	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

func (r *RoleRouter) InitRoleRouter(g *gin.RouterGroup) {
	roleGateway := g.Group("role").Use(interceptor.MiddleCommon())
	roleGateway.POST("role", RBAC.CreateRole)
}
