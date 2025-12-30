package controller

import (
	_ItemManagingService "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemManaging/service"
)

type itemManagingControllerImpl struct {
	itemManagingService _ItemManagingService.ItemManagingService
}

func NewItemManagingControllerImpl(itemManagingService _ItemManagingService.ItemManagingService) ItemManagingController {
	return &itemManagingControllerImpl{itemManagingService}
}
