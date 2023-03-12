package model

import (
	"server/internal/consts"
	"server/internal/model/entity"
)

type NoteCreateInput struct {
	Content string
}

type NoteQuery struct {
	UserId   *int64
	NoteId   *int64
	NoteType consts.NoteTypeEnum
}

type NoteVO struct {
	Note *entity.Note
	User *entity.User
}
