package controller

import (
	"context"
	"net/http"
	"pikachu/model"
	"pikachu/service"
	"pikachu/util"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// User ...
type User struct {
	userSvc service.UserService
}

// NewUserController ...
func NewUserController(userSvc service.UserService) UserController {
	return &User{
		userSvc: userSvc,
	}
}

// GetUser ...
func (u *User) GetUser(c echo.Context) (err error) {
	ctx := c.Request().Context()
	uid := c.Param("uid")
	zlog.With(ctx).Infow("[New request]", "uid", uid)
	intCtx, cancel := context.WithTimeout(ctx, util.CtxTimeOut)
	defer cancel()

	if _, err = uuid.Parse(uid); err != nil {
		zlog.With(intCtx).Warnw("ID is not valid", "uid", uid, "err", err)
		return response(c, http.StatusBadRequest, "User is not valid")
	}

	var user *model.User
	if user, err = u.userSvc.GetUser(intCtx, uid); err != nil {
		zlog.With(intCtx).Warnw("UserSvc GetUser failed", "uid", uid, "err", err)
		return response(c, http.StatusInternalServerError, err.Error())
	}

	return response(c, http.StatusOK, "GetUser OK", user)
}
