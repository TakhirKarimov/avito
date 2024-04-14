package repository

import (
	"avito/db/orm"
	"avito/pkg/di"
)

const statusActiveBanner = 1

type ContentRepo interface {
	GetContentByTagAndFeatureId(tagId, featureId string) (*orm.Content, error)
	GetContentByTagOrFeatureId(tagId, featureId string, offset, limit int) ([]orm.Content, error)
	CreateBanner(tagId, featureId string, content string, isActive int) (int, error)
	UpdateBanner(id int, updates map[string]interface{}) error
	DeleteBannerById(id int) error
}

type ContentRepoImpl struct {
	Db         *Connect
	Id         int
	TagID      string
	FeatureID  string
	UserToken  string
	AdminToken string
	Content    string
	IsActive   int
}

func NewContentRepo() ContentRepo {
	return &ContentRepoImpl{
		Db: di.Get("db").(*Connect),
	}
}

func (c *ContentRepoImpl) GetContentByTagAndFeatureId(tagId, featureId string) (*orm.Content, error) {
	var content orm.Content
	err := c.Db.Conn.
		Model(&content).
		Where(
			"tag_id = ? AND feature_id = ? AND is_active = ?",
			tagId,
			featureId,
			statusActiveBanner,
		).
		Take(&content).
		Error
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func (c *ContentRepoImpl) GetContentByTagOrFeatureId(tagId, featureId string, offset, limit int) ([]orm.Content, error) {
	var model orm.Content
	var results []orm.Content
	query := c.Db.Conn.Model(&model)
	if offset != 0 {
		query.Offset(offset)
	}
	if limit != 0 {
		query.Limit(limit)
	}
	if tagId != "" {
		query.Where("tag_id = ?", tagId)
	}
	if featureId != "" {
		query.Where("feature_id = ?", featureId)
	}
	err := query.Find(&results).Error
	return results, err
}

func (c *ContentRepoImpl) CreateBanner(tagId, featureId string, content string, isActive int) (int, error) {
	data := orm.Content{
		TagID:      tagId,
		FeatureID:  featureId,
		IsActive:   isActive,
		UserToken:  defaultUserToken,
		AdminToken: defaultAdminToken,
		Content:    content,
	}
	err := c.Db.Conn.Create(&data).Error
	if err != nil {
		return 0, err
	}

	var res orm.Content
	err = c.Db.Conn.
		Model(&orm.Content{}).
		Where("tag_id = ? AND feature_id = ?",
			data.TagID,
			data.FeatureID).
		Take(&res).
		Error

	return res.Id, err
}

func (c *ContentRepoImpl) UpdateBanner(id int, updates map[string]interface{}) error {
	updateMap := make(map[string]interface{})
	for field, value := range updates {
		if value != "" {
			updateMap[field] = value
		}
	}
	if len(updateMap) > 0 {
		if err := c.Db.Conn.
			Model(orm.Content{}).
			Where("id = ?", id).
			Updates(updateMap).
			Error; err != nil {
			return err
		}
	}
	return nil
}

func (c *ContentRepoImpl) DeleteBannerById(id int) error {
	return c.Db.Conn.Delete(orm.Content{}, id).Error
}
