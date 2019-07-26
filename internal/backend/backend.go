package backend

import (
	"github.com/logingood/gooz/internal/schema"
)

type Store interface {
	ReadTickets() ([]*schema.Ticket, error)
	ReadOrganizations() ([]*schema.Organization, error)
	ReadUsers() ([]*schema.User, error)
}
