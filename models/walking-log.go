package models

import entity "github.com/anwarul/healthtracking_engine/Framework/Entity"

type WalkingLog[T any] struct {
	StepCount int `json:"stepCount"`
	HealthTrackingCommon[T]
	entity.EntityBase[T]
}
