package selection

import (
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	selectionProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/selection/v1"
)

func ProtoToDto(selection *selectionProto.Selection) *dto.Selection {
	return &dto.Selection{
		GroupId: selection.GroupId,
		BaanId:  selection.BaanId,
		Order:   int(selection.Order),
	}
}

func DtoToProto(selection *dto.Selection) *selectionProto.Selection {
	return &selectionProto.Selection{
		GroupId: selection.GroupId,
		BaanId:  selection.BaanId,
		Order:   int32(selection.Order),
	}
}

func ProtoToDtoList(selections []*selectionProto.Selection) []*dto.Selection {
	out := make([]*dto.Selection, 0, len(selections))
	for _, selection := range selections {
		out = append(out, ProtoToDto(selection))
	}
	return out
}

func ProtoToDtoBaanCounts(baanCounts []*selectionProto.BaanCount) []*dto.BaanCount {
	out := make([]*dto.BaanCount, 0, len(baanCounts))
	for _, baanCount := range baanCounts {
		out = append(out, &dto.BaanCount{
			BaanId: baanCount.BaanId,
			Count:  int(baanCount.Count),
		})
	}
	return out
}
