// Package client defines a GopherJS RPC client.  This speaks the Flynn dialect
// of JSON-RPC over a WebSocket connection from a web browser.
package client

import (
	"github.com/gopherjs/websocket"
	"github.com/shutej/flynn/pkg/rpcplus"
	"github.com/shutej/flynn/pkg/rpcplus/jsonrpc"
)

type Client struct {
	*rpcplus.Client
}

func New(url string) (Client, error) {
	conn, err := websocket.Dial(url)
	if err != nil {
		return Client{nil}, err
	}
	self := Client{rpcplus.NewClientWithCodec(jsonrpc.NewClientCodec(conn))}
	return self, nil
}
