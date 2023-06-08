package mqtt

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	log "github.com/go-ozzo/ozzo-log"
	"sync"
	"time"
)

// SubscribeHandler mqtt 消息订阅拦截函数
type SubscribeHandler func(topic, payload string)

// Topic 主题结构
type Topic struct {
	Name     string
	Qos      byte
	Retained bool
}

type Client struct {
	sync.RWMutex
	subscribeHandler SubscribeHandler
	logger           *log.Logger // 引入新的日志框架
	config           *MqttConfig
	cli              MQTT.Client
	isReady          bool
	subscribedMap    map[string]bool // 订阅列表
	topicsMap        map[string]Topic
}

func NewClient(cfg *MqttConfig, args ...interface{}) *Client {
	c := &Client{
		logger:        log.NewLogger(),
		subscribedMap: make(map[string]bool),
		topicsMap:     make(map[string]Topic),
	}

	cfg.Opts.SetDefaultPublishHandler(c.DefaultMessageHanlder)
	c.config = cfg

	if len(args) > 0 {
		h, ok := args[0].(SubscribeHandler)
		if ok {
			c.SetSubscribeHandler(h)
		}
	}
	c.Run()
	return c
}

// Subscribe tells mqtt client to subscribe
func (c *Client) Subscribe(topics ...Topic) error {
	if !c.isReady {
		return fmt.Errorf("mqtt client is not ready")
	}

	c.Lock()
	defer c.Unlock()

	for _, t := range topics {
		// 1. 如果已经在订阅的列表里面，先取消订阅
		if _, ok := c.subscribedMap[t.Name]; ok {
			c.cli.Unsubscribe(t.Name)
			delete(c.subscribedMap, t.Name)
		}

		// 2. 添加到订阅列表
		if _, ok := c.topicsMap[t.Name]; !ok {
			c.topicsMap[t.Name] = t
		}

		// 订阅 topic ，并判断结果
		token := c.cli.Subscribe(t.Name, t.Qos, nil)
		if !token.WaitTimeout(3 * time.Second) {
			c.logger.Error("subscribe wait timeout")
		}
		if token.Error() != nil {
			c.logger.Error("mqtt client:[%s], sub topic:[%s] err:%v", c.config.Opts.ClientID, t.Name, token.Error())
			return token.Error()
		}
		// 已经订阅的列表中添加
		c.subscribedMap[t.Name] = true
		c.logger.Error("mqtt client:[%s], sub topics:[%s] success", c.config.Opts.ClientID, t.Name)

	}

	return nil
}
func (c *Client) Unsubscribe(topics ...Topic) {
	c.Lock()
	defer c.Unlock()
	for _, t := range topics {
		delete(c.subscribedMap, t.Name)
		delete(c.topicsMap, t.Name)
		c.cli.Unsubscribe(t.Name)
	}
}

// Publish  发布消息
func (c *Client) Publish(t Topic, p []byte) error {
	if !c.isReady {
		return fmt.Errorf("mqtt client is not ready")
	}
	token := c.cli.Publish(t.Name, t.Qos, t.Retained, p)
	if !token.WaitTimeout(3 * time.Second) {
		return fmt.Errorf("subscribe wait timeout")
	}
	return token.Error()
}

func (c *Client) Close() {
	c.cli.Disconnect(10)
}

// DefaultMessageHanlder 默认消息拦截
func (c *Client) DefaultMessageHanlder(client MQTT.Client, msg MQTT.Message) {
	topic := msg.Topic()
	payload := string(msg.Payload())
	if c.subscribeHandler != nil {
		c.subscribeHandler(topic, payload)
	}
}

// SetSubscribeHandler 设置消息订阅拦截函数
func (c *Client) SetSubscribeHandler(handler SubscribeHandler) {
	c.subscribeHandler = handler
}

// GetSubscribeHandler 获取消息订阅拦截函数
func (c *Client) GetSubscribeHandler() SubscribeHandler {
	return c.subscribeHandler
}

// OnConnectHandler 连接回调函数
func (c *Client) OnConnectHandler(client MQTT.Client) {
	c.isReady = true
	c.RLock()
	newMap := make(map[string]Topic)
	for k, v := range c.topicsMap {
		newMap[k] = v
	}
	c.RUnlock()
	// 在这里处理订阅
	for _, t := range newMap {
		_ = c.Subscribe(t)
	}
}

// ConnectionLostHandler 断开回调函数
func (c *Client) ConnectionLostHandler(client MQTT.Client, err error) {
	c.logger.Error("Mqtt client %s lost, err: %v", c.config.Opts.ClientID, err)
	c.isReady = false
	c.Lock()
	defer c.Unlock()
	c.subscribedMap = make(map[string]bool)
}
func (c *Client) Run() {
	// Set connection/reconnection/disconnection callback
	c.config.Opts.SetOnConnectHandler(c.OnConnectHandler)
	c.config.Opts.SetConnectionLostHandler(c.ConnectionLostHandler)

	// connect mqtt broker
	client := MQTT.NewClient(c.config.Opts)
	token := client.Connect()
	for token.Wait() && token.Error() != nil {
		c.logger.Error("%s cannot connect: %s", c.config.Opts.ClientID, token.Error())
		time.Sleep(c.config.Opts.MaxReconnectInterval)
		token = client.Connect()
	}
	c.isReady = true
	c.cli = client
}
