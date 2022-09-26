package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	gomqtt "github.com/eclipse/paho.mqtt.golang"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	MqttClient *Client
	tracking   = `{
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

type CarData struct {
	CarId string
	X     float64
	Y     float64
	T     float64
}

func fileReader(path string) (map[string][]CarData, float64) {
	file, err := os.OpenFile(path, 0, 0777)

	if err != nil {
		panic(err.Error())
	}
	datas, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err.Error())
	}
	lanes := strings.Split(string(datas), "\n")
	re := regexp.MustCompile(`carId=([^,]+),xl=([^,]+),yl=([^,]+),timesstmap\s*=([^,]+)`)
	res := map[string][]CarData{}

	minTime := math.MaxFloat64
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
			fmt.Println()
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
		timestamp, err := strconv.ParseFloat(strings.TrimSpace(fields[4]), 64)
		if err != nil {
			fmt.Println("field time with unexpected type")
			continue
		}
		fmt.Println(carId, x, y, timestamp)
		//payload := fmt.Sprintf(tracking,
		//	// lat, lon, theta, vin, time, vin
		//	y, x, 3.88, carId, timestamp/1e3, carId)
		//DataPub(payload)
		//fmt.Printf("payload:%s", payload)
		if _, ok := res[carId]; ok {
			res[carId] = append(res[carId], CarData{
				CarId: carId, X: x, Y: y, T: timestamp,
			})
		} else {
			res[carId] = make([]CarData, 0, 0)
			res[carId] = append(res[carId], CarData{
				CarId: carId, X: x, Y: y, T: timestamp,
			})
		}
		if minTime > timestamp {
			minTime = timestamp
		}
	}
	return res, minTime
}

// GetTimestampFloat 获取浮点型时间戳
func GetTimestampFloat() float64 {
	return ConvertTime2float64(time.Now())
}

// ConvertTime2float64 将时间类型转浮点型时间戳
func ConvertTime2float64(t time.Time) float64 {
	return float64(t.UnixNano()) / 1e9
}

var (
	filePath = flag.String("f", ".\\carId.yaml", "the file path")
	mqttAddr = flag.String("a", "10.8.0.60:1883", "the address of mqtt")
	interval = flag.Int64("i", 500, "the interval of mqtt data")
)

func pubRoutine(datas []CarData, delay, interval int64, sig, over chan struct{}) {
	<-sig
	time.Sleep(time.Duration(delay) * time.Millisecond)
	tt := time.NewTicker(time.Duration(interval) * time.Millisecond)
	fmt.Printf("start with delay(%v) and interval(%v)\n", delay, interval)
	for _, cur := range datas {
		fmt.Printf("%v send one data to mqtt\n", cur.CarId)
		select {
		case <-tt.C:
			// cur := datas[0]
			payload := fmt.Sprintf(tracking,
				// lat, lon, theta, vin, time, vin
				cur.Y, cur.X, 3.88, cur.CarId, GetTimestampFloat(), cur.CarId)
			DataPub(payload)
		}
	}
	over <- struct{}{}
}

func main() {
	// todo: do
	flag.Parse()
	Init(*mqttAddr)
	datas, mint := fileReader("D:\\gosummaryCode\\gosummary\\gomqtt\\carId.yaml")
	sigs := make([]chan struct{}, 0)
	overs := make([]chan struct{}, 0)
	for _, v := range datas {
		single := make(chan struct{})
		over := make(chan struct{})
		// 排序
		sort.Slice(v, func(i, j int) bool {
			return v[i].T < v[j].T
		})
		go pubRoutine(v, int64(v[0].T-mint), *interval, single, over)
		sigs = append(sigs, single)
		overs = append(overs, over)
	}
	for _, s := range sigs {
		s <- struct{}{}
	}
	//time.Sleep(1 * time.Second)
	for _, o := range overs {
		<-o
	}
}
