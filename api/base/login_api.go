/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-22 17:27:56
 * @LastEditTime: 2022-11-15 12:52:32
 * @LastEditors: Gavin 850680822@qq.com
 */
package base

import (
	"Artemis-admin-web/model/base/request"
	"Artemis-admin-web/model/base/response"
	"Artemis-admin-web/model/global"
	"Artemis-admin-web/service/base_core"
	"Artemis-admin-web/service/rbac_core"
	"Artemis-admin-web/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var login request.Login
	err := c.ShouldBindJSON(&login)

	if err != nil {
		utils.Fail(c)
		return
	}
	if utils.CaptchaVerify(c, login.Code) {
		b := base_core.Base{}
		if sul, err2 := b.Login(login); err2 != nil {
			utils.FailM(err2.Error(), c)

		} else {

			api := new(rbac_core.User)
			pk := global.Primarykey{
				ID: sul.ID,
			}
			sui, err2 := api.GetItem(&pk)

			jwt := new(utils.JWT)
			s, err2 := jwt.InitJWT(sui)
			if err2 == nil {
				utils.OkDM(response.Token{
					Token: s,
				}, "success", c)
			}

		}

		// if err2 == nil {
		// 	utils.OkDM(token, "success", c)

		// } else {
		// 	utils.FailDM("", err2.Error(), c)
		// }

		return
	}
	utils.FailDM("", "Verification code error", c)
	return

}
