package repository

import (
	"github.com/Lilith-zny/isekai-shop-api-tut-V2/entities"
	_itemShopModel "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemShop/model"
)

type ItemShopRepository interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error)
	Counting(itemFilter *_itemShopModel.ItemFilter) (int64, error)
	FindByID(itemID uint64) (*entities.Item, error)
}
