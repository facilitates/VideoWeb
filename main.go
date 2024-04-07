// Code generated by hertz generator.

package main

import (
	"flag"
	"github.com/cloudwego/hertz/pkg/app/server"
	"videoweb/conf"
)

func main() {
	conf.Init()
	flag.Parse()
	h := server.Default(server.WithHostPorts(conf.HttpPort), server.WithMaxRequestBodySize(200*1024*1024))
	register(h)
	h.Spin()
}
