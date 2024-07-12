package db

import (
	"github.com/isd-sgcu/rpkm67-gateway/internal/context"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Handler interface {
	CleanDb(c context.Ctx)
}

type handlerImpl struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewHandler(db *gorm.DB, log *zap.Logger) Handler {
	return &handlerImpl{
		db:  db,
		log: log,
	}
}

// RefreshToken godoc
// @Summary Clean all data in database (only in development environment)
// @Description must be used only in development environment
// @Tags db
// @Produce json
// @Success 200 {object} dto.Credential
// @Failure 500 {object} apperror.AppError
// @Router /clean-db [get]
func (h *handlerImpl) CleanDb(c context.Ctx) {
	err := h.db.Exec("TRUNCATE TABLE groups, users, selections, stamps, check_ins RESTART IDENTITY CASCADE").Error
	if err != nil {
		h.log.Named("CleanDb").Error("Failed to clean database", zap.Error(err))
		c.InternalServerError("Failed to clean database")
		return
	}

	c.JSON(200, "Database cleaned successfully")
}
