package test

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
