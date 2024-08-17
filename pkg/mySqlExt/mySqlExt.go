package mySqlExt

import (
	"context"
	"fmt"
	"time"

	"api-example/constants"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type IMySqlExt interface {
	Close() error
	QueryContext(
		ctx context.Context,
		query string,
		args ...interface{},
	) (IMySqlRows, error)
	ExecContext(
		ctx context.Context,
		query string,
		args ...interface{},
	) (bool, error)
	NamedExecContext(
		ctx context.Context,
		query string,
		args interface{},
	) (bool, error)
	GetContext(
		ctx context.Context,
		dest interface{},
		query string,
		args ...interface{},
	) error
	Ping() error
	SelectContext(
		ctx context.Context,
		dest interface{},
		query string,
		args ...interface{},
	) error
	ExecTx(ctx context.Context, fn func(tx IMySqlExt) error) error
	GetSchema() string
}

type Config struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DBName       string
	MaxIdleConns int
	MaxIdleTime  int
	MaxLifeTime  int
	MaxOpenConns int
}

type mySqlExt struct {
	db         *sqlx.DB
	schemaName string
}

type IMySqlRows interface {
	Next() bool
	Close() error
	Scan(dest ...any) error
	Err() error
}

func New(config Config) (IMySqlExt, error) {
	db, err := sqlx.Connect(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			config.Username,
			config.Password,
			config.Host,
			config.Port,
			config.DBName,
		),
	)
	if err != nil {
		return nil, err
	}

	setDBConfig(db, config)

	return &mySqlExt{db, config.DBName}, nil
}

func (m *mySqlExt) Close() error {
	return m.db.Close()
}

func setDBConfig(db *sqlx.DB, config Config) {
	if config.MaxIdleConns == 0 {
		config.MaxIdleConns = 15
	}

	if config.MaxOpenConns == 0 {
		config.MaxOpenConns = 25
	}

	if config.MaxIdleTime == 0 {
		config.MaxIdleTime = 300 // 5 Mins
	}

	if config.MaxLifeTime == 0 {
		config.MaxLifeTime = 300 // 5 Mins
	}

	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetConnMaxIdleTime(time.Duration(config.MaxIdleTime) * time.Second)
	db.SetConnMaxLifetime(time.Duration(config.MaxLifeTime) * time.Second)
}

func (m *mySqlExt) getTableName(ctx context.Context) string {
	if ctx.Value(constants.CtxSQLTableNameKey) != nil {
		return ctx.Value(constants.CtxSQLTableNameKey).(string)
	}
	return ""
}

func (m *mySqlExt) QueryContext(
	ctx context.Context,
	query string,
	args ...interface{},
) (IMySqlRows, error) {

	stmt, err := m.db.PreparexContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryxContext(ctx, args...)
	if err != nil {
		return nil, err
	}

	return rows.Rows, nil
}

func (m *mySqlExt) SelectContext(
	ctx context.Context,
	dest interface{},
	query string,
	args ...interface{},
) error {

	stmt, err := m.db.PreparexContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.SelectContext(ctx, dest, args...)
}
func (m *mySqlExt) ExecContext(
	ctx context.Context,
	query string,
	args ...interface{},
) (bool, error) {

	stmt, err := m.db.PreparexContext(ctx, query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	sqlResults, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return false, err
	}

	affected, err := sqlResults.RowsAffected()
	return affected != 0, err
}

func (m *mySqlExt) NamedExecContext(
	ctx context.Context,
	query string,
	args interface{},
) (bool, error) {

	stmt, err := m.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	sqlResults, err := stmt.ExecContext(ctx, args)
	if err != nil {
		return false, err
	}

	affected, err := sqlResults.RowsAffected()
	return affected != 0, err
}

func (m *mySqlExt) GetContext(
	ctx context.Context,
	dest interface{},
	query string,
	args ...interface{},
) error {

	stmt, err := m.db.PreparexContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.GetContext(ctx, dest, args...)
}
func (m *mySqlExt) ExecTx(ctx context.Context, fn func(tx IMySqlExt) error) error {

	tx, err := m.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	err = fn(m)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

func (m *mySqlExt) Ping() error {
	return m.db.Ping()
}

func (m *mySqlExt) GetSchema() string {
	return m.schemaName
}
