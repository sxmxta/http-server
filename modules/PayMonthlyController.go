package modules

import (
	"fmt"
	"gorm.io/gorm"
	"http_server/common"
	"http_server/database"
	"http_server/server"
)

type PayMonthly struct {
	gorm.Model
	QQ         string `json:"qq"`
	RedeemCode string `json:"redeem_code"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	Money      string `json:"money"`
}

func init() {
	database.GetDB().AutoMigrate(&PayMonthly{})
}

func list(ctx *server.Context) {
	fmt.Println(ctx.Get("name"))
	bean := &[]PayMonthly{}
	database.GetDB().Find(bean)
	ctx.Write(bean)
}

func add(ctx *server.Context) {
	qq := ctx.Get("qq")
	date := ctx.Get("date")
	money := ctx.Get("money")
	fmt.Println(qq, date)
	bean := &PayMonthly{}
	database.GetDB().Find(bean, "qq = ?", qq)
	bean.QQ = qq
	bean.RedeemCode = common.RandomString(6)
	bean.StartDate = date
	bean.EndDate = date
	bean.Money = money
	if bean.ID > 0 {
		database.GetDB().Model(bean).Updates(bean)
	} else {
		database.GetDB().Create(bean)
	}
	ctx.Write("ok")
}

func delete(ctx *server.Context) {
	id := ctx.Get("id")
	bean := &PayMonthly{}
	database.GetDB().Delete(bean, id)
}
