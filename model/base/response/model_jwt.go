/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-21 18:23:47
 * @LastEditTime: 2022-07-23 13:17:29
 * @LastEditors: Gavin
 */
package response

import (
	"Artemis-admin-web/model/RBAC/request"

	"github.com/golang-jwt/jwt/v4"
)

// JTW模版
type MyClaims struct {
	UserInfo request.SysUserInfo
	jwt.StandardClaims
}
