package main

import (
	"log"

	"work-automation/cmd/auth"
	"work-automation/pkg/element"

	"github.com/sclevine/agouti"
)

const (
	AUTH_CONFIG  = "../../config/config.yml"
	PublicKeyPem = "../../config/automationPublicKey.pem"
)

func main() {
	//init
	c := auth.Controller{}
	c.Driver = agouti.ChromeDriver(
		agouti.ChromeOptions(
			"args", []string{
				"--headless",
			}),
	)
	defer c.Driver.Stop()
	if err := c.Driver.Start(); err != nil {
		log.Fatal(err)
	}

	var err error
	c.Page, err = c.Driver.NewPage()
	if err != nil {
		log.Fatal(err)
	}

	// login
	auth.Login("https://id.jobcan.jp/users/sign_in?app_key=atd&redirect_to=https://ssl.jobcan.jp/jbcoauth/callback", c, AUTH_CONFIG, PublicKeyPem)

	attach := c.Page.FindByID(element.Attach)
	attach.Click()
}
