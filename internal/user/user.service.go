package user

import (
	"context"
	"time"

	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	userProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/user/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service interface {
	FindOne(req *dto.FindOneUserRequest) (*dto.FindOneUserResponse, *apperror.AppError)
	Update(req *dto.UpdateUserRequest) (*dto.UpdateUserResponse, *apperror.AppError)
}

type serviceImpl struct {
	client userProto.UserServiceClient
	log    *zap.Logger
}

func NewService(client userProto.UserServiceClient, log *zap.Logger) Service {
	return &serviceImpl{
		client: client,
		log:    log,
	}
}

func (s *serviceImpl) FindOne(req *dto.FindOneUserRequest) (*dto.FindOneUserResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.FindOne(ctx, &userProto.FindOneUserRequest{
		Id: req.Id,
	})
	if err != nil {
		s.log.Named("FindOne").Error("FindOne: ", zap.Error(err))
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

	return &dto.FindOneUserResponse{
		User: ProtoToDto(res.User),
	}, nil
}

func (s *serviceImpl) Update(req *dto.UpdateUserRequest) (*dto.UpdateUserResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updateReq := s.createUpdateUserRequestProto(req)

	res, err := s.client.Update(ctx, updateReq)
	if err != nil {
		s.log.Named("Update").Error("Update: ", zap.Error(err))
		st, ok := status.FromError(err)
		if !ok {
			return nil, apperror.InternalServer
		}
		switch st.Code() {
		case codes.NotFound:
			return nil, apperror.NotFoundError("User not found")
		case codes.Internal:
			return nil, apperror.InternalServerError(err.Error())
		default:
			return nil, apperror.ServiceUnavailable
		}
	}

	return &dto.UpdateUserResponse{
		User: ProtoToDto(res.User),
	}, nil
}

func (s *serviceImpl) createUpdateUserRequestProto(req *dto.UpdateUserRequest) *userProto.UpdateUserRequest {
	return &userProto.UpdateUserRequest{
		Id:          req.Id,
		Nickname:    req.Nickname,
		Title:       req.Title,
		Firstname:   req.Firstname,
		Lastname:    req.Lastname,
		Year:        int32(req.Year),
		Faculty:     req.Faculty,
		Tel:         req.Tel,
		ParentTel:   req.ParentTel,
		Parent:      req.Parent,
		FoodAllergy: req.FoodAllergy,
		DrugAllergy: req.DrugAllergy,
		Illness:     req.Illness,
		PhotoKey:    req.PhotoKey,
		PhotoUrl:    req.PhotoUrl,
		Baan:        req.Baan,
		ReceiveGift: int32(req.ReceiveGift),
		GroupId:     req.GroupId,
	}
}
