// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// InviteUser is the golang structure for table invite_user.
type InviteUser struct {
	Id       int64 `json:"id"       description:""`
	UserId   int64 `json:"userId"   description:"用户id"`
	InviteId int64 `json:"inviteId" description:"邀请id"`
	IsDelete bool  `json:"isDelete" description:""`
}
