package system

import "github.com/google/wire"

// Set of providers for system components.
var Providers = wire.NewSet(NewStore, HttpHandler)
