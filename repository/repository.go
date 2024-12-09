package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"pikachu/config"
	"pikachu/model"
	"pikachu/model/queryfilter"
	"pikachu/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var zlog *util.Logger

type dbLogger struct {
	*util.Logger
}

func (dl *dbLogger) LogMode(l logger.LogLevel) logger.Interface {
	return dl
}

func (dl *dbLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	dl.Logger.With(ctx).Info(msg, data)
}

func (dl *dbLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	dl.Logger.With(ctx).Warn(msg, data)
}

func (dl *dbLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	dl.Logger.With(ctx).Error(msg, data)
}

func (dl *dbLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil {
		dl.Logger.With(ctx).Infow(err.Error(), "elapsed", fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6), "rows", rows, "sql", sql)
	} else {
		dl.Logger.With(ctx).Infow("", "elapsed", fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6), "rows", rows, "sql", sql)
	}
}

func init() {
	var err error
	zlog, err = util.NewLogger()
	if err != nil {
		log.Fatalf("InitLog module[service] err[%s]", err.Error())
		os.Exit(1)
	}
}

// Repository ...
type Repository struct {
	User                UserRepository
	Invoice             InvoiceRepository
	CompanyReadOnly     CompanyReadOnlyRepository
	InvoiceReadOnly     InvoiceReadOnlyRepository
	FeeReadOnly         FeeReadOnlyRepository
	BankAccountReadOnly BankAccountReadOnlyRepository
}

// Init ...
func Init(pikachu *config.ViperConfig) (*Repository, error) {
	mysqlConn, err := mysqlConnect(pikachu, "database")
	if err != nil {
		return nil, err
	}

	mysqlReadOnlyConn, err := mysqlConnect(pikachu, "readOnlyDatabase")
	if err != nil {
		return nil, err
	}

	db := &model.DB{
		MainDB: mysqlConn,
		ReadDB: mysqlReadOnlyConn,
	}

	userRepo := NewGormUserRepository(db.MainDB)
	invoiceRepo := NewGormInvoiceRepository(db.MainDB)
	companyReadRepo := NewGormCompanyReadOnlyRepository(db.ReadDB)
	invoiceReadRepo := NewGormInvoiceReadOnlyRepository(db.ReadDB)
	feeReadRepo := NewGormFeeReadOnlyRepository(db.ReadDB)
	bankAccountReadRepo := NewGormBankAccountReadOnlyRepository(db.ReadDB)

	return &Repository{
		User:                userRepo,
		Invoice:             invoiceRepo,
		CompanyReadOnly:     companyReadRepo,
		InvoiceReadOnly:     invoiceReadRepo,
		FeeReadOnly:         feeReadRepo,
		BankAccountReadOnly: bankAccountReadRepo,
	}, nil
}

func mysqlConnect(pikachu *config.ViperConfig, prefix string) (mysql *gorm.DB, err error) {
	return gorm.Open(getDialector(pikachu, prefix), getConfig())
}

func getDialector(pikachu *config.ViperConfig, prefix string) gorm.Dialector {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?&charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=UTC",
		pikachu.GetString(fmt.Sprintf("%s.username", prefix)),
		pikachu.GetString(fmt.Sprintf("%s.password", prefix)),
		pikachu.GetString(fmt.Sprintf("%s.host", prefix)),
		pikachu.GetInt(fmt.Sprintf("%s.port", prefix)),
		pikachu.GetString(fmt.Sprintf("%s.dbname", prefix)),
	)

	return mysql.Open(dbURI)
}

func getConfig() (gConfig *gorm.Config) {
	dbLogger := &dbLogger{zlog}
	gConfig = &gorm.Config{
		Logger:      dbLogger,
		PrepareStmt: true,
	}

	return gConfig
}

// UserRepository ...
type UserRepository interface {
	NewUser(ctx context.Context, user *model.User) (ruser *model.User, err error)
	GetUser(ctx context.Context, uid string) (ruser *model.User, err error)
	GetUserByEmail(ctx context.Context, email string) (ruser *model.User, err error)
}

// CompanyReadOnlyRepository ...
type CompanyReadOnlyRepository interface {
	GetCompanyByUserID(ctx context.Context, uid string) (rcompany *model.Company, err error)
	GetCompany(ctx context.Context, id uint64) (rcompany *model.Company, err error)
}

// InvoiceRepository ...
type InvoiceRepository interface {
	NewInvoice(ctx context.Context, invoice *model.InvoiceAggregate) (rinvoice *model.InvoiceAggregate, err error)
}

// InvoiceReadOnlyRepository ...
type InvoiceReadOnlyRepository interface {
	GetInvoices(ctx context.Context, filter queryfilter.QueryFilter) (invoices model.InvoiceAggreates, err error)
}

// FeeReadOnlyRepository ...
type FeeReadOnlyRepository interface {
	GetFeeByCountryCode(ctx context.Context, CountryCode util.CountryCode) (rfee *model.Fee, err error)
}

// BankAccountReadOnlyRepository ...
type BankAccountReadOnlyRepository interface {
	GetBankAccountByCompanyID(ctx context.Context, companyID uint64) (rbankAccount *model.BankAccount, err error)
}

// TxRepository ...
type TxRepository interface {
	Begin(ctx context.Context, fn func(ctx context.Context) error) error
}
