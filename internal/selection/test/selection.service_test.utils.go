package test

import (
	"github.com/bxcodec/faker/v4"
	selectionProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/selection/v1"
)

func MockSelectionsProto() []*selectionProto.Selection {
	var selections []*selectionProto.Selection
	for i := 0; i < 10; i++ {
		selection := &selectionProto.Selection{
			GroupId: faker.UUIDDigit(),
			BaanId:  faker.UUIDDigit(),
		}
		selections = append(selections, selection)
	}
	return selections
}
