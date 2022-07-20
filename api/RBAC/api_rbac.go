/*
 * @Description: api层 RBAC
 * @Author: Gavin
 * @Date: 2022-07-19 17:56:36
 * @LastEditTime: 2022-07-20 17:23:02
 * @LastEditors: Gavin
 */
package RBAC

import (
	"Artemis-admin-web/model/RBAC/request"
	"Artemis-admin-web/service/rbac_core"
	"Artemis-admin-web/utils"

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
	//声明一个SysRole
	var newSysRole request.SysRole
	//成功JSON化
	err := ctx.ShouldBindJSON(&newSysRole)
	if err != nil {
		utils.Fail(ctx)
		return
	}

	//载入api
	r2 := new(rbac_core.RBACApi)
	res, err2 := r2.CreateItem(newSysRole)
	if err2 != nil {
		utils.Fail(ctx)
		return
	}
	utils.OkDM(res.ID, "操作成功", ctx)

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
