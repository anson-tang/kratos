package http

import (
	"context"
	"net"
	"net/http"
	"time"

	pb "github.com/go-kratos/kratos/v2/api/kratos/config/http"
	"github.com/go-kratos/kratos/v2/server"
	transport "github.com/go-kratos/kratos/v2/transport/http"
)

var _ server.Server = (*Server)(nil)

// Option is HTTP server option.
type Option func(o *options)

type options struct {
	transport *transport.Server
	network   string
	address   string
	timeout   time.Duration
}

// Network with server network.
func Network(network string) Option {
	return func(o *options) {
		o.network = network
	}
}

// Address with server address.
func Address(addr string) Option {
	return func(o *options) {
		o.address = addr
	}
}

// Timeout with timeout.
func Timeout(timeout time.Duration) Option {
	return func(o *options) {
		o.timeout = timeout
	}
}

// Transport with server transport.
func Transport(trans *transport.Server) Option {
	return func(o *options) {
		o.transport = trans
	}
}

// Apply apply config.
func Apply(c *pb.ServerConfig) Option {
	return func(o *options) {
		o.network = c.Network
		o.address = c.Address
		if c.Timeout != nil {
			o.timeout = c.Timeout.AsDuration()
		}
	}
}

// Server is a HTTP server wrapper.
type Server struct {
	*http.Server
	opts options
}

// NewServer creates a HTTP server by options.
func NewServer(opts ...Option) *Server {
	options := options{
		network: "tcp",
		address: ":8000",
		timeout: time.Second,
	}
	for _, o := range opts {
		o(&options)
	}
	return &Server{
		opts: options,
		Server: &http.Server{
			Handler:      options.transport,
			ReadTimeout:  options.timeout,
			WriteTimeout: options.timeout,
		},
	}
}

// Start start the HTTP server.
func (s *Server) Start(ctx context.Context) error {
	lis, err := net.Listen(s.opts.network, s.opts.address)
	if err != nil {
		return err
	}
	return s.Serve(lis)
}

// Stop stop the HTTP server.
func (s *Server) Stop(ctx context.Context) error {
	return s.Shutdown(ctx)
}
