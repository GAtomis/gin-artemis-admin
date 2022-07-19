/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-19 11:15:53
 * @LastEditTime: 2022-07-19 14:26:19
 * @LastEditors: Gavin
 */
package RPC

type RESTful interface {
	QueryItemByGet() (error, any)
	CreateItemByPost() (error, any)
	UpdateItemByPut() (error, any)
	ChangeItembyDelete() (error, any)
}
