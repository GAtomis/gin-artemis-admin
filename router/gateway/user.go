/*
 * @Description: user模块
 * @Author: Gavin
 * @Date: 2022-07-22 13:55:29
 * @LastEditTime: 2022-07-27 22:05:47
 * @LastEditors: Gavin
 */
package gateway

import (
	"Artemis-admin-web/api/RBAC"
	"Artemis-admin-web/router/interceptor"

	"github.com/gin-gonic/gin"
)

func (r *Router) InitUserRouter(g *gin.RouterGroup) {
	roleGateway := g.Group("user").Use(interceptor.JWTAuth())
	roleGateway.POST("user", RBAC.CreateUser)
	roleGateway.PUT("user", RBAC.UpdateUser)
	roleGateway.GET("getUserInfo", RBAC.GetUserInfo)
}
