/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-22 16:00:03
 * @LastEditTime: 2022-11-15 15:05:41
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

type USER_INFO_API struct {
}

func (uia *USER_INFO_API) CreateUser(ctx *gin.Context) {

	var newUser request.SysUserInfo

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
	utils.OkDM(res.ID, "success", ctx)

}

func (uia *USER_INFO_API) GetUserInfo(ctx *gin.Context) {
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
		utils.OkDM(res, "success", ctx)

	} else {

		if tokenInfo != nil {
			utils.OkDM(tokenInfo.UserInfo, "success", ctx)
		} else {
			utils.Fail(ctx)
		}
	}

}
func (uia *USER_INFO_API) UpdateUser(ctx *gin.Context) {
	jwt := new(utils.JWT)
	tokenInfo := jwt.GetUserInfo(ctx)
	var user request.SysUserInfo
	ctx.ShouldBindJSON(&user)

	if tokenInfo != nil {
		api := new(rbac_core.User)
		res, err2 := api.UpdateItem(user)
		if err2 != nil {
			utils.Fail(ctx)
			return
		}
		utils.OkDM(res.ID, "success", ctx)
	} else {
		utils.Fail(ctx)
	}

}
