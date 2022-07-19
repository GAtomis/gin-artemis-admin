/*
 * @Description: api层 RBAC
 * @Author: Gavin
 * @Date: 2022-07-19 17:56:36
 * @LastEditTime: 2022-07-19 21:37:34
 * @LastEditors: Gavin
 */
package RBAC

import (
	"Artemis-admin-web/model/RBAC/request"
	"fmt"

	"github.com/gin-gonic/gin"
)

// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {object} response.Response{data=systemRes.SysMenusResponse,msg=string} "获取用户动态路由,返回包括系统菜单详情列表"
// @Router /menu/getMenu [post]
func CreateRole(ctx *gin.Context) {

	var r request.SysRole
	err := ctx.ShouldBindJSON(&r)

	if err != nil {
		ctx.JSON(200, gin.H{
			"msg":  "报错了",
			"data": []any{},
		})
	} else {
		fmt.Printf("r: %v\n", r)
		ctx.JSON(200, gin.H{
			"msg": "success",
		})
	}

	// if menus, err := menuService.GetMenuTree(utils.GetUserAuthorityId(c)); err != nil {
	// 	global.GVA_LOG.Error("获取失败!", zap.Error(err))
	// 	response.FailWithMessage("获取失败", c)
	// } else {
	// 	if menus == nil {
	// 		menus = []system.SysMenu{}
	// 	}
	// 	response.OkWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "获取成功", c)
	// }
}
