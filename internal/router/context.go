package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	"github.com/isd-sgcu/rpkm67-gateway/internal/dto"
)

type Context interface {
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
}

type contextImpl struct {
	*gin.Context
}

func NewContext(c *gin.Context) Context {
	return &contextImpl{c}
}

func (c *contextImpl) JSON(statusCode int, obj interface{}) {
	c.Context.JSON(statusCode, obj)
}

func (c *contextImpl) ResponseError(err *apperror.AppError) {
	c.JSON(err.HttpCode, gin.H{"error": err.Error()})
}

func (c *contextImpl) BadRequestError(err string) {
	c.JSON(apperror.BadRequest.HttpCode, gin.H{"error": err})
}

func (c *contextImpl) UnauthorizedError(err string) {
	c.JSON(apperror.Unauthorized.HttpCode, gin.H{"error": err})
}

func (c *contextImpl) ForbiddenError(err string) {
	c.JSON(apperror.Forbidden.HttpCode, gin.H{"error": err})
}

func (c *contextImpl) NotFoundError(err string) {
	c.JSON(apperror.NotFound.HttpCode, gin.H{"error": err})
}

func (c *contextImpl) InternalServerError(err string) {
	c.JSON(apperror.InternalServer.HttpCode, gin.H{"error": err})
}

func (c *contextImpl) ServiceUnavailableError(err string) {
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
