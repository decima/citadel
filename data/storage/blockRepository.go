package storage

import (
	"citadel-api/data/model"
	"citadel-api/utils/container"
	"gorm.io/gorm"
)

func init() {
	container.Add[BlockRepositoryInterface](NewBlockRepository())
}

type BlockRepositoryInterface interface {
	Get(uuid string) (*model.Block, error)
	GetAll() ([]model.Block, error)
	GetChildren(parentId *string) ([]model.Block, error)
	GetByType(blockType string) ([]model.Block, error)
	CountByType(blockType string) (int, error)
	Create(block *model.Block) (*model.Block, error)
	Update(block *model.Block) (*model.Block, error)
	Delete(uuid string) error
}

type BlockRepository struct {
}

func (r *BlockRepository) Get(uuid string) (*model.Block, error) {
	db := GetDB()
	var block model.Block
	result := db.Preload("Children").First(&block, "id = ?", uuid)
	if result.Error != nil {
		return nil, result.Error
	}
	return &block, nil
}

func (r *BlockRepository) GetAll() ([]model.Block, error) {
	db := GetDB()
	var blocks []model.Block
	result := db.Find(&blocks)
	if result.Error != nil {
		return nil, result.Error
	}
	return blocks, nil

}

func (r *BlockRepository) GetChildren(parentId *string) ([]model.Block, error) {
	db := GetDB()
	var blocks []model.Block
	var result *gorm.DB
	if parentId == nil {
		result = db.Where("parent_id IS NULL").Find(&blocks)

	} else {
		result = db.Where("parent_id = ?", parentId).Find(&blocks)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return blocks, nil
}

func (r *BlockRepository) Create(block *model.Block) (*model.Block, error) {
	db := GetDB()
	result := db.Save(block)
	if result.Error != nil {
		return nil, result.Error
	}
	return block, nil
}

func (r *BlockRepository) Update(block *model.Block) (*model.Block, error) {
	db := GetDB()
	result := db.Save(block)
	if result.Error != nil {
		return nil, result.Error
	}
	return block, nil
}

func (r *BlockRepository) Delete(uuid string) error {
	db := GetDB()
	result := db.Delete(&model.Block{}, "id = ?", uuid)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *BlockRepository) GetByType(blockType string) ([]model.Block, error) {
	db := GetDB()
	var blocks []model.Block
	result := db.Where("type = ?", blockType).Find(&blocks)
	if result.Error != nil {
		return nil, result.Error
	}
	return blocks, nil
}

func (r *BlockRepository) CountByType(blockType string) (int, error) {
	db := GetDB()
	var count int64
	result := db.Model(&model.Block{}).Where("type = ?", blockType).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(count), nil
}

func NewBlockRepository() BlockRepositoryInterface {
	return &BlockRepository{}
}
