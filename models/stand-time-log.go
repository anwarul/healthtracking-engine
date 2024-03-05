package models

import entity "github.com/anwarul/healthtracking_engine/Framework/Entity"

type StandTimeBasic[T any] struct {
	Value float32 `json:"value"`
	HealthTrackingCommon[T]
}
type StandTimeLog[T any] struct {
	StandTimeBasic[T]
	entity.EntityBase[T]
}
