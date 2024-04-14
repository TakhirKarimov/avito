package repository

import (
	"avito/pkg/config"
	"avito/pkg/di"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connect struct {
	Conn *gorm.DB
}

func NewConnect() (*Connect, error) {
	cfg := di.Get("config").(*config.Config)
	conn, err := gorm.Open(mysql.Open(cfg.DBConfig.Dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Connect{Conn: conn}, nil
}
