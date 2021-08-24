package logger

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
  "log"
)

func logQuery(query string,
	args ...interface{}) {
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

	log.Printf(q, a...)
}
