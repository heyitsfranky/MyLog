package myLog

import (
	"testing"
)

func TestInit(t *testing.T) {
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
			err := Init("../configs/" + tt.config)
			if err == nil && tt.expectErr {
				t.Errorf("Expected an error but got nil")
			} else if err != nil && !tt.expectErr {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

func TestCreateSyncEvent(t *testing.T) {
	Init("../configs/config.yaml")
	tests := []struct {
		name      string
		body      interface{}
		givenType string
		level     int
		async     bool
		expectErr bool
	}{
		{"ValidEventSync", map[string]interface{}{"key1": "value1", "key2": 42}, "testType", 1, false, false},
		{"ValidEventAsync", "test body", "testType", 1, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := createEvent(tt.body, tt.givenType, tt.level, tt.async)
			if err == nil && tt.expectErr {
				t.Errorf("Expected an error but got nil")
			} else if err != nil && !tt.expectErr {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}
