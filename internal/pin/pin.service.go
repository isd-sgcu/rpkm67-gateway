package pin

import (
	"context"
	"time"

	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	pinProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/pin/v1"
	"go.uber.org/zap"
)

type Service interface {
	FindAll(req *dto.FindAllPinRequest) (*dto.FindAllPinResponse, *apperror.AppError)
	ResetPin(req *dto.ResetPinRequest) (*dto.ResetPinResponse, *apperror.AppError)
	CheckPin(req *dto.CheckPinRequest) (*dto.CheckPinResponse, *apperror.AppError)
}

type serviceImpl struct {
	client pinProto.PinServiceClient
	log    *zap.Logger
}

func NewService(client pinProto.PinServiceClient, log *zap.Logger) Service {
	return &serviceImpl{
		client: client,
		log:    log,
	}
}

func (s *serviceImpl) FindAll(req *dto.FindAllPinRequest) (*dto.FindAllPinResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.FindAll(ctx, &pinProto.FindAllPinRequest{})
	if err != nil {
		s.log.Named("FindAllBaan").Error("FindAllBaan: ", zap.Error(err))
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.FindAllPinResponse{
		Pins: ProtoToDtoList(res.Pins),
	}, nil
}

func (s *serviceImpl) ResetPin(req *dto.ResetPinRequest) (*dto.ResetPinResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.ResetPin(ctx, &pinProto.ResetPinRequest{
		ActivityId: req.ActivityId,
	})
	if err != nil {
		s.log.Named("FindOneBaan").Error("FindOneBaan: ", zap.Error(err))
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.ResetPinResponse{
		Success: res.Success,
	}, nil
}

func (s *serviceImpl) CheckPin(req *dto.CheckPinRequest) (*dto.CheckPinResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.CheckPin(ctx, &pinProto.CheckPinRequest{
		ActivityId: req.ActivityId,
		Code:       req.Code,
	})
	if err != nil {
		s.log.Named("CheckPin").Error("CheckPin: ", zap.Error(err))
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.CheckPinResponse{
		IsMatch: res.IsMatch,
	}, nil
}
