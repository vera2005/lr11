package main

import (
	"flag"
	"log"

	_ "github.com/lib/pq"
	"github.com/vera2005/lr10/internal/count/api"
	"github.com/vera2005/lr10/internal/count/config"
	"github.com/vera2005/lr10/internal/count/provider"
	"github.com/vera2005/lr10/internal/count/usecase"
)

func main() {
	// Считываем аргументы командной строки
	configPath := flag.String("config-path", "D:\\Go\\lr10\\\\configs\\count.yaml", "путь к файлу конфигурации")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	//Инициализация провайдера
	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	//Инициализация бизнес-логики
	use := usecase.NewUsecase(cfg.Usecase.DefaultCount, prv)
	//Инициализация сервера
	srv := api.NewServer(cfg.IP, cfg.Port, cfg.API.MaxNum, use)

	srv.Run()
}
