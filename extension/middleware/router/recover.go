package router

import (
	"github.com/sllt/arpc"
	"github.com/sllt/arpc/util"
)

// Recover returns the recovery middleware handler.
func Recover() arpc.HandlerFunc {
	return func(ctx *arpc.Context) {
		defer util.Recover()
		ctx.Next()
	}
}
