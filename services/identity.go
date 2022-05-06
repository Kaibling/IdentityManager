package services

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Kaibling/IdentityManager/lib/config"
	g "github.com/Kaibling/IdentityManager/lib/generator"
	"github.com/Kaibling/IdentityManager/models"
	"github.com/Kaibling/IdentityManager/repositories"
)

var IdentityServiceI *IdentityService = nil

type IdentityService struct {
	identities map[string]models.PersonFull
}

func InitIdentityService() {
	is := &IdentityService{identities: map[string]models.PersonFull{}}
	is.ReadFromFile(config.Configuration.DBFilePath)
	IdentityServiceI = is
}

func (s *IdentityService) ReadFromFile(filePath string) error {
	data, err := repositories.ReadFromFile(filePath)
	if err != nil {
		return err
	}
	for i := range data {
		var person models.PersonFull
		err = json.Unmarshal([]byte(data[i][1]), &person)
		if err != nil {
			return err
		}
		s.identities[data[i][0]] = person
	}
	return nil
}

func (s *IdentityService) GetList() map[string]string {
	m := map[string]string{}
	for k, v := range s.identities {
		m[k] = v.Email
	}
	return m
}
func (s *IdentityService) NewIdentity(domain string) error {
	if _, ok := s.identities[domain]; ok {
		return errors.New("entry already exists")
	}
	newIdentity := g.NewRandomPerson()
	line := []string{domain, newIdentity.ToString()}
	err := repositories.WriteData(config.Configuration.DBFilePath, line)
	if err != nil {
		return fmt.Errorf("write Error: %s", err.Error())
	}
	s.identities[domain] = *newIdentity
	return nil
}

func (s *IdentityService) ShowIdentity(domain string, verbose bool) error {
	if p, ok := s.identities[domain]; ok {
		if verbose {
			b, _ := json.MarshalIndent(p, "", " ")
			fmt.Println(string(b))
		} else {
			b, _ := json.MarshalIndent(p.ToReduced(), "", " ")
			fmt.Println(string(b))
		}
		return nil
	}
	return errors.New("no entry found")
}

func (s *IdentityService) Delete(domain string) error {
	if _, ok := s.identities[domain]; ok {
		return repositories.RemoveLine(domain, config.Configuration.DBFilePath)
	}
	return errors.New("not found")
}

func (s *IdentityService) Renew(domain string) error {
	if val, ok := s.identities[domain]; ok {
		val.Password = g.RandomPassword()
		err := repositories.ReplaceLine([]string{domain, val.ToString()}, config.Configuration.DBFilePath)
		if err != nil {
			return err
		}
		s.identities[domain] = val
		return nil
	}
	return errors.New("not found")
}

func (s *IdentityService) Change(domain, property, data string) error {
	if val, ok := s.identities[domain]; ok {
		pm, err := val.ToMap()
		if err != nil {
			return err
		}
		pm[property] = data
		newPerson := models.PersonFull{}
		newPerson.FromMap(pm)
		err = repositories.ReplaceLine([]string{domain, newPerson.ToString()}, config.Configuration.DBFilePath)
		if err != nil {
			return err
		}
		s.identities[domain] = newPerson
		return nil
	}
	return errors.New("not found")
}
