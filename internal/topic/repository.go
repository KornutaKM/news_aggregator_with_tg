package topic

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(topic *Topic) error {
	return r.db.Create(topic).Error
}

func (r *Repository) GetAll() ([]Topic, error) {
	var topics []Topic
	err := r.db.Find(&topics).Error
	return topics, err
}

func (r *Repository) GetByName(name string) (*Topic, error) {
	var topic Topic
	err := r.db.Where("name = ?", name).First(&topic).Error
	return &topic, err
}

func (r *Repository) GetByID(id uint) (*Topic, error) {
	var topic Topic
	err := r.db.First(&topic, id).Error
	return &topic, err
}

func (r *Repository) Delete(id uint) error {
	return r.db.Delete(&Topic{}, id).Error
}
