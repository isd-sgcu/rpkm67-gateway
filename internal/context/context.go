package context

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/constant"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	"github.com/isd-sgcu/rpkm67-gateway/internal/metrics"
)

type Ctx interface {
	JSON(statusCode int, obj interface{})
	ResponseError(err *apperror.AppError)
	BadRequestError(err string)
	InternalServerError(err string)
	NewUUID() uuid.UUID
	Bind(obj interface{}) error
	Param(key string) string
	Query(key string) string
	PostForm(key string) string
	FormFile(key string, allowedContentType map[string]struct{}, maxFileSize int64) (*dto.DecomposedFile, error)
	GetString(key string) string
	GetHeader(key string) string
}

type contextImpl struct {
	*gin.Context
	httpMethod constant.Method
	path       string
	reqMetrics metrics.RequestMetrics
}

func NewContext(c *gin.Context, httpMethod constant.Method, path string, reqMetrics metrics.RequestMetrics) Ctx {
	return &contextImpl{
		Context:    c,
		httpMethod: httpMethod,
		path:       path,
		reqMetrics: reqMetrics,
	}
}

func (c *contextImpl) JSON(statusCode int, obj interface{}) {
	c.reqMetrics.AddRequest(c.path, c.httpMethod, statusCode)
	c.Context.JSON(statusCode, obj)
}

func (c *contextImpl) ResponseError(err *apperror.AppError) {
	c.reqMetrics.AddRequest(c.path, c.httpMethod, err.HttpCode)
	c.JSON(err.HttpCode, gin.H{"error": err.Error()})
}

func (c *contextImpl) BadRequestError(err string) {
	c.reqMetrics.AddRequest(c.path, c.httpMethod, apperror.BadRequest.HttpCode)
	c.JSON(apperror.BadRequest.HttpCode, gin.H{"error": err})
}

func (c *contextImpl) UnauthorizedError(err string) {
	c.reqMetrics.AddRequest(c.path, c.httpMethod, apperror.Unauthorized.HttpCode)
	c.JSON(apperror.Unauthorized.HttpCode, gin.H{"error": err})
}

func (c *contextImpl) ForbiddenError(err string) {
	c.reqMetrics.AddRequest(c.path, c.httpMethod, apperror.Forbidden.HttpCode)
	c.JSON(apperror.Forbidden.HttpCode, gin.H{"error": err})
}

func (c *contextImpl) NotFoundError(err string) {
	c.reqMetrics.AddRequest(c.path, c.httpMethod, apperror.NotFound.HttpCode)
	c.JSON(apperror.NotFound.HttpCode, gin.H{"error": err})
}

func (c *contextImpl) InternalServerError(err string) {
	c.reqMetrics.AddRequest(c.path, c.httpMethod, apperror.InternalServer.HttpCode)
	c.JSON(apperror.InternalServer.HttpCode, gin.H{"error": err})
}

func (c *contextImpl) ServiceUnavailableError(err string) {
	c.reqMetrics.AddRequest(c.path, c.httpMethod, apperror.ServiceUnavailable.HttpCode)
	c.JSON(apperror.ServiceUnavailable.HttpCode, gin.H{"error": err})
}

func (c *contextImpl) NewUUID() uuid.UUID {
	return uuid.New()
}

func (c *contextImpl) Bind(obj interface{}) error {
	return c.Context.Bind(obj)
}

func (c *contextImpl) Param(key string) string {
	return c.Context.Param(key)
}

func (c *contextImpl) Query(key string) string {
	return c.Context.Query(key)
}

func (c *contextImpl) PostForm(key string) string {
	return c.Context.PostForm(key)
}

func (c *contextImpl) FormFile(key string, allowedContentType map[string]struct{}, maxFileSize int64) (*dto.DecomposedFile, error) {
	file, err := c.Context.FormFile(key)
	if err != nil {
		return nil, err
	}

	data, err := ExtractFile(file, allowedContentType, maxFileSize)
	if err != nil {
		return nil, err
	}

	return &dto.DecomposedFile{
		Filename: file.Filename,
		Data:     data,
	}, nil
}

func (c *contextImpl) GetString(key string) string {
	return c.Context.GetString(key)
}

func (c *contextImpl) GetHeader(key string) string {
	return c.Context.GetHeader(key)
}
