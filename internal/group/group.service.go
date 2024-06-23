package group

import (
	"context"
	"time"

	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	groupProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/group/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service interface {
	FindOne(req *dto.FindOneGroupRequest) (*dto.FindOneGroupResponse, *apperror.AppError)
	FindByToken(req *dto.FindByTokenGroupRequest) (*dto.FindByTokenGroupResponse, *apperror.AppError)
	Update(req *dto.UpdateGroupRequest) (*dto.UpdateGroupResponse, *apperror.AppError)
	Join(req *dto.JoinGroupRequest) (*dto.JoinGroupResponse, *apperror.AppError)
	DeleteMember(req *dto.DeleteMemberGroupRequest) (*dto.DeleteMemberGroupResponse, *apperror.AppError)
	Leave(req *dto.LeaveGroupRequest) (*dto.LeaveGroupResponse, *apperror.AppError)
	SelectBaan(req *dto.SelectBaanRequest) (*dto.SelectBaanResponse, *apperror.AppError)
}

type serviceImpl struct {
	client groupProto.GroupServiceClient
	log    *zap.Logger
}

func NewService(client groupProto.GroupServiceClient, log *zap.Logger) Service {
	return &serviceImpl{
		client: client,
		log:    log,
	}
}

func (s *serviceImpl) DeleteMember(req *dto.DeleteMemberGroupRequest) (*dto.DeleteMemberGroupResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.DeleteMember(ctx, &groupProto.DeleteMemberGroupRequest{
		UserId:   req.UserId,
		LeaderId: req.LeaderId,
	})
	if err != nil {
		s.log.Named("DeleteMember").Error("DeleteMember: ", zap.Error(err))
		st, ok := status.FromError(err)
		if !ok {
			return nil, apperror.InternalServer
		}
		switch st.Code() {
		case codes.InvalidArgument:
			return nil, apperror.BadRequestError("")
		case codes.Internal:
			return nil, apperror.InternalServerError(err.Error())
		default:
			return nil, apperror.ServiceUnavailable
		}
	}

	return &dto.DeleteMemberGroupResponse{
		Group: GroupProtoToDto(res.Group),
	}, nil
}

func (s *serviceImpl) FindByToken(req *dto.FindByTokenGroupRequest) (*dto.FindByTokenGroupResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.FindByToken(ctx, &groupProto.FindByTokenGroupRequest{
		Token: req.Token,
	})
	if err != nil {
		s.log.Named("FindByToken").Error("FindByToken: ", zap.Error(err))
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

	return &dto.FindByTokenGroupResponse{
		Id:     res.Id,
		Token:  res.Token,
		Leader: UserInfoProtoToDto(res.Leader),
	}, nil
}

func (s *serviceImpl) FindOne(req *dto.FindOneGroupRequest) (*dto.FindOneGroupResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.FindOne(ctx, &groupProto.FindOneGroupRequest{
		UserId: req.UserId,
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

	return &dto.FindOneGroupResponse{
		Group: GroupProtoToDto(res.Group),
	}, nil
}

func (s *serviceImpl) Join(req *dto.JoinGroupRequest) (*dto.JoinGroupResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.Join(ctx, &groupProto.JoinGroupRequest{
		Token:  req.Token,
		UserId: req.UserId,
	})

	if err != nil {
		s.log.Named("Join").Error("Join: ", zap.Error(err))
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

	return &dto.JoinGroupResponse{
		Group: GroupProtoToDto(res.Group),
	}, nil
}

func (s *serviceImpl) Leave(req *dto.LeaveGroupRequest) (*dto.LeaveGroupResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.Leave(ctx, &groupProto.LeaveGroupRequest{
		UserId: req.UserId,
	})
	if err != nil {
		s.log.Named("Leave").Error("Leave: ", zap.Error(err))
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

	return &dto.LeaveGroupResponse{
		Group: GroupProtoToDto(res.Group),
	}, nil
}

func (s *serviceImpl) SelectBaan(req *dto.SelectBaanRequest) (*dto.SelectBaanResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.SelectBaan(ctx, &groupProto.SelectBaanRequest{
		UserId: req.UserId,
		Baans:  req.Baans,
	})
	if err != nil {
		s.log.Named("SelectBaan").Error("SelectBaan: ", zap.Error(err))
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

	return &dto.SelectBaanResponse{
		Success: res.Success,
	}, nil
}

func (s *serviceImpl) Update(req *dto.UpdateGroupRequest) (*dto.UpdateGroupResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.Update(ctx, &groupProto.UpdateGroupRequest{
		Group:    GroupDtoToProto(req.Group),
		LeaderId: req.LeaderId,
	})
	if err != nil {
		s.log.Named("Update").Error("Update: ", zap.Error(err))
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

	return &dto.UpdateGroupResponse{
		Group: GroupProtoToDto(res.Group),
	}, nil
}
