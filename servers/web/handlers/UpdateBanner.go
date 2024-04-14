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

type RequestBodyUpdate struct {
	TagIDs    []int                  `json:"tag_ids"`
	FeatureID int                    `json:"feature_id"`
	Content   map[string]interface{} `json:"content"`
	IsActive  bool                   `json:"is_active"`
}

func UpdateBanner(c echo.Context) error {
	lg := di.Get("logger").(*logger.Logger)
	ctx := logging.ContextWithLogger(context.Background(), lg.Log)

	adminToken := c.Request().Header.Get("token")
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	traceId := uuid.New().String()

	logging.WithAttrs(ctx,
		logging.StringAttr("trace_id", traceId),
		logging.AnyAttr("admin_token", adminToken),
		logging.IntAttr("id", id),
	).Info("start processing update_banner")

	if err != nil {
		logging.WithAttrs(ctx).Error("failed get id", err)
		return c.JSON(http.StatusBadRequest, "incorrect id")
	}
	if adminToken == "" {
		lg.Log.Error("user is not authorized. Empty token")
		return c.JSON(http.StatusUnauthorized, custom_errors.ErrUserUnAuthorized.Error())
	}
	reqBody := new(RequestBodyUpdate)
	if err := c.Bind(reqBody); err != nil {
		logging.WithAttrs(ctx).Error("error reading body")
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	reqDto := scheduler.NewUpdateBannerDto(id, reqBody.TagIDs, reqBody.FeatureID, reqBody.Content, reqBody.IsActive)
	err = scheduler.UpdateBanner(ctx, traceId, adminToken, reqDto)
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

	return c.JSON(http.StatusOK, "")
}
