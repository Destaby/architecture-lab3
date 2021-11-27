//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/Destaby/architecture-lab3/server/system"
)

// ComposeApiServer will create an instance of SystemApiServer according to providers defined in this file.
func ComposeApiServer(port HttpPortNumber) (*SystemApiServer, error) {
	wire.Build(
		// DB connection provider (defined in main.go).
		NewDbConnection,
		// Add providers from system package.
		system.Providers,
		// Provide SystemApiServer instantiating the structure and injecting system handler and port number.
		wire.Struct(new(SystemApiServer), "Port", "SystemHandler"),
	)
	return nil, nil
}