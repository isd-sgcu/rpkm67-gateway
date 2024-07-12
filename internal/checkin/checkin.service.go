package checkin

import (
	"context"
	"time"

	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	checkinProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/checkin/checkin/v1"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type Service interface {
	Create(ctx context.Context, req *dto.CreateCheckInRequest) (*dto.CreateCheckInResponse, *apperror.AppError)
	FindByUserID(ctx context.Context, req *dto.FindByUserIdCheckInRequest) (*dto.FindByUserIdCheckInResponse, *apperror.AppError)
	FindByEmail(ctx context.Context, req *dto.FindByEmailCheckInRequest) (*dto.FindByEmailCheckInResponse, *apperror.AppError)
}

type serviceImpl struct {
	client checkinProto.CheckInServiceClient
	log    *zap.Logger
	tracer trace.Tracer
}

func NewService(client checkinProto.CheckInServiceClient, log *zap.Logger, tracer trace.Tracer) Service {
	return &serviceImpl{
		client: client,
		log:    log,
		tracer: tracer,
	}
}

func (s *serviceImpl) Create(ctx context.Context, req *dto.CreateCheckInRequest) (*dto.CreateCheckInResponse, *apperror.AppError) {
	ctx, span := s.tracer.Start(ctx, "service.checkin.Create")
	defer span.End()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := s.client.Create(ctx, &checkinProto.CreateCheckInRequest{
		Email:  req.Email,
		UserId: req.UserID,
		Event:  req.Event,
	})
	if err != nil {
		s.log.Named("Create").Error("Create: ", zap.Error(err))
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.CreateCheckInResponse{
		CheckIn: &dto.CheckIn{
			ID:     res.CheckIn.Id,
			UserID: res.CheckIn.UserId,
			Email:  res.CheckIn.Email,
			Event:  res.CheckIn.Event,
		},
	}, nil
}

func (s *serviceImpl) FindByEmail(ctx context.Context, req *dto.FindByEmailCheckInRequest) (*dto.FindByEmailCheckInResponse, *apperror.AppError) {
	ctx, span := s.tracer.Start(ctx, "service.checkin.FindByEmail")
	defer span.End()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := s.client.FindByEmail(ctx, &checkinProto.FindByEmailCheckInRequest{
		Email: req.Email,
	})
	if err != nil {
		s.log.Named("FindByEmail").Error("FindByEmail: ", zap.Error(err))
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.FindByEmailCheckInResponse{
		CheckIns: ProtoToDtos(res.CheckIns),
	}, nil
}

func (s *serviceImpl) FindByUserID(ctx context.Context, req *dto.FindByUserIdCheckInRequest) (*dto.FindByUserIdCheckInResponse, *apperror.AppError) {
	ctx, span := s.tracer.Start(ctx, "service.checkin.FindByUserId")
	defer span.End()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	md := metadata.Pairs(
		"timestamp", time.Now().Format(time.StampNano),
		"req.user.id", req.UserID,
	)
	ctx = metadata.NewOutgoingContext(ctx, md)

	res, err := s.client.FindByUserId(ctx, &checkinProto.FindByUserIdCheckInRequest{
		UserId: req.UserID,
	})
	if err != nil {
		s.log.Named("FindByUserID").Error("FindByUserID: ", zap.Error(err))
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.FindByUserIdCheckInResponse{
		CheckIns: ProtoToDtos(res.CheckIns),
	}, nil
}
