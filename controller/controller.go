package controller

import (
	"log"
	"os"
	"pikachu/util"
	"strconv"

	"github.com/labstack/echo/v4"
)

var zlog *util.Logger

type pickachuStatus struct {
	TRID       string      `json:"trid"`
	ResultCode string      `json:"resultCode"`
	ResultMsg  string      `json:"resultMsg"`
	ResultData interface{} `json:"resultData,omitempty"`
}

func init() {
	var err error
	zlog, err = util.NewLogger()
	if err != nil {
		log.Fatalf("InitLog module[controller] err[%s]", err.Error())
		os.Exit(1)
	}
}

func response(c echo.Context, code int, resultMsg string, result ...interface{}) error {
	strCode := strconv.Itoa(code)
	trid, ok := c.Request().Context().Value(util.TRID).(string)
	if !ok {
		trid = util.GetTRID()
	}

	res := pickachuStatus{
		TRID:       trid,
		ResultCode: strCode,
		ResultMsg:  resultMsg,
	}

	if result != nil {
		res.ResultData = result[0]
	}

	return c.JSON(code, res)
}

// UserController ...
type UserController interface {
	GetUser(c echo.Context) (err error)
}

// AuthController ...
type AuthController interface {
	Signup(c echo.Context) (err error)
	Signin(c echo.Context) (err error)
}

// InvoiceController ...
type InvoiceController interface {
	NewInvoice(c echo.Context) (err error)
	GetInvoices(c echo.Context) (err error)
}
