package MyLog

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/heyitsfranky/MyConfig"
	"github.com/segmentio/kafka-go"
)

var Data *InitData

type InitData struct {
	KafkaBroker  string `yaml:"kafka-broker"`
	ClientOrigin string `yaml:"client-origin"`
}

type LogData struct {
	Body   interface{} `json:"body"`
	Origin string      `json:"origin"`
	Caller string      `json:"caller"`
	Level  int         `json:"level"`
}

func Init(configPath string) error {
	if Data == nil {
		err := MyConfig.Init(configPath, &Data)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateEvent(body interface{}, caller string, level int, async bool) error {
	if Data == nil {
		return fmt.Errorf("must first call Init() to initialize the kafka settings")
	}
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{Data.KafkaBroker},
		Topic:    "create_log",
		Balancer: &kafka.LeastBytes{},
	})
	if !async { //necessary as otherwise the writer is closed as soon as the function ends
		defer writer.Close()
	}
	logData := LogData{
		Body:   body,
		Origin: Data.ClientOrigin,
		Caller: caller,
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
			writer.Close()
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
func CreateInfoEvent(body string, caller string) {
	CreateEvent(body, caller, 0, true)
}
func CreateWarningEvent(body string, caller string) {
	CreateEvent(body, caller, 1, true)
}
func CreateErrorEvent(body string, caller string) {
	CreateEvent(body, caller, 2, true)
}
func CreateCriticalEvent(body string, caller string) {
	CreateEvent(body, caller, 3, true)
}
