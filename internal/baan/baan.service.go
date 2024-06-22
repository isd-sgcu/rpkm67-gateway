package baan

import (
	"context"
	"time"

	"github.com/isd-sgcu/rpkm67-gateway/apperrors"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/utils"
	baanProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/baan/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service interface {
	FindAllBaan(req *dto.FindAllBaanRequest) (*dto.FindAllBaanResponse, *apperrors.AppError)
	FindOneBaan(req *dto.FindOneBaanRequest) (*dto.FindOneBaanResponse, *apperrors.AppError)
}

type serviceImpl struct {
	client baanProto.BaanServiceClient
	log    *zap.Logger
}

func NewService(client baanProto.BaanServiceClient, log *zap.Logger) Service {
	return &serviceImpl{
		client: client,
		log:    log,
	}
}

func (s *serviceImpl) FindAllBaan(req *dto.FindAllBaanRequest) (*dto.FindAllBaanResponse, *apperrors.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.FindAllBaan(ctx, &baanProto.FindAllBaanRequest{})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return nil, apperrors.InternalServer
		}
		switch st.Code() {
		case codes.NotFound:
			return nil, apperrors.NotFoundError("Baans not found")
		case codes.Internal:
			return nil, apperrors.InternalServerError(err.Error())
		default:
			return nil, apperrors.ServiceUnavailable
		}
	}

	return &dto.FindAllBaanResponse{
		Baans: utils.ProtoToDtoList(res.Baans),
	}, nil
}

func (s *serviceImpl) FindOneBaan(req *dto.FindOneBaanRequest) (*dto.FindOneBaanResponse, *apperrors.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.FindOneBaan(ctx, &baanProto.FindOneBaanRequest{
		Id: req.Id,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return nil, apperrors.InternalServer
		}
		switch st.Code() {
		case codes.NotFound:
			return nil, apperrors.NotFoundError("Baan not found")
		case codes.Internal:
			return nil, apperrors.InternalServerError(err.Error())
		default:
			return nil, apperrors.ServiceUnavailable
		}
	}

	return &dto.FindOneBaanResponse{
		Baan: *utils.ProtoToDto(res.Baan),
	}, nil
}
