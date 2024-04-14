package handlers

import (
	"avito/internal/scheduler"
	"avito/model/custom_errors"
	"avito/pkg/di"
	"avito/pkg/logger"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/theartofdevel/logging"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func DeleteBanner(c echo.Context) error {
	lg := di.Get("logger").(*logger.Logger)
	ctx := logging.ContextWithLogger(context.Background(), lg.Log)

	adminToken := c.Request().Header.Get("token")
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	traceId := uuid.New().String()

	if err != nil {
		logging.WithAttrs(ctx).Error("incorrect id")
		return c.JSON(http.StatusBadRequest, "incorrect id")
	}
	if adminToken == "" {
		lg.Log.Error("user is not authorized. Empty token")
		return c.JSON(http.StatusUnauthorized, custom_errors.ErrUserUnAuthorized.Error())
	}

	err = scheduler.DeleteBanner(ctx, traceId, id, adminToken)
	switch {
	case err == nil:
	case errors.Is(err, gorm.ErrRecordNotFound):
		logging.WithAttrs(ctx).Error("wrong data", logging.ErrAttr(err))
		return c.JSON(http.StatusNotFound, custom_errors.ErrStatusNotFound.Error())
	default:
		logging.WithAttrs(ctx).Error("internal error: ", logging.ErrAttr(err))
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	logging.WithAttrs(ctx,
		logging.StringAttr("trace_id", traceId),
	).Info("task completed successfully")

	return c.String(http.StatusNoContent, "")
}
