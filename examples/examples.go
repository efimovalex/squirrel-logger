package examples

import (
	"os"

	sq "github.com/Masterminds/squirrel"
	"github.com/efimovalex/sqlogger"
	"github.com/rs/zerolog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapgrpc"
)

func ExampleLogger() {
	// Output:
	// LogQuery: "SELECT * FROM users WHERE arg1 = ? AND arg2 = ?"
	stmt, args, err := sq.Select("*").
		From("table").
		Where(sq.Eq{"arg1": "test1"}).
		Where(sq.Eq{"arg2": "test2"}).
		OrderBy("arg1").ToSql()

	if err != nil {
		//...
		return
	}

	sqlogger.LogQuery(stmt, args...)
	// Output: "SELECT * FROM users WHERE arg1 = "test1" AND arg2 = "test2"

	return
}

func ExamplePostgresLogger() {
	// Output:
	// LogQuery: "SELECT * FROM users WHERE arg1 = $1 AND arg2 = $2"
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	stmt, args, err := psql.Select("*").
		From("table").
		Where(sq.Eq{"arg1": "test1"}).
		Where(sq.Eq{"arg2": "test2"}).
		OrderBy("arg1").ToSql()

	if err != nil {
		//...
		return
	}

	sqlogger.LogQuery(stmt, args...)
	// Output: "SELECT * FROM users WHERE arg1 = "test1" AND arg2 = "test2"

	return
}

func ExampleZerolog() {
	// Output:
	// LogQuery: "SELECT * FROM users WHERE arg1 = $1 AND arg2 = $2"

	// Set your own logger that implements the Printf method. eg Zerolog logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zl := zerolog.New(os.Stderr).With().Timestamp().Logger()
	sqlogger.SetLogger(&zl)

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	stmt, args, err := psql.Select("*").
		From("table").
		Where(sq.Eq{"arg1": "test1"}).
		Where(sq.Eq{"arg2": "test2"}).
		OrderBy("arg1").ToSql()

	if err != nil {
		//...
		return
	}

	sqlogger.LogQuery(stmt, args...)
	// Output: "SELECT * FROM users WHERE arg1 = "test1" AND arg2 = "test2"

	return
}

func ExampleZapLogger() {
	// Output:
	// LogQuery: "SELECT * FROM users WHERE arg1 = $1 AND arg2 = $2"

	// Set your own logger that implements the Printf method. eg Zerolog logger

	zl := zapgrpc.NewLogger(zap.NewExample())
	sqlogger.SetLogger(zl)

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	stmt, args, err := psql.Select("*").
		From("table").
		Where(sq.Eq{"arg1": "test1"}).
		Where(sq.Eq{"arg2": "test2"}).
		OrderBy("arg1").ToSql()

	if err != nil {
		//...
		return
	}

	sqlogger.LogQuery(stmt, args...)
	// Output: "SELECT * FROM users WHERE arg1 = "test1" AND arg2 = "test2"

	return
}
