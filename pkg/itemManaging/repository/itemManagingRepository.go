package repository

import "github.com/Lilith-zny/isekai-shop-api-tut-V2/entities"

type ItemManagingRepository interface {
	Creating(itemEntity *entities.Item) (*entities.Item, error)
}
