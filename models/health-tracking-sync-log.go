package models

import entity "github.com/anwarul/healthtracking_engine/Framework/Entity"

type HealthTrackingSyncLog[T any] struct {
	SyncDataDate T      `json:"syncDataDate"`
	DeviceMacId  string `json:"deviceMacId"`
	entity.EntityBase[T]
}
