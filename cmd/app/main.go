package main

import (
	"backend/config"
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/server"
	"backend/internal/service"
	"log"
	"os"
)

//	@title			hackaton
//	@version		1.0
//	@description	My swagger doc.

//	@contact.name	karl
//	@contact.email	foreverwantlive@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:9090
// @BasePath	/
// @schemes	http
func main() {
	var cfgPath string
	switch len(os.Args[1:]) {
	case 1:
		cfgPath = os.Args[1]
	case 0:
		cfgPath = "./.env"
	default:
		log.Print("USAGE: go run [CONFIG_PATH]")
		return
	}

	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		log.Print(err)
		return
	}

	db, err := repository.LoadDB(cfg.MongoDB)
	if err != nil {
		log.Print(err)
		return
	}

	r := repository.NewRepository(db)

	s := service.NewService(*r, cfg.OpenAiToken)

	h := handler.NewHandler(s)

	if err := server.StartServer(cfg.Host, cfg.Port, h.GenerateRoutes()); err != nil {
		log.Print(err)
		return
	}

}
