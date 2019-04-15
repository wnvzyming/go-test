package mian

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"math/rand"
	"time"
)


type lotteryController struct {
	Ctx iris.Context
}

func newApp() *iris.Application  {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	return app

}

func main()  {
	app := newApp()
	app.Run(iris.Addr(":8088"))

}

func (c *lotteryController) Get() string {
	var str string
	send := time.Now().UnixNano()
	prizeCode := rand.New(rand.NewSource(send)).Intn(10)
	switch {
	case prizeCode == 1:
		str = "一等奖"
	case prizeCode == 2:
		str = "二等奖"
	case prizeCode == 3:
		str = "三等奖"
	default:
		str = "未中奖"
	}
	return fmt.Sprint(str)
}

func (c *lotteryController) GetPrize()  string{

	//prizeArr := []int{}

	var prizeArr [7]int
	send := time.Now().UnixNano()
	code := rand.New(rand.NewSource(send))
	for i:=0; i<6 ;i++  {
		prizeArr[i] = code.Intn(33)+1
	}
	//sort.Ints(prizeArr)
	prizeArr[6] = code.Intn(16)+1

	return fmt.Sprint("双色球中奖号码：%v",prizeArr)

}