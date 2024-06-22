package utils

import (
	"math/rand"

	"github.com/bxcodec/faker/v4"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	baanProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/baan/v1"
)

func ProtoToDto(in *baanProto.Baan) *dto.Baan {
	return &dto.Baan{
		Id:            in.Id,
		NameTH:        in.NameTH,
		DescriptionTH: in.DescriptionTH,
		NameEN:        in.NameEN,
		DescriptionEN: in.DescriptionEN,
		Size:          dto.BaanSize(in.Size),
		Facebook:      in.Facebook,
		FacebookUrl:   in.FacebookUrl,
		Instagram:     in.Instagram,
		InstagramUrl:  in.InstagramUrl,
		Line:          in.Line,
		LineUrl:       in.LineUrl,
		ImageUrl:      in.ImageUrl,
	}
}

func ProtoToDtoList(in []*baanProto.Baan) []*dto.Baan {
	var out []*dto.Baan
	for _, b := range in {
		out = append(out, ProtoToDto(b))
	}
	return out
}

func MockBaansProto() []*baanProto.Baan {
	var baans []*baanProto.Baan
	for i := 0; i < 10; i++ {
		baan := &baanProto.Baan{
			Id:            faker.UUIDHyphenated(),
			NameTH:        faker.Name(),
			DescriptionTH: faker.Sentence(),
			NameEN:        faker.Name(),
			DescriptionEN: faker.Sentence(),
			Size:          baanProto.BaanSize(rand.Intn(6)),
			Facebook:      faker.URL(),
			FacebookUrl:   faker.URL(),
			Instagram:     faker.URL(),
			InstagramUrl:  faker.URL(),
			Line:          faker.URL(),
			LineUrl:       faker.URL(),
			ImageUrl:      faker.URL(),
		}
		baans = append(baans, baan)
	}
	return baans
}
