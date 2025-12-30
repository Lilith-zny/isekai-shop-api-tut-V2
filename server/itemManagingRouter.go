package server

import (
	_itemManagingController "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemManaging/controller"
	_itemManagingRepository "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemManaging/repository"
	_ItemManagingService "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemManaging/service"
	_itemShopRepository "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemShop/repository"
)

func (s *echoServer) initItemManagingRouter() {
	router := s.app.Group("/v1/item-managing")

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemManagingRepository := _itemManagingRepository.NewItemManagingRepositoryImpl(s.db, s.app.Logger)
	itemManagingService := _ItemManagingService.NewItemManagingServiceImpl(itemManagingRepository, itemShopRepository)
	itemManagingController := _itemManagingController.NewItemManagingControllerImpl(itemManagingService)

	router.POST("", itemManagingController.Creating)
	router.PATCH("/:itemID", itemManagingController.Editing)
	router.DELETE("/:itemID", itemManagingController.Archiving)
}
