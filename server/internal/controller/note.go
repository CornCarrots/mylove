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
	noteList, err := service.Note().QueryNote(ctx, query)
	if err != nil {
		return nil, err
	}
	uidList := make([]int64, 0)
	for _, note := range noteList {
		uidList = append(uidList, note.UserId)
	}
	userMap, err := service.User().MGetUser(ctx, uidList)
	if err != nil {
		return nil, err
	}
	noteVOList := make([]*model.NoteVO, 0, len(noteList))
	for _, note := range noteList {
		noteVOList = append(noteVOList, &model.NoteVO{
			Note: note,
			User: userMap[note.UserId],
		})
	}
	res.NoteList = noteVOList
	return
}
