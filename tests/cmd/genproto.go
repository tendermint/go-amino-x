package main

import (
	"github.com/tendermint/go-amino-x/genproto"
	"github.com/tendermint/go-amino-x/tests"
)

func main() {
	pkg := tests.Package
	genproto.WriteProto3Schema(pkg)
	genproto.MakeProtoFolder(pkg, "proto")
	genproto.RunProtoc(pkg, "proto")
	genproto.WriteProtoBindings(pkg)
}
