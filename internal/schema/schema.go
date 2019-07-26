package schema

import (
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type ZTime struct {
	time.Time
}

func (self *ZTime) UnmarshalJSON(b []byte) (err error) {
	// Date doesn't satisfy RFC3339 because of spce in the layout
	// e.g. 2006-01-02T15:04:05 -01:00
	// we format it to 2006-01-02T15:04:05-01:00 (no space)
	s, err := strconv.Unquote(string(b))
	if err != nil {
		log.Errorf("STRConv problems", err)
	}

	s = strings.Join(strings.Fields(s), "")
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		// default date
		log.Errorf("TIME problems", err)
		t, err = time.Parse(time.RFC3339, "2001-01-01T00:00:00-00:00")
	}
	self.Time = t
	return
}

type User struct {
	Id             int64     `json:"_id"`
	URL            string    `json:"url"`
	ExtId          uuid.UUID `json:"external_id"`
	Name           string    `json:"name"`
	Alias          string    `json:"alias"`
	CreatedAt      ZTime     `json:"created_at"`
	Active         bool      `json:"active"`
	Verified       bool      `json:"verified"`
	Shared         bool      `json:"shared"`
	Locale         string    `json:"locale"`
	TZ             string    `json:"timezone"`
	LastLoginAt    ZTime     `json:"last_loging_at"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Signature      string    `json:"signature"`
	OrganizationId int64     `json:"organization_id"`
	Tags           []string  `json:"tags"`
	Suspended      bool      `json:"suspended"`
	Role           string    `json:"role"`
}

type Ticket struct {
	Id             uuid.UUID `json:"_id"`
	URL            string    `json:"url"`
	ExtId          uuid.UUID `json:"external_id"`
	CreatedAt      ZTime     `json:"created_at"`
	Type           string    `json:"type"`
	Subject        string    `json:"subject"`
	Description    string    `json:"description"`
	Priority       string    `json:"priority"`
	Status         string    `json:"status"`
	SubmitterId    int64     `json:"submitter_id"`
	AssigneeId     int64     `json:"assignee_id"`
	OrganizationId int64     `json:"organization_id"`
	Tags           []string  `json:"tags"`
	HasIncidents   bool      `json:"has_incidents"`
	DueAt          ZTime     `json:"due_at"`
	Channel        string    `json:"via"`
}

type Organization struct {
	Id            int64     `json:"_id"`
	URL           string    `json:"url"`
	ExtId         uuid.UUID `json:"external_id"`
	Name          string    `json:"name"`
	DomainNames   []string  `json:"domain_names"`
	CreatedAt     ZTime     `json:"created_at"`
	Details       string    `json:"details"`
	SharedTickets bool      `json:"shared_tickets"`
	Tags          []string  `json:"tags"`
}
