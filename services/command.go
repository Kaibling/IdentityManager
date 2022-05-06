package services

import (
	"errors"
	"fmt"

	"github.com/Kaibling/IdentityManager/lib/cmd"
)

func List(args []string) error {
	if len(args) > 0 {
		//show single entry
		err := IdentityServiceI.ShowIdentity(args[0], cmd.Flags["a"]) // TODO enum
		if err != nil {
			return err
		}
		return nil
	}
	l := IdentityServiceI.GetList()
	for k, v := range l {
		fmt.Printf("%s: %s\n", k, v)
	}
	return nil
}

func New(args []string) error {
	if len(args) < 1 {
		Help()
		return errors.New("not enough arguments")
	}
	newDomain := args[0]
	err := IdentityServiceI.NewIdentity(newDomain)
	if err != nil {
		return err
	}
	err = IdentityServiceI.ShowIdentity(newDomain, cmd.Flags["a"])
	if err != nil {
		return err
	}
	return nil

}

func Del(args []string) error {
	if len(args) < 1 {
		Help()
		return errors.New("not enough arguments")
	}
	err := IdentityServiceI.Delete(args[0])
	if err != nil {
		return err
	}
	return nil
}

func Renew(args []string) error {
	if len(args) < 1 {
		Help()
		return errors.New("not enough arguments")
	}
	err := IdentityServiceI.Renew(args[0])
	if err != nil {
		return err
	}
	err = IdentityServiceI.ShowIdentity(args[0], cmd.Flags["a"])
	if err != nil {
		return err
	}
	return nil
}

func Change(args []string) error {
	if len(args) < 2 {
		Help()
		return errors.New("not enough arguments")
	}
	domain := args[0]
	property := args[1]
	data := args[2]
	err := IdentityServiceI.Change(domain, property, data)
	if err != nil {
		return err
	}
	err = IdentityServiceI.ShowIdentity(domain, cmd.Flags["a"])
	if err != nil {
		return err
	}
	return nil
}

func Help() {
	fmt.Println("help") // TODO
}
