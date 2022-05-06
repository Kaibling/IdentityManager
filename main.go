package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Kaibling/IdentityManager/cmd"
	"github.com/Kaibling/IdentityManager/config"
	"github.com/Kaibling/IdentityManager/services"
)

func list(args []string) error {
	if len(args) > 0 {
		//show single entry
		err := services.IdentityServiceI.ShowIdentity(args[0], cmd.Flags["a"]) // TODO enum
		if err != nil {
			return err
		}
		return nil
	}
	l := services.IdentityServiceI.GetList()
	for k, v := range l {
		fmt.Printf("%s: %s\n", k, v)
	}
	return nil
}

func new(args []string) error {
	if len(args) < 1 {
		help()
		return errors.New("not enough arguments")
	}
	newDomain := args[0]
	err := services.IdentityServiceI.NewIdentity(newDomain)
	if err != nil {
		return err
	}
	err = services.IdentityServiceI.ShowIdentity(newDomain, cmd.Flags["a"])
	if err != nil {
		return err
	}
	return nil

}

func del(args []string) error {
	if len(args) < 1 {
		help()
		return errors.New("not enough arguments")
	}
	err := services.IdentityServiceI.Delete(args[0])
	if err != nil {
		return err
	}
	return nil
}

func renew(args []string) error {
	if len(args) < 1 {
		help()
		return errors.New("not enough arguments")
	}
	err := services.IdentityServiceI.Renew(args[0])
	if err != nil {
		return err
	}
	err = services.IdentityServiceI.ShowIdentity(args[0], cmd.Flags["a"])
	if err != nil {
		return err
	}
	return nil
}

func help() {
	fmt.Println("help") // TODO
}

func main() {
	config.InitConfig()
	services.InitIdentityService()
	args := os.Args[1:]
	args = cmd.ParseFlags(args)
	if len(args) == 0 {
		// show help
		help()
		return
	}

	c := cmd.NewCommands()
	c.AddCommand("list", list)
	c.AddCommand("new", new)
	c.AddCommand("del", del)
	c.AddCommand("renew", renew)

	err := c.Exec(args)
	if err != nil {
		fmt.Println(err.Error())
	}
	// lock  encrypt data
	// unlock decrypts data
	// renew <domain> -> renews password
}
