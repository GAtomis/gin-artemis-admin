/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-08-03 10:58:11
 * @LastEditTime: 2022-08-08 23:00:46
 * @LastEditors: Gavin
 */
package gateway

import (
	"Artemis-admin-web/api/RBAC"
	"Artemis-admin-web/router/interceptor"

	"github.com/gin-gonic/gin"
)

func (r *Router) InitPermissionRouter(g *gin.RouterGroup) {
	roleGateway := g.Group("permission").Use(interceptor.MiddleCommon()).Use(interceptor.JWTAuth())

	api := new(RBAC.PERMISSION_API)
	roleGateway.GET("getListByRoleId", api.GetPermissionByRoleId)
	roleGateway.GET("list", api.GetPermissionList)
	roleGateway.GET("permission", api.GetPermission)
	roleGateway.POST("permission", api.CreatePermission)
	roleGateway.PUT("permission", api.UpdatePermission)
	roleGateway.DELETE("permission", api.DeletePermission)

}
