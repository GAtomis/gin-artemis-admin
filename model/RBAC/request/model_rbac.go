/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-18 16:58:15
 * @LastEditTime: 2022-11-15 11:37:57
 * @LastEditors: Gavin 850680822@qq.com
 */
package request

import (
	"Artemis-admin-web/model/business/request"
	"Artemis-admin-web/model/global"
)

type SysUserLogin struct {
	global.DBModel
	Username string `json:"username" gorm:"type:varchar(64);comment:姓名;unique_index"`
	Password string `json:"password" gorm:"type:varchar(32);comment:密码;"`
	Salt     string `json:"salt" gorm:"type:varchar(128);comment:盐;"`
	Locked   bool   `json:"locked" gorm:"comment:账号是否锁定,1:锁定:0未锁定;default:false"`
}

type SysUserInfo struct {
	global.DBModel
	Username    string               `json:"username" gorm:"type:varchar(64);comment:姓名;unique_index"`
	Name        string               `json:"name" gorm:"type:varchar(32);comment:姓名;"`
	Avatar      string               `json:"avatar" gorm:"type:varchar(255);comment:头像;default:https://avatars.githubusercontent.com/u/40788938?v=4;"`
	JobType     string               `json:"jobType" gorm:"type:varchar(32);comment:职位;"`
	Emall       string               `json:"emall" gorm:"type:varchar(32);comment:邮箱;"`
	Company     string               `json:"company" gorm:"type:varchar(32);comment:公司;"`
	CatchPhrase string               `json:"catchPhrase" gorm:"type:varchar(255);comment:个人简介;"`
	Gender      bool                 `json:"gender" gorm:"comment:性别;default:true"`
	RoleId      string               `json:"roleId" gorm:"comment:权限ID;default:1"`
	BizComments []request.BizComment `json:"bizComments" gorm:"foreignKey:UserId;comment:评论条目;"`
}

type SysRole struct {
	global.DBModel
	Name           string          `json:"name" gorm:"type:varchar(128);comment:角色名;"`
	Available      bool            `json:"available" gorm:"comment:角色是否可用 1 Y ,0: N;default:true;"`
	Describe       string          `json:"describe" gorm:"type:varchar(255);comment:描述;"`
	Level          int             `json:"level" gorm:"comment:级别 0-系统管理员 1-组带表;default:1"`
	SysUsers       []SysUserInfo   `json:"sysUsers" gorm:"foreignKey:RoleId;comment:拥有用户;"`
	SysPermissions []SysPermission `json:"sysPermissions"  gorm:"many2many:sys_role_permission;comment:角色&权限;"`
}

type SysPermission struct {
	global.DBModel
	Name       string    `json:"name" gorm:"type:varchar(128);comment:资源名称;"`
	Type       string    `json:"type" gorm:"type:varchar(32);comment:资源类型：menu,link;"`
	Url        string    `json:"url" gorm:"type:varchar(128);comment:访问url地址;"`
	Percode    string    `json:"percode" gorm:"type:varchar(128);comment:权限代码字符串';"`
	Parentid   string    `json:"parentid" gorm:"type:varchar(255);comment:父结点id;"`
	Parentids  string    `json:"parentids" gorm:"type:varchar(128);comment:父结点id列表串;"`
	Sortstring string    `json:"sortstring" gorm:"type:varchar(128);comment:排序号;"`
	Available  bool      `json:"available" gorm:"comment:是否可用 1 Y ,0: N;default:true;"`
	SysRoles   []SysRole `json:"sysRoles" gorm:"many2many:sys_role_permission;comment:角色&权限;"`
}
