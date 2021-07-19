package dingtalk

import (
	"github.com/dtapps/go-dingtalk/dingtalk"
	"github.com/dtapps/go-dingtalk/dingtalk/message"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	bot := dingtalk.DingBot{
		Secret:      "SEC05342ba24a7eb2e1dbeae61b3df997eb1a97b7cda414566876e983f1db0fec0b",
		AccessToken: "aad81de7f6b218bb7d085264d4885714c805cc80a460690a0d19db91a05dd174",
	}
	msg := message.Message{
		MsgType: message.TextStr,
		Text: message.Text_{
			Content: "测试",
		},
	}
	send, err := bot.Send(msg)
	if err != nil {
		log.Printf("err：%v\n", err)
		return
	}
	log.Printf("send：%v\n", send)
}
