package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/fn-code/hexagonal-arcitec/storage"
	_ "github.com/lib/pq"
)

// postgresConn return database connection
type postgresConn struct {
	conn   *sql.DB
	ctx    context.Context
	cancel func()
}

// NewPostgres is set database config
func NewPostgres(ctx context.Context, username, pass, port, dbname string, opt ...func(*Postgres)) (storage.Database, error) {
	cfg := &Postgres{
		username: username,
		password: pass,
		port:     port,
		dbName:   dbname,
		sslMode:  "disable",
		host:     "localhost",
	}

	for _, f := range opt {
		f(cfg)
	}

	db, err := cfg.run(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// run is running database connection
// and return Database interface and error
func (db *Postgres) run(ctx context.Context) (*postgresConn, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?port=%v&sslmode=%s&connect_timeout=90", db.username, db.password, db.host, db.dbName, db.port, db.sslMode)
	sqlConn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	sqlConn.SetConnMaxLifetime(db.maxLivetimeCons)
	sqlConn.SetMaxIdleConns(db.maxIdleConnection)
	sqlConn.SetMaxOpenConns(db.maxIdleConnection)

	ctx, cancel := context.WithCancel(ctx)

	return &postgresConn{sqlConn, ctx, cancel}, nil
}

// Close is to close database connection
// and returning an error
func (db *postgresConn) Close() error {
	db.cancel()
	err := db.closedb()
	if err != nil {
		return err
	}
	return nil
}

func (db *postgresConn) closedb() error {
	<-db.ctx.Done()
	err := db.conn.Close()
	if err != nil {
		return err
	}
	return nil
}
