/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-21 18:20:29
 * @LastEditTime: 2022-08-08 18:29:03
 * @LastEditors: Gavin
 */
package utils

import (
	reqRBAC "Artemis-admin-web/model/RBAC/request"
	"Artemis-admin-web/model/base/response"
	"errors"

	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)
var mySigningKey []byte

/**
 * @description: 创建一个令牌 jwt
 * @param {string} u 用户名
 * @return {string, error} 返回一个token
 * @Date: 2022-07-21 18:43:59
 */
func (j *JWT) InitJWT(u reqRBAC.SysUserInfo) (string, error) {

	c := response.MyClaims{
		UserInfo: u,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,      //开始时间
			ExpiresAt: time.Now().Unix() + 60*60*6, //过期时间
			Issuer:    u.Name,                      //戳
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	s, err := token.SignedString(mySigningKey) //加密
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return "", err
	}

	return s, nil

}

//解析一个令牌
// func (j *JWT) ParseToken(tokenString string) (*response.MyClaims, error) {

// 	token, err := jwt.ParseWithClaims(tokenString, &response.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return mySigningKey, nil
// 	})

// 	if claims, ok := token.Claims.(*response.MyClaims); ok && token.Valid {
// 		return claims, err
// 	} else {
// 		return nil, err
// 	}

// }

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*response.MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &response.MyClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return mySigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*response.MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}

// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户角色id
func (j *JWT) GetUserInfo(c *gin.Context) *response.MyClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := j.CheckClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(response.MyClaims)
		return &waitUse
	}
}

// 检查请求头中是否有token
func (j *JWT) CheckClaims(c *gin.Context) (*response.MyClaims, error) {
	token := c.Request.Header.Get("Authorization")

	claims, err := j.ParseToken(token)
	if err != nil {
		return nil, errors.New("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}
