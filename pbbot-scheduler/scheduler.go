package pbbot_scheduler

import (
	"errors"
	"github.com/ProtobufBot/go-pbbot"
	"github.com/ProtobufBot/go-pbbot/proto_gen/onebot"
)

type Scheduler struct {
	*CmdGroup
}

type HandleFunc func(*Context)

func New() *Scheduler {
	scheduler := &Scheduler{
		CmdGroup: &CmdGroup{
			Name:         []string{""},
			Handlers:     []HandleFunc{},
			subCmdGroups: make([]*CmdGroup, 0),
		},
	}
	scheduler.CmdGroup.scheduler = scheduler
	return scheduler
}

func (s *Scheduler) createContext() *Context {
	return &Context{
		scheduler:   s,
		handlers:    make([]HandleFunc, 0),
		index:       0,
		shouldReply: false,
	}
}

func (s *Scheduler) Process(bot *pbbot.Bot, event interface{}) error {
	var rawMessage string
	if privateEvent, ok := event.(*onebot.PrivateMessageEvent); ok {
		rawMessage = privateEvent.RawMessage
	} else if groupEvent, ok := event.(*onebot.GroupMessageEvent); ok {
		rawMessage = groupEvent.RawMessage
	} else {
		return errors.New("event类型错误!必须为*onebot.PrivateMessageEvent或*onebot.GroupMessageEvent")
	}
	c := s.createContext()
	c.RawMessage = rawMessage
	handlerChain, content := s.findHandler(rawMessage)
	c.handlers = handlerChain
	c.PretreatedMessage = content
	for c.index < len(handlerChain) {
		handlerChain[c.index](c)
	}

	if c.ShouldReply {

	}

	return nil
}

func (s *Scheduler) findHandler(message string) ([]HandleFunc, string) {
	return s.CmdGroup.SearchHandlerChain(message)
}
