/*
 * @Description: 工具
 * @Author: Gavin
 * @Date: 2022-07-20 11:14:40
 * @LastEditTime: 2022-11-15 12:42:04
 * @LastEditors: Gavin 850680822@qq.com
 */
package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code   int    `json:"code"`
	Result any    `json:"result"`
	Msg    string `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

/**
 * @description: 返回一个状态
 * @param {int} code 状态
 * @param {interface{}} data 数据
 * @param {string} msg
 * @param {*gin.Context} c
 * @return {*}
 * @Date: 2022-07-20 11:26:25
 */
func Result(code int, result any, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		result,
		msg,
	})
}

/**
 * @description: 成功状态 直接返回一个空
 * @param {*gin.Context} ctx gin上下文
 * @return {*}
 * @Date: 2022-07-20 11:26:23
 */
func Ok(ctx *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "success", ctx)
}

/**
 * @description: 成功 参数为data和msg
 * @param {interface{}} data 返回数据
 * @param {string} message 消息提醒
 * @param {*gin.Context} c gin上下文
 * @return {*}
 * @Date: 2022-07-20 11:26:19
 */
func OkDM(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

/**
 * @description:失败状态 直接返回一个空
 * @param {*gin.Context} ctx
 * @return {*}
 * @Date: 2022-07-20 11:31:34
 */
func Fail(ctx *gin.Context) {

	Result(ERROR, map[string]interface{}{}, "fail", ctx)

}

/**
 * @description: 失败状态 参数为data和msg
 * @param {interface{}} data 返回数据
 * @param {string} message 消息提醒
 * @param {*gin.Context} c gin上下文
 * @return {*}
 * @Date: 2022-07-20 11:37:41
 */
func FailDM(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

/**
 * @description: 失败状态 参数为data和msg
 * @param {interface{}} data 返回数据
 * @param {string} message 消息提醒
 * @param {*gin.Context} c gin上下文
 * @return {*}
 * @Date: 2022-07-20 11:37:41
 */
func FailM(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}
