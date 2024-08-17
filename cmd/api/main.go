package main

import (
	repository_user "api-example/internal/repository/user"
	"api-example/internal/server"
	service_user "api-example/internal/service/user"
	"api-example/pkg/mySqlExt"
	"fmt"
	"os"
)

func main() {

	mysqlDB, err := mySqlExt.New(mySqlExt.Config{
		Host:         os.Getenv("DB_HOST"),
		Port:         os.Getenv("DB_PORT"),
		Username:     os.Getenv("DB_USERNAME"),
		Password:     os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_DATABASE"),
		MaxIdleConns: 50,
		MaxOpenConns: 50,
		MaxLifeTime:  0,
		MaxIdleTime:  10,
	})
	if err != nil {
		panic(err)
	}

	// setup repositories
	userRepo := repository_user.New(mysqlDB)

	// setup services
	userService := service_user.New(userRepo)

	server := server.NewServer(server.Services{
		UserService: userService,
	})

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
