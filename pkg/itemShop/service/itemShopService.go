package service

import _itemShopModel "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemShop/model"

type ItemShopService interface {
	Listing() ([]*_itemShopModel.Item, error)
}
