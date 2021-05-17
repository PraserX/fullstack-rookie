package main

import (
	"fmt"
	"os"

	"github.com/PraserX/fullstack-rookie/pkg/webserver"
)

func main() {
	var err error
	var server *webserver.Server

	var serverParams []webserver.Option
	serverParams = append(serverParams, webserver.OptionDevMode(true))

	if server, err = webserver.New(serverParams...); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	server.Serve()
}
