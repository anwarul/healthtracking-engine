package models

import (
	entity "github.com/anwarul/healthtracking_engine/Framework/Entity"
)

type ActiveEnergyBurnedLog[T any] struct {
	Value float32 `json:"value"`
	Type  string  `json:"type"`
	HealthTrackingCommon[T]
	entity.EntityBase[T]
}
