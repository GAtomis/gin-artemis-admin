/*
 * @Description: api层 RBAC
 * @Author: Gavin
 * @Date: 2022-07-19 17:56:36
 * @LastEditTime: 2022-07-31 09:55:10
 * @LastEditors: Gavin
 */
package RBAC

import (
	"Artemis-admin-web/model/RBAC/request"
	"Artemis-admin-web/model/global"
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
	r2 := new(rbac_core.Role)
	res, err2 := r2.CreateItem(newSysRole)
	if err2 != nil {
		utils.Fail(ctx)
		return
	}
	utils.OkDM(res.ID, "操作成功", ctx)
}
func GetRoleList(ctx *gin.Context) {
	var pageInfo global.PageInfo
	_ = ctx.ShouldBindQuery(&pageInfo)

	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		utils.FailM(err.Error(), ctx)
		return
	}
	var role request.SysRole
 	rbac_core.GetRoleList(&role, &pageInfo) {

	}

}
