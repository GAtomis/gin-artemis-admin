/*
 * @Description:权限
 * @Author: Gavin
 * @Date: 2022-08-03 10:53:48
 * @LastEditTime: 2022-08-13 14:39:11
 * @LastEditors: Gavin
 */
package RBAC

import (
	"Artemis-admin-web/model/RBAC/request"
	"Artemis-admin-web/model/global"
	"Artemis-admin-web/service/rbac_core"
	"Artemis-admin-web/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type PERMISSION_API struct {
}

func (api *PERMISSION_API) GetPermissionByRoleId(ctx *gin.Context) {

	jwt := new(utils.JWT)
	MyClaims := jwt.GetUserInfo(ctx)
	//载入api
	r2 := new(rbac_core.Permission)
	sp, err := r2.GetItemByRoleId(MyClaims.UserInfo.RoleId)
	if err != nil {
		utils.FailM("查询失败", ctx)
		return
	}

	utils.OkDM(sp, "success", ctx)

}

/**
 * @description: 、获得所有路由权限列表
 * @param {*gin.Context} ctx
 * @return {*}
 * @Date: 2022-08-04 17:57:34
 */
func (api *PERMISSION_API) GetPermissionList(ctx *gin.Context) {
	var pageInfo global.PageInfo
	_ = ctx.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		utils.FailM(err.Error(), ctx)
		return
	}
	var permission request.SysPermission
	r2 := new(rbac_core.Permission)
	if res, err2 := r2.GetList(&permission, &pageInfo); err2 != nil {
		utils.FailDM(res, err2.Error(), ctx)
	} else {
		utils.OkDM(res, "success", ctx)
	}

}

/**
 * @description: 修改角色
 * @param {*gin.Context} ctx
 * @return {*}
 * @Date: 2022-08-04 18:02:51
 */
func (api *PERMISSION_API) UpdatePermission(ctx *gin.Context) {
	//声明一个SysPermission
	var newSysPermission request.SysPermission
	//成功JSON化
	err := ctx.ShouldBindJSON(&newSysPermission)
	if err != nil {
		utils.Fail(ctx)
		return
	}
	// fmt.Printf("newSysPermission.SysPermissions: %v\n", newSysPermission.SysPermissions)
	//载入api
	fmt.Printf("%+v", newSysPermission)
	r2 := new(rbac_core.Permission)
	res, err2 := r2.UpdateItem(&newSysPermission)
	if err2 != nil {
		utils.Fail(ctx)
		return
	}
	utils.OkDM(res.ID, "success", ctx)

}

// @Tags AuthorityMenu
// @Summary 创建用户角色
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {object} response.Response{data=systemRes.SysMenusResponse,msg=string} "获取用户动态路由,返回包括系统菜单详情列表"
// @Router /menu/getMenu [post]
func (api *PERMISSION_API) CreatePermission(ctx *gin.Context) {
	//声明一个SysPermission
	var newSysPermission request.SysPermission
	//成功JSON化
	err := ctx.ShouldBindJSON(&newSysPermission)
	if err != nil {
		utils.FailM(err.Error(), ctx)
		return
	}
	//载入api
	r2 := new(rbac_core.Permission)
	res, err2 := r2.CreateItem(&newSysPermission)
	if err2 != nil {
		utils.FailM(err2.Error(), ctx)
		return
	}
	utils.OkDM(res.ID, "success", ctx)
}
func (api *PERMISSION_API) DeletePermission(ctx *gin.Context) {
	//声明一个SysPermission
	var newSysPermission request.SysPermission
	//成功JSON化
	err := ctx.ShouldBindJSON(&newSysPermission)
	if err != nil {
		utils.Fail(ctx)
		return
	}
	// fmt.Printf("newSysPermission.SysPermissions: %v\n", newSysPermission.SysPermissions)
	//载入api
	r2 := new(rbac_core.Permission)
	res, err2 := r2.DelItem(&newSysPermission)
	if err2 != nil {
		utils.Fail(ctx)
		return
	}
	utils.OkDM(res.ID, "success", ctx)
}

func (api *PERMISSION_API) GetPermission(ctx *gin.Context) {
	var Primarykey global.Primarykey
	_ = ctx.ShouldBindQuery(&Primarykey)
	if err := utils.Verify(Primarykey, utils.Primarykey); err != nil {
		utils.FailM(err.Error(), ctx)
		return
	}
	//载入api
	r2 := new(rbac_core.Permission)

	res, err2 := r2.GetItem(&Primarykey)
	if err2 != nil {
		utils.Fail(ctx)
		return
	}
	utils.OkDM(*res, "success", ctx)
}
