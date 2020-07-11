package bucontracts

import "time"

type ContactBO struct {
	Id            uint
	Contact       string
	ContactTypeId uint
	UpdatedAt     time.Time
}
