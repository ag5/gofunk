package dbfunk

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresLogin struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
}

func (login *PostgresLogin) String() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", login.User, login.Password, login.Host, login.Port, login.Database)
}

func ConnectionDo[R any](ctx context.Context, login PostgresLogin, fn func(connPool *pgxpool.Pool) R) R {
	url := login.String()
	connPool, err := pgxpool.Connect(ctx, url)
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}
	defer connPool.Close()
	r := fn(connPool)
	return r
}
