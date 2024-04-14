package scheduler

import (
	"avito/db/repository"
	"avito/model/custom_errors"
	"context"
	"github.com/theartofdevel/logging"
)

func DeleteBanner(ctx context.Context, traceId string, id int, adminToken string) error {
	logging.L(ctx).With(logging.StringAttr("trace_id", traceId)).Info("task being processed")
	tokenRepo := repository.NewTokensRepo()
	tableToken, err := tokenRepo.GetAdminToken()
	if err != nil {
		return err
	}
	if adminToken != tableToken {
		return custom_errors.ErrUserHasNotAccess
	}
	repo := repository.NewContentRepo()

	return repo.DeleteBannerById(id)
}
