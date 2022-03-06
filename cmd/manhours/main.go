package main

import (
	"log"
	"time"

	"work-automation/cmd/auth"
	"work-automation/pkg/element"

	"github.com/sclevine/agouti"
)

const (
	AUTH_CONFIG  = "config/config.yml"
	PublicKeyPem = "config/automationPublicKey.pem"
	Mytmpl01     = "テスト"
)

var err error

func main() {
	c := auth.Controller{}
	c.Driver = agouti.ChromeDriver()
	if err := c.Driver.Start(); err != nil {
		log.Fatal(err)
	}
	c.Page, err = c.Driver.NewPage()
	if err != nil {
		log.Fatal(err)
	}

	auth.Login("https://id.jobcan.jp/users/sign_in?app_key=atd&redirect_to=https://ssl.jobcan.jp/jbcoauth/callback", c, AUTH_CONFIG, PublicKeyPem)

	c.Page.Navigate("https://ssl.jobcan.jp/employee/man-hour-manage")

	edit := c.Page.FindByXPath(element.TodayXpath)
	edit.Click()

	time.Sleep(time.Second * 1)

	tmpl := c.Page.FindByName(element.Tmpl)
	tmpl.Select(Mytmpl01)
}
