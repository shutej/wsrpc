// Package client defines a GopherJS RPC client.
//
// See [Heroku](http://bit.ly/1As5KrE) for background reading about the
// authentication strategy.
package client

import (
	"github.com/gopherjs/websocket"
	"github.com/shutej/flynn/pkg/rpcplus"
	"github.com/shutej/flynn/pkg/rpcplus/jsonrpc"
)

// New creates a new client.
//
// It sends a frame as authentication.  Then, it uses the WebSocket connection
// to speak Flynn JSON-RPC.
func New(url, frame string) (*rpcplus.Client, error) {
	conn, err := websocket.Dial(url)
	if err != nil {
		return nil, err
	}
	if err := conn.Send(frame); err != nil {
		return nil, err
	}
	return rpcplus.NewClientWithCodec(jsonrpc.NewClientCodec(conn)), nil
}
