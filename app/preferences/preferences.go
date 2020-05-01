package preferences

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

var ErrCantParse = errors.New("Can't parse config")

type Preferences struct {
	DatabaseURL                string `envconfig:"DATABASE_URL"`
	DatabaseMaxConnLifetimeSec int    `envconfig:"DATABASE_MAX_CONN_LIFETIME_SEC"`
	DatabaseMaxIdleConns       int    `envconfig:"DATABASE_MAX_IDLE_CONNS"`
	DatabaseMaxOpenConns       int    `envconfig:"DATABASE_MAX_OPEN_CONNS"`
}

func Get() (*Preferences, error) {
	var p Preferences
	if err := envconfig.Process("", &p); err != nil {
		return nil, ErrCantParse
	}
	return &p, nil
}
