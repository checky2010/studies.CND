package entities

import (
	"evaluation/domain/model"
	"time"

	"gorm.io/gorm"
)

type DatapointEntity struct {
	gorm.Model
	Value int       `json:"value"`
	Date  time.Time `json:"date"`
}

func (entity *DatapointEntity) ToDatapoint() *model.Datapoint {
	return &model.Datapoint{
		Value: entity.Value,
		Date:  entity.Date,
	}
}
