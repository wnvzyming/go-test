package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"math/rand"
	"strings"
	"sync"
	"time"
)

var userList []string
var mu sync.Mutex //互斥锁，在改变共享变量值的时候，加锁，保证线程安全

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
	userList = make([]string,0)
	mu = sync.Mutex{}
	app.Run(iris.Addr(":8088"))
	
}

func (c *lotteryController) Get() string {
	count := len(userList)
	return fmt.Sprintf("当前参与抽奖用户总数： %d\n", count)
}

func (c *lotteryController) PostImport() string {
	str_user := c.Ctx.FormValue("users")
	users := strings.Split(str_user,",")

	mu.Lock()
	defer mu.Unlock()

	count1 := len(users)
	for _ , u := range users {
		u = strings.TrimSpace(u)
		if len(u) > 0 {
			userList = append(userList,u)
		}
	}
	count2 := len(userList)
	return fmt.Sprintf("当前参与抽奖用户总数： %d\n , 成功导入总人数", count1,count2)

}

func (c *lotteryController) GetLuck() string {
	mu.Lock()
	defer mu.Unlock()
	count3 := len(userList)
	if count3 > 1 {
		seed := time.Now().UnixNano()
		index := rand.New(rand.NewSource(seed)).Int31n(int32(count3))
		user := userList[index]
		userList = append(userList[0:index],userList[index+1:]...)
		return fmt.Sprintf("当前中奖用户： %s\n ", user)

	}else if count3 ==1 {
		user := userList[0:0]
		userList = []string {}
		return fmt.Sprintf("当前中奖用户： %s\n ", user )

	} else {
		return fmt.Sprintf("没有人了")
	}
}
