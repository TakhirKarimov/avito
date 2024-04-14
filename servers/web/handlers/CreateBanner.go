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

type RequestBody struct {
	TagIDs    []int                  `json:"tag_ids"`
	FeatureID int                    `json:"feature_id"`
	Content   map[string]interface{} `json:"content"`
	IsActive  bool                   `json:"is_active"`
}

func CreateBanner(c echo.Context) error {
	lg := di.Get("logger").(*logger.Logger)
	ctx := logging.ContextWithLogger(context.Background(), lg.Log)

	adminToken := c.Request().Header.Get("token")
	if adminToken == "" {
		lg.Log.Error("user is not authorized. Empty token")
		return c.String(http.StatusUnauthorized, custom_errors.ErrUserUnAuthorized.Error())
	}
	reqBody := new(RequestBody)
	if err := c.Bind(reqBody); err != nil {
		logging.WithAttrs(ctx).Error("error reading body")
		return c.String(http.StatusBadRequest, err.Error())
	}
	traceId := uuid.New().String()
	reqDto := scheduler.NewCreateBannerDto(reqBody.TagIDs, reqBody.FeatureID, reqBody.Content, reqBody.IsActive)
	logging.WithAttrs(ctx,
		logging.StringAttr("trace_id", traceId),
		logging.AnyAttr("tag_id", reqBody.TagIDs),
		logging.IntAttr("feature_id", reqBody.FeatureID),
		logging.AnyAttr("use_last_revision", reqBody.Content),
		logging.BoolAttr("is_active", reqBody.IsActive),
		logging.StringAttr("admin_token", adminToken),
	).Info("start processing create_banner")

	res, err := scheduler.CreateBanner(ctx, traceId, adminToken, reqDto)
	switch {
	case err == nil:
	case errors.Is(err, gorm.ErrRecordNotFound):
		logging.WithAttrs(ctx).Error("wrong data", logging.ErrAttr(err))
		return c.String(http.StatusNotFound, custom_errors.ErrStatusNotFound.Error())
	default:
		logging.WithAttrs(ctx).Error("internal error: ", logging.ErrAttr(err))
		return c.String(http.StatusInternalServerError, err.Error())
	}
	mapRes := map[string]interface{}{
		"IDs": res,
	}

	logging.WithAttrs(ctx,
		logging.StringAttr("trace_id", traceId),
	).Info("task completed successfully")

	return c.JSON(http.StatusOK, mapRes)
}
