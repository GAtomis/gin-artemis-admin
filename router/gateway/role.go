/*
 * @Description: 路由层 RBAC
 * @Author: Gavin
 * @Date: 2022-07-19 16:30:11
 * @LastEditTime: 2022-11-10 12:10:12
 * @LastEditors: Gavin 850680822@qq.com
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
	roleGateway.GET("role", RBAC.GetRole)
	roleGateway.PUT("role", RBAC.UpdateRole)
	roleGateway.GET("list", RBAC.GetRoleList)
	roleGateway.PUT("updateRoleOfPermission", RBAC.UpdateRoleOfPermission)

}
