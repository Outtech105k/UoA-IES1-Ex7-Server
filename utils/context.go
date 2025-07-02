package utils

import (
	"github.com/Outtech105k/UoA-IES1-Ex7-Server/models"
	"github.com/jmoiron/sqlx"
)

type AppContext struct {
	Config models.AppConfig
	DB     *sqlx.DB
}
