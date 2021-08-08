// source: api.proto
package api

import (
	"context"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/go-kratos/kratos/pkg/net/http/blademaster/binding"
)
import google_protobuf1 "google.golang.org/protobuf/types/known/emptypb"

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathDemoPing = "/demo.service.v1.Demo/Ping"
var PathDemoSayHello = "/demo.service.v1.Demo/SayHello"
var PathDemoSayHelloURL = "/kratos-demo/say_hello"

// DemoBMServer is the server API for Demo service.
type DemoBMServer interface {
	Ping(ctx context.Context, req *google_protobuf1.Empty) (resp *google_protobuf1.Empty, err error)

	SayHello(ctx context.Context, req *HelloReq) (resp *google_protobuf1.Empty, err error)

	SayHelloURL(ctx context.Context, req *HelloReq) (resp *HelloResp, err error)
}

var DemoSvc DemoBMServer

func demoPing(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := DemoSvc.Ping(c, p)
	c.JSON(resp, err)
}

func demoSayHello(c *bm.Context) {
	p := new(HelloReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := DemoSvc.SayHello(c, p)
	c.JSON(resp, err)
}

func demoSayHelloURL(c *bm.Context) {
	p := new(HelloReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := DemoSvc.SayHelloURL(c, p)
	c.JSON(resp, err)
}

// RegisterDemoBMServer Register the blademaster route
func RegisterDemoBMServer(e *bm.Engine, server DemoBMServer) {
	DemoSvc = server
	e.GET("/demo.service.v1.Demo/Ping", demoPing)
	e.GET("/demo.service.v1.Demo/SayHello", demoSayHello)
	e.GET("/kratos-demo/say_hello", demoSayHelloURL)
}
