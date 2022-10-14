package mqtt

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/zeromicro/go-zero/core/threading"
)

// mqtt 消息订阅拦截函数
type SubscribeHandler func(topic, payload string)

// Topic 主题结构
type Topic struct {
	Name     string
	Qos      byte
	Retained bool
}

// subscribe 订阅结构
type subscribe struct {
	topics []Topic
	add    bool
}

// 发布结构
type publish struct {
	topic   Topic
	payload []byte
}

type Client struct {
	config           *MqttConfig
	subscriber       chan subscribe
	publisher        chan publish
	connection       chan bool
	quit             chan struct{}
	subScribeHandler SubscribeHandler
	isReady          bool
}

func (c *Client) New(cfg *MqttConfig) {
	c.config = cfg
	c.subscriber = make(chan subscribe)
	c.publisher = make(chan publish)
	c.connection = make(chan bool)
	c.quit = make(chan struct{})
}

func NewClient(cfg *MqttConfig, args ...interface{}) *Client {
	c := new(Client)
	cfg.Opts.SetDefaultPublishHandler(c.DefaultMessageHanlder)
	c.New(cfg)

	if len(args) > 0 {
		h, ok := args[0].(SubscribeHandler)
		if ok {
			c.SetSubscribeHandler(h)
		}
	}
	// Run
	threading.GoSafe(c.Run)

	return c
}

//  默认消息拦截
func (c *Client) DefaultMessageHanlder(client MQTT.Client, msg MQTT.Message) {

	topic := msg.Topic()
	payload := string(msg.Payload())
	if c.subScribeHandler != nil {
		c.subScribeHandler(topic, payload)
	}

}

//设置消息订阅拦截函数
func (c *Client) SetSubscribeHandler(handler SubscribeHandler) {
	c.subScribeHandler = handler
}

//获取消息订阅拦截函数
func (c *Client) GetSubscribeHandler() SubscribeHandler {
	return c.subScribeHandler
}

// 是否准备好了，run起来了。
func (c *Client) Ready() bool {
	return c.isReady
}

func (c *Client) Run() {
	// 1. 配置config

	// 2. 初始化NewClient
	client := MQTT.NewClient(c.config.Opts)

	// 3. connect

	// 4. 监听
	var connected bool = false
	for {
		select {
		case connected = <-c.connection:
			fmt.Println()
		case sub := <-c.subscriber:
			fmt.Println(sub)
		case pub := <-c.publisher:
			fmt.Println(pub)
		case <-c.quit:
			client.Disconnect(10)
			return
		}
	}

}
