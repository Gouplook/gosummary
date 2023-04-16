/*
 * @Author: yinjinlin yinjinlin_uplook@163.com
 * @Date: 2023-04-16 14:05:25
 * @LastEditors: yinjinlin yinjinlin_uplook@163.com
 * @LastEditTime: 2023-04-16 21:12:22
 * @FilePath: \gosummary\mqtt\exmqq\mqtt.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package exmqq

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const broker = "127.0.0.1"
const port = 1883
const topic = "t/1"
const username = ""
const password = ""

// 发布
func Pubilsh(client mqtt.Client) {

	qos := 0
	msgCount := 0
	for {
		payload := fmt.Sprintf("message: %d!", msgCount)
		pbyte, _ := json.Marshal(payload)
		if token := client.Publish(topic, byte(qos), false, pbyte); token.Wait() && token.Error() != nil {
			fmt.Printf("publish failed, topic: %s, payload: %s\n", topic, payload)
		} else {
			fmt.Printf("publish success, topic: %s, payload: %s\n", topic, payload)
		}
		msgCount++
		time.Sleep(time.Second * 1)
	}

}

// 订阅
func Subscribe(client mqtt.Client) {
	qos := 0
	client.Subscribe(topic, byte(qos), func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received %s from %s topic\n", msg.Payload(), msg.Topic())
	})
}

func CreateMqttClient() mqtt.Client {
	connectAddress := fmt.Sprintf("tcp://%s:%d", broker, port)
	client_id := fmt.Sprintf("go-client-%d", rand.Int())

	fmt.Println("connect address: ", connectAddress)
	opts := mqtt.NewClientOptions()
	opts.AddBroker(connectAddress)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetClientID(client_id)
	opts.SetKeepAlive(60)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.WaitTimeout(3*time.Second) && token.Error() != nil {
		log.Fatal(token.Error())
	}
	return client
}
