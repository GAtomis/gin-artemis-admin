/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-21 15:57:48
 * @LastEditTime: 2022-08-10 16:33:37
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
func (r *Router) InitComment(g *gin.RouterGroup) {
	roleGateway := g.Group("comment").Use(interceptor.MiddleCommon()).Use(interceptor.JWTAuth())
	ca := new(base.COMMENT_API)
	roleGateway.GET("getComments", ca.GetCommentsByUid)
	roleGateway.POST("comment", ca.CreateComment)
}
