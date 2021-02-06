package main

import (
	pbbot "github.com/ProtobufBot/go-pbbot"
	"github.com/ProtobufBot/go-pbbot/proto_gen/onebot"
	"net/http"
	pbbot_scheduler "pbbot-test/pbbot-scheduler"
)

func main() {
	scheduler := pbbot_scheduler.New()

	middleWare := func() {}
	handler := func() {}

	scheduler.Use(middleWare)
	scheduler.Bind(handler, "keyword1", "keyword2")
	scheduler.Group()

	pbbot.HandleGroupMessage = func(bot *pbbot.Bot, event *onebot.GroupMessageEvent) {
		err := scheduler.Process(bot, event)
		if err != nil {
			//TODO

		}
	}

	pbbot.HandlePrivateMessage = func(bot *pbbot.Bot, event *onebot.PrivateMessageEvent) {
		err := scheduler.Process(bot, event)
		if err != nil {
			//TODO

		}
	}
	http.HandleFunc("/ws/test/", func(w http.ResponseWriter, r *http.Request) {
		pbbot.UpgradeWebsocket(w, r)
	})
	http.ListenAndServe(":8081", nil)

}
