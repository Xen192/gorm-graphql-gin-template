package ctxparser

import (
	"context"
	"mygpt/models"
)

type ContextKey string

const (
	CTXUser ContextKey = "user"
)

func GetCTXUser(ctx context.Context) *models.User {
	return ctx.Value(CTXUser).(*models.User)
}
