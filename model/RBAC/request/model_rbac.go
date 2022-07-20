/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-18 16:58:15
 * @LastEditTime: 2022-07-20 13:57:59
 * @LastEditors: Gavin
 */
package request

import (
	"gorm.io/gorm"
)

type SysUser struct {
	gorm.Model
	Usercode string `gorm:"type:varchar(32);comment:账号;"`
	Username string `gorm:"type:varchar(64);comment:姓名;"`
	Password string `gorm:"type:varchar(32);comment:密码;"`
	Salt     string `gorm:"type:varchar(128);comment:盐;"`
	Locked   bool   `gorm:"comment:账号是否锁定，1：锁定，0未锁定;"`
	RoleId   uint
}

type SysRole struct {
	gorm.Model
	Name           string          `gorm:"type:varchar(128);comment:角色名;"`
	Available      bool            `gorm:"comment:角色是否可用 1 Y ,0: N;"`
	SysUsers       []SysUser       `gorm:"foreignKey:RoleId;comment:拥有用户;"`
	SysPermissions []SysPermission `gorm:"many2many:sys_role_permission;comment:角色&权限;"`
}

type SysPermission struct {
	gorm.Model
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
