package auth

import (
	"io/ioutil"
	"log"
	"work-automation/pkg/auth"
	"work-automation/pkg/element"

	"github.com/sclevine/agouti"
	"gopkg.in/yaml.v2"
)

type Controller struct {
	Driver *agouti.WebDriver
	Page   *agouti.Page
}

type User struct {
	Email string
}

func Login(url string, c Controller, authconfig string, publickeypem string) {
	buf, err := ioutil.ReadFile(authconfig)
	if err != nil {
		log.Fatal(err)
	}

	var user User
	err = yaml.Unmarshal([]byte(buf), &user)
	if err != nil {
		log.Fatal(err)
	}

	// start automation
	c.Page.Navigate(url)
	email := c.Page.FindByID(element.LoginUser)
	password := c.Page.FindByID(element.LoginPassword)

	email.SendKeys(user.Email)
	password.SendKeys(string(auth.Decoder(publickeypem)))

	submit := c.Page.FindByName(element.LoginSubmit)
	submit.Click()
}
