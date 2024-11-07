package main

import (
	"fmt"
	"log"
	"ps-direct/internal/database"
	"ps-direct/internal/request"
	"ps-direct/internal/service"
	"ps-direct/pkg/logger"
)

func main() {

	db, err := database.PostgresConnect()
	if err != nil {
		log.Fatalln(err)
	}

	redis, err := database.RedisConnect()
	if err != nil {
		log.Fatalln(err)
	}

	loggerJournal, err := logger.NewLogger(db)
	if err != nil {
		log.Fatalln(err)
	}

	loggerJournal.LoggerBasic(logger.INFO_LOG, "Parser Statistics Yandex start. Redis connected & Postgres connected")

	store := service.NewStore(db, redis)

	data, err := request.GetStatApi()
	if err != nil {
		fmt.Println(err)
		loggerJournal.LoggerBasic(logger.ERROR_LOG, err.Error())
	}

	err = store.InsertStatistics(data)
	if err != nil {
		fmt.Println(err)
		loggerJournal.LoggerBasic(logger.ERROR_LOG, err.Error())
	}

}
