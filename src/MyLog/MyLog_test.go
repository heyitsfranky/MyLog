package myLog

import (
	"testing"
)

func Test_Init(t *testing.T) {
	tests := []struct {
		name      string
		config    string
		expectErr bool
	}{
		{"ValidConfig", "config.yaml", false},
		{"MissingClientOrigin", "missing_clientorigin.yaml", true},
		{"MissingKafkaBroker", "missing_kafkabroker.yaml", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Data = nil
			err := Init("../../configs/" + tt.config)
			if err == nil && tt.expectErr {
				t.Errorf("Expected an error but got nil")
			} else if err != nil && !tt.expectErr {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

func Test_CreateEvent(t *testing.T) {
	Init("../../configs/config.yaml")
	tests := []struct {
		name      string
		body      interface{}
		caller    string
		level     int
		async     bool
		expectErr bool
	}{
		{"ValidEventSync", map[string]interface{}{"key1": "value1", "key2": 42}, "callerfunction", 1, false, false},
		{"ValidEventAsync", "test body", "callerfunction", 1, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateEvent(tt.body, tt.caller, tt.level, tt.async)
			if err == nil && tt.expectErr {
				t.Errorf("Expected an error but got nil")
			} else if err != nil && !tt.expectErr {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}
