package test

import (
	pinProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/pin/v1"
	"strconv"
)
func MockPinsProto() []*pinProto.Pin {
	var Pins []*pinProto.Pin
	for i := 0; i < 10; i++ {
		Pin := &pinProto.Pin{
			ActivityId: strconv.Itoa(i),
			Code:       strconv.Itoa(i + 100000),
		}
		Pins = append(Pins, Pin)
	}
	return Pins
}