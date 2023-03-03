package invite

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/model/do"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/guid"
)

type (
	sInvite struct{}
)

//func init() {
//	service.RegisterNote(New())
//}

func New() sInvite {
	return sInvite{}
}

func (s sInvite) CreateInvite(ctx context.Context, in model.InviteCreateInput) (err error) {
	if gstr.LenRune(in.InviteCode) == 0 {
		return gerror.New(`InviteCode is empty`)
	}
	return dao.Invite.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.Invite.Ctx(ctx).Data(do.Invite{
			InviteId:   guid.S(),
			InviteCode: in.InviteCode,
			InviteType: consts.InviteTypeEnumPrivate,
		}).Insert()
		return err
	})
}

func (s sInvite) QueryInvite(ctx context.Context) (list []*entity.Invite, err error) {
	err = dao.Invite.Ctx(ctx).OrderDesc(dao.Invite.Columns().Id).Scan(&list)
	return
}

func (s sInvite) InviteUser(ctx context.Context, in model.InviteUserInput) (err error) {
	if in.InviteId == nil || in.UserId == nil {
		return gerror.New(`InviteCode is empty`)
	}
	return dao.InviteUser.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.InviteUser.Ctx(ctx).Data(do.InviteUser{
			InviteId: *in.InviteId,
			UserId:   *in.UserId,
		}).Insert()
		return err
	})
}
