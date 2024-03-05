package entity

import "github.com/google/uuid"

type EntityIdentity struct {
	ItemId uuid.UUID `json:"itemId"`
}

type EntityBase[T any] struct {
	TenantId            uuid.UUID `json:"tenantId"`
	SiteId              uuid.UUID `json:"siteId"`
	IsMarkedToDelete    bool      `json:"isMarkedToDelete"`
	LockId              int64     `json:"lockId"`
	LockTime            T         `json:"lockTime"`
	Tags                []string  `json:"tags"`
	CreatedByUserId     uuid.UUID `json:"createdByUserId"`
	LastUpdatedByUserId uuid.UUID `json:"lastUpdatedByUserId"`
	TimeStamp           int64     `json:"timeStamp"`
	CorrelationId       uuid.UUID `json:"correlationId"`
	CreateDate          T         `json:"createDate"`
	LastUpdateDate      T         `json:"lastUpdateDate"`
	EntityIdentity
}
