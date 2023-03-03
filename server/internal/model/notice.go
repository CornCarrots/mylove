package model

type NoteCreateInput struct {
	UserId  int64
	Content string
}

type NoteQuery struct {
	UserId *int64
	NoteId *int64
}
