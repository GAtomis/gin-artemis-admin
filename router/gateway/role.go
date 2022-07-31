/*
 * @Description: 路由层 RBAC
 * @Author: Gavin
 * @Date: 2022-07-19 16:30:11
 * @LastEditTime: 2022-07-31 09:55:37
 * @LastEditors: Gavin
 */
package gateway

import (
	"Artemis-admin-web/api/RBAC"
	"Artemis-admin-web/router/interceptor"

	"github.com/gin-gonic/gin"
)

type Router struct{}

func (r *Router) InitRoleRouter(g *gin.RouterGroup) {
	roleGateway := g.Group("role").Use(interceptor.MiddleCommon()).Use(interceptor.JWTAuth())
	roleGateway.POST("role", RBAC.CreateRole)
	roleGateway.GET("role", RBAC.GetRoleList)

}
