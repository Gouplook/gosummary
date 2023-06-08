package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"io/ioutil"
	"time"
)

type MqttClientConfig struct {
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
	CaFilePath        string `json:",default=/conf/ca/root_cert.crt"`
	CrtFilePath       string `json:",default=/conf/ca/device_cert.crt"`
	KeyFilePath       string `json:",default=/conf/ca/device_key.pem"`
}

type MqttConfig struct {
	Opts *MQTT.ClientOptions
}

func NewMqttConfig(config MqttClientConfig) (*MqttConfig, error) {
	// 初始参数
	opt := MQTT.NewClientOptions()
	opt.AddBroker(config.EndPoint)
	clientID := config.ClientIdPrefix + time.Now().Format(time.RFC3339Nano)
	opt.SetClientID(clientID)
	opt.SetUsername(config.Username)
	opt.SetPassword(config.Password)
	opt.SetAutoReconnect(true)
	if config.ReconnectInterval > 0 {
		opt.SetMaxReconnectInterval(time.Second * time.Duration(config.ReconnectInterval))
	}

	if config.KeepAlive > 0 {
		opt.SetKeepAlive(time.Second * time.Duration(config.KeepAlive))
	}

	if config.PingTimeout > 0 {
		opt.SetPingTimeout(time.Second * time.Duration(config.PingTimeout))
	}

	// 传输层可使用TLS加密
	if config.TlsEnable {
		tlsConfig := NewTLSConfig(config)

		opt.SetTLSConfig(tlsConfig)
	}

	return &MqttConfig{
		Opts: opt,
	}, nil
}

// NewTLSConfig 做一层加密
func NewTLSConfig(config MqttClientConfig) *tls.Config {
	// Import trusted certificates from CAfile.pem.
	// Alternatively, manually add CA certificates to
	// default openssl CA bundle.
	certpool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile(config.CaFilePath)
	if err == nil {
		certpool.AppendCertsFromPEM(pemCerts)
	}

	// Import client certificate/key pair
	cert, err := tls.LoadX509KeyPair(config.CrtFilePath, config.KeyFilePath)
	if err != nil {
		panic(err)
	}

	// Just to print out the client certificate..
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		panic(err)
	}
	// fmt.Println(cert.Leaf)

	// Create tls.Config with desired tls properties
	return &tls.Config{
		// RootCAs = certs used to verify server cert.
		RootCAs: certpool,
		// ClientAuth = whether to request cert from server.
		// Since the server is set up for SSL, this happens
		// anyways.
		ClientAuth: tls.NoClientCert,
		// ClientCAs = certs used to validate client cert.
		ClientCAs: nil,
		// InsecureSkipVerify = verify that cert contents
		// match server. IP matches what is in cert etc.
		InsecureSkipVerify: true,
		// Certificates = list of certs client sends to server.
		Certificates: []tls.Certificate{cert},
	}
}
