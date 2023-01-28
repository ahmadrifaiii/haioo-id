package swagger

import "haioo.id/cart/pkg/swagger/docs"

// @termsOfService http://swagger.io/terms/

func Init() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Telkom Indonesia - Shopping Cart"
	docs.SwaggerInfo.Description = "for Telkom Indonesia"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
