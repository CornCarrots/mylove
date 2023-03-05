package note

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/model/do"
	"server/internal/model/entity"
	"server/internal/packed"
	"server/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
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
	user := service.Session().GetUser(ctx)
	if user == nil {
		return gerror.New(`user don't login`)
	}
	return dao.Note.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.Note.Ctx(ctx).Data(do.Note{
			NoteId:     packed.IDGenerator.Generate().Int64(),
			Content:    in.Content,
			UserId:     user.UserId,
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
		if query.NoteType > 0 {
			db = db.Where(do.Note{NoteType: query.NoteType})
		}
	}
	err = db.OrderDesc(dao.Note.Columns().CreateTime).Scan(&list)
	return
}
