package v1

import (
	"server/internal/consts"
	"server/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateNoteReq struct {
	g.Meta  `path:"/note/create" method:"post" tags:"NoteService" summary:"Create the note"`
	Content string `v:"required" json:"content"`
}
type CreateNoteRes struct {
}

type QueryNoteReq struct {
	g.Meta   `path:"/note/query" method:"get" tags:"NoteService" summary:"query the note"`
	NoteType consts.NoteTypeEnum `json:"note_type"`
	CommonPaginationReq
}
type QueryNoteRes struct {
	NoteList []*model.NoteVO
}
