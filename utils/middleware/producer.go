package middleware

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

// NSQ Producer Demo

var producer *nsq.Producer

// 初始化生产者
func initProducer(str string) (err error) {
	config := nsq.NewConfig()
	producer, err = nsq.NewProducer(str, config)
	if err != nil {
		fmt.Printf("create producer failed, err:%v\n", err)
		return err
	}
	return nil
}

func Producer(topic, content string) {
	nsqAddress := ""
	err := initProducer(nsqAddress)
	if err != nil {
		fmt.Printf("init producer failed, err:%v\n", err)
		return
	}
	err = producer.Publish(topic, []byte(content))
	if err != nil {
		fmt.Printf("publish msg to nsq failed, err:%v\n", err)
	}
}
