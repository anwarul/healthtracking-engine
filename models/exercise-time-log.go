package models

import (
	entity "github.com/anwarul/healthtracking_engine/Framework/Entity"
	"github.com/anwarul/healthtracking_engine/enums"
)

type ExerciseTimeLogBasic[T any] struct {
	Value float32                        `json:"value"`
	Type  string                         `json:"type"`
	Os    enums.ExerciseTrackingDeviceOs `json:"os"`
	HealthTrackingCommon[T]
}

type ExerciseTimeLog[T any] struct {
	ExerciseTimeLogBasic[T]
	entity.EntityBase[T]
}
