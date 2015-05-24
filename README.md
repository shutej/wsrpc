# wsrpc

A functional implementation of JSON-RPC over WebSockets.  GopherJS clients
should use the `client` package and `net/http` servers should use the `server`
package.

flynn's rpcplus implementation passes
[one](https://github.com/shutej/flynn/blob/master/pkg/rpcplus/server.go#L576)
context object to
[all](https://github.com/shutej/flynn/blob/master/pkg/rpcplus/server.go#L438)
methods called.

This significantly impacts the design of an authentication/session mechanism.

One way to design an authentication/session mechanism would be to try to make
an RPC service that changed the context itself.  However, contexts are designed
to be copied, not mutated, and changing the type of context would be anathema
to the static analysis tools that support context.

The other way is to bake authentication into the protocol directly.
