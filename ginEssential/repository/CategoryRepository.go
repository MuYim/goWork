package repository

import (
	"ginEssential/common"
	"ginEssential/model"
	"github.com/jinzhu/gorm"
)

type CateGoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRespository() CateGoryRepository {
	return CateGoryRepository{
		DB: common.GetDB(),
	}
}

func (c CateGoryRepository) Create(name string) (*model.Category, error) {
	category := model.Category{
		Name: name,
	}

	if err := c.DB.Create(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (c CateGoryRepository) Update(category model.Category, name string) (*model.Category, error) {
	if err := c.DB.Model(&category).Update("name", name).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CateGoryRepository) SelectById(id int) (*model.Category, error) {
	var category model.Category
	if err := c.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CateGoryRepository) DeleteById(id int) error {
	if err := c.DB.Delete(model.Category{}).Error; err != nil {
		return err
	}
	return nil
}
