package auth

import (
	"haioo.id/cart/config"
)

type Module struct {
	Config config.Configuration
}

func InitModule(conf config.Configuration) *Module {
	return &Module{
		Config: conf,
	}
}
