package graphql

import (
	"evaluation/domain"
	"evaluation/domain/model"
	"time"
)

type StatisticServiceImpl struct {
	DatapointService domain.DatapointService
}

func (s *StatisticServiceImpl) AllDatapoints(start, end *time.Time) []*model.Datapoint {
	return s.DatapointService.GetDatapointsForTime(start, end)
}

func (s *StatisticServiceImpl) AverageValue(start, end *time.Time) float64 {
	datapoints := s.DatapointService.GetDatapointsForTime(start, end)
	sum := 0.0
	for _, dp := range datapoints {
		sum += float64(dp.Value)
	}
	return sum / float64(len(datapoints))
}

func (s *StatisticServiceImpl) MaxDatapoint(start, end *time.Time) *model.Datapoint {
	datapoints := s.DatapointService.GetDatapointsForTime(start, end)
	var max *model.Datapoint
	for _, dp := range datapoints {
		if max == nil || dp.Value > max.Value {
			max = dp
		}
	}
	return max
}

func (s *StatisticServiceImpl) MinDatapoint(start, end *time.Time) *model.Datapoint {
	datapoints := s.DatapointService.GetDatapointsForTime(start, end)
	var min *model.Datapoint
	for _, dp := range datapoints {
		if min == nil || dp.Value < min.Value {
			min = dp
		}
	}
	return min
}
