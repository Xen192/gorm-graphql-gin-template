package ctxparser

import (
	"context"
	"mygpt/model"
)

type ContextKey string

const (
	CTXUser ContextKey = "user"
)

func GetCTXUser(ctx context.Context) *model.User {
	return ctx.Value(CTXUser).(*model.User)
}
