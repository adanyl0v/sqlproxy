package sqlx

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/adanyl0v/sqlproxy"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"time"
)

const driverName = "pgx"

type Stmt struct {
	x *sqlx.Stmt
}

func (s *Stmt) Close() error {
	return s.x.Close()
}

func (s *Stmt) Get(dest interface{}, args ...interface{}) error {
	return s.x.Get(dest, args...)
}

func (s *Stmt) GetContext(ctx context.Context, dest interface{}, args ...interface{}) error {
	return s.x.GetContext(ctx, dest, args...)
}

func (s *Stmt) Exec(args ...interface{}) (sql.Result, error) {
	return s.x.Exec(args...)
}

func (s *Stmt) ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error) {
	return s.x.ExecContext(ctx, args...)
}

func (s *Stmt) QueryRow(args ...interface{}) sqlproxy.Row {
	return s.x.QueryRowx(args...)
}

func (s *Stmt) QueryRowContext(ctx context.Context, args ...interface{}) sqlproxy.Row {
	return s.x.QueryRowxContext(ctx, args...)
}

func (s *Stmt) Query(args ...interface{}) (sqlproxy.Rows, error) {
	return s.x.Queryx(args...)
}

func (s *Stmt) QueryContext(ctx context.Context, args ...interface{}) (sqlproxy.Rows, error) {
	return s.x.QueryxContext(ctx, args...)
}

func (s *Stmt) Select(dest interface{}, args ...interface{}) error {
	return s.x.Select(dest, args...)
}

func (s *Stmt) SelectContext(ctx context.Context, dest interface{}, args ...interface{}) error {
	return s.x.SelectContext(ctx, dest, args...)
}

type NamedStmt struct {
	x *sqlx.NamedStmt
}

func (s *NamedStmt) Close() error {
	return s.x.Close()
}

func (s *NamedStmt) Exec(arg interface{}) (sql.Result, error) {
	return s.x.Exec(arg)
}

func (s *NamedStmt) ExecContext(ctx context.Context, arg interface{}) (sql.Result, error) {
	return s.x.ExecContext(ctx, arg)
}

func (s *NamedStmt) Get(dest interface{}, arg interface{}) error {
	return s.x.Get(dest, arg)
}

func (s *NamedStmt) GetContext(ctx context.Context, dest interface{}, arg interface{}) error {
	return s.x.GetContext(ctx, dest, arg)
}

func (s *NamedStmt) MustExec(arg interface{}) sql.Result {
	return s.x.MustExec(arg)
}

func (s *NamedStmt) MustExecContext(ctx context.Context, arg interface{}) sql.Result {
	return s.x.MustExecContext(ctx, arg)
}

func (s *NamedStmt) Query(arg interface{}) (sqlproxy.Rows, error) {
	return s.x.Queryx(arg)
}

func (s *NamedStmt) QueryContext(ctx context.Context, arg interface{}) (sqlproxy.Rows, error) {
	return s.x.QueryxContext(ctx, arg)
}

func (s *NamedStmt) QueryRow(arg interface{}) sqlproxy.Row {
	return s.x.QueryRowx(arg)
}

func (s *NamedStmt) QueryRowContext(ctx context.Context, arg interface{}) sqlproxy.Row {
	return s.x.QueryRowxContext(ctx, arg)
}

func (s *NamedStmt) Select(dest interface{}, arg interface{}) error {
	return s.x.Select(dest, arg)
}

func (s *NamedStmt) SelectContext(ctx context.Context, dest interface{}, arg interface{}) error {
	return s.x.SelectContext(ctx, dest, arg)
}

type Tx struct {
	x *sqlx.Tx
}

func (tx *Tx) Get(dest interface{}, query string, args ...interface{}) error {
	return tx.x.Get(dest, query, args...)
}

func (tx *Tx) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return tx.x.GetContext(ctx, dest, query, args...)
}

func (tx *Tx) Select(dest interface{}, query string, args ...interface{}) error {
	return tx.x.Select(dest, query, args...)
}

func (tx *Tx) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return tx.x.SelectContext(ctx, dest, query, args...)
}

func (tx *Tx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return tx.x.Exec(query, args...)
}

func (tx *Tx) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return tx.x.ExecContext(ctx, query, args...)
}

func (tx *Tx) NamedExec(query string, arg interface{}) (sql.Result, error) {
	return tx.x.NamedExec(query, arg)
}

func (tx *Tx) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return tx.x.NamedExecContext(ctx, query, arg)
}

func (tx *Tx) QueryRow(query string, args ...interface{}) sqlproxy.Row {
	return tx.x.QueryRowx(query, args...)
}

func (tx *Tx) QueryRowContext(ctx context.Context, query string, args ...interface{}) sqlproxy.Row {
	return tx.x.QueryRowxContext(ctx, query, args...)
}

func (tx *Tx) Query(query string, args ...interface{}) (sqlproxy.Rows, error) {
	return tx.x.Queryx(query, args...)
}

func (tx *Tx) QueryContext(ctx context.Context, query string, args ...interface{}) (sqlproxy.Rows, error) {
	return tx.x.QueryxContext(ctx, query, args...)
}

func (tx *Tx) Prepare(query string) (sqlproxy.Stmt, error) {
	stmt, err := tx.x.Preparex(query)
	if err != nil {
		return nil, err
	}

	return &Stmt{x: stmt}, nil
}

func (tx *Tx) PrepareContext(ctx context.Context, query string) (sqlproxy.Stmt, error) {
	stmt, err := tx.x.PreparexContext(ctx, query)
	if err != nil {
		return nil, err
	}

	return &Stmt{x: stmt}, nil
}

func (tx *Tx) PrepareNamed(query string) (sqlproxy.NamedStmt, error) {
	stmt, err := tx.x.PrepareNamed(query)
	if err != nil {
		return nil, err
	}

	return &NamedStmt{x: stmt}, nil
}

func (tx *Tx) PrepareNamedContext(ctx context.Context, query string) (sqlproxy.NamedStmt, error) {
	stmt, err := tx.x.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, err
	}

	return &NamedStmt{x: stmt}, nil
}

func (tx *Tx) Commit() error {
	return tx.x.Commit()
}

func (tx *Tx) Rollback() error {
	return tx.x.Rollback()
}

type DB struct {
	x *sqlx.DB
}

func (db *DB) Close() error {
	return db.x.Close()
}

func (db *DB) Get(dest interface{}, query string, args ...interface{}) error {
	return db.x.Get(dest, query, args...)
}

func (db *DB) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return db.x.GetContext(ctx, dest, query, args...)
}

func (db *DB) Select(dest interface{}, query string, args ...interface{}) error {
	return db.x.Select(dest, query, args...)
}

func (db *DB) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return db.x.SelectContext(ctx, dest, query, args...)
}

func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.x.Exec(query, args...)
}

func (db *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.x.ExecContext(ctx, query, args...)
}

func (db *DB) NamedExec(query string, arg interface{}) (sql.Result, error) {
	return db.x.NamedExec(query, arg)
}

func (db *DB) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return db.x.NamedExecContext(ctx, query, arg)
}

func (db *DB) Query(query string, args ...interface{}) (sqlproxy.Rows, error) {
	return db.x.Queryx(query, args...)
}

func (db *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (sqlproxy.Rows, error) {
	return db.x.QueryxContext(ctx, query, args...)
}

func (db *DB) NamedQuery(query string, arg interface{}) (sqlproxy.Rows, error) {
	return db.x.NamedQuery(query, arg)
}

func (db *DB) NamedQueryContext(ctx context.Context, query string, arg interface{}) (sqlproxy.Rows, error) {
	return db.x.NamedQueryContext(ctx, query, arg)
}

func (db *DB) QueryRow(query string, args ...interface{}) sqlproxy.Row {
	return db.x.QueryRowx(query, args...)
}

func (db *DB) QueryRowContext(ctx context.Context, query string, args ...interface{}) sqlproxy.Row {
	return db.x.QueryRowxContext(ctx, query, args...)
}

func (db *DB) Prepare(query string) (sqlproxy.Stmt, error) {
	stmt, err := db.x.Preparex(query)
	if err != nil {
		return nil, err
	}

	return &Stmt{x: stmt}, nil
}

func (db *DB) PrepareContext(ctx context.Context, query string) (sqlproxy.Stmt, error) {
	stmt, err := db.x.PreparexContext(ctx, query)
	if err != nil {
		return nil, err
	}

	return &Stmt{x: stmt}, nil
}

func (db *DB) PrepareNamed(query string) (sqlproxy.NamedStmt, error) {
	stmt, err := db.x.PrepareNamed(query)
	if err != nil {
		return nil, err
	}

	return &NamedStmt{x: stmt}, nil
}

func (db *DB) PrepareNamedContext(ctx context.Context, query string) (sqlproxy.NamedStmt, error) {
	stmt, err := db.x.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, err
	}

	return &NamedStmt{x: stmt}, nil
}

func (db *DB) DriverName() string {
	return driverName
}

func (db *DB) Begin() (sqlproxy.Tx, error) {
	tx, err := db.x.Beginx()
	if err != nil {
		return nil, err
	}

	return &Tx{x: tx}, nil
}

type Config struct {
	Host         string
	Port         int
	User         string
	Password     string
	Database     string
	SslMode      string
	MaxOpenConns int
	MaxIdleConns int
	ConnLifetime time.Duration
	ConnIdleTime time.Duration
}

func (c *Config) DataSourceName() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.Database, c.SslMode)
}

func Open(dataSourceName string) (*DB, error) {
	db, err := sqlx.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	return &DB{x: db}, nil
}

func OpenWithConfig(c *Config) (*DB, error) {
	db, err := sqlx.Open(driverName, c.DataSourceName())
	if err != nil {
		return nil, err
	}

	return &DB{x: db}, nil
}

func Connect(dataSourceName string) (*DB, error) {
	db, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	return &DB{x: db}, nil
}

func ConnectContext(ctx context.Context, dataSourceName string) (*DB, error) {
	db, err := sqlx.ConnectContext(ctx, driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	return &DB{x: db}, nil
}

func ConnectWithConfig(c *Config) (*DB, error) {
	db, err := sqlx.Connect(driverName, c.DataSourceName())
	if err != nil {
		return nil, err
	}

	setConfig(db, c)
	return &DB{x: db}, nil
}

func ConnectContextWithConfig(ctx context.Context, c *Config) (*DB, error) {
	db, err := sqlx.ConnectContext(ctx, driverName, c.DataSourceName())
	if err != nil {
		return nil, err
	}

	setConfig(db, c)
	return &DB{x: db}, nil
}

func setConfig(db *sqlx.DB, c *Config) {
	if c.MaxOpenConns > 0 {
		db.SetMaxOpenConns(c.MaxOpenConns)
	}
	if c.MaxIdleConns > 0 {
		db.SetMaxIdleConns(c.MaxIdleConns)
	}
	if c.ConnLifetime > 0 {
		db.SetConnMaxLifetime(c.ConnLifetime)
	}
	if c.ConnIdleTime > 0 {
		db.SetConnMaxIdleTime(c.ConnIdleTime)
	}
}
