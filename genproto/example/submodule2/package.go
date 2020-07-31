package submodule2

import (
	"github.com/tendermint/go-amino-x"
)

var Package = amino.RegisterPackage(
	amino.NewPackage(
		"github.com/tendermint/go-amino-x/genproto/example/submodule2",
		"submodule2",
		amino.GetCallersDirname(),
	).WithTypes(
		StructSM2{},
	),
)
