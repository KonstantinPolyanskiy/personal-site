package common_data

import "time"

type IdentityInformation struct {
	Id   int
	Name *string
}

type CreationInformation struct {
	CreatedAt time.Time
	UpdatedAt *time.Time
}
