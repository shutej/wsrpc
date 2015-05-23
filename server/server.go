// Package server defines a GopherJS RPC server.  This speaks the Flynn dialect
// of JSON-RPC over a WebSocket connection to the client.
package server

import (
	"log"
	"net/http"

	"github.com/shutej/flynn/pkg/rpcplus"
	"github.com/shutej/flynn/pkg/rpcplus/jsonrpc"
	"golang.org/x/net/websocket"
)

func defaultFactory(*websocket.Conn) interface{} {
	return nil
}

type server struct {
	server  *rpcplus.Server
	factory Factory
}

type Option func(server *server)
type Factory func(*websocket.Conn) interface{}

func ContextFactory(factory Factory) Option {
	return func(self *server) {
		self.factory = factory
	}
}

func Register(rcvr interface{}) Option {
	return func(self *server) {
		if err := self.server.Register(rcvr); err != nil {
			log.Panic(err)
		}
	}
}

func RegisterName(name string, rcvr interface{}) Option {
	return func(self *server) {
		if err := self.server.RegisterName(name, rcvr); err != nil {
			log.Panic(err)
		}
	}
}

func Handler(options ...Option) http.Handler {
	self := &server{
		factory: defaultFactory,
		server:  rpcplus.NewServer(),
	}

	for _, option := range options {
		option(self)
	}

	return websocket.Handler(func(conn *websocket.Conn) {
		codec := jsonrpc.NewServerCodec(conn)
		context := self.factory(conn)
		self.server.ServeCodecWithContext(codec, context)
	})
}
