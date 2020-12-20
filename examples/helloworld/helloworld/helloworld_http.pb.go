// Code generated by protoc-gen-go-http. DO NOT EDIT.

package helloworld

import (
	context "context"
	http1 "github.com/go-kratos/kratos/v2/transport/http"
	http "net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
// context./http.
const _ = http1.SupportPackageIsVersion1

type GreeterHTTPServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterGreeterHTTPServer(s http1.ServiceRegistrar, srv GreeterHTTPServer) {
	s.RegisterService(&_HTTP_Greeter_serviceDesc, srv)
}

func _HTTP_Greeter_SayHello_0(srv interface{}, ctx context.Context, dec func(interface{}) error, req *http.Request) (interface{}, error) {
	var in HelloRequest

	if err := http1.PopulateVars(&in, req, []string{
		"name",
	}, nil); err != nil {
		return nil, err
	}

	if err := http1.PopulateForm(&in, req, nil, []string{
		"name",
	}); err != nil {
		return nil, err
	}

	out, err := srv.(GreeterServer).SayHello(ctx, &in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _HTTP_Greeter_serviceDesc = http1.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterHTTPServer)(nil),
	Methods: []http1.MethodDesc{

		{
			Path:    "/helloworld/{name}",
			Method:  "GET",
			Handler: _HTTP_Greeter_SayHello_0,
		},
	},
	Metadata: "helloworld.proto",
}