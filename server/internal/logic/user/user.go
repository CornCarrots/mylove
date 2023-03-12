package user

import (
	"context"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/model/do"
	"server/internal/model/entity"
	"server/internal/packed"
	"server/internal/service"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

// Create creates user account.
func (s *sUser) Create(ctx context.Context, in model.UserCreateInput) (err error) {
	// If Nickname is not specified, it then uses Passport as its default Nickname.
	if in.Nickname == "" {
		in.Nickname = in.Passport
	}
	var (
		available bool
	)
	// Nickname checks.
	available, err = s.IsNicknameAvailable(ctx, in.Nickname)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`Nickname "%s" is already token by others`, in.Nickname)
	}
	// Passport checks.
	available, err = s.IsPassportAvailable(ctx, in.Passport)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`Passport "%s" is already token by others`, in.Passport)
	}
	encryptSt, err := gmd5.EncryptString(in.Password)
	if err != nil {
		return err
	}
	return dao.User.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.User.Ctx(ctx).Data(do.User{
			UserId:     packed.IDGenerator.Generate().Int64(),
			Passport:   in.Passport,
			Password:   encryptSt,
			Nickname:   in.Nickname,
			CreateTime: gtime.Now(),
			UpdateTime: gtime.Now(),
		}).Insert()
		return err
	})
}

// IsSignedIn checks and returns whether current user is already signed-in.
func (s *sUser) IsSignedIn(ctx context.Context) bool {
	if v := service.BizCtx().Get(ctx); v != nil && v.User != nil {
		return true
	}
	return false
}

// SignIn creates session for given user account.
func (s *sUser) SignIn(ctx context.Context, in model.UserSignInInput) (err error) {
	var user *entity.User
	encryptSt, err := gmd5.EncryptString(in.Password)
	if err != nil {
		return err
	}
	err = dao.User.Ctx(ctx).Where(do.User{
		Passport: in.Passport,
		Password: encryptSt,
	}).Scan(&user)
	if err != nil {
		return err
	}
	if user == nil {
		return gerror.New(`Passport or Password not correct`)
	}
	if err = service.Session().SetUser(ctx, user); err != nil {
		return err
	}
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		UserId:   user.UserId,
		Passport: user.Passport,
		Nickname: user.Nickname,
	})
	return nil
}

// SignOut removes the session for current signed-in user.
func (s *sUser) SignOut(ctx context.Context) error {
	return service.Session().RemoveUser(ctx)
}

// IsPassportAvailable checks and returns given passport is available for signing up.
func (s *sUser) IsPassportAvailable(ctx context.Context, passport string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{
		Passport: passport,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// IsNicknameAvailable checks and returns given nickname is available for signing up.
func (s *sUser) IsNicknameAvailable(ctx context.Context, nickname string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{
		Nickname: nickname,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// GetProfile retrieves and returns current user info in session.
func (s *sUser) GetProfile(ctx context.Context) *entity.User {
	return service.Session().GetUser(ctx)
}

func (s *sUser) MGetUser(ctx context.Context, uidList []int64) (res map[int64]*entity.User, err error) {
	var list []*entity.User
	err = dao.User.Ctx(ctx).WhereIn("user_id", uidList).Scan(&list)
	if err != nil {
		return nil, err
	}
	res = make(map[int64]*entity.User, len(list))
	for _, user := range list {
		res[user.UserId] = user
	}
	return res, nil
}
