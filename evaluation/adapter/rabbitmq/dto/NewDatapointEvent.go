package dto

import "time"

type NewDatapointEvent struct {
	Value int       `json:"value"`
	Date  time.Time `json:"date"`
}
