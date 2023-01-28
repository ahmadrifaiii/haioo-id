package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"haioo.id/cart/config/env"
	"haioo.id/cart/utl/middleware/logging"
	"haioo.id/cart/utl/middleware/secure"

	"io/ioutil"
	"net/http"

	"os"
	"os/signal"
	"time"

	"haioo.id/cart/pkg/swagger"
	_ "haioo.id/cart/pkg/swagger/docs"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	MysqlMigrate "haioo.id/cart/utl/database/migrate/migration_mysql"
)

const DefaultPort = 8080

// HTTPServerMain main function for serving services over http

func (s *Service) HTTPServerMain() *echo.Echo {
	// to active swagger
	swagger.Init()

	e := echo.New()

	// e.Use(middleware.ServerHeader, middleware.Logger)

	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORS())
	e.Use(secure.Headers())
	e.Use(logging.Logging())

	// administrator group
	adm := e.Group("/api/v1")

	// auth module
	ModuleAuth := adm.Group("/auth")
	s.AuthModule.HandleRest(ModuleAuth)

	// module identity access management
	ModuleUser := adm.Group("/user", s.MiddlewareAuth.BearerVerify())
	s.UserModule.HandleRest(ModuleUser)

	// domain module order
	ModuleOrder := adm.Group("/order", s.MiddlewareAuth.BearerVerify())
	s.OrderModule.HandleRest(ModuleOrder)

	// domain module product
	ProductModule := adm.Group("/product", s.MiddlewareAuth.BearerVerify())
	s.ProductModule.HandleRest(ProductModule)

	e.GET("/docs/*", echoSwagger.WrapHandler)
	e.GET("/migration", MysqlMigrate.Migration)

	data, _ := json.MarshalIndent(e.Routes(), "", "  ")

	ioutil.WriteFile("routes.json", data, 0644)

	return e
}

func (s *Service) StartServer() {
	server := s.HTTPServerMain()
	listenerPort := fmt.Sprintf(":%v", env.Conf.HttpPort)
	if err := server.StartServer(&http.Server{
		Addr:         listenerPort,
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
	}); err != nil {
		server.Logger.Fatal(err.Error())
	}
}

func (s *Service) ShutdownServer() {
	server := s.HTTPServerMain()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err.Error())
	}
}
