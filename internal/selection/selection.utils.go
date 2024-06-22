package selection

import (
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	selectionProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/selection/v1"
)

func ProtoToDto(selection *selectionProto.Selection) *dto.Selection {
	return &dto.Selection{
		Id:      selection.Id,
		GroupId: selection.GroupId,
		BaanIds: selection.BaanIds,
	}
}

func DtoToProto(selection *dto.Selection) *selectionProto.Selection {
	return &selectionProto.Selection{
		Id:      selection.Id,
		GroupId: selection.GroupId,
		BaanIds: selection.BaanIds,
	}
}
