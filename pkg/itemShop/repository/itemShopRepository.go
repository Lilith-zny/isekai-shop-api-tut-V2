package repository

import "github.com/Lilith-zny/isekai-shop-api-tut-V2/entities"

type ItemShopRepository interface {
	Listing() ([]*entities.Item, error)
}
