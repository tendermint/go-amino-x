package main

import (
	"fmt"

	amino "github.com/tendermint/go-amino-x"
	"github.com/tendermint/go-amino-x/tests"
)

func main() {
	// Paste an example "quoted" string from tests/fuzz/json/crashers/* here.
	// NOTE: You may want to set printLog = true.
	bz := []byte("TODO")
	cdc := amino.NewCodec()
	cst := tests.ComplexSt{}
	err := cdc.UnmarshalJSON(bz, &cst)
	fmt.Printf("Expected a panic but did not. (err: %v)", err)
}
