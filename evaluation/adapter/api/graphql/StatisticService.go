package graphql

import (
	"evaluation/domain/model"
	"time"
)

type StatisticService interface {
	AverageValue(start, end *time.Time) float64
	MaxDatapoint(start, end *time.Time) *model.Datapoint
	MinDatapoint(start, end *time.Time) *model.Datapoint
}
