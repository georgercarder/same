package same

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/big"
	"reflect"

	gthCommon "github.com/ethereum/go-ethereum/common"
)

func Same(a, b interface{}) (same bool) {
	if reflect.TypeOf(a).Name() != reflect.TypeOf(b).Name() {
		return
	}
	if (IsNil(a) && !IsNil(b)) || (!IsNil(a) && IsNil(b)) {
		return
	}
	switch a.(type) {
	case nil:
		same = true
		return
	case []byte:
		return SameSlice(a.([]byte), b.([]byte))
	case string:
		return SameString(a.(string), b.(string))
	case int:
		return SameInt(a.(int), b.(int))
	case uint64:
		return SameUint64(a.(uint64), b.(uint64))
	case *big.Int:
		return SameBigInt(a.(*big.Int), b.(*big.Int))
	case *ecdsa.PrivateKey:
		return SameECDSAPrivateKey(
			a.(*ecdsa.PrivateKey), b.(*ecdsa.PrivateKey))
	case elliptic.Curve:
		return SameEllipticCurve(
			a.(elliptic.Curve), b.(elliptic.Curve))
	case gthCommon.Address:
		aa := a.(gthCommon.Address)
		bb := b.(gthCommon.Address)
		return SameSlice(aa[:], bb[:])
	}
	return
}

func IsNil(t ...interface{}) (tf bool) {
	for _, tt := range t {
		if tt == nil {
			tf = true
			break
		}
		kind := reflect.TypeOf(tt).Kind()
		// from golang docs:
		// The argument must be a chan, func, interface, map, pointer,
		// or slice value; if it is not, IsNil panics.
		if kind == reflect.Chan ||
			kind == reflect.Func ||
			kind == reflect.Interface ||
			kind == reflect.Map ||
			kind == reflect.Ptr ||
			kind == reflect.Slice {
			if reflect.ValueOf(tt).IsNil() {
				tf = true
				break
			}
		}
	}
	return
}

func IsZero(t ...interface{}) (tf bool) {
	tf = IsNil(t)
	if tf {
		return
	}
	for _, tt := range t {
		switch tt.(type) {
		case uint64:
			if tt.(uint64) == uint64(0) {
				tf = true
				return
			}
			// TODO OTHER TYPES
		}
	}
	return
}

func IsPowerOf2(n int) (tf bool) {
	switch true {
	case n > 2:
		if n%2 == 0 {
			return IsPowerOf2(n / 2)
		}
		return
	case n == 2 || n == 1:
		tf = true
		return
	}
	return
}

func SameSlice(s1, s2 []byte) (same bool) {
	if len(s1) != len(s2) {
		return
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return
		}
	}
	same = true
	return
}

func SameString(a, b string) (same bool) {
	if a == b {
		same = true
	}
	return
}

func SameInt(a, b int) (same bool) {
	if a == b {
		same = true
	}
	return
}

func SameUint64(a, b uint64) (same bool) {
	if a == b {
		same = true
	}
	return
}

func SameBigInt(b1, b2 *big.Int) (same bool) {
	cmp := b1.Cmp(b2)
	if cmp != 0 {
		return
	}
	same = true
	return
}

func SameECDSAPrivateKey(k1, k2 *ecdsa.PrivateKey) (same bool) {
	// check PublicKey
	// elliptic.Curve
	if !Same(k1.PublicKey.Curve, k2.PublicKey.Curve) {
		return
	}
	// X,Y
	if !Same(k1.PublicKey.X, k2.PublicKey.X) {
		return
	}
	if !Same(k1.PublicKey.Y, k2.PublicKey.Y) {
		return
	}
	// check "D"
	if !Same(k1.D, k2.D) {
		return
	}
	same = true
	return
}

func SameEllipticCurve(c1, c2 elliptic.Curve) (same bool) {
	p1 := c1.Params()
	p2 := c2.Params()
	if !Same(p1.Name, p2.Name) {
		return
	}
	// P N B (Gx,Gy) BitSize Name
	if !Same(p1.P, p2.P) {
		return
	}
	if !Same(p1.N, p2.N) {
		return
	}
	if !Same(p1.B, p2.B) {
		return
	}
	if !Same(p1.Gx, p2.Gx) {
		return
	}
	if !Same(p1.Gy, p2.Gy) {
		return
	}
	if !Same(p1.BitSize, p2.BitSize) {
		return
	}
	same = true
	return
}
