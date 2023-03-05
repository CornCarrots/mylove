package controller

import (
	"context"
	v1 "server/api/v1"
	"server/internal/model"
	"server/internal/service"
)

var NoteController = noteController{}

type noteController struct{}

func (c *noteController) CreateNote(ctx context.Context, req *v1.CreateNoteReq) (res *v1.CreateNoteRes, err error) {
	err = service.Note().CreateNote(ctx, model.NoteCreateInput{
		Content: req.Content,
	})
	return
}

func (c *noteController) QueryNote(ctx context.Context, req *v1.QueryNoteReq) (res *v1.QueryNoteRes, err error) {
	query := &model.NoteQuery{
		NoteType: req.NoteType,
	}
	res = &v1.QueryNoteRes{}
	res.NoteList, err = service.Note().QueryNote(ctx, query)
	return
}
