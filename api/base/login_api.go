/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-22 17:27:56
 * @LastEditTime: 2022-07-25 12:01:15
 * @LastEditors: Gavin
 */
package base

import (
	"Artemis-admin-web/model/base/request"
	"Artemis-admin-web/service/base_core"
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
		token, err2 := b.Login(login)
		if err2 == nil {
			utils.OkDM(token, "验证成功", c)

		} else {
			utils.FailDM("", err2.Error(), c)
		}

		return
	}
	utils.FailDM("", "验证码错误", c)
	return

}
