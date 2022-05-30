package dbfunk

import (
	"context"
	"errors"
	"fmt"
	"github.com/ag5/gofunk/funk"
	"github.com/jackc/pgx/v4"
	"os"
	"reflect"
)

type Queryable interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
}

func QueryIntoStruct[S any](ctx context.Context, conn Queryable, sql string, args ...interface{}) ([]*S, error) {
	var coll []*S
	rows, err := conn.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	sampleType := reflect.ValueOf(new(S)).Type().Elem()
	if sampleType.Kind() != reflect.Struct {
		return nil, errors.New("sample is not a struct")
	}

	for rows.Next() {
		obj := reflect.New(sampleType)
		fields := funk.MapRange(0, sampleType.NumField(), func(i int) interface{} {
			return obj.Elem().Field(i).Addr().Interface()
		})
		coll = append(coll, obj.Interface().(*S))
		err = rows.Scan(fields...)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			os.Exit(1)
		}

	}
	return coll, nil
}
