<!--
 * @Description: 文档
 * @Author: Gavin
 * @Date: 2022-07-18 15:07:50
 * @LastEditTime: 2022-11-14 22:48:58
 * @LastEditors: Gavin 850680822@qq.com
-->
# gin-artemis-admin
后台管理基础模版,通过golang+gin+gorm,默认是连接的Mysql数据库
## 联动项目[vue-artemis-admin](https://github.com/GAtomis/vue-artemis-admin)


## 功能添加
为现在的模版添加功能
* 验证码
* RBAC鉴权
* 评论

## 功能补充
* 日志记录
* redis 缓存
* 更新菜单是否可用状态bug
* 更新了权限的等级制度
* 修改了一些bug,将主键全部替换为uuid

## 数据库配置
在config目录下建立自己的数据库配置
```
package config

import "fmt"

const (
	SQL_URL = "xxxxx"
	// SQL_URL    = "127.0.0.1"
	SQL_NAME   = "xxxx"
	SQL_CONFIG = "?charset=utf8mb4&parseTime=True&loc=Local"
	USERNAME   = "root"
	PASSWORD   = "xxxxx"
	SQL_PORT   = "3306"
	PROTOCOL   = "tcp"
	TIME_OUT   = "30s"
)

/**
 * @description: 获取sql连接
 * @return {*}
 * @Date: 2022-07-20 14:06:48
 */
func GetDsn() (dsn string) {
	dsn = fmt.Sprintf("%s:%s@%s(%s:%s)/%s%s&%s", USERNAME, PASSWORD, PROTOCOL, SQL_URL, SQL_PORT, SQL_NAME, SQL_CONFIG, TIME_OUT)
	fmt.Println(dsn)
	return
}


```