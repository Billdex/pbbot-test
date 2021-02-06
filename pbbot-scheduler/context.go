package pbbot_scheduler

import "github.com/ProtobufBot/go-pbbot"

const abortIndex = 64

type Context struct {
	scheduler *Scheduler
	handlers  []HandleFunc
	index     int

	RawMessage        string
	PretreatedMessage string

	shouldReply  bool
	replyMessage *pbbot.Msg
}

func (c *Context) Next() {
	c.index++
	for c.index < (len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

func (c *Context) IsAborted() bool {
	return c.index >= abortIndex
}

func (c *Context) Abort() {
	c.index = abortIndex
}

func (c *Context) Reply(msg *pbbot.Msg) {
	c.shouldReply = true
	c.replyMessage = msg
}
