package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest"

	"xlsqac/go-zeor-demo/demo/internal/config"
	"xlsqac/go-zeor-demo/demo/internal/handler"
	"xlsqac/go-zeor-demo/demo/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/demo-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
