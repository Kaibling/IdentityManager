package sqlite

import (
	"strings"

	"github.com/Kaibling/IdentityManager/models"
	"gorm.io/gorm"
)

type person struct {
	ID                uint `gorm:"primaryKey"`
	FirstName         string
	LastName          string
	Domain            string // TODO not null
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
	SecurityQuestions string
}

type SQLiteRepo struct {
	db *gorm.DB
}

func NewSQLiteRepo(db *gorm.DB) (*SQLiteRepo, error) {
	return &SQLiteRepo{db: db}, nil
}

func (s *SQLiteRepo) ReadAll() ([]models.Person, error) {
	var persons []person
	result := s.db.Model(&person{}).Find(&persons)
	if result.Error != nil {
		return nil, result.Error
	}
	return personArrayUnmarshal(persons), nil
}
func (s *SQLiteRepo) Create(p models.Person) error {
	person := personMarshal(p)
	result := s.db.Create(&person)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (s *SQLiteRepo) Delete(domain string) error {
	result := s.db.Where("domain = ?", domain).Delete(&person{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (s *SQLiteRepo) Update(p models.Person) error {
	person := personMarshal(p)
	result := s.db.Where("domain = ?", p.Domain).Updates(&person)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func personMarshal(p models.Person) person {
	sq := strings.Join(p.SecurityQuestions, ";")
	return person{
		FirstName:         p.FirstName,
		LastName:          p.LastName,
		Domain:            p.Domain,
		Email:             p.Email,
		Username:          p.Username,
		Password:          p.Password,
		BirthDate:         p.BirthDate,
		Gender:            p.Gender,
		Street:            p.Street,
		Country:           p.Country,
		PhoneNumber:       p.PhoneNumber,
		CreditCard:        p.CreditCard,
		HomePage:          p.HomePage,
		Company:           p.Company,
		Job:               p.Job,
		SecurityQuestions: sq,
	}
}

func personUnmarshal(p person) models.Person {
	sq := strings.Split(p.SecurityQuestions, ";")
	return models.Person{
		FirstName:         p.FirstName,
		LastName:          p.LastName,
		Domain:            p.Domain,
		Email:             p.Email,
		Username:          p.Username,
		Password:          p.Password,
		BirthDate:         p.BirthDate,
		Gender:            p.Gender,
		Street:            p.Street,
		Country:           p.Country,
		PhoneNumber:       p.PhoneNumber,
		CreditCard:        p.CreditCard,
		HomePage:          p.HomePage,
		Company:           p.Company,
		Job:               p.Job,
		SecurityQuestions: sq,
	}
}

func personArrayUnmarshal(p []person) []models.Person {
	personArray := []models.Person{}
	for i := range p {
		mp := personUnmarshal(p[i])
		personArray = append(personArray, mp)
	}
	return personArray
}

type sqliteDBMigrator struct {
	db *gorm.DB
}

func (s sqliteDBMigrator) Migrate() error {
	err := s.db.AutoMigrate(&person{})
	if err != nil {
		return err
	}
	return nil
}

func NewSQLiteMigrator(db *gorm.DB) *sqliteDBMigrator {
	return &sqliteDBMigrator{db: db}
}
