/*
 * @Description: user模块
 * @Author: Gavin
 * @Date: 2022-07-22 13:55:29
 * @LastEditTime: 2022-11-15 15:06:59
 * @LastEditors: Gavin 850680822@qq.com
 */
package gateway

import (
	"Artemis-admin-web/api/RBAC"
	"Artemis-admin-web/router/interceptor"

	"github.com/gin-gonic/gin"
)

func (r *Router) InitUserRouter(g *gin.RouterGroup) {
	api := new(RBAC.USER_INFO_API)
	roleGateway := g.Group("user").Use(interceptor.JWTAuth())
	roleGateway.POST("user", api.CreateUser)
	roleGateway.PUT("user", api.UpdateUser)
	roleGateway.GET("getUserInfo", api.GetUserInfo)

}
