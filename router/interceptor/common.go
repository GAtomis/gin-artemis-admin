/*
 * @Description: 公共拦截器
 * @Author: Gavin
 * @Date: 2022-07-19 10:44:07
 * @LastEditTime: 2022-07-20 17:16:51
 * @LastEditors: Gavin
 */
//模拟中间件2

package interceptor

import (
	"github.com/gin-gonic/gin"
)

/**
 * @description: 拦截器
 * @param  [] func {*gin.Context} c before和after 请求拦截器
 * @return {*}
 * @Date: 2022-07-19 17:22:56
 */
func MiddleCommon(args ...func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {

		if len(args) > 0 {
			before := args[0]
			before(c)

		}

		c.Next()
		if len(args) > 1 {
			after := args[1]
			after(c)
		}

	}
}
