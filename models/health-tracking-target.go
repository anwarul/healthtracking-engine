package models

import (
	"github.com/anwarul/healthtracking_engine/enums"
	"github.com/anwarul/healthtracking_engine/framework/entity"
)

type HealthTrackingTarget[T any] struct {
	DeviceMacId string                    `json:"deviceMacId"`
	TargetValue float32                   `json:"targetValue"`
	TargetPoint enums.HealthTrackingPoint `json:"targetPoint"`
	entity.EntityBase[T]
}
