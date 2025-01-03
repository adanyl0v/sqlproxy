package sqlproxy

import (
	"context"
	"database/sql"
)

type Row interface {
	Err() error
	Scan(dest ...interface{}) error
}

type Rows interface {
	Row
	Next() bool
	SliceScan() ([]interface{}, error)
	MapScan(dest map[string]interface{}) error
	StructScan(dest interface{}) error
	Close() error
}

type Stmt interface {
	closer
	Get(dest interface{}, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, args ...interface{}) error
	Exec(args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error)
	QueryRow(args ...interface{}) Row
	QueryRowContext(ctx context.Context, args ...interface{}) Row
	Query(args ...interface{}) (Rows, error)
	QueryContext(ctx context.Context, args ...interface{}) (Rows, error)
	Select(dest interface{}, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, args ...interface{}) error
}

type NamedStmt interface {
	closer
	Exec(arg interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, arg interface{}) (sql.Result, error)
	Get(dest interface{}, arg interface{}) error
	GetContext(ctx context.Context, dest interface{}, arg interface{}) error
	MustExec(arg interface{}) sql.Result
	MustExecContext(ctx context.Context, arg interface{}) sql.Result
	Query(arg interface{}) (Rows, error)
	QueryContext(ctx context.Context, arg interface{}) (Rows, error)
	QueryRow(arg interface{}) Row
	QueryRowContext(ctx context.Context, arg interface{}) Row
	Select(dest interface{}, arg interface{}) error
	SelectContext(ctx context.Context, dest interface{}, arg interface{}) error
}

type Tx interface {
	getter
	selector
	executor
	namedExecutor
	rowQuerier
	querier
	preparer
	namedPreparer
	Commit() error
	Rollback() error
}

type DB interface {
	closer
	getter
	selector
	executor
	namedExecutor
	querier
	namedQuerier
	rowQuerier
	preparer
	namedPreparer
	DriverName() string
	Begin() (Tx, error)
}

type closer interface {
	Close() error
}

type executor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type namedExecutor interface {
	NamedExec(query string, arg interface{}) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}

type rowQuerier interface {
	QueryRow(query string, args ...interface{}) Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) Row
}

type querier interface {
	Query(query string, args ...interface{}) (Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (Rows, error)
}

type namedQuerier interface {
	NamedQuery(query string, arg interface{}) (Rows, error)
	NamedQueryContext(ctx context.Context, query string, arg interface{}) (Rows, error)
}

type getter interface {
	Get(dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type selector interface {
	Select(dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type preparer interface {
	Prepare(query string) (Stmt, error)
	PrepareContext(ctx context.Context, query string) (Stmt, error)
}

type namedPreparer interface {
	PrepareNamed(query string) (NamedStmt, error)
	PrepareNamedContext(ctx context.Context, query string) (NamedStmt, error)
}
