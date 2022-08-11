/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-08-10 13:33:40
 * @LastEditTime: 2022-08-10 16:09:24
 * @LastEditors: Gavin
 */

package base

import (
	"Artemis-admin-web/model/business/request"
	"Artemis-admin-web/model/global"
	"Artemis-admin-web/service/base_core"
	"Artemis-admin-web/utils"

	"github.com/gin-gonic/gin"
)

type COMMENT_API struct {
}

func (c *COMMENT_API) GetCommentsByUid(ctx *gin.Context) {
	var Primarykey global.Primarykey

	_ = ctx.ShouldBindQuery(&Primarykey)

	if err := utils.Verify(Primarykey, utils.Primarykey); err != nil {
		utils.FailM(err.Error(), ctx)
		return
	}

	bc := new(base_core.Comment)
	if bc2, err := bc.GetCommentsByUid(&Primarykey); err != nil {
		utils.FailM(err.Error(), ctx)
		return
	} else {
		utils.OkDM(bc2, "success", ctx)
	}

}
func (c *COMMENT_API) CreateComment(ctx *gin.Context) {
	var body request.BizComment

	_ = ctx.ShouldBindJSON(&body)

	bc := new(base_core.Comment)
	if bc2, err := bc.CreateItem(&body); err != nil {
		utils.FailM(err.Error(), ctx)
		return
	} else {
		utils.OkDM(bc2.ID, "success", ctx)
	}

}
