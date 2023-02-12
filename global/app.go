package global

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"project/config"
)

type Application struct {
	Config config.Configuration
	DB     *sqlx.DB
	Log    *zap.Logger
}

var App = new(Application)
