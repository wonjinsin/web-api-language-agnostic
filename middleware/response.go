package middleware

import (
	"pikachu/util"
	"strconv"

	"github.com/labstack/echo/v4"
)

type pickachuStatus struct {
	TRID       string      `json:"trid"`
	ResultCode string      `json:"resultCode"`
	ResultMsg  string      `json:"resultMsg"`
	ResultData interface{} `json:"resultData,omitempty"`
}

func response(c echo.Context, code int, resultMsg string) error {
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

	return c.JSON(code, res)
}
