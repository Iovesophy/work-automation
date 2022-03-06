package main

import (
	"io/fs"
	"io/ioutil"
	"log"

	"work-automation/cmd/auth"
	"work-automation/pkg/element"

	"github.com/sclevine/agouti"
)

const (
	AUTH_CONFIG  = "config/config.yml"
	PublicKeyPem = "config/automationPublicKey.pem"
	STATUS       = "config/touch.log"
)

var perm fs.FileMode = 600

func main() {
	status, err := ioutil.ReadFile(STATUS)
	if err != nil {
		log.Fatal(err)
	}
	if string(status) == "1" {
		c := auth.Controller{}
		c.Driver = agouti.ChromeDriver(agouti.ChromeOptions("args", []string{"--headless"}))
		defer c.Driver.Stop()
		if err := c.Driver.Start(); err != nil {
			log.Fatal(err)
		}
		c.Page, err = c.Driver.NewPage()
		if err != nil {
			log.Fatal(err)
		}

		auth.Login("https://id.jobcan.jp/users/sign_in?app_key=atd&redirect_to=https://ssl.jobcan.jp/jbcoauth/callback", c, AUTH_CONFIG, PublicKeyPem)

		attach := c.Page.FindByID(element.Attach)
		attach.Click()

		if err := ioutil.WriteFile(STATUS, []byte("0"), perm); err != nil {
			log.Fatal(err)
		}
	}
}
