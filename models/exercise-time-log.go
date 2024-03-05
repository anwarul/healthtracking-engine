package models

import (
	"github.com/anwarul/healthtracking_engine/enums"
	"github.com/anwarul/healthtracking_engine/framework/entity"
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
