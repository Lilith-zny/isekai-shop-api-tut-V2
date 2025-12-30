package repository

import (
	"github.com/Lilith-zny/isekai-shop-api-tut-V2/entities"
	_itemManagingException "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemManaging/exception"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type itemManagingRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemManagingRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemManagingRepository {
	return &itemManagingRepositoryImpl{db, logger}
}

func (r *itemManagingRepositoryImpl) Creating(itemEntity *entities.Item) (*entities.Item, error) {
	item := new(entities.Item)

	if err := r.db.Create(itemEntity).Scan(item).Error; err != nil {
		r.logger.Errorf("Creating item failed: %s", err.Error())
		return nil, &_itemManagingException.ItemCreating{}
	}

	return item, nil
}
