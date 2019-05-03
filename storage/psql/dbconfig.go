package psql

import (
	"time"
)

// Postgres return postgres config
type Postgres struct {
	username          string
	password          string
	port              string
	host              string
	dbName            string
	sslMode           string
	maxOpenConnection int
	maxIdleConnection int
	maxLivetimeCons   time.Duration
}

// MaxOpenCons return func options
func MaxOpenCons(o int) func(*Postgres) {
	return func(cfg *Postgres) {
		cfg.maxOpenConnection = o
	}
}

// MaxIdleCons return func options
func MaxIdleCons(i int) func(*Postgres) {
	return func(cfg *Postgres) {
		cfg.maxIdleConnection = i
	}
}

// MaxLivetimeCons return func options
func MaxLivetimeCons(t time.Duration) func(*Postgres) {
	return func(cfg *Postgres) {
		cfg.maxLivetimeCons = t
	}
}
