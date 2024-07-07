package group

import (
	"context"
	"time"

	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	groupProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/backend/group/v1"
	"go.uber.org/zap"
)

type Service interface {
	FindByUserId(req *dto.FindByUserIdGroupRequest) (*dto.FindByUserIdGroupResponse, *apperror.AppError)
	FindByToken(req *dto.FindByTokenGroupRequest) (*dto.FindByTokenGroupResponse, *apperror.AppError)
	UpdateConfirm(req *dto.UpdateConfirmGroupRequest) (*dto.UpdateConfirmGroupResponse, *apperror.AppError)
	Join(req *dto.JoinGroupRequest) (*dto.JoinGroupResponse, *apperror.AppError)
	Leave(req *dto.LeaveGroupRequest) (*dto.LeaveGroupResponse, *apperror.AppError)
	DeleteMember(req *dto.DeleteMemberGroupRequest) (*dto.DeleteMemberGroupResponse, *apperror.AppError)
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

func (s *serviceImpl) FindByUserId(req *dto.FindByUserIdGroupRequest) (*dto.FindByUserIdGroupResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.FindByUserId(ctx, &groupProto.FindByUserIdGroupRequest{
		UserId: req.UserId,
	})
	if err != nil {
		s.log.Named("FindOne").Error("FindOne: ", zap.Error(err))
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.FindByUserIdGroupResponse{
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
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.FindByTokenGroupResponse{
		Id:     res.Id,
		Token:  res.Token,
		Leader: UserInfoProtoToDto(res.Leader),
	}, nil
}

func (s *serviceImpl) UpdateConfirm(req *dto.UpdateConfirmGroupRequest) (*dto.UpdateConfirmGroupResponse, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.UpdateConfirm(ctx, &groupProto.UpdateConfirmGroupRequest{
		// Group:    GroupDtoToProto(req.Group),
		LeaderId: req.LeaderId,
	})
	if err != nil {
		s.log.Named("Update").Error("Update: ", zap.Error(err))
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.UpdateConfirmGroupResponse{
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
		return nil, apperror.HandleServiceError(err)
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
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.LeaveGroupResponse{
		Group: GroupProtoToDto(res.Group),
	}, nil
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
		return nil, apperror.HandleServiceError(err)
	}

	return &dto.DeleteMemberGroupResponse{
		Group: GroupProtoToDto(res.Group),
	}, nil
}
