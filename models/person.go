package models

import "encoding/json"

type PersonFull struct {
	FirstName         string
	LastName          string
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

func (p *PersonFull) ToString() string {
	byt, _ := json.Marshal(p)
	return string(byt)
}

func (p *PersonFull) ToReduced() PersonReduced {
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
