package models

import entity "github.com/anwarul/healthtracking_engine/Framework/Entity"

type HeartRateBasic[T any] struct {
	MinValue float32 `json:"minValue"`
	MaxValue float32 `json:"maxValue"`
	AvgValue float32 `json:"avgValue"`
	HealthTrackingCommon[T]
}
type HeartRateLog[T any] struct {
	HeartRateBasic[T]
	entity.EntityBase[T]
}
