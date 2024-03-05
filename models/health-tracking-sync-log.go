package models

import "github.com/anwarul/healthtracking_engine/framework/entity"

type HealthTrackingSyncLog[T any] struct {
	SyncDataDate T      `json:"syncDataDate"`
	DeviceMacId  string `json:"deviceMacId"`
	entity.EntityBase[T]
}
