package service

import (
	"context"
	"log"
	"os"
	"pikachu/config"
	"pikachu/model/dao"
	"pikachu/model/queryfilter"
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
	User    UserService
	Auth    AuthService
	Invoice InvoiceService
}

// Init ...
func Init(conf *config.ViperConfig, repo *repository.Repository) (*Service, error) {
	userSvc := NewUserService(repo.User)
	authSvc := NewAuthService(conf, repo.User)
	invoiceSvc := NewInvoiceService(repo.Invoice, repo.CompanyReadOnly, repo.InvoiceReadOnly, repo.FeeReadOnly, repo.BankAccountReadOnly)
	return &Service{
		User:    userSvc,
		Auth:    authSvc,
		Invoice: invoiceSvc,
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

// InvoiceService ...
type InvoiceService interface {
	NewInvoice(ctx context.Context, uid string, newInvoice *dao.NewInvoice) (err error)
	GetInvoices(ctx context.Context, uid string, filter *queryfilter.InvoiceQueryFilter) (invoices model.InvoiceAggreates, err error)
}
