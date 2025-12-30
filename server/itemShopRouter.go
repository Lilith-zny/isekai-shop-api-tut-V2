package server

import (
	_itemShopController "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemShop/controller"
	_itemShopRepository "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemShop/repository"
	_itemShopService "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemShop/service"
)

func (s *echoServer) initItemShopRouter() {
	router := s.app.Group("/v1/item-shop")

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemShopService := _itemShopService.NewItemShopServicImpl(itemShopRepository)
	itemShopController := _itemShopController.NewItemShopControllerImpl(itemShopService)

	router.GET("/", itemShopController.Listing)
}
