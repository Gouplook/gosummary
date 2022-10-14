package mqtt

import mqtt "github.com/eclipse/paho.mqtt.golang"

type MqttClientConf struct {
	TopicPrefix       string
	ClientIdPrefix    string
	EndPoint          string
	Username          string `json:",optional"`
	Password          string `json:",optional"`
	ReconnectInterval int    `json:",default=3"`
	KeepAlive         int    `json:",default=3"`
	PingTimeout       int    `json:",default=10"`
	TlsEnable         bool   `json:",default=false"`
	IsGrayRelease     bool   `json:",default=false"`
}

type MqttConfig struct {
	Opts *mqtt.ClientOptions
}

func NewMqttConfig(config MqttClientConf) (*MqttConfig, error) {
	//cfg := &MqttClientConf{}
	//opt,err :=
	return nil, nil
}
