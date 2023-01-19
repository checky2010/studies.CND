package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestDatapoint_ToDatapoint(t *testing.T) {
	dp := &DatapointEntity{
		Model: gorm.Model{
			ID:        0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now().Add(time.Minute),
			DeletedAt: gorm.DeletedAt{},
		},
		Value: 25,
		Date:  time.Now().Add(time.Hour),
	}

	model := dp.ToDatapoint()

	assert.Equal(t, dp.Value, model.Value)
	assert.Equal(t, dp.Date, model.Date)
}
