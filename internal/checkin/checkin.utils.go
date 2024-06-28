package checkin

import (
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	checkinProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/checkin/checkin/v1"
)

func ProtoToDto(in *checkinProto.CheckIn) *dto.CheckIn {
	return &dto.CheckIn{
		ID:     in.Id,
		UserID: in.UserId,
		Email:  in.Email,
		Event:  in.Event,
	}
}

func ProtoToDtos(in []*checkinProto.CheckIn) []*dto.CheckIn {
	var out []*dto.CheckIn
	for _, v := range in {
		out = append(out, ProtoToDto(v))
	}
	return out
}
