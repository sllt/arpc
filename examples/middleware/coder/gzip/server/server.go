package main

import (
	"github.com/sllt/arpc"
	"github.com/sllt/arpc/extension/middleware/coder/gzip"
	"github.com/sllt/arpc/log"
)

func main() {
	svr := arpc.NewServer()

	svr.Handler.UseCoder(gzip.New(1024))

	// register router
	svr.Handler.Handle("/echo", func(ctx *arpc.Context) {
		ctx.Write(ctx.Body())
		log.Info("/echo")
	})

	svr.Run("localhost:8888")
}
