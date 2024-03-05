package main

import (
	"fmt"

	"github.com/natchapol2347/woodman_web_back/adaptor/storage"
	"github.com/natchapol2347/woodman_web_back/internal/config"
	"github.com/natchapol2347/woodman_web_back/internal/database"
	"github.com/natchapol2347/woodman_web_back/service"
)

func main() {
	// time.DateOnly
	fmt.Println(config.LoadConfig())
	db, _ := database.NewPostgres(config.LoadConfig())

	_ = service.NewService(storage.NewStorage(db))

}
