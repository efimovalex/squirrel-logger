package sqlogger

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
)

// Logger interface enables you to use your own logger.
type Logger interface {
	Printf(format string, v ...interface{})
}

var logger Logger

func init() {
	logger = &log.Logger{}
}

// Set your own logger that implements the Printf method.
func SetLogger(l Logger) {
	logger = l
}

// LogQuery replaces ? with %v in the query and logs it with the apropriate args.
func LogQuery(query string, args ...interface{}) {
	query = regexp.MustCompile(`\s+`).ReplaceAllString(query, " ")
	q := strings.Replace(query, "?", "%v", -1)
	a := []interface{}{}
	for _, i := range args {
		rv := reflect.ValueOf(i)
		if rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
		}

		if rv.Kind() == reflect.Bool {
			if rv.Bool() {
				a = append(a, 1)
			} else {
				a = append(a, 0)
			}
			continue
		}

		if rv.Kind() == reflect.String {
			a = append(a, fmt.Sprintf(`"%s"`, rv))
			continue
		}

		a = append(a, rv)

	}

	logger.Printf(q, a...)
}

// LogPostgresQuery replaces $1, $2, ... with %v in the query and logs it with the apropriate args.
func LogPostgresQuery(query string,
	args ...interface{}) {
	query = regexp.MustCompile(`\s+`).ReplaceAllString(query, " ")
	q := regexp.MustCompile(`\$\d`).ReplaceAllString(query, "%v")
	a := []interface{}{}
	for _, i := range args {
		rv := reflect.ValueOf(i)
		if rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
		}

		if rv.Kind() == reflect.Bool {
			if rv.Bool() {
				a = append(a, 1)
			} else {
				a = append(a, 0)
			}
			continue
		}

		if rv.Kind() == reflect.String {
			a = append(a, fmt.Sprintf(`"%s"`, rv))
			continue
		}

		a = append(a, rv)

	}

	logger.Printf(q, a...)
}
