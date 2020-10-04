package same

import (
	"math/big"
)

func CompareHex(aHex, bHex string) (cmp int) {
	// -1,0,1
	// -1 : aHex < bHex
	//  0 : aHex == bHex
	//  1 : aHex > bHex
	aHex = emptyHexGuard(aHex)
	bHex = emptyHexGuard(bHex)
	a := new(big.Int)
	a.SetString(aHex[2:], 16)
	b := new(big.Int)
	b.SetString(bHex[2:], 16)
	cmp = a.Cmp(b)
	return
}

func emptyHexGuard(maybeHex string) (hex string) {
	if maybeHex == "" {
		hex = "0x0"
		return
	}
	hex = maybeHex
	return
}

func CompareHashAsSlice(a, b []byte) (cmp int) {
	// -1,0,1
	// -1 : a < b
	//  0 : a == b
	//  1 : a > b
	aI := new(big.Int)
	aI.SetBytes(a)
	bI := new(big.Int)
	bI.SetBytes(b)
	cmp = aI.Cmp(bI)
	return
}
