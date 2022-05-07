package main

import (
	"fmt"
	"os"

	"github.com/Kaibling/IdentityManager/lib/cmd"
	"github.com/Kaibling/IdentityManager/lib/config"
	"github.com/Kaibling/IdentityManager/lib/database"
	"github.com/Kaibling/IdentityManager/repositories/csv"
	"github.com/Kaibling/IdentityManager/repositories/sqlite"
	"github.com/Kaibling/IdentityManager/services"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var r services.IdentityRepo
	switch config.Configuration.Dialect {
	case "SQLITE":
		db, err := database.InitDBConnection()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = database.Migrate(db, []database.DBMigrator{sqlite.NewSQLiteMigrator(db)})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		r, err = sqlite.NewSQLiteRepo(db)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	case "CSV":
		r = csv.NewCSVRepo(config.Configuration.DBFilePath)
	default:
		fmt.Printf("DB Dialect '%s' not valid\n", config.Configuration.Dialect)
		return
	}

	services.InitIdentityService(r)
	args := os.Args[1:]
	args = cmd.ParseFlags(args)
	if len(args) == 0 {
		// show help
		services.Help()
		return
	}

	c := cmd.NewCommands()
	c.AddCommand("list", services.List)
	c.AddCommand("new", services.New)
	c.AddCommand("del", services.Del)
	c.AddCommand("renew", services.Renew)
	c.AddCommand("change", services.Change)

	err = c.Exec(args)
	if err != nil {
		fmt.Println(err.Error())
	}
}
