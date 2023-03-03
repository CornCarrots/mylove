package consts

type NoteTypeEnum int

const (
	NoteTypeEnumWishNote = 1 // 礼物墙
	NoteTypeEnumTodoNote = 2
)

type NoteStatusEnum int

const (
	NoteStatusEnumDraft  = 1 // 礼物墙
	NoteStatusEnumActive = 2 // 礼物墙
	NoteStatusEnumFinish = 3
)
