package services

import (
	"encoding/json"
	"errors"
	"fmt"

	g "github.com/Kaibling/IdentityManager/lib/generator"
	"github.com/Kaibling/IdentityManager/models"
)

type identityRepo interface {
	ReadAll() ([]models.PersonFull, error)
	Create(p models.PersonFull) error
	Delete(domain string) error
	Update(p models.PersonFull) error
}

var IdentityServiceI *IdentityService = nil

type IdentityService struct {
	identities map[string]models.PersonFull
	repo       identityRepo
}

func InitIdentityService(repo identityRepo) error {
	is := &IdentityService{identities: map[string]models.PersonFull{}, repo: repo}
	p, err := is.repo.ReadAll()
	if err != nil {
		return err
	}
	for i := range p {
		is.identities[p[i].Domain] = p[i]
	}
	IdentityServiceI = is
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
	newIdentity.Domain = domain
	err := s.repo.Create(*newIdentity)
	if err != nil {
		return err
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
		return s.repo.Delete(domain)
	}
	return errors.New("not found")
}

func (s *IdentityService) Renew(domain string) error {
	if val, ok := s.identities[domain]; ok {
		val.Password = g.RandomPassword()
		err := s.repo.Update(val)
		//err := csv.ReplaceLine([]string{domain, val.ToString()}, config.Configuration.DBFilePath)
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
		err = s.repo.Update(newPerson)
		if err != nil {
			return err
		}
		s.identities[domain] = newPerson
		return nil
	}
	return errors.New("not found")
}
