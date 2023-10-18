package myLog

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/heyitsfranky/MyConfig/src/myConfig"
	"github.com/segmentio/kafka-go"
)

var data *InitData

type InitData struct {
	KafkaBroker  string `yaml:"kafka-broker"`
	ClientOrigin string `yaml:"client-origin"`
}

type LogData struct {
	Body   interface{} `json:"body"`
	Origin string      `json:"origin"`
	Type   string      `json:"type"`
	Level  int         `json:"level"`
}

func Init(configPath string) error {
	if data == nil {
		err := myConfig.Init(configPath, &data)
		if err != nil {
			return err
		}
	}
	return nil
}

func createEvent(body interface{}, givenType string, level int, async bool) error {
	if data == nil {
		return fmt.Errorf("must first call Init() to initialize the kafka settings")
	}
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{data.KafkaBroker},
		Topic:    "create_log",
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()
	logData := LogData{
		Body:   body,
		Origin: data.ClientOrigin,
		Type:   givenType,
		Level:  level,
	}
	jsonString, err := json.Marshal(logData)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return err
	}
	message := kafka.Message{
		Key:   []byte(fmt.Sprintf("key-%d", time.Now().Unix())),
		Value: []byte(jsonString),
	}

	if async {
		go func() {
			err := writer.WriteMessages(context.Background(), message)
			if err != nil {
				fmt.Println("Error sending message asynchronously:", err)
			}
		}()
	} else {
		err := writer.WriteMessages(context.Background(), message)
		if err != nil {
			fmt.Println("Error sending message synchronously:", err)
			return err
		}
	}
	return nil
}

// these functions are just used to create simple, fast logs
func createInfoEvent(body string, givenType string) {
	createEvent(body, givenType, 0, true)
}
func createWarningEvent(body string, givenType string) {
	createEvent(body, givenType, 1, true)
}
func createErrorEvent(body string, givenType string) {
	createEvent(body, givenType, 2, true)
}
func createCriticalEvent(body string, givenType string) {
	createEvent(body, givenType, 3, true)
}
