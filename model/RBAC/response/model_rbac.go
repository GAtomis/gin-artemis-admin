/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-18 16:58:15
 * @LastEditTime: 2022-07-19 15:31:58
 * @LastEditors: Gavin
 */
package response

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
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

func initSQL(fn func(*gorm.DB)) {

	dsn := "root:asdADSDC!!!UJNk@tcp(ec2-3-112-56-234.ap-northeast-1.compute.amazonaws.com:3306)/artemis?charset=utf8mb4&parseTime=True&loc=Local"

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	db.AutoMigrate(&SysUser{}, &SysRole{}, &SysPermission{})
	// i := IDCard{
	// 	Num: 123456,
	// }
	// s := Student{
	// 	StudentName: "zhounan",
	// 	IDCard:      i,
	// }
	// t := Teacher{
	// 	TeacherName: "老师傅",
	// 	Students: []Student{
	// 		s,
	// 	},
	// }
	// c := Class{
	// 	ClassName: "迈瑞中国",
	// 	Students: []Student{
	// 		s,
	// 	},
	// }
	// _ = db.Create(&c).Error
	// _ = db.Create(&t).Error
	fn(db)

}

//模拟中间件2
func middleTow() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before next2")
		c.Next()
		fmt.Println("after next2")
	}
}
func roleByPost(s *SysRole) {
	initSQL(func(db *gorm.DB) {
		db.Create(s)
	})
}
func OnServer() {
	// r := gin.Default()
	initSQL(func(db *gorm.DB) {

	})
	// testModel := r.Group("test").Use(middleTow())
	// testModel.POST("student", func(ctx *gin.Context) {
	// 	var s Student
	// 	err := ctx.ShouldBindJSON(&s)

	// 	if err != nil {
	// 		ctx.JSON(200, gin.H{
	// 			"msg":  "报错了",
	// 			"data": []any{},
	// 		})
	// 	} else {
	// 		studentByPost(&s)
	// 		ctx.JSON(200, gin.H{
	// 			"msg":  "success",
	// 			"data": s,
	// 		})
	// 	}

	// })

	// testModel.GET("student", func(ctx *gin.Context) {

	// 	id := ctx.Query("id")

	// 	if id == "" {
	// 		ctx.JSON(200, gin.H{
	// 			"msg":  "报错了",
	// 			"data": []any{},
	// 		})
	// 	} else {
	// 		s := studentByGet(id)
	// 		ctx.JSON(200, gin.H{
	// 			"msg":  "success",
	// 			"data": s,
	// 		})
	// 	}

	// })
	// testModel.GET("class", func(ctx *gin.Context) {

	// 	id := ctx.Query("id")

	// 	if id == "" {
	// 		ctx.JSON(200, gin.H{
	// 			"msg":  "报错了",
	// 			"data": []any{},
	// 		})
	// 	} else {
	// 		s := classByGet(id)
	// 		ctx.JSON(200, gin.H{
	// 			"msg":  "success",
	// 			"data": s,
	// 		})
	// 	}

	// })

	// testModel.POST("role", func(ctx *gin.Context) {
	// 	var c SysRole
	// 	err := ctx.ShouldBindJSON(&c)

	// 	if err != nil {
	// 		ctx.JSON(200, gin.H{
	// 			"msg":  "报错了",
	// 			"data": []any{},
	// 		})
	// 	} else {
	// 		fmt.Printf("c: %v\n", c)
	// 		roleByPost(&c)
	// 		ctx.JSON(200, gin.H{
	// 			"msg": "success",
	// 		})
	// 	}

	// })

	// testModel.GET("token", func(ctx *gin.Context) {
	// 	s := ctx.GetHeader("Authorization")

	// 	parseToken(s)
	// 	ctx.JSON(200, gin.H{
	// 		"msg":  "报错了",
	// 		"data": s,
	// 	})
	// })

	// r.Run(":9527")

}
