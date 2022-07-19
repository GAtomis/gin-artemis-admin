/*
 * @Description: 公共拦截器
 * @Author: Gavin
 * @Date: 2022-07-19 10:44:07
 * @LastEditTime: 2022-07-19 18:26:45
 * @LastEditors: Gavin
 */
//模拟中间件2

package interceptor

import (
	"fmt"

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
		before := args[0]
		if before != nil {
			before(c)
		}
		fmt.Println("before next2")
		c.Next()
		after := args[1]
		if after != nil {
			after(c)
		}
		fmt.Println("after next2")
	}
}
