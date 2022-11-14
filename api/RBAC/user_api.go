/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-22 16:00:03
 * @LastEditTime: 2022-11-14 17:16:09
 * @LastEditors: Gavin 850680822@qq.com
 */
package RBAC

import (
	"Artemis-admin-web/model/RBAC/request"
	"Artemis-admin-web/model/global"
	"Artemis-admin-web/service/rbac_core"
	"Artemis-admin-web/utils"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {

	var newUser request.SysUser

	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		utils.FailM(err.Error(), ctx)
		return
	}
	api := new(rbac_core.User)
	res, err2 := api.CreateItem(newUser)
	if err2 != nil {
		utils.FailM(err2.Error(), ctx)
		return
	}
	utils.OkDM(res.ID, "操作成功", ctx)

}

func GetUserInfo(ctx *gin.Context) {
	jwt := new(utils.JWT)
	tokenInfo := jwt.GetUserInfo(ctx)

	if noCache, _ := ctx.GetQuery("noCache"); noCache == "1" {

		var pk global.Primarykey
		pk.ID = tokenInfo.UserInfo.ID
		api := new(rbac_core.User)

		res, err2 := api.GetItem(&pk)
		if err2 != nil {
			utils.Fail(ctx)
			return
		}
		utils.OkDM(res, "操作成功", ctx)

	} else {

		if tokenInfo != nil {
			utils.OkDM(tokenInfo.UserInfo, "操作成功", ctx)
		} else {
			utils.Fail(ctx)
		}
	}

}
func UpdateUser(ctx *gin.Context) {
	jwt := new(utils.JWT)
	tokenInfo := jwt.GetUserInfo(ctx)
	var user request.SysUser
	ctx.ShouldBindJSON(&user)

	if tokenInfo != nil {
		api := new(rbac_core.User)
		res, err2 := api.UpdateItem(user)
		if err2 != nil {
			utils.Fail(ctx)
			return
		}
		utils.OkDM(res.ID, "操作成功", ctx)
	} else {
		utils.Fail(ctx)
	}

}
