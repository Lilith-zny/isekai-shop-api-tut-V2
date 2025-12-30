package main

import (
	"github.com/Lilith-zny/isekai-shop-api-tut-V2/config"
	"github.com/Lilith-zny/isekai-shop-api-tut-V2/databases"
	"github.com/Lilith-zny/isekai-shop-api-tut-V2/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db.ConnectionGetting())

	server.Start()
}
