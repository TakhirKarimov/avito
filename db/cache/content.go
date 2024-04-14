package cache

import (
	"avito/db/orm"
	"avito/db/repository"
	"avito/pkg/di"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

const expirationCacheTimeInMinute = 5

type ContentRepoCacheImpl struct {
	Repo      repository.ContentRepo
	CacheRepo *CacheConnect
}

func NewContentRepoCache() repository.ContentRepo {
	return &ContentRepoCacheImpl{
		CacheRepo: di.Get("cache_db").(*CacheConnect),
		Repo:      repository.NewContentRepo(),
	}
}

func (c *ContentRepoCacheImpl) GetContentByTagAndFeatureId(tagId, featureId string) (*orm.Content, error) {
	var content *orm.Content
	key := fmt.Sprintf("%s_%s", tagId, featureId)
	cacheValue, err := c.CacheRepo.Conn.Get(context.Background(), key).Result()
	switch {
	case err == nil:
	case err == redis.Nil:
		content, err = c.Repo.GetContentByTagAndFeatureId(tagId, featureId)
		if err != nil {
			return nil, err
		}
		serialized, err := json.Marshal(content)
		if err != nil {
			return nil, err
		}
		c.CacheRepo.Conn.Set(context.Background(), key, serialized, expirationCacheTimeInMinute*time.Minute)
		return content, nil
	default:
		return nil, fmt.Errorf("failed fetch cache data: %v", err)
	}

	if err = json.Unmarshal([]byte(cacheValue), &content); err != nil {
		return nil, fmt.Errorf("failed decode cache data: %v", err)
	}

	return content, nil
}

func (c *ContentRepoCacheImpl) GetContentByTagOrFeatureId(tagId, featureId string, offset, limit int) ([]orm.Content, error) {
	return c.Repo.GetContentByTagOrFeatureId(tagId, featureId, offset, limit)
}

func (c *ContentRepoCacheImpl) CreateBanner(tagId, featureId string, content string, isActive int) (int, error) {
	return c.Repo.CreateBanner(tagId, featureId, content, isActive)
}

func (c *ContentRepoCacheImpl) UpdateBanner(id int, updates map[string]interface{}) error {
	return c.Repo.UpdateBanner(id, updates)
}

func (c *ContentRepoCacheImpl) DeleteBannerById(id int) error {
	return c.Repo.DeleteBannerById(id)
}
