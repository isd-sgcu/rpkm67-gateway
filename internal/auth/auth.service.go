package auth

import (
	"context"
	"time"

	"github.com/isd-sgcu/rpkm67-gateway/apperrors"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	authProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/auth/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service interface {
	Validate()
	RefreshToken()
}

type serviceImpl struct {
	client authProto.AuthServiceClient
}

func NewService(client authProto.AuthServiceClient) Service {
	return &serviceImpl{
		client: client,
	}
}

func (s *serviceImpl) Validate() {
}
func (s *serviceImpl) RefreshToken() {
}

func (s *serviceImpl) SignUp(ctx context.Context, req *dto.SignUpRequest) (*dto.Credential, *apperrors.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.SignUp(ctx, &authProto.SignUpRequest{
		StudentId: req.StudentId,
		Password:  req.Password,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Tel:       req.Tel,
		Role:      "student",
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return nil, apperrors.InternalServer
		}
		switch st.Code() {
		case codes.AlreadyExists:
			return nil, apperrors.BadRequestError("User already exists")
		case codes.Internal:
			return nil, apperrors.InternalServer
		default:
			return nil, apperrors.ServiceUnavailable
		}
	}

	return &dto.Credential{
		AccessToken:  res.Credential.AccessToken,
		RefreshToken: res.Credential.RefreshToken,
		ExpiresIn:    int(res.Credential.ExpiresIn),
	}, nil
}

func (s *serviceImpl) SignIn(ctx context.Context, req *dto.SignInRequest) (*dto.Credential, *apperrors.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.SignIn(ctx, &authProto.SignInRequest{
		StudentId: req.StudentId,
		Password:  req.Password,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return nil, apperrors.InternalServer
		}
		switch st.Code() {
		case codes.AlreadyExists:
			return nil, apperrors.BadRequestError("User already exists")
		case codes.Internal:
			return nil, apperrors.InternalServer
		default:
			return nil, apperrors.ServiceUnavailable
		}
	}

	return &dto.Credential{
		AccessToken:  res.Credential.AccessToken,
		RefreshToken: res.Credential.RefreshToken,
		ExpiresIn:    int(res.Credential.ExpiresIn),
	}, nil
}

func (s *serviceImpl) SignOut(ctx context.Context, req *dto.TokenPayloadAuth) (*dto.SignOutResponse, *apperrors.AppError) {
	return nil, nil
}

func (s *serviceImpl) ForgotPassword(ctx context.Context, req *dto.ForgotPasswordRequest) (*dto.ForgotPasswordResponse, *apperrors.AppError) {
	return nil, nil
}

func (s *serviceImpl) ResetPassword(ctx context.Context, req *dto.ResetPasswordRequest) (*dto.ResetPasswordResponse, *apperrors.AppError) {
	return nil, nil
}
