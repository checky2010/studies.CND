package gorm

import (
	"evaluation/domain/model"
	"time"
)

type DatapointRepository interface {
	Save(datapoint *model.Datapoint)
	FindForTime(start, end *time.Time) []*model.Datapoint
}
