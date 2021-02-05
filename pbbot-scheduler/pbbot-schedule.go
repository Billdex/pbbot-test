package pbbot_scheduler

import (
	"errors"
	"github.com/ProtobufBot/go-pbbot"
	"github.com/ProtobufBot/go-pbbot/proto_gen/onebot"
)

type Scheduler struct {
}

type Context struct {
	scheduler *Scheduler
}

func New() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) Use(handleFunc ...HandleFunc) {

}

func (s *Scheduler) Bind(handler HandleFunc, keywords ...string) {

}

func (s *Scheduler) Process(bot *pbbot.Bot, event interface{}) error {
	if privateEvent, ok := event.(*onebot.PrivateMessageEvent); ok {
		// TODO
		s.MessageHandler(privateEvent.RawMessage)
		if c.shouldReply {

		}
		return nil
	} else if groupEvent, ok := event.(*onebot.GroupMessageEvent); ok {
		// TODO

		return nil
	} else {
		return errors.New("event类型错误!必须为*onebot.PrivateMessageEvent或*onebot.GroupMessageEvent")
	}
}

type HandleFunc func()
