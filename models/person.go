package models

import "encoding/json"

type Person struct {
	FirstName         string
	LastName          string
	Domain            string
	Email             string
	Username          string
	Password          string
	BirthDate         string
	Gender            string
	Street            string
	Country           string
	PhoneNumber       string
	CreditCard        string
	HomePage          string
	Company           string
	Job               string
	SecurityQuestions []string
}

func (p *Person) ToString() string {
	byt, _ := json.Marshal(p)
	return string(byt)
}

func (p *Person) ToMap() (map[string]interface{}, error) {
	byt, _ := json.Marshal(p)
	var m map[string]interface{}
	err := json.Unmarshal(byt, &m)
	if err != nil {
		return map[string]interface{}{}, err
	}
	return m, nil
}

func (p *Person) FromMap(m map[string]interface{}) error {
	byt, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byt, p)
	if err != nil {
		return err
	}
	return nil
}

func (p *Person) ToReduced() PersonReduced {
	return PersonReduced{
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,
		Username:  p.Username,
		Password:  p.Password,
	}
}

type PersonReduced struct {
	FirstName string
	LastName  string
	Email     string
	Username  string
	Password  string
}
