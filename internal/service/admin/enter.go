package adminsrv

import "github.com/google/wire"

var AdminServiceProvider = wire.NewSet(
	NewEmployeeService,
)
