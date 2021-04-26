package main

import (
	"github.com/PraserX/fullstack-rookie/pkg/webserver"
)

func main() {
	var serverParams []webserver.Option
	serverParams = append(serverParams, webserver.OptionDevMode(true))

	server := webserver.New(serverParams...)
	server.Serve()
}
