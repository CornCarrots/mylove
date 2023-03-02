// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Note is the golang structure for table note.
type Note struct {
	Id         int64       `json:"id"         description:""`
	NoteId     int64       `json:"noteId"     description:"文章id"`
	UserId     int64       `json:"userId"     description:"用户ID"`
	Content    string      `json:"content"    description:"内容"`
	Status     int         `json:"status"     description:"状态 0-生效中 1-已完成"`
	NoteType   int         `json:"noteType"   description:"心愿类型 1-想要的礼物 2-想做的事情"`
	IsDelete   bool        `json:"isDelete"   description:""`
	CreateTime *gtime.Time `json:"createTime" description:"创建时间"`
	UpdateTime *gtime.Time `json:"updateTime" description:"更新时间"`
}
