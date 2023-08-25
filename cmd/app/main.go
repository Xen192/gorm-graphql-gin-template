package main

import (
	"fmt"
	"mygpt/pkg/infrastructure/datastore"
	"mygpt/pkg/infrastructure/graphql"
	"mygpt/pkg/infrastructure/router"
	"mygpt/pkg/utils"
	"mygpt/query"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	utils.InitializeLogger()
	err := godotenv.Load()
	if err != nil {
		logrus.Warn("Failed To Load ENV: ", err)
	}

	db := datastore.GetInstance()
	query.SetDefault(db)

	loc, err := time.LoadLocation("Asia/Kathmandu")
	if err != nil {
		logrus.Fatal(err)
	}
	time.Local = loc
	datastore.InitializeClerk()
}

func main() {
	srv := graphql.NewServer()
	e := router.New(srv)
	fmt.Println("\n\033[35mStarting Server On: \033[31mhttp://localhost:" + os.Getenv("BE_PORT") + "\033[0m\n")
	logrus.Error(e.Run(":" + os.Getenv("BE_PORT")))
}
