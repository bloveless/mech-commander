package commons

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type GameApi interface {
	Move(id, x, y int) bool
}

// UserImplementation is the interface that we're exposing as a plugin.
type UserImplementation interface {
	Tick(GameApi) bool
}

// Here is an implementation that talks over RPC
type UserImplementationRPC struct{ client *rpc.Client }

func (g *UserImplementationRPC) Tick(gameApi GameApi) bool {
	var resp bool
	err := g.client.Call("Plugin.Tick", gameApi, &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}

	return resp
}

// Here is the RPC server that UserImplementationRPC talks to, conforming to
// the requirements of net/rpc
type UserImplementationRPCServer struct {
	// This is the real implementation
	Impl UserImplementation
}

func (s *UserImplementationRPCServer) Tick(gameApi GameApi, resp *bool) error {
	*resp = s.Impl.Tick(gameApi)
	return nil
}

// This is the implementation of plugin.Plugin so we can serve/consume this
//
// This has two methods: Server must return an RPC server for this plugin
// type. We construct a UserImplementationRPCServer for this.
//
// Client must return an implementation of our interface that communicates
// over an RPC client. We return UserImplementationRPC for this.
//
// Ignore MuxBroker. That is used to create more multiplexed streams on our
// plugin connection and is a more advanced use case.
type UserImplementationPlugin struct {
	// Impl Injection
	Impl UserImplementation
}

func (p *UserImplementationPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &UserImplementationRPCServer{Impl: p.Impl}, nil
}

func (UserImplementationPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &UserImplementationRPC{client: c}, nil
}
