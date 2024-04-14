package scheduler

import (
	"avito/db/repository"
	"avito/model/custom_errors"
	"context"
	"encoding/json"
	"github.com/theartofdevel/logging"
)

type UpdateBannerDto struct {
	Id        int                    `json:"id"`
	TagIDs    []int                  `json:"tag_ids"`
	FeatureID int                    `json:"feature_id"`
	Content   map[string]interface{} `json:"content"`
	IsActive  bool                   `json:"is_active"`
}

func NewUpdateBannerDto(id int, tagIDs []int, featureID int, content map[string]interface{}, isActive bool) *UpdateBannerDto {
	return &UpdateBannerDto{
		Id:        id,
		TagIDs:    tagIDs,
		FeatureID: featureID,
		Content:   content,
		IsActive:  isActive,
	}
}

func UpdateBanner(ctx context.Context, traceId, adminToken string, req *UpdateBannerDto) error {
	logging.L(ctx).With(logging.StringAttr("trace_id", traceId)).Info("task being processed")
	tokenRepo := repository.NewTokensRepo()
	tableToken, err := tokenRepo.GetAdminToken()
	if err != nil {
		return err
	}
	if adminToken != tableToken {
		return custom_errors.ErrUserHasNotAccess
	}
	var isActiveInt int
	if req.IsActive == true {
		isActiveInt = 1
	}
	contentJSON, err := json.Marshal(req.Content)
	if err != nil {
		return err
	}
	repo := repository.NewContentRepo()
	for _, tagId := range req.TagIDs {
		m := map[string]interface{}{
			"tag_id":     tagId,
			"feature_id": req.FeatureID,
			"content":    string(contentJSON),
			"is_active":  isActiveInt,
		}
		err := repo.UpdateBanner(req.Id, m)
		if err != nil {
			return err
		}
	}

	return nil
}
