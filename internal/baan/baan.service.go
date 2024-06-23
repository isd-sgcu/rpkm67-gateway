package baan

import (
	"context"
	"time"

	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	baanProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/baan/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service interface {
	FindAllBaan(req *dto.FindAllBaanRequest) (*dto.FindAllBaanResponse, *apperror.AppError)
	FindOneBaan(req *dto.FindOneBaanRequest) (*dto.FindOneBaanResponse, *apperror.AppError)
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

func (s *serviceImpl) FindAllBaan(req *dto.FindAllBaanRequest) (*dto.FindAllBaanResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.FindAllBaan(ctx, &baanProto.FindAllBaanRequest{})
	if err != nil {
		s.log.Named("FindAllBaan").Error("FindAllBaan: ", zap.Error(err))
		st, ok := status.FromError(err)
		if !ok {
			return nil, apperror.InternalServer
		}
		switch st.Code() {
		case codes.InvalidArgument:
			return nil, apperror.BadRequestError("Invalid argument")
		case codes.Internal:
			return nil, apperror.InternalServerError(err.Error())
		default:
			return nil, apperror.ServiceUnavailable
		}
	}

	return &dto.FindAllBaanResponse{
		Baans: ProtoToDtoList(res.Baans),
	}, nil
}

func (s *serviceImpl) FindOneBaan(req *dto.FindOneBaanRequest) (*dto.FindOneBaanResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.FindOneBaan(ctx, &baanProto.FindOneBaanRequest{
		Id: req.Id,
	})
	if err != nil {
		s.log.Named("FindOneBaan").Error("FindOneBaan: ", zap.Error(err))
		st, ok := status.FromError(err)
		if !ok {
			return nil, apperror.InternalServer
		}
		switch st.Code() {
		case codes.NotFound:
			return nil, apperror.NotFoundError("Baan not found")
		case codes.Internal:
			return nil, apperror.InternalServerError(err.Error())
		default:
			return nil, apperror.ServiceUnavailable
		}
	}

	return &dto.FindOneBaanResponse{
		Baan: ProtoToDto(res.Baan),
	}, nil
}
