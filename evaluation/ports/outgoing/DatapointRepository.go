package outgoing

import (
	"evaluation/domain/model"
	"time"
)

type DatapointRepository interface {
	FindForTime(start, end *time.Time) []*model.Datapoint
	Save(datapoint *model.Datapoint)
}
