package cmd

import (
	"errors"
	"testing"
)

func TestCommands_AddCommand(t *testing.T) {
	c := NewCommands()
	c.AddCommand("test", func(s []string) error { return nil })
	err := c.Exec([]string{"test"})
	if err != nil {
		t.Error()
	}
	c.AddCommand("test2", func(s []string) error { return errors.New("") })
	err = c.Exec([]string{"test2"})
	if err == nil {
		t.Error("should be error")
	}
	err = c.Exec([]string{"test3"})
	if err == nil {
		t.Error("should be error")
	}
}
