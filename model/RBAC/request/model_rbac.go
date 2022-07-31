/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-18 16:58:15
 * @LastEditTime: 2022-07-31 10:26:09
 * @LastEditors: Gavin
 */
package request

import (
	"Artemis-admin-web/model/global"
)

type SysUser struct {
	global.DBModel
	Name        string `json:"name" gorm:"type:varchar(32);comment:姓名;"`
	Username    string `json:"username" gorm:"type:varchar(64);comment:姓名;"`
	Password    string `json:"password" gorm:"type:varchar(32);comment:密码;"`
	Avatar      string `json:"avatar" gorm:"type:varchar(255);comment:头像;default:https://avatars.githubusercontent.com/u/40788938?v=4;"`
	JobType     string `json:"jobType" gorm:"type:varchar(32);comment:职位;"`
	Company     string `json:"company" gorm:"type:varchar(32);comment:公司;"`
	CatchPhrase string `json:"catchPhrase" gorm:"type:varchar(255);comment:个人简介;"`

	Salt   string `json:"salt" gorm:"type:varchar(128);comment:盐;"`
	Gender bool   `json:"gender" gorm:"comment:性别;default:true"`
	Locked bool   `json:"locked" gorm:"comment:账号是否锁定，1：锁定，0未锁定;default:false"`
	RoleId uint   `json:"roleId"`
}

type SysRole struct {
	global.DBModel
	Name           string          `json:"name" gorm:"type:varchar(128);comment:角色名;"`
	Available      bool            `json:"available" gorm:"comment:角色是否可用 1 Y ,0: N;"`
	SysUsers       []SysUser       `gorm:"foreignKey:RoleId;comment:拥有用户;"`
	SysPermissions []SysPermission `gorm:"many2many:sys_role_permission;comment:角色&权限;"`
}

type SysPermission struct {
	global.DBModel
	Name       string    `gorm:"type:varchar(128);comment:资源名称;"`
	Type       string    `gorm:"type:varchar(32);comment:资源类型：menu,button;"`
	Url        string    `gorm:"type:varchar(128);comment:访问url地址;"`
	Percode    string    `gorm:"type:varchar(128);comment:权限代码字符串';"`
	Parentid   string    `gorm:"type:varchar(20);comment:父结点id;"`
	Parentids  string    `gorm:"type:varchar(128);comment:父结点id列表串;"`
	Sortstring string    `gorm:"type:varchar(128);comment:排序号;"`
	Available  bool      `gorm:"comment:是否可用 1 Y ,0: N;"`
	SysRoles   []SysRole `gorm:"many2many:sys_role_permission;comment:角色&权限;"`
}
