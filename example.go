package dingtalk

import (
	"github.com/dtapps/go-dingtalk/dingtalk"
	"github.com/dtapps/go-dingtalk/dingtalk/message"
	"log"
	"testing"
)

func main(t *testing.T) {
	bot := dingtalk.DingBot{
		Secret:      "",
		AccessToken: "",
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
