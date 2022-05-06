package repositories

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

func ensureFile(filePath string) error {
	var f *os.File
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		f, err = os.Create(filePath)
		if err != nil {
			return err
		}
	}
	defer f.Close()
	return nil
}

func ReadFromFile(filePath string) ([][]string, error) {
	ensureFile(filePath)
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	r := csv.NewReader(f)

	return r.ReadAll()
}

func WriteData(filePath string, line []string) error {
	ensureFile(filePath)
	var f *os.File

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0655)
	if err != nil {
		return err
	}

	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()
	if err := w.Write(line); err != nil {
		return err
	}

	return nil
}

func RemoveLine(domain, filePath string) error {
	ensureFile(filePath)
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0655)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()

	r := csv.NewReader(f)
	data, err := r.ReadAll()
	if err != nil {
		fmt.Println(err.Error())
	}
	line := 0
	for i := range data {
		if data[i][0] == domain {
			line = i
		}
	}
	newData := data[0:line]
	newData = append(newData, data[line+1:]...)

	f.Close()
	f, err = os.Create(filePath)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	w := csv.NewWriter(f)
	err = w.WriteAll(newData)
	if err != nil {
		return err
	}
	return nil

}
