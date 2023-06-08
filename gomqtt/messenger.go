package mqtt

import (
	"github.com/zeromicro/go-zero/core/threading"
	"regexp"
	"strings"
	"sync"
)

// Messenger 消息处理接口
type Messenger interface {
	Subscribe(topics []Topic, handler *SubscribeHandler) error
	Unsubscribe(topics []Topic, handler *SubscribeHandler)
	Publish(topic Topic, payload []byte) error
	Close()
	Ready() bool
	GetConfig() *MqttConfig
	GetClient() *Client
}

// MessengerImpl 消息处理实现
type MessengerImpl struct {
	sync.RWMutex
	c          *Client
	handlerMap map[string]map[*SubscribeHandler]SubscribeHandler
	cfg        *MqttConfig
}

// NewMessengerImpl 实例化
func NewMessengerImpl(conf MqttClientConfig) Messenger {
	cfg, err := NewMqttConfig(conf)
	if err != nil {
		return nil
	}

	m := &MessengerImpl{}
	err = m.new(cfg)
	if err != nil {
		return nil
	}

	return m
}

// Subscribe 订阅
func (m *MessengerImpl) Subscribe(topics []Topic, handler *SubscribeHandler) error {
	m.Lock()

	for _, topic := range topics {
		handlers, ok := m.handlerMap[topic.Name]
		if !ok {
			handlers = make(map[*SubscribeHandler]SubscribeHandler)
		}

		_, exists := handlers[handler]
		if !exists {
			handlers[handler] = *handler
		}

		if strings.HasPrefix(topic.Name, "$share/") {
			arr := strings.Split(topic.Name, "/")
			topic.Name = strings.Join(arr[2:], "/")
		}
		m.handlerMap[topic.Name] = handlers
	}

	m.Unlock()
	return m.c.Subscribe(topics...)
}

// Unsubscribe 取消订阅
func (m *MessengerImpl) Unsubscribe(topics []Topic, handler *SubscribeHandler) {
	m.Lock()

	for _, topic := range topics {
		handlers, ok := m.handlerMap[topic.Name]
		if !ok {
			continue
		}

		_, exists := handlers[handler]
		if !exists {
			continue
		}

		delete(handlers, handler)

		if len(handlers) == 0 {
			delete(m.handlerMap, topic.Name)
			m.c.Unsubscribe(topic)
		} else {
			m.handlerMap[topic.Name] = handlers
		}
	}

	m.Unlock()
}

// Publish 发布
func (m *MessengerImpl) Publish(topic Topic, payload []byte) error {
	return m.c.Publish(topic, payload)
}

// Close 关闭
func (m *MessengerImpl) Close() {
	m.Lock()

	m.handlerMap = make(map[string]map[*SubscribeHandler]SubscribeHandler)

	m.Unlock()
	m.c.Close()
}

// Ready 是否可用
func (m *MessengerImpl) Ready() bool {
	return m.c.isReady
}

// GetConfig 获取配置
func (m *MessengerImpl) GetConfig() *MqttConfig {
	return m.cfg
}

// GetClient 获取配置
func (m *MessengerImpl) GetClient() *Client {
	return m.c
}

// invokeSubscribeHandler 调用订阅处理函数
func (m *MessengerImpl) invokeSubscribeHandler(topic, payload string) {
	m.RLock()

	defer m.RUnlock()

	handlers, ok := m.handlerMap[topic]
	if ok {
		for _, handler := range handlers {
			threading.GoSafe(func() {
				handler(topic, payload)
			})
		}
		return
	}

	for subTopic, handlers := range m.handlerMap {
		pattern := strings.Replace(subTopic, "+", "\\S+", -1)
		pattern = strings.Replace(pattern, "#", ".*", -1)
		pattern = strings.Replace(pattern, "$", "\\$", -1)
		matched, _ := regexp.MatchString(pattern, topic)
		if !matched {
			continue
		}

		for _, handler := range handlers {
			threading.GoSafe(func() {
				handler(topic, payload)
			})
		}
	}
}

// new 初始化
func (m *MessengerImpl) new(cfg *MqttConfig) (err error) {
	m.cfg = cfg
	m.handlerMap = make(map[string]map[*SubscribeHandler]SubscribeHandler)
	m.c = NewClient(m.cfg, SubscribeHandler(m.invokeSubscribeHandler))
	return nil
}
