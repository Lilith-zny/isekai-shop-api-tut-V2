package repository

import (
	"github.com/Lilith-zny/isekai-shop-api-tut-V2/entities"
	_itemManagingModel "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemManaging/model"
)

type ItemManagingRepository interface {
	Creating(itemEntity *entities.Item) (*entities.Item, error)
	Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (uint64, error)
}
