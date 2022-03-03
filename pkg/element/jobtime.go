package element

import (
	"fmt"
	"time"
)

const (
	DayXpath = "/html/body/div/div/div[2]/main/div/div[2]/div/div[2]/table/tbody/tr[%s]/td[4]/div"
	Tmpl     = "template"
)

var (
	Today      = time.Now().Format("02")
	TodayXpath = fmt.Sprintf(DayXpath, Today)
)
