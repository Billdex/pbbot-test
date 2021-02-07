package main

import (
	"fmt"
	pbbot "github.com/ProtobufBot/go-pbbot"
	"github.com/ProtobufBot/go-pbbot/proto_gen/onebot"
	log "github.com/sirupsen/logrus"
	"net/http"
	pbbot_scheduler "pbbot-test/pbbot-scheduler"
	"strconv"
	"time"
)

func main() {
	scheduler := pbbot_scheduler.New()

	group1 := scheduler.Group("*")
	{
		group1.Use(logTime)
		group1.Bind("admin", pbbot_scheduler.MustGroupAdmin, echoAdmin).Alias("staff", "administrator").IgnoreCase()
	}
	group2 := group1.Group("cal").Alias("math").IgnoreCase()
	{
		group2.Use(limitNum(20))
		group2.Bind("sum", calculateSum)
		group2.Bind("fib", calculateFib).IgnoreCase()
	}

	pbbot.HandleGroupMessage = func(bot *pbbot.Bot, event *onebot.GroupMessageEvent) {
		err := scheduler.Process(bot, event)
		if err != nil {
			log.Error(err)
		}
	}

	pbbot.HandlePrivateMessage = func(bot *pbbot.Bot, event *onebot.PrivateMessageEvent) {
		err := scheduler.Process(bot, event)
		if err != nil {
			log.Error(err)
		}
	}
	http.HandleFunc("/ws/test/", func(w http.ResponseWriter, r *http.Request) {
		pbbot.UpgradeWebsocket(w, r)
	})
	http.ListenAndServe(":8081", nil)
}

func logTime(c *pbbot_scheduler.Context) {
	start := time.Now()
	c.Next()
	end := time.Now()
	strStart := start.Format("2006-01-02 15:04:05")
	totalTime := (float64(end.UnixNano() - start.UnixNano())) / 1e6
	log.Info(fmt.Sprintf("[%s]请求内容:%s, 处理请求耗时:%.2fms\n", strStart, c.GetRawMessage(), totalTime))
}

func limitNum(n int) pbbot_scheduler.HandleFunc {
	return func(c *pbbot_scheduler.Context) {
		content := c.PretreatedMessage
		num, err := strconv.Atoi(content)
		if err != nil {
			log.Error(fmt.Sprintf("'%s' not a num", c.PretreatedMessage))
			c.Abort()
			return
		}
		if num > n {
			msg := pbbot.NewMsg().Text("数字过大")
			c.Reply(msg, false)
			c.Abort()
		} else if num < 1 {
			msg := pbbot.NewMsg().Text("必须取正整数")
			c.Reply(msg, false)
			c.Abort()
		}
	}
}

func echoAdmin(c *pbbot_scheduler.Context) {
	msg := pbbot.NewMsg().Text("is admin")
	_, err := c.Reply(msg, false)
	if err != nil {
		log.Error("echoAdmin Fail!")
	}
}

// 为了使logTime输出的时间更为明显，采用了运算效率较低的实现（然鹅主要耗时还是在发送消息上）
func calculateSum(c *pbbot_scheduler.Context) {
	num, err := strconv.Atoi(c.PretreatedMessage)
	if err != nil {
		msg := pbbot.NewMsg().Text("参数不是数字")
		_, err = c.Reply(msg, false)
		if err != nil {
			log.Error("reply failed:", err)
		}
		return
	}
	result := make([]int, 0)
	for i := 1; i <= num; i++ {
		result = append(result, sum(i))
	}
	msg := pbbot.NewMsg().Text(fmt.Sprintf("%v", result))
	_, err = c.Reply(msg, false)
	if err != nil {
		log.Error("reply failed:", err)
	}
}

func sum(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}
	return result
}

func calculateFib(c *pbbot_scheduler.Context) {
	num, err := strconv.Atoi(c.PretreatedMessage)
	if err != nil {
		msg := pbbot.NewMsg().Text("参数不是数字")
		_, err = c.Reply(msg, false)
		if err != nil {
			log.Error("reply failed:", err)
		}
		return
	}
	result := make([]int, 0)
	for i := 1; i <= num; i++ {
		result = append(result, fib(i))
	}
	msg := pbbot.NewMsg().Text(fmt.Sprintf("%v", result))
	_, err = c.Reply(msg, false)
	if err != nil {
		log.Error("reply failed:", err)
	}
}

func fib(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
