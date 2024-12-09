package service

import (
	"context"
	"errors"
	"pikachu/model"
	"pikachu/model/dao"
	"pikachu/model/queryfilter"
	"pikachu/repository"
	"pikachu/util"
)

type invoiceUsecase struct {
	invoiceRepo         repository.InvoiceRepository
	companyReadRepo     repository.CompanyReadOnlyRepository
	invoiceReadRepo     repository.InvoiceReadOnlyRepository
	feeReadRepo         repository.FeeReadOnlyRepository
	bankAccountReadRepo repository.BankAccountReadOnlyRepository
}

// NewInvoiceService ...
func NewInvoiceService(
	invoiceRepo repository.InvoiceRepository,
	companyReadRepo repository.CompanyReadOnlyRepository,
	invoiceReadRepo repository.InvoiceReadOnlyRepository,
	feeReadRepo repository.FeeReadOnlyRepository,
	bankAccountReadRepo repository.BankAccountReadOnlyRepository,
) InvoiceService {
	return &invoiceUsecase{
		invoiceRepo:         invoiceRepo,
		companyReadRepo:     companyReadRepo,
		invoiceReadRepo:     invoiceReadRepo,
		feeReadRepo:         feeReadRepo,
		bankAccountReadRepo: bankAccountReadRepo,
	}
}

// NewInvoice ...
func (u *invoiceUsecase) NewInvoice(ctx context.Context, uid string, newInvoice *dao.NewInvoice) (err error) {
	zlog.With(ctx).Infow("[New Service Request]", "uid", uid, "newInvoice", newInvoice)

	if !newInvoice.Validate() {
		zlog.With(ctx).Errorw("NewInvoice Validate Failed", "newInvoice", newInvoice)
		return errors.New("NewInvoice Validate Failed")
	}

	company, err := u.companyReadRepo.GetCompanyByUserID(ctx, uid)
	if err != nil {
		zlog.With(ctx).Errorw("CompanyRepo GetCompanyByUserID Failed", "uid", uid, "err", err)
		return err
	}
	if !company.SameID(*newInvoice.ApplicantCompanyID) {
		zlog.With(ctx).Errorw("CompanyRepo GetCompanyByUserID Failed", "uid", uid, "err", err)
		return errors.New("CompanyRepo GetCompanyByUserID Failed")
	}

	_, err = u.companyReadRepo.GetCompany(ctx, *newInvoice.RecipientCompanyID)
	if err != nil {
		zlog.With(ctx).Errorw("CompanyRepo GetCompany Failed", "uid", uid, "err", err)
		return err
	}

	bankAccount, err := u.bankAccountReadRepo.GetBankAccountByCompanyID(ctx, *newInvoice.RecipientCompanyID)
	if err != nil {
		zlog.With(ctx).Errorw("BankAccountRepo GetBankAccountByCompanyID Failed", "uid", uid, "err", err)
		return err
	}

	fee, err := u.feeReadRepo.GetFeeByCountryCode(ctx, util.CountryCodeJP)
	if err != nil {
		zlog.With(ctx).Errorw("FeeRepo GetFeeByCountryCode Failed", "uid", uid, "err", err)
		return err
	}

	invoice := model.NewInvoiceAggregate(
		uid,
		*newInvoice.ApplicantCompanyID,
		*newInvoice.RecipientCompanyID,
		*newInvoice.Amount,
		*newInvoice.DueDate,
		fee,
		bankAccount,
	)

	if _, err = u.invoiceRepo.NewInvoice(ctx, invoice); err != nil {
		zlog.With(ctx).Errorw("InvoiceRepo NewInvoice Failed", "uid", uid, "err", err)
		return err
	}

	return nil
}

// GetInvoices ...
func (u *invoiceUsecase) GetInvoices(ctx context.Context, uid string, filter *queryfilter.InvoiceQueryFilter) (invoices model.InvoiceAggreates, err error) {
	zlog.With(ctx).Infow("[New Service Request]", "uid", uid, "filter", filter)

	company, err := u.companyReadRepo.GetCompanyByUserID(ctx, uid)
	if err != nil {
		zlog.With(ctx).Errorw("CompanyRepo GetCompanyByUserID Failed", "uid", uid, "err", err)
		return nil, err
	}

	filter.CompanyID = &company.ID
	if invoices, err = u.invoiceReadRepo.GetInvoices(ctx, filter); err != nil {
		zlog.With(ctx).Errorw("InvoiceRepo GetInvoices Failed", "uid", uid, "err", err)
		return nil, err
	}

	return invoices, nil
}
