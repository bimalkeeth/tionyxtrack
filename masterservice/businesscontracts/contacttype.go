package bucontracts

import uuid "github.com/satori/go.uuid"

type ContactTypeBO struct {
	Id          uuid.UUID
	ContactType string
}
