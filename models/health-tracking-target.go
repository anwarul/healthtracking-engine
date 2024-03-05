package models

import (
	entity "github.com/anwarul/healthtracking_engine/Framework/Entity"
	"github.com/anwarul/healthtracking_engine/enums"
)

type HealthTrackingTarget[T any] struct {
	DeviceMacId string                    `json:"deviceMacId"`
	TargetValue float32                   `json:"targetValue"`
	TargetPoint enums.HealthTrackingPoint `json:"targetPoint"`
	entity.EntityBase[T]
}
