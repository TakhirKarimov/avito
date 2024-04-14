package repository

const defaultAdminToken = "admin"
const defaultUserToken = "user"

type TokensRepo interface {
	GetAdminToken() (string, error)
	GetUserToken() (string, error)
}

type TokensRepoImpl struct {
}

func NewTokensRepo() TokensRepo {
	return &TokensRepoImpl{}
}

func (t *TokensRepoImpl) GetAdminToken() (string, error) {
	return defaultAdminToken, nil
}

func (t *TokensRepoImpl) GetUserToken() (string, error) {
	return defaultUserToken, nil
}
