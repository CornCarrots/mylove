package model

type InviteCreateInput struct {
	InviteCode string
}

type InviteUserInput struct {
	UserId   *int64
	InviteId *int64
}
