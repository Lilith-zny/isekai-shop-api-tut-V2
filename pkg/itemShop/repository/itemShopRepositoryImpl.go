package repository

import (
	"github.com/Lilith-zny/isekai-shop-api-tut-V2/databases"
	"github.com/Lilith-zny/isekai-shop-api-tut-V2/entities"
	_itemShopException "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemShop/exception"
	_itemShopModel "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemShop/model"
	"github.com/labstack/echo/v4"
)

type itemShopRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db databases.Database, logger echo.Logger) ItemShopRepository {
	return &itemShopRepositoryImpl{db, logger}
}

func (r *itemShopRepositoryImpl) Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error) {
	itemList := make([]*entities.Item, 0)

	query := r.db.Connect().Model(&entities.Item{}).Where("is_archive = ?", false) // select * from items

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%")
	}

	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}

	// offset
	// 1 2 3 4 5 6 7 8 9 10
	// 0       | 5       |

	// (page - 1) * size
	// (1 - 1) * 5 = 0
	// (2 - 1) * 5 = 5
	offset := int((itemFilter.Page - 1) * itemFilter.Size)
	limit := int(itemFilter.Size)

	if err := query.Offset(offset).Limit(limit).Find(&itemList).Order("id asc").Error; err != nil {
		r.logger.Errorf("Failed to list items: %s", err.Error())
		return nil, &_itemShopException.ItemListing{}
	}
	return itemList, nil
}

func (r *itemShopRepositoryImpl) Counting(itemFilter *_itemShopModel.ItemFilter) (int64, error) {
	query := r.db.Connect().Model(&entities.Item{}).Where("is_archive = ?", false) // select * from items

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%")
	}

	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}

	var count int64

	if err := query.Count(&count).Error; err != nil {
		r.logger.Errorf("Counting items failed: %s", err.Error())
		return -1, &_itemShopException.ItemCounting{}
	}
	return count, nil
}

func (r *itemShopRepositoryImpl) FindByID(itemID uint64) (*entities.Item, error) {
	item := new(entities.Item)

	if err := r.db.Connect().First(item, itemID).Error; err != nil {
		r.logger.Errorf("Failed to find item by ID: %s", err.Error())
		return nil, &_itemShopException.ItemNotFound{}
	}

	return item, nil
}
