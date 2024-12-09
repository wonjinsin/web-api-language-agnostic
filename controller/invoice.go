package controller

import (
	"context"
	"net/http"
	"pikachu/model"
	"pikachu/model/dao"
	"pikachu/model/queryfilter"
	"pikachu/service"
	"pikachu/util"

	"github.com/labstack/echo/v4"
)

// Invoice ...
type Invoice struct {
	invoiceSvc service.InvoiceService
}

// NewInvoiceController ...
func NewInvoiceController(invoiceSvc service.InvoiceService) InvoiceController {
	return &Invoice{
		invoiceSvc: invoiceSvc,
	}
}

// NewInvoice ...
func (I *Invoice) NewInvoice(c echo.Context) (err error) {
	ctx := c.Request().Context()
	uid, ok := ctx.Value(util.UUID).(string)
	if !ok {
		zlog.With(ctx).Warnw("uuid is not exist")
		return response(c, http.StatusBadRequest, "uuid is not exist")
	}

	newInvoice := &dao.NewInvoice{}
	if err := c.Bind(newInvoice); err != nil {
		zlog.With(ctx).Warnw("echo Bind error", "err", err)
		return response(c, http.StatusBadRequest, err.Error())
	}

	if err = I.invoiceSvc.NewInvoice(ctx, uid, newInvoice); err != nil {
		zlog.With(ctx).Warnw("InvoiceSvc NewInvoice failed", "uid", uid, "err", err)
		return response(c, http.StatusInternalServerError, err.Error())
	}

	return response(c, http.StatusOK, "NewInvoice OK")
}

// GetInvoices ...
func (I *Invoice) GetInvoices(c echo.Context) (err error) {
	ctx := c.Request().Context()
	uid, ok := ctx.Value(util.UUID).(string)
	if !ok {
		zlog.With(ctx).Warnw("uuid is not exist")
		return response(c, http.StatusBadRequest, "uuid is not exist")
	}

	zlog.With(ctx).Infow("[New request]", "uid", uid)
	intCtx, cancel := context.WithTimeout(ctx, util.CtxTimeOut)
	defer cancel()

	invoiceQueryFilter := &queryfilter.InvoiceQueryFilter{}
	if err := c.Bind(invoiceQueryFilter); err != nil {
		zlog.With(ctx).Warnw("echo Bind error", "err", err)
		return response(c, http.StatusBadRequest, err.Error())
	}

	var invoices model.InvoiceAggreates
	if invoices, err = I.invoiceSvc.GetInvoices(intCtx, uid, invoiceQueryFilter); err != nil {
		zlog.With(intCtx).Warnw("InvoiceSvc GetInvoices failed", "uid", uid, "err", err)
		return response(c, http.StatusInternalServerError, err.Error())
	}

	return response(c, http.StatusOK, "GetInvoices OK", invoices)
}
