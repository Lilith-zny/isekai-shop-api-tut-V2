package repository

import (
	"github.com/Lilith-zny/isekai-shop-api-tut-V2/entities"
	_itemShopModel "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemShop/model"
)

type ItemShopRepository interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error)
}
