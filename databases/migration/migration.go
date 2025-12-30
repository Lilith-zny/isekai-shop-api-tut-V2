package main

import (
	"fmt"

	"github.com/Lilith-zny/isekai-shop-api-tut-V2/config"
	"github.com/Lilith-zny/isekai-shop-api-tut-V2/databases"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)

	fmt.Println(db.ConnectionGetting())
}
