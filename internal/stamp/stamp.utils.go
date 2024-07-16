package stamp

import (
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	stampProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/stamp/v1"
)

func ProtoToDto(in *stampProto.Stamp) *dto.Stamp {
	return &dto.Stamp{
		ID:     in.Id,
		UserID: in.UserId,
		PointA: in.PointA,
		PointB: in.PointB,
		PointC: in.PointC,
		PointD: in.PointD,
		Stamp:  in.Stamp,
	}
}
