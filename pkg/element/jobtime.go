package element

import (
	"fmt"
	"time"
)

const (
	EditorTitle = "edit-menu-title"
	DayXpath    = "/html/body/div/div/div[2]/main/div/div[2]/div/div[2]/table/tbody/tr[%s]/td[4]/div"
	Tmpl        = "template"
	Mytmpl01    = "テスト"

	Input01 = "/html/body/div[1]/div/div[2]/main/div/div[1]/div/div[2]/form/div[2]/div[3]/table/tbody/tr[2]/td[4]/input[1]"
	Input02 = "/html/body/div[1]/div/div[2]/main/div/div[1]/div/div[2]/form/div[2]/div[3]/table/tbody/tr[3]/td[4]/input[1]"
	Input03 = "/html/body/div[1]/div/div[2]/main/div/div[1]/div/div[2]/form/div[2]/div[3]/table/tbody/tr[4]/td[4]/input[1]"

	Save = "save"
)

var (
	Today      = time.Now().Format("02")
	TodayXpath = fmt.Sprintf(DayXpath, Today)
)
