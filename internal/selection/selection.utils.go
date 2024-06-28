package selection

import (
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	selectionProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/selection/v1"
)

func ProtoToDto(selection *selectionProto.Selection) *dto.Selection {
	return &dto.Selection{
		Id:      selection.Id,
		GroupId: selection.GroupId,
		BaanId:  selection.BaanId,
		Order:   int(selection.Order),
	}
}

func DtoToProto(selection *dto.Selection) *selectionProto.Selection {
	return &selectionProto.Selection{
		Id:      selection.Id,
		GroupId: selection.GroupId,
		BaanId:  selection.BaanId,
		Order:   int32(selection.Order),
	}
}

func ProtoToDtoList(selections []*selectionProto.Selection) []*dto.Selection {
	var out []*dto.Selection
	for _, selection := range selections {
		out = append(out, ProtoToDto(selection))
	}
	return out
}
