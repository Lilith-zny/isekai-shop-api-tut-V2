package server

import (
	_itemManagingController "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemManaging/controller"
	_itemManagingRepository "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemManaging/repository"
	_ItemManagingService "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemManaging/service"
)

func (s *echoServer) initItemManagingRouter() {
	router := s.app.Group("/v1/item-managing")

	itemManagingRepository := _itemManagingRepository.NewItemManagingRepositoryImpl(s.db, s.app.Logger)
	itemManagingService := _ItemManagingService.NewItemManagingServiceImpl(itemManagingRepository)
	itemManagingController := _itemManagingController.NewItemManagingControllerImpl(itemManagingService)

	router.POST("", itemManagingController.Creating)
}
