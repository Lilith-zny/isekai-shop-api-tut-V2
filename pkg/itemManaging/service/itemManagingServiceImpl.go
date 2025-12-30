package service

import (
	_itemManagingRepository "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemManaging/repository"
)

type itemManagingServiceImpl struct {
	itemManagingRepository _itemManagingRepository.ItemManagingRepository
}

func NewItemManagingServiceImpl(itemManagingRepository _itemManagingRepository.ItemManagingRepository) ItemManagingService {
	return &itemManagingServiceImpl{itemManagingRepository}
}
