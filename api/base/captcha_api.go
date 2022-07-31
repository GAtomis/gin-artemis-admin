/*
 * @Description: api层 验证码
 * @Author: Gavin
 * @Date: 2022-07-21 14:15:00
 * @LastEditTime: 2022-07-22 19:12:13
 * @LastEditors: Gavin
 */
package base

import (
	"Artemis-admin-web/utils"

	"github.com/gin-gonic/gin"
)

func GetCaptcha(c *gin.Context) {
	utils.Captcha(c, 4)

}
