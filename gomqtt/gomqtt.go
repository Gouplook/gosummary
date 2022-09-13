package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	gomqtt "github.com/eclipse/paho.mqtt.golang"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	MqttClient *Client
)

func Init(addr string) {
	optsTracking := GetOptions(addr)
	MqttClient = new(Client)
	err := MqttClient.Init(optsTracking)
	if err != nil {
		fmt.Println(fmt.Sprintf("mqtt init error with %s", err.Error()))
		panic(fmt.Sprintf("mqtt init error with %s", err.Error()))
	}
	fmt.Println("mqtt start successfully...")
}

type SubscribeType struct {
	Topic      string
	Qos        byte
	Callback   gomqtt.MessageHandler
	RetryTimes int // 为0表示无限重试
}

// MqttClientId 生成clientId
func MqttClientId() string {
	uuidStr := strings.Join(strings.Split(uuid.NewV4().String(), "-"), "")
	return fmt.Sprintf("traffic-test|%s", uuidStr)
}

// GetOptions 获取MQTT连接配置项
func GetOptions(addr string) *gomqtt.ClientOptions {
	opts := gomqtt.NewClientOptions()
	opts.SetAutoReconnect(true)
	opts.AddBroker(fmt.Sprintf("tcp://%s", addr))
	opts.SetUsername("compass")
	opts.SetPassword("compass")
	opts.SetKeepAlive(60 * time.Second)
	opts.SetPingTimeout(5 * time.Second)
	opts.SetOnConnectHandler(func(c gomqtt.Client) {})
	opts.SetConnectionLostHandler(func(c gomqtt.Client, e error) {})
	opts.SetTLSConfig(&tls.Config{
		ClientAuth:         tls.NoClientCert,
		ClientCAs:          nil,
		InsecureSkipVerify: true,
	})
	opts.SetClientID(MqttClientId())
	return opts
}

// NewClient 获取MQTT连接
func NewClient(opts *gomqtt.ClientOptions) (client gomqtt.Client, err error) {
	client = gomqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		err = token.Error()
	}
	return
}

// Client MQTT连接
type Client struct {
	client      gomqtt.Client   // 实际连接
	subscribers []SubscribeType // 订阅监听器
}

// Init 初始化监听器
func (m *Client) Init(opts *gomqtt.ClientOptions) (err error) {
	opts.SetOnConnectHandler(m.onConnectHandler(opts.OnConnect))
	opts.SetConnectionLostHandler(m.onConnectionLostHandler(opts.OnConnectionLost))
	m.client, err = NewClient(opts)
	return
}

// Close 关闭监视器
func (m *Client) Close() {
	m.client.Disconnect(0)
}

// 连接上服务器的操作(启动所有订阅)
func (m *Client) onConnectHandler(handler gomqtt.OnConnectHandler) gomqtt.OnConnectHandler {
	return func(c gomqtt.Client) {
		for _, item := range m.subscribers {
			m.subscribe(item)
		}
		handler(c)
	}
}

// 丢失连接的操作
func (m *Client) onConnectionLostHandler(handler gomqtt.ConnectionLostHandler) gomqtt.ConnectionLostHandler {
	return func(c gomqtt.Client, e error) {
		handler(c, e)
	}
}

// 通用发布消息接口
func (m *Client) publish(topic string, payload interface{}, qos byte, retained bool) (err error) {
	token := m.client.Publish(topic, qos, retained, payload)
	if token.Wait() && token.Error() != nil {
		err = token.Error()
	}
	return
}

func (m *Client) subscribeItem(item SubscribeType) (token gomqtt.Token, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	token = m.client.Subscribe(item.Topic, item.Qos, item.Callback)
	return
}

func (m *Client) subscribe(item SubscribeType) {
	times := 0
	for {
		token, err := m.subscribeItem(item)
		if err != nil {
			if item.RetryTimes == 0 || times < item.RetryTimes {
				times++
				time.Sleep(3 * time.Second)
				continue
			} else {
				panic(err)
			}
		}
		if token.Wait() && token.Error() != nil {
			if item.RetryTimes == 0 || times < item.RetryTimes {
				times++
				time.Sleep(3 * time.Second)
				continue
			} else {
				panic(token.Error())
			}
		}
		break
	}
}

func DataPub(payload string) {
	topic := "/efence_k8s_backend/tracking"
	if err := MqttClient.publish(topic, payload, 1, false); err != nil {
		fmt.Printf("local_path_publish error: %v\n", err.Error())
		return
	}
}

var tracking = `{
   "channel": "tracking",
   "params": {
      "loc": {
         "state": 3,
         "lat": %v,
         "lon": %v
      },
      "theta": %v
   },
   "vin": "%v",
   "timestamp": %v,
   "name": "%v"
}`

func fileReader(path string) {
	file, err := os.OpenFile(path, 0, 0777)
	if err != nil {
		panic(err.Error())
	}
	datas, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err.Error())
	}
	lanes := strings.Split(string(datas), "\n")
	fmt.Println(lanes)
	re := regexp.MustCompile(`carId=([^,]+),xl=([^,]+),yl=([^,]+),timesstmap\s*=([^,]+)`)
	for _, line := range lanes {
		if strings.TrimSpace(line) == "" {
			//fmt.Println("read space line from file.")
			continue
		}
		fields := re.FindStringSubmatch(line)

		if len(fields) < 5 {
			fmt.Println(fields)
			fmt.Println("read space field from line.")
			continue
		}
		carId := strings.TrimSpace(fields[1])
		x, err := strconv.ParseFloat(strings.TrimSpace(fields[2]), 64)
		if err != nil {
			fmt.Println("field x with unexpected type")
			continue
		}
		y, err := strconv.ParseFloat(strings.TrimSpace(fields[3]), 64)
		if err != nil {
			fmt.Println("field y with unexpected type")
			continue
		}
		// 时间根据实时发送生成。
		// 问题， 时间
		timestamp, err := strconv.ParseFloat(strings.TrimSpace(fields[4]), 64)
		if err != nil {
			fmt.Println("field time with unexpected type")
			continue
		}
		fmt.Println(carId, x, y, timestamp)
		payload := fmt.Sprintf(tracking,
			// lat, lon, theta, vin, time, vin
			y, x, 3.88, carId, timestamp/1e3, carId)
		DataPub(payload)
		fmt.Printf("payload:%s", payload)
	}
}

var (
	filePath = flag.String("f", "./carId.yml", "the file path")
	mqttAddr = flag.String("a", "10.8.0.60:1883", "the address of mqtt")
)

func main() {
	// todo: do
	flag.Parse()
	Init(*mqttAddr)
	fileReader(*filePath)
}
