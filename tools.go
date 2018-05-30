package boomer

import (
	"context"
)

var Key4Orm string = "key4Orm"

type Data4Orm struct {
	UserId int
}

func CtxInitialize(ctx *context.Context) {

}

func CtxSetValues(ctx *context.Context, key string, value interface{}) {
	*ctx = context.WithValue(*ctx, key, value)
}

func CtxGetValues(ctx *context.Context, key string) interface{} {
	return (*ctx).Value(key)
}
