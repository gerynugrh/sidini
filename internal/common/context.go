package common

import (
	"context"
	"github.com/pocketbase/pocketbase/daos"
)

type contextKey string

var dao = contextKey("daos")

func Dao(ctx context.Context) *daos.Dao {
	return ctx.Value(dao).(*daos.Dao)
}
