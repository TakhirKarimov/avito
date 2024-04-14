package scheduler

import (
	"avito/db/repository"
	"avito/model/custom_errors"
	"context"
	"encoding/json"
	"github.com/theartofdevel/logging"
	"strconv"
)

type CreateBannerDto struct {
	TagIDs    []int                  `json:"tag_ids"`
	FeatureID int                    `json:"feature_id"`
	Content   map[string]interface{} `json:"content"`
	IsActive  bool                   `json:"is_active"`
}

func NewCreateBannerDto(tagIDs []int, featureID int, content map[string]interface{}, isActive bool) *CreateBannerDto {
	return &CreateBannerDto{
		TagIDs:    tagIDs,
		FeatureID: featureID,
		Content:   content,
		IsActive:  isActive,
	}
}

func CreateBanner(ctx context.Context, traceId, adminToken string, req *CreateBannerDto) ([]int, error) {
	logging.L(ctx).With(logging.StringAttr("trace_id", traceId)).Info("task being processed")
	tokenRepo := repository.NewTokensRepo()
	tableToken, err := tokenRepo.GetAdminToken()
	if err != nil {
		return nil, err
	}
	if adminToken != tableToken {
		return nil, custom_errors.ErrUserHasNotAccess
	}
	var isActiveInt int
	if req.IsActive == true {
		isActiveInt = 1
	}
	contentJSON, err := json.Marshal(req.Content)
	if err != nil {
		return nil, err
	}
	repo := repository.NewContentRepo()
	res := make([]int, 0, len(req.TagIDs))
	for _, tagId := range req.TagIDs {
		id, err := repo.CreateBanner(strconv.Itoa(tagId), strconv.Itoa(req.FeatureID), string(contentJSON), isActiveInt)
		if err != nil {
			return res, err
		}
		res = append(res, id)
	}
	return res, nil
}
