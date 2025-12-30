package service

import (
	_itemShopModel "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemShop/model"
	_itemShopRepository "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemShop/repository"
)

type itemShopServiceImpl struct {
	itemShopRepository _itemShopRepository.ItemShopRepository
}

func NewItemShopServicImpl(itemShopRepository _itemShopRepository.ItemShopRepository) ItemShopService {
	return &itemShopServiceImpl{itemShopRepository}
}

func (s *itemShopServiceImpl) Listing() ([]*_itemShopModel.Item, error) {
	// entity
	itemList, err := s.itemShopRepository.Listing()
	if err != nil {
		return nil, err
	}

	// model
	itemModelList := make([]*_itemShopModel.Item, 0)
	for _, item := range itemList {
		itemModelList = append(itemModelList, item.ToItemModel())
	}

	return itemModelList, nil
}
