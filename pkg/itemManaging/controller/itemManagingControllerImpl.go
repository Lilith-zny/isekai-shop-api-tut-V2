package controller

import (
	"net/http"

	"github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/custom"
	_itemManagingModel "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemManaging/model"
	_ItemManagingService "github.com/Lilith-zny/isekai-shop-api-tut-V2/pkg/itemManaging/service"
	"github.com/labstack/echo/v4"
)

type itemManagingControllerImpl struct {
	itemManagingService _ItemManagingService.ItemManagingService
}

func NewItemManagingControllerImpl(itemManagingService _ItemManagingService.ItemManagingService) ItemManagingController {
	return &itemManagingControllerImpl{itemManagingService}
}

func (c *itemManagingControllerImpl) Creating(pctx echo.Context) error {
	itemCreatingReq := new(_itemManagingModel.ItemCreatingReq)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(itemCreatingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err.Error())
	}

	item, err := c.itemManagingService.Creating(itemCreatingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusCreated, item)
}
