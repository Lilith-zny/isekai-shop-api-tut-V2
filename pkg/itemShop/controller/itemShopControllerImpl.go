package controller

import (
	"net/http"

	"github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/custom"
	_itemShopModel "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemShop/model"
	_itemShopService "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemShop/service"
	"github.com/labstack/echo/v4"
)

type itemShopControllerImpl struct {
	itemShopService _itemShopService.ItemShopService
}

func NewItemShopControllerImpl(itemShopService _itemShopService.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}

func (c *itemShopControllerImpl) Listing(pctx echo.Context) error {
	itemFilter := new(_itemShopModel.ItemFilter)

	// if err := pctx.Bind(itemFilter); err != nil {
	// 	return custom.Error(pctx, http.StatusBadRequest, err.Error())
	// }

	// // เรา validate ว่า ข้อมูลมีมั้ยที่ ItemFilter Model ได้ส่งมาป่าว ถ้าทำแบบนี้มันต้องเขียนทุกรอบเราสามารถ custom ได้
	// validating := validator.New()

	// if err := validating.Struct(itemFilter); err != nil {
	// 	return custom.Error(pctx, http.StatusBadRequest, err.Error())
	// }

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(itemFilter); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err.Error())
	}

	itemModelList, err := c.itemShopService.Listing(itemFilter)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusOK, itemModelList)
}
