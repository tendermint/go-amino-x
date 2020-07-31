package main

import (
	"fmt"

	amino "github.com/tendermint/go-amino-x"
	"github.com/tendermint/go-amino-x/tests"
)

func main() {
	// Paste an example "quoted" string from tests/fuzz/binary/crashers/* here.
	// NOTE: You may want to set printLog = true.
	bz := []byte("\a\x1a\x05\x1a\x01\x80\xf7\x00")
	cdc := amino.NewCodec()
	cst := tests.ComplexSt{}
	err := cdc.UnmarshalBinaryLengthPrefixed(bz, &cst)
	fmt.Printf("Expected a panic but did not. (err: %v)", err)
}
