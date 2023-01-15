package dto

import "time"

type Datapoint struct {
	Value int       `json:"value"`
	Date  time.Time `json:"date"`
}
