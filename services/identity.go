package services

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Kaibling/IdentityManager/config"
	g "github.com/Kaibling/IdentityManager/generator"
	"github.com/Kaibling/IdentityManager/repositories"
)

var IdentityServiceI = NewIdentityService()

type IdentityService struct {
	identities map[string]g.Person
}

func NewIdentityService() *IdentityService {
	is := &IdentityService{identities: map[string]g.Person{}}
	is.ReadFromFile(config.Configuration.DBFilePath)
	return is
}

func (s *IdentityService) ReadFromFile(filePath string) error {
	data, err := repositories.ReadFromFile(filePath)
	if err != nil {
		return err
	}
	for i := range data {
		var person g.Person
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

func (s *IdentityService) ShowIdentity(domain string) error {
	if p, ok := s.identities[domain]; ok {
		b, _ := json.MarshalIndent(p, "", " ")
		fmt.Println(string(b))
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
