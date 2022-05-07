package csv

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/Kaibling/IdentityManager/models"
)

type CSVRepo struct {
	filePath string
}

func NewCSVRepo(filePath string) *CSVRepo {
	return &CSVRepo{filePath: filePath}
}

func (c *CSVRepo) ReadAll() ([]models.Person, error) {
	ensureFile(c.filePath)
	f, err := os.Open(c.filePath)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	persons := []models.Person{}
	for i := range records {
		var person models.Person
		err = json.Unmarshal([]byte(records[i][1]), &person)
		if err != nil {
			return nil, err
		}

		persons = append(persons, person)
	}
	return persons, nil
}

func (c *CSVRepo) Create(p models.Person) error {
	line := []string{p.Domain, p.ToString()}

	ensureFile(c.filePath)
	var f *os.File

	f, err := os.OpenFile(c.filePath, os.O_APPEND|os.O_WRONLY, 0655)
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

func (c *CSVRepo) Delete(domain string) error {
	ensureFile(c.filePath)
	f, err := os.OpenFile(c.filePath, os.O_RDONLY, 0655)
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
	f, err = os.Create(c.filePath)
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

func (c *CSVRepo) Update(p models.Person) error {
	ensureFile(c.filePath)
	f, err := os.OpenFile(c.filePath, os.O_RDONLY, 0655)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()

	r := csv.NewReader(f)
	data, err := r.ReadAll()
	if err != nil {
		fmt.Println(err.Error())
	}

	rline := []string{p.Domain, p.ToString()}

	for i := range data {
		if data[i][0] == rline[0] {
			data[i][1] = rline[1]
		}
	}
	f.Close()
	f, err = os.Create(c.filePath)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	w := csv.NewWriter(f)
	err = w.WriteAll(data)
	if err != nil {
		return err
	}
	return nil

}

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
