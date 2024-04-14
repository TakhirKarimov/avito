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
)

func GetUserBanner(c echo.Context) error {
	lg := di.Get("logger").(*logger.Logger)
	ctx := logging.ContextWithLogger(context.Background(), lg.Log)
	tagId := c.QueryParam("tag_id")
	featureId := c.QueryParam("feature_id")
	useLastRevision := c.QueryParam("use_last_revision")
	userToken := c.Request().Header.Get("token")
	traceId := uuid.New().String()

	logging.WithAttrs(ctx,
		logging.StringAttr("trace_id", traceId),
		logging.StringAttr("tag_id", tagId),
		logging.StringAttr("feature_id", featureId),
		logging.StringAttr("use_last_revision", useLastRevision),
		logging.StringAttr("user_token", userToken),
	).Info("start processing create_user_banner")

	if userToken == "" {
		logging.WithAttrs(ctx).Error("invalid token", logging.ErrAttr(custom_errors.ErrUserUnAuthorized))
		return c.String(http.StatusUnauthorized, custom_errors.ErrUserUnAuthorized.Error())
	}
	content, err := scheduler.GetUserBannerHandler(ctx, traceId, userToken, tagId, featureId, useLastRevision)
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
