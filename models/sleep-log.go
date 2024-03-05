package models

import entity "github.com/anwarul/healthtracking_engine/Framework/Entity"

type SleepRawData[T any] struct {
	StartDate T      `json:"startDate"`
	EndDate   T      `json:"endDate"`
	Type      string `json:"type"`
}
type SleepLogBasic[T any] struct {
	TotalTime     float32           `json:"totalTime"`
	TotalInBed    float32           `json:"totalInBed"`
	TotalIsAwake  float32           `json:"totalIsAwake"`
	TotalIsAsleep float32           `json:"totalIsAsleep"`
	Data          []SleepRawData[T] `json:"data"`
	HealthTrackingCommon[T]
}
type SleepLog[T any] struct {
	SleepLogBasic[T]
	entity.EntityBase[T]
}
