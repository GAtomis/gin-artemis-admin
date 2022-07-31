/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-21 15:57:48
 * @LastEditTime: 2022-07-22 19:12:38
 * @LastEditors: Gavin
 */

package gateway

import (
	"Artemis-admin-web/api/base"
	"Artemis-admin-web/router/interceptor"

	"github.com/gin-gonic/gin"
)

func (r *Router) InitCaptcha(g *gin.RouterGroup) {
	roleGateway := g.Group("base").Use(interceptor.Session("golang-tech-stack"))
	roleGateway.GET("code", base.GetCaptcha)
	roleGateway.POST("login", base.Login)
}
