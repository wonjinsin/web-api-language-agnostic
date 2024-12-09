package router

import (
	ct "pikachu/controller"
	"pikachu/service"

	"github.com/labstack/echo/v4"
)

// Init ...
func Init(e *echo.Echo, svc *service.Service) {
	api := e.Group("/api")

	makeV1AuthRoute(api, svc)
	makeV1UserRoute(api, svc)
	makeV1InvoiceRoute(api, svc)
}

func makeV1AuthRoute(ver *echo.Group, svc *service.Service) {
	auth := ver.Group("/auths")
	authCt := ct.NewAuthController(svc.Auth)
	auth.POST("/signup", authCt.Signup)
	auth.POST("/signin", authCt.Signin)
}

func makeV1UserRoute(ver *echo.Group, svc *service.Service) {
	user := ver.Group("/users")
	userCt := ct.NewUserController(svc.User)
	user.GET("/:uid", userCt.GetUser)
}

func makeV1InvoiceRoute(ver *echo.Group, svc *service.Service) {
	invoice := ver.Group("/invoices")
	invoiceCt := ct.NewInvoiceController(svc.Invoice)
	invoice.GET("", invoiceCt.GetInvoices)
	invoice.POST("", invoiceCt.NewInvoice)
}
