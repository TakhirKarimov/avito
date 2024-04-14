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

func GetBanner(c echo.Context) error {
	lg := di.Get("logger").(*logger.Logger)
	ctx := logging.ContextWithLogger(context.Background(), lg.Log)

	featureId := c.QueryParam("feature_id")
	tagId := c.QueryParam("tag_id")
	limitStr := c.QueryParam("limit")
	offsetStr := c.QueryParam("offset")
	adminToken := c.Request().Header.Get("token")
	traceId := uuid.New().String()

	if adminToken == "" {
		lg.Log.Error("user is not authorized. Empty token")
		return c.String(http.StatusUnauthorized, custom_errors.ErrUserUnAuthorized.Error())
	}

	logging.WithAttrs(ctx,
		logging.StringAttr("trace_id", traceId),
		logging.StringAttr("tag_id", tagId),
		logging.StringAttr("feature_id", featureId),
		logging.StringAttr("limit", limitStr),
		logging.StringAttr("offset", offsetStr),
		logging.StringAttr("admin_token", adminToken),
	).Info("start processing get_banner")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		logging.WithAttrs(ctx).Error("failed get limit", err)
		return c.String(http.StatusBadRequest, "incorrect limit")
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		logging.WithAttrs(ctx).Error("failed get offset", err)
		return c.String(http.StatusBadRequest, "incorrect offset")
	}

	content, err := scheduler.GetBanner(ctx, traceId, adminToken, featureId, tagId, limit, offset)
	switch {
	case err == nil:
	case errors.Is(err, gorm.ErrRecordNotFound):
		logging.WithAttrs(ctx).Error("wrong data", logging.ErrAttr(err))
		return c.String(http.StatusNotFound, custom_errors.ErrStatusNotFound.Error())
	default:
		logging.WithAttrs(ctx).Error("internal error: ", logging.ErrAttr(err))
		return c.String(http.StatusInternalServerError, err.Error())
	}

	logging.WithAttrs(ctx,
		logging.StringAttr("trace_id", traceId),
	).Info("task completed successfully")

	return c.String(http.StatusOK, content)
}
