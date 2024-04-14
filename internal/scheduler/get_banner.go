package scheduler

import (
	"avito/db/orm"
	"avito/db/repository"
	"avito/model/custom_errors"
	"context"
	"github.com/theartofdevel/logging"
)

func GetBanner(ctx context.Context, traceId, adminToken, featureId, tagId string, limit, offset int) (string, error) {
	logging.L(ctx).With(logging.StringAttr("trace_id", traceId)).Info("task being processed")
	tokenRepo := repository.NewTokensRepo()
	tableToken, err := tokenRepo.GetAdminToken()
	if err != nil {
		return "", err
	}
	if adminToken != tableToken {
		return "", custom_errors.ErrUserHasNotAccess
	}
	repo := repository.NewContentRepo()
	contents, err := repo.GetContentByTagOrFeatureId(tagId, featureId, offset, limit)
	if err != nil {
		return "", err
	}
	return buildResponse(contents), nil
}

func buildResponse(contents []orm.Content) string {
	var response string
	for _, c := range contents {
		response += c.Content + "\n"
	}
	return response
}
