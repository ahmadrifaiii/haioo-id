package handler

import (
	"haioo.id/cart/config"
	"haioo.id/cart/config/database"
	"haioo.id/cart/domain/v1/auth"
	"haioo.id/cart/domain/v1/order"
	"haioo.id/cart/domain/v1/product"
	"haioo.id/cart/domain/v1/user"
	authMid "haioo.id/cart/utl/middleware/auth"
)

type Service struct {
	MiddlewareAuth *authMid.Handle
	AuthModule     *auth.Module
	UserModule     *user.Module
	OrderModule    *order.Module
	ProductModule  *product.Module
}

func InitHandler() *Service {

	// mysql init
	MySQLConnection, err := database.MysqlDB()
	if err != nil {
		panic(err)
	}

	config := config.Configuration{
		MysqlDB: MySQLConnection,
	}

	// set service modular
	middlewareAuth := authMid.InitAuthMiddleware(config)
	moduleAuth := auth.InitModule(config)
	moduleUser := user.InitModule(config)
	moduleOrder := order.InitModule(config)
	moduleProduct := product.InitModule(config)

	return &Service{
		AuthModule:     moduleAuth,
		UserModule:     moduleUser,
		MiddlewareAuth: middlewareAuth,
		OrderModule:    moduleOrder,
		ProductModule:  moduleProduct,
	}
}
