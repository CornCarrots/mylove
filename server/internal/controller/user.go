package controller

import (
	"context"
	v1 "server/api/v1"
	"server/internal/model"
	"server/internal/service"

	"github.com/gogf/gf/v2/encoding/gjson"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gerror"
)

// UserController is the controller for user.
var UserController = userController{}

type userController struct{}

// SignUp is the API for user sign up.
func (c *userController) SignUp(ctx context.Context, req *v1.UserSignUpReq) (res *v1.UserSignUpRes, err error) {
	err = service.User().Create(ctx, model.UserCreateInput{
		Passport: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
	})
	return
}

// SignIn is the API for user sign in.
func (c *userController) SignIn(ctx context.Context, req *v1.UserSignInReq) (res *v1.UserSignInRes, err error) {
	err = service.User().SignIn(ctx, model.UserSignInInput{
		Passport: req.Passport,
		Password: req.Password,
	})
	g.Log().Infof(ctx, "[SignIn] %s", gjson.MustEncodeString(service.BizCtx().Get(ctx)))
	return
}

// IsSignedIn checks and returns whether the user is signed in.
func (c *userController) IsSignedIn(ctx context.Context, req *v1.UserIsSignedInReq) (res *v1.UserIsSignedInRes, err error) {
	res = &v1.UserIsSignedInRes{
		OK: service.User().IsSignedIn(ctx),
	}
	return
}

// SignOut is the API for user sign out.
func (c *userController) SignOut(ctx context.Context, req *v1.UserSignOutReq) (res *v1.UserSignOutRes, err error) {
	err = service.User().SignOut(ctx)
	return
}

// CheckPassport checks and returns whether the user passport is available.
func (c *userController) CheckPassport(ctx context.Context, req *v1.UserCheckPassportReq) (res *v1.UserCheckPassportRes, err error) {
	available, err := service.User().IsPassportAvailable(ctx, req.Passport)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf(`Passport "%s" is already token by others`, req.Passport)
	}
	return
}

// CheckNickName checks and returns whether the user nickname is available.
func (c *userController) CheckNickName(ctx context.Context, req *v1.UserCheckNickNameReq) (res *v1.UserCheckNickNameRes, err error) {
	available, err := service.User().IsNicknameAvailable(ctx, req.Nickname)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf(`Nickname "%s" is already token by others`, req.Nickname)
	}
	return
}

// Profile returns the user profile.
func (c *userController) Profile(ctx context.Context, req *v1.UserProfileReq) (res *v1.UserProfileRes, err error) {
	res = &v1.UserProfileRes{
		User: service.User().GetProfile(ctx),
	}
	return
}
