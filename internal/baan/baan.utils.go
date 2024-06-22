package baan

import (
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
