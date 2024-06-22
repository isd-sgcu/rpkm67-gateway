package dto

type Credential struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL3BiZX..."`
	RefreshToken string `json:"refresh_token" example:"e7e84d54-7518-4..."`
	ExpiresIn    int    `json:"expires_in" example:"3600"`
}

type SignInRequest struct {
	StudentId string `json:"student_id" validate:"required"`
	Password  string `json:"password" validate:"required,gte=6,lte=30"`
}
type SignUpRequest struct {
	StudentId string `json:"student_id" validate:"required"`
	Password  string `json:"password" validate:"required,gte=6,lte=30"`
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Tel       string `json:"tel" validate:"required"`
}

type TokenPayloadAuth struct {
	UserId string `json:"user_id"`
	Role   string `json:"role"`
}

type SignupResponse struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type SignOutResponse struct {
	IsSuccess bool `json:"is_success"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ForgotPasswordResponse struct {
	IsSuccess bool `json:"is_success"`
}

type ResetPasswordRequest struct {
	Token    string `json:"token" validate:"required"`
	Password string `json:"password" validate:"required,gte=6,lte=30"`
}

type ResetPasswordResponse struct {
	IsSuccess bool `json:"is_success"`
}
