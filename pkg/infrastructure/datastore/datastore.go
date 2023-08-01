package datastore

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"sync"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var instance *gorm.DB
var once sync.Once

func GetConnectionString(custom_db *string) string {
	pgOptions := url.Values{}
	if os.Getenv("DB_PROFILE") == "local" {
		pgOptions.Add("sslmode", "disable")
	}
	pgOptions.Add("connect_timeout", "10")
	// pgOptions.Add("client_encoding", "utf8")

	pgURL := &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD")),
		Host:     fmt.Sprintf("%s:%s", os.Getenv("PG_HOST"), os.Getenv("PG_PORT")),
		Path:     os.Getenv("PG_DB"),
		RawQuery: pgOptions.Encode(),
	}
	if custom_db != nil {
		pgURL.Path = *custom_db
	}

	return pgURL.String()
}

func Connect() *gorm.DB {
	dsn := GetConnectionString(nil)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}

func GetInstance() *gorm.DB {
	once.Do(func() {
		instance = Connect()
		switch os.Getenv("APP_ENV") {
		case "prod":
			instance.Logger.LogMode(logger.Error)
		case "dev":
			instance.Logger = logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					LogLevel:                  logger.Info,
					IgnoreRecordNotFoundError: true,
					ParameterizedQueries:      true,
					Colorful:                  true,
				})
		}
	})
	return instance
}

func InitializeClerk() {
	client, err := clerk.NewClient(os.Getenv("CLERK_SECRET_KEY"))
	if err != nil {
		logrus.Panic("Failed To Initiate Clerk Client: ", err)
	}
	ClerkClient = client
}
