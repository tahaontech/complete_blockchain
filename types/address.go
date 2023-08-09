package types

import "fmt"

type Address [20]uint8

func (a Address) IsZero() bool {
	for i := 0; i < 32; i++ {
		if a[i] != 0 {
			return false
		}
	}
	return true
}

func (a Address) ToSlice() []byte {
	b := make([]byte, 20)
	for i := 0; i < 20; i++ {
		b[i] = a[i]
	}
	return b
}

func NewAddressFromBytes(b []byte) Address {
	if len(b) != 20 {
		msg := fmt.Sprintf("given bytes with length %d should be 20\n", len(b))
		panic(msg)
	}

	var value [20]uint8
	for i := 0; i < 20; i++ {
		value[i] = b[i]
	}

	return Address(value)
}
