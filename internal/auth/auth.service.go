package auth

import (
	"context"
	"time"

	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	authProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/auth/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service interface {
	Validate()
	RefreshToken()
	SignUp(req *dto.SignUpRequest) (*dto.Credential, *apperror.AppError)
	SignIn(req *dto.SignInRequest) (*dto.Credential, *apperror.AppError)
	SignOut(req *dto.TokenPayloadAuth) (*dto.SignOutResponse, *apperror.AppError)
	ForgotPassword(req *dto.ForgotPasswordRequest) (*dto.ForgotPasswordResponse, *apperror.AppError)
	ResetPassword(req *dto.ResetPasswordRequest) (*dto.ResetPasswordResponse, *apperror.AppError)
}

type serviceImpl struct {
	client authProto.AuthServiceClient
	log    *zap.Logger
}

func NewService(client authProto.AuthServiceClient, log *zap.Logger) Service {
	return &serviceImpl{
		client: client,
		log:    log,
	}
}

func (s *serviceImpl) Validate() {
}
func (s *serviceImpl) RefreshToken() {
}

func (s *serviceImpl) SignUp(req *dto.SignUpRequest) (*dto.Credential, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.SignUp(ctx, &authProto.SignUpRequest{
		Email:     req.Email,
		Password:  req.Password,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Role:      "student",
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return nil, apperror.InternalServer
		}
		switch st.Code() {
		case codes.AlreadyExists:
			return nil, apperror.BadRequestError("User already exists")
		case codes.Internal:
			return nil, apperror.InternalServerError(err.Error())
		default:
			return nil, apperror.ServiceUnavailable
		}
	}

	return &dto.Credential{
		AccessToken:  res.Credential.AccessToken,
		RefreshToken: res.Credential.RefreshToken,
		ExpiresIn:    int(res.Credential.ExpiresIn),
	}, nil
}

func (s *serviceImpl) SignIn(req *dto.SignInRequest) (*dto.Credential, *apperror.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.SignIn(ctx, &authProto.SignInRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return nil, apperror.InternalServer
		}
		switch st.Code() {
		case codes.AlreadyExists:
			return nil, apperror.BadRequestError("User already exists")
		case codes.Internal:
			return nil, apperror.InternalServerError(err.Error())
		default:
			return nil, apperror.ServiceUnavailable
		}
	}

	return &dto.Credential{
		AccessToken:  res.Credential.AccessToken,
		RefreshToken: res.Credential.RefreshToken,
		ExpiresIn:    int(res.Credential.ExpiresIn),
	}, nil
}

func (s *serviceImpl) SignOut(req *dto.TokenPayloadAuth) (*dto.SignOutResponse, *apperror.AppError) {
	return nil, nil
}

func (s *serviceImpl) ForgotPassword(req *dto.ForgotPasswordRequest) (*dto.ForgotPasswordResponse, *apperror.AppError) {
	return nil, nil
}

func (s *serviceImpl) ResetPassword(req *dto.ResetPasswordRequest) (*dto.ResetPasswordResponse, *apperror.AppError) {
	return nil, nil
}
