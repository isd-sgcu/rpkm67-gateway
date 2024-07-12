package pin

import (
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	pinProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/pin/v1"
)

func ProtoToDto(in *pinProto.Pin) *dto.Pin {
	return &dto.Pin{
		ActivityId: in.ActivityId,
		Code:       in.Code,
	}
}

func ProtoToDtoList(in []*pinProto.Pin) []*dto.Pin {
	var out []*dto.Pin
	for _, b := range in {
		out = append(out, ProtoToDto(b))
	}
	return out
}
