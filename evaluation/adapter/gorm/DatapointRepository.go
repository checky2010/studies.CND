package gorm

import (
	"evaluation/domain/model"
	"time"
)

type DatapointRepository interface {
	Save(datapoint *model.Datapoint)
	FindForTime(start time.Time, end time.Time) []*model.Datapoint
}
