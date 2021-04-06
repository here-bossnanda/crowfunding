package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindById(userId int) ([]Campaign, error)
}

type repositiory struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repositiory {
	return &repositiory{db}
}

func (r *repositiory) FindAll() ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Find(campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return  campaigns, nil
}

func (r *repositiory) FindById(userId int) ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Where("userId = ?", userId).Preload("CampaignImages", "campaign_images.is_thumbnail = 1").Find(campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return  campaigns, nil
}