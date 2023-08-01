package main

import (
	"fmt"
	"mygpt/pkg/infrastructure/datastore"
	"mygpt/pkg/utils"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Warn("Failed To Load ENV: ", err)
	}
	utils.InitializeLogger()
}

func main() {
	args := os.Args
	if len(args) == 0 || (!utils.Contains(args, "gen", false) && !utils.Contains(args, "apply", false)) {
		logrus.Fatal("No Stages Passed")
	}

	if utils.Contains(args, "gen", false) {
		fmt.Println("---------------------------------------------------------------------------------------")
		dsn := datastore.GetConnectionString(utils.Pointer(os.Getenv("PG_MIGRATION_DB")))
		db, err := pq.Open(dsn)
		if err != nil {
			logrus.Warn(err)
			logrus.Info("Switching to DockerDB")
			dsn = "docker://postgres/15/test?search_path=public"
		} else {
			defer db.Close()
		}

		cmd := exec.Command("atlas", "migrate", "diff", "--dir", "file://db/migration", "--to", "file://db/db-schema.sql", "--dev-url", dsn)
		out, err := cmd.CombinedOutput()
		fmt.Println(strings.Trim(string(out), "\n"))
		if err != nil {
			logrus.Fatal("Gen: ", err)
		} else {
			fmt.Println("Migration Plan Updated")
		}
	}

	if utils.Contains(args, "apply", false) {
		fmt.Println("---------------------------------------------------------------------------------------")
		dsn := datastore.GetConnectionString(nil)
		db, err := pq.Open(dsn)
		if err != nil {
			fmt.Println(dsn)
			logrus.Fatal(err)
		} else {
			defer db.Close()
		}

		cmd := exec.Command("atlas", "migrate", "apply", "--dir", "file://db/migration", "--url", dsn)
		if utils.Contains(args, "dry", false) {
			mg_dsn := datastore.GetConnectionString(utils.Pointer(os.Getenv("PG_MIGRATION_DB")))
			db, err := pq.Open(mg_dsn)
			if err != nil {
				logrus.Warn(err)
				logrus.Info("Switching to DockerDB")
				mg_dsn = "docker://postgres/15/test?search_path=public"
			} else {
				defer db.Close()
			}
			cmd = exec.Command("atlas", "schema", "apply", "--to", "file://db/db-schema.sql", "--url", dsn, "--dev-url", mg_dsn, "--dry-run")
		}
		out, err := cmd.CombinedOutput()
		fmt.Println(strings.Trim(string(out), "\n"))
		if err != nil {
			logrus.Fatal("Apply: ", err)
		} else {
			fmt.Println("Migration Plan Applied")
		}
	}
	fmt.Println("---------------------------------------------------------------------------------------")
}
