package note

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/model/do"
	"server/internal/model/entity"
	"server/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/guid"
)

type (
	sNote struct{}
)

func init() {
	service.RegisterNote(New())
}

func New() sNote {
	return sNote{}
}

func (s sNote) CreateNote(ctx context.Context, in model.NoteCreateInput) (err error) {
	if gstr.LenRune(in.Content) == 0 {
		return gerror.New(`content is empty`)
	}
	if in.UserId == 0 {
		return gerror.New(`UserId is invalid`)
	}
	return dao.Note.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.Note.Ctx(ctx).Data(do.Note{
			NoteId:     guid.S(),
			Content:    in.Content,
			UserId:     in.UserId,
			NoteType:   consts.NoteTypeEnumWishNote,
			Status:     consts.NoteStatusEnumActive,
			CreateTime: gtime.Now(),
			UpdateTime: gtime.Now(),
		}).Insert()
		return err
	})
}

func (s sNote) QueryNote(ctx context.Context, query *model.NoteQuery) (list []*entity.Note, err error) {
	db := dao.Note.Ctx(ctx).Where(do.Note{
		Status:   consts.NoteStatusEnumActive,
		IsDelete: false,
	})
	if query != nil {
		if query.UserId != nil {
			db = db.Where(do.Note{UserId: *query.UserId})
		}
		if query.NoteId != nil {
			db = db.Where(do.Note{NoteId: *query.NoteId})
		}
	}
	err = db.OrderDesc(dao.Note.Columns().CreateTime).Scan(&list)
	return
}
