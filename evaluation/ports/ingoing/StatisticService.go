package ingoing

import (
	"evaluation/domain/model"
	"time"
)

type StatisticService interface {
	AllDatapoints(start, end *time.Time) []*model.Datapoint
	AverageValue(start, end *time.Time) float64
	MaxDatapoint(start, end *time.Time) *model.Datapoint
	MinDatapoint(start, end *time.Time) *model.Datapoint
}
