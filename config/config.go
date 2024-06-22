package config

import (
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
)

// RabbitMQ type declaration
type RabbitMQ struct {
	Host         string `koanf:"host"`
	Port         string `koanf:"port"`
	User         string `koanf:"user"`
	Password     string `koanf:"password"`
	ExchangeName string `koanf:"exchange_name"`
	ExchangeType string `koanf:"exchange_type"`
}

// Global koanf instance. Use "." as the key path delimiter. This can be "/" or any character.
var k = koanf.New(".")

// loads the config file into the koanf instance
func loadConfig() {
	if err := k.Load(file.Provider("/home/amit/Desktop/RabbitMQ_Trial/Twitter_Tweet_Delivery_Fanout_Architecture/config/config.json"), json.Parser()); err != nil {
		log.Fatalf("Error loading config.json : %v", err)
	}
}

var RabbitMQConfig RabbitMQ

// Unmarshal RabbitMQ stanza from config.json into RabbitMQ type
func NewRabbitMQConfig() *RabbitMQ {
	loadConfig()
	k.Unmarshal("RabbitMQ", &RabbitMQConfig)
	return &RabbitMQConfig
}

// Building RabbitMQ url
func (r *RabbitMQ) GetURL() string {
	url := "amqp://" + r.User + ":" + r.Password + "@" + r.Host + ":" + r.Port + "/"
	return url
}

// Getter method for exchange name
func (r *RabbitMQ) GetExchangeName() string {
	return r.ExchangeName
}

// Getter method for exchange type
func (r *RabbitMQ) GetExchangeType() string {
	return r.ExchangeType
}
