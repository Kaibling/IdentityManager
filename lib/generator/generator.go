package generator

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/Kaibling/IdentityManager/lib/config"
	"github.com/Kaibling/IdentityManager/models"
	"github.com/jaswdr/faker"
)

var faaker = faker.New()
var passwordlength = 40

func NewRandomPerson() *models.PersonFull {

	gender := faaker.Person().Gender()
	var firstName string
	p := faaker.Person()
	if gender == "Male" {
		firstName = p.FirstNameMale()
	} else {
		firstName = p.FirstNameFemale()
	}
	lastName := faaker.Person().LastName()
	address := faaker.Address()
	street := strings.ReplaceAll(address.Address(), "\n", ",")
	country := "United States of America"
	birthDate := randomDate()
	homePage := "https://" + faaker.Internet().Domain()
	company := faaker.Company().Name()
	job := faaker.Company().JobTitle()
	phoneNumber := faaker.Phone().Number()
	creditCard := faaker.Payment().CreditCardNumber()
	return &models.PersonFull{
		FirstName:         firstName,
		LastName:          lastName,
		Email:             strings.ToLower(fmt.Sprintf("%s.%s%s", firstName, lastName, config.Configuration.Email)),
		Username:          generateUserName(firstName, lastName, birthDate),
		BirthDate:         birthDate,
		Street:            street,
		Country:           country,
		Password:          RandomPassword(),
		HomePage:          homePage,
		Gender:            gender,
		Company:           company,
		Job:               job,
		SecurityQuestions: generateSecurityQuestions(),
		PhoneNumber:       phoneNumber,
		CreditCard:        creditCard,
	}
}

func generateSecurityQuestions() []string {
	return []string{
		"First Pet: " + faaker.Pet().Name(),
		"Mother Maiden name: " + faaker.Person().LastName(),
		"Favorite Color: " + faaker.Color().SafeColorName(),
		"Favorite Beer: " + faaker.Beer().Name(),
		"Favorite Car: " + faaker.Car().Maker(),
	}
}

func generateUserName(firstname, lastname, BirthDate string) string {
	var name string
	rand.Seed(time.Now().UnixNano())
	c := rand.Intn(3) + 1
	switch c {
	case 1:
		name = strings.ToLower(fmt.Sprintf("%s.%s", firstname[:1], lastname))
	case 2:
		name = strings.ToLower(fmt.Sprintf("%s.%s%s", firstname[:1], lastname, BirthDate[len(BirthDate)-2:]))
	default:
		color := faaker.Color().ColorName()
		thing := ""
		if rand.Intn(1) == 1 {
			thing = faaker.Food().Fruit()
		} else {
			thing = faaker.Food().Vegetable()
		}
		if rand.Intn(1) == 1 {
			name = strings.ToLower(fmt.Sprintf("%s%s", color, thing))
		} else {
			name = strings.ToLower(fmt.Sprintf("%s%s%s", color, thing, BirthDate[len(BirthDate)-2:]))
		}
	}
	return strings.ReplaceAll(name, " ", "")
}

func randomDate() string {
	min := time.Date(1960, 1, 2, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Now().AddDate(-21, 1, 0).Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	t := time.Unix(sec, 0)
	return t.Format("02-Jan-2006")
}

func RandomPassword() string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	specials := "~=+%^*/()[]{}/!@#$?|"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		digits + specials
	length := passwordlength
	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = specials[rand.Intn(len(specials))]
	for i := 2; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	return string(buf)
}
