package main

import (
	"fmt"
	"os"

	"github.com/Kaibling/IdentityManager/lib/cmd"
	"github.com/Kaibling/IdentityManager/lib/config"
	"github.com/Kaibling/IdentityManager/repositories/csv"
	"github.com/Kaibling/IdentityManager/services"
)

func main() {
	config.InitConfig()
	csvRepo := csv.NewCSVRepo(config.Configuration.DBFilePath)
	services.InitIdentityService(csvRepo)
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

	err := c.Exec(args)
	if err != nil {
		fmt.Println(err.Error())
	}
}
