package schema

type User struct {
	Id             int64    `json:"_id"`
	URL            string   `json:"url"`
	ExtId          string   `json:"external_id"`
	Name           string   `json:"name"`
	Alias          string   `json:"alias"`
	CreatedAt      string   `json:"created_at"`
	Active         bool     `json:"active"`
	Verified       bool     `json:"verified"`
	Shared         bool     `json:"shared"`
	Locale         string   `json:"local"`
	TZ             string   `json:"timezone"`
	LastLoginAt    string   `json:"last_loging_at"`
	Email          string   `json:"email"`
	Phone          string   `json:"phone"`
	Signature      string   `json:"signature"`
	OrganizationId int64    `json:"organization_id"`
	Tags           []string `json:"tags"`
	Suspended      bool     `json:"suspended"`
	Role           string   `json:"role"`
}

type Ticket struct {
	Id             string   `json:"_id"`
	URL            string   `json:"url"`
	ExtId          string   `json:"external_id"`
	CreatedAt      string   `json:"created_at"`
	Type           string   `json:"type"`
	Subject        string   `json:"subject"`
	Description    string   `json:"description"`
	Priority       string   `json:"priority"`
	Status         string   `json:"status"`
	SubmitterId    int64    `json:"submitter_id"`
	AssigneeId     int64    `json:"assignee_id"`
	OrganizationId int64    `json:"organization_id"`
	Tags           []string `json:"tags"`
	HasIncidents   bool     `json:"has_incidents"`
	DueAt          string   `json:"due_at"`
	Channel        string   `json:"via"`
}

type Organization struct {
	Id            int64    `json:"_id"`
	URL           string   `json:"url"`
	ExtId         string   `json:"external_id"`
	Name          string   `json:"name"`
	DomainNames   []string `json:"domain_names"`
	CreatedAt     string   `json:"created_at"`
	Details       string   `json:"details"`
	SharedTickets bool     `json:"shared_tickets"`
	Tags          []string `json:"tags"`
}
