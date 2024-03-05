package models

import entity "github.com/anwarul/healthtracking_engine/Framework/Entity"

type DistanceWalkingRunningLog[T any] struct {
	Value float32 `json:"value"`
	HealthTrackingCommon[T]
	entity.EntityBase[T]
}
