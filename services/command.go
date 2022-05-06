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
		fmt.Println("new <domain>")
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
		fmt.Println("del <domain>")
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
		fmt.Println("renew <domain>")
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
		fmt.Println("change <domain> <property> <value>")
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
	fmt.Printf("generate and store identities\ncommands:\n")
	fmt.Printf("list : lists all domains and corresponding email\n")
	fmt.Printf("list <domain> : shows details to an identity\n")
	fmt.Printf("del : removes identity and domain\n")
	fmt.Printf("change : change properties of identity\n")
	fmt.Printf("new : create random identity\n")
	fmt.Printf("renew : recrate password of identity\n")
}
