package service

import (
	"context"
	"log"
	"os"
	"pikachu/config"
	"pikachu/repository"
	"pikachu/util"

	"pikachu/model"
)

var zlog *util.Logger

func init() {
	var err error
	zlog, err = util.NewLogger()
	if err != nil {
		log.Fatalf("InitLog module[service] err[%s]", err.Error())
		os.Exit(1)
	}
}

// Service ...
type Service struct {
	User UserService
	Auth AuthService
}

// Init ...
func Init(conf *config.ViperConfig, repo *repository.Repository) (*Service, error) {
	userSvc := NewUserService(repo.User)
	authSvc := NewAuthService(conf, repo.User)
	return &Service{
		User: userSvc,
		Auth: authSvc,
	}, nil
}

// UserService ...
type UserService interface {
	GetUser(ctx context.Context, id string) (ruser *model.User, err error)
}

// AuthService ...
type AuthService interface {
	Signup(ctx context.Context, signup *model.Signup) (token *model.Token, err error)
	Signin(ctx context.Context, signin *model.Signin) (token *model.Token, err error)
}
