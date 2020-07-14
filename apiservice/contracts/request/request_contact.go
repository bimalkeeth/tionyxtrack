package request

import "time"

type ContactRequest struct {
	Id            uint64    `json:"id"`
	Contact       string    `json:"contact"`
	ContactTypeId uint64    `json:"contacttypeid"`
	UpdatedAt     time.Time `json:"updatedat"`
}
