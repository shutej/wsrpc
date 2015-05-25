// Package server defines a GopherJS RPC server.
//
// See [Heroku](http://bit.ly/1As5KrE) for background reading about the
// authentication strategy.
package server

import (
	"log"
	"net/http"
	"reflect"

	"github.com/shutej/flynn/pkg/rpcplus"
	"github.com/shutej/flynn/pkg/rpcplus/jsonrpc"
	"golang.org/x/net/context"
	"golang.org/x/net/websocket"
)

var (
	background = context.Background()
	key        = &struct{}{}
)

// Auth gets the result of auth from ctx.
func Auth(ctx context.Context) interface{} {
	return ctx.Value(key)
}

func init() {
	// We just want a bare `context.Context` interface in our code.
	contextType := reflect.TypeOf((*context.Context)(nil)).Elem()
	rpcplus.DefaultServer.SetContextType(contextType)
}

// New creates a new HTTP handler.
//
// It receives a frame as authentication, and converts it using the auth
// function.  Then, it uses the WebSocket connection to speak Flynn JSON-RPC.
func New(auth func(string) interface{}) http.Handler {
	return websocket.Handler(func(conn *websocket.Conn) {
		codec := jsonrpc.NewServerCodec(conn)

		ctx := background

		frame := ""
		if err := websocket.Message.Receive(conn, &frame); err != nil {
			log.Printf("wsrpc/server: %v", err)
		} else {
			ctx = context.WithValue(ctx, key, auth(frame))
		}

		rpcplus.DefaultServer.ServeCodecWithContext(codec, ctx)
	})
}
