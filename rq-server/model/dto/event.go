package dto

import (
	"time"
)

type Event struct {
	Type      string            `json:"eventType"`
	EventTime time.Time         `json:"eventTime"`
	MetaData  map[string]string `json:"metadata"`
}
