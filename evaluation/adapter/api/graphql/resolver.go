package graphql

import (
	"evaluation/ports/ingoing"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	StatisticService ingoing.StatisticService
}
