package model

import "server/internal/consts"

type NoteCreateInput struct {
	Content string
}

type NoteQuery struct {
	UserId   *int64
	NoteId   *int64
	NoteType consts.NoteTypeEnum
}
