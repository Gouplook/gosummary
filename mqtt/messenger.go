package mqtt

import (
	"strings"
	"sync"
)

// Messenger 消息处理接口
type Messenger interface {
	Subscribe(topics []Topic, handler *SubscribeHandler)
	UnSubscribe(topics []Topic, handler *SubscribeHandler)
	Publish(topic Topic, payload []byte)
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

// 订阅
func (m *MessengerImpl) Subscribe(topics []Topic, handler *SubscribeHandler) {
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
	// client 实现
}
