package main

import (
	"github.com/tendermint/go-amino-x"
	"github.com/tendermint/go-amino-x/genproto/example/submodule"
)

var Package = amino.RegisterPackage(
	amino.NewPackage(
		"main",
		"main",
		amino.GetCallersDirname(),
	).WithP3GoPkgPath(
		"github.com/tendermint/go-amino-x/genproto/example/pb",
	).WithDependencies(
		submodule.Package,
	).WithTypes(
		StructA{},
		StructB{},
	),
)
