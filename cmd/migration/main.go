package main

import (
	"api-example/pkg/mySqlExt"
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/pressly/goose/v3"
)

func main() {
	migrationPath := "./migrations/"
	argsWithoutProg := os.Args[1:]
	command := argsWithoutProg[0]
	gooseArgs := []string{}

	mysqlConfig := mySqlExt.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_DATABASE"),
	}

	dbConnection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local&multiStatements=true",
		mysqlConfig.Username,
		mysqlConfig.Password,
		mysqlConfig.Host+":"+mysqlConfig.Port,
		mysqlConfig.DBName,
	)

	params := []string{}
	if len(argsWithoutProg) > 1 {
		params = argsWithoutProg[1:]
	}

	slog.Info("Migration", "command", command, "params", params)
	// Check db connection
	db, err := goose.OpenDBWithDriver("mysql", dbConnection)
	if err != nil {
		panic(fmt.Sprintf("goose: failed to open DB: %v\n", err))
	}

	defer func() {
		if err := db.Close(); err != nil {
			panic(fmt.Sprintf("goose: failed to close DB: %v\n", err))
		}
	}()

	if command == "create" && len(params) == 0 {
		panic(fmt.Errorf("goose: migration name is required"))
	}

	switch command {
	case "create":
		migrationType := "sql"
		migrationName := params[0]
		gooseArgs = append(gooseArgs, migrationName, migrationType)
	case "up", "down":
		gooseArgs = append(gooseArgs, params...)
	}

	// executing actual goose
	if err := goose.RunContext(context.Background(), command, db, migrationPath, gooseArgs...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
