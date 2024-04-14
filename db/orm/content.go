package orm

type Content struct {
	Id         int    `gorm:"primaryKey"`
	TagID      string `gorm:"column:tag_id"`
	FeatureID  string `gorm:"column:feature_id"`
	UserToken  string `gorm:"column:user_token"`
	AdminToken string `gorm:"column:admin_token"`
	Content    string `gorm:"column:content"`
	IsActive   int    `gorm:"column:is_active"`
}

func (с *Content) TableName() string {
	return "content"
}
