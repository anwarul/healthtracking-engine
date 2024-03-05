package models

type HealthTrackingCommon[T any] struct {
	SyncDataDate T      `json:"syncDataDate"`
	StartDate    T      `json:"startDate"`
	EndDate      T      `json:"endDate"`
	DeviceMacId  string `json:"deviceMacId"`
}
