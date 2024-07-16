package test

import (
	//"github.com/bxcodec/faker/v4"
	pinProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/pin/v1"
	"strconv"
)
// Assuming the Pin struct is defined in this package

/*type Pin struct {
	ActivityId string `protobuf:"bytes,1,opt,name=activityId,proto3" json:"activityId,omitempty"`
	Code       string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}
*/
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

/*package test

import (
	//"github.com/bxcodec/faker/v4"
	pinProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/pin/v1"
	"strconv"
)

func MockPinsProto() []*pinProto.Pin {
	var Pins []*pinProto.Pin
	for i := 0; i < 10; i++ {
		Pin := &pinProto.Pin{
			ActivityId: strconv.Itoa(i),
			Code: strconv.Itoa(i + 100000),
		}
		Pins = append(Pins, Pin)
	}
	return Pins
}
*/
// import (
// 	"math/rand"

// 	"github.com/bxcodec/faker/v4"
// 	baanProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/baan/v1"
// )

// func MockBaansProto() []*baanProto.Baan {
// 	var baans []*baanProto.Baan
// 	for i := 0; i < 10; i++ {
// 		baan := &baanProto.Baan{
// 			Id:            faker.UUIDHyphenated(),
// 			NameTH:        faker.Name(),
// 			DescriptionTH: faker.Sentence(),
// 			NameEN:        faker.Name(),
// 			DescriptionEN: faker.Sentence(),
// 			Size:          baanProto.BaanSize(rand.Intn(6)),
// 			Facebook:      faker.URL(),
// 			FacebookUrl:   faker.URL(),
// 			Instagram:     faker.URL(),
// 			InstagramUrl:  faker.URL(),
// 			Line:          faker.URL(),
// 			LineUrl:       faker.URL(),
// 			ImageUrl:      faker.URL(),
// 		}
// 		baans = append(baans, baan)
// 	}
// 	return baans
// }
