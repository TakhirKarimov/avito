package scheduler

import (
	"avito/db/cache"
	"avito/db/repository"
	"avito/model/custom_errors"
	"context"
	"github.com/theartofdevel/logging"
	"strings"
)

func GetUserBannerHandler(ctx context.Context, traceId, token, tagId, featureId, useLastRevision string) (string, error) {
	logging.L(ctx).With(logging.StringAttr("trace_id", traceId)).Info("task being processed")
	tokenRepo := repository.NewTokensRepo()
	tableToken, err := tokenRepo.GetUserToken()
	if err != nil {
		return "", err
	}
	if token != tableToken {
		return "", custom_errors.ErrUserHasNotAccess
	}

	var repo repository.ContentRepo
	if strings.ToLower(useLastRevision) == "true" {
		repo = cache.NewContentRepoCache()
	} else {
		repo = repository.NewContentRepo()
	}

	content, err := repo.GetContentByTagAndFeatureId(tagId, featureId)
	if err != nil {
		return "", err
	}
	return content.Content, nil
}
