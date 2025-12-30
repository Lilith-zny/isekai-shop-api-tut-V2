package controller

import "github.com/labstack/echo/v4"

type ItemShopController interface {
	// pctx == c
	Listing(pctx echo.Context) error
}
