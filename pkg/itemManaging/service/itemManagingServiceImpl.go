package service

import (
	"github.com/Lilith-zny/isekai-shop-api-tut-V2/entities"
	_itemManagingModel "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemManaging/model"
	_itemManagingRepository "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemManaging/repository"
	_itemShopModel "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemShop/model"
)

type itemManagingServiceImpl struct {
	itemManagingRepository _itemManagingRepository.ItemManagingRepository
}

func NewItemManagingServiceImpl(itemManagingRepository _itemManagingRepository.ItemManagingRepository) ItemManagingService {
	return &itemManagingServiceImpl{itemManagingRepository}
}

func (s *itemManagingServiceImpl) Creating(itemCreatingReq *_itemManagingModel.ItemCreatingReq) (*_itemShopModel.Item, error) {
	itemEntity := &entities.Item{
		Name:        itemCreatingReq.Name,
		Description: itemCreatingReq.Description,
		Picture:     itemCreatingReq.Picture,
		Price:       itemCreatingReq.Price,
	}

	itemEntityResult, err := s.itemManagingRepository.Creating(itemEntity)
	if err != nil {
		return nil, err
	}

	return itemEntityResult.ToItemModel(), nil
}
