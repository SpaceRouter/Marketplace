package server

import (
	"github.com/spacerouter/marketplace/config"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) error {
	configs := config.GetConfig()
	r := NewRouter(db)
	return r.Run(configs.GetString("server.host") + ":" + configs.GetString("server.port"))
}
