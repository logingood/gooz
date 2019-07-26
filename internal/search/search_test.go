package search_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/logingood/gooz/internal/backend/zfile"
	"github.com/logingood/gooz/internal/config"
	. "github.com/logingood/gooz/internal/search"
	"github.com/spf13/afero"
)

var Config = &config.Config{}
var TestUsers = []byte(`[
	{
		"_id": 1,
		"url": "http://test.zendesk.com/api/v2/users/75.json",
		"external_id": "0db0c1da-8901-4dc3-a469-fe4b500d0fca",
		"name": "Vasya Pupking",
		"alias": "Ivan Karavai",
		"created_at": "2016-06-07T09:18:00 -10:00",
		"active": false,
		"verified": true,
		"shared": true,
		"locale": "ru-RU",
		"timezone": "US Minor Outlying Islands",
		"last_login_at": "2012-10-15T12:36:41 -11:00",
		"email": "rosannasimpson@example.com",
		"phone": "8615-883-123",
		"signature": "Don't Worry Be Happy!",
		"organization_id": 222,
		"tags": [
			"tag5",
			"tag6",
			"anothertag"
		],
		"suspended": true,
		"role": "agent"
	},
	{
		"_id": 2,
		"url": "http://test.zendesk.com/api/v2/users/75.json",
		"external_id": "0db0c1ff-8901-4dc3-a469-fe4b500d0fca",
		"name": "John Smith",
		"alias": "Agent Smith",
		"created_at": "2016-06-07T09:18:00 -10:00",
		"active": false,
		"verified": true,
		"shared": true,
		"locale": "ru-RU",
		"timezone": "US Minor Outlying Islands",
		"last_login_at": "2012-11-05T12:36:41 -11:00",
		"email": "rosannasimpson@domain.com",
		"phone": "8615-883-789",
		"signature": "Don't Worry Be Happy!",
		"organization_id": 222,
		"tags": [
			"tag1",
			"tag2",
			"anothertag"
		],
		"suspended": false,
		"role": "agent"
	}]`)

var TestOrgs = []byte(`[
		{
			"_id": 111,
			"url": "http://initech.zendesk.com/api/v2/organizations/125.json",
			"external_id": "42a1a845-70cf-40ed-a762-acb27fd606cc",
			"name": "Petya",
			"domain_names": [
				"goodwebsite.com",
				"betterwebsite.com"
			],
			"created_at": "2016-02-21T06:11:51 -11:00",
			"details": "MegaGOOS",
			"shared_tickets": false,
			"tags": [
				"Silver",
				"Cat",
				"Eats"
			]
		},
		{
			"_id": 222,
			"url": "http://initech.zendesk.com/api/v2/organizations/125.json",
			"external_id": "42a1a845-70cf-40ed-a762-acb27fd606cc",
			"name": "Kolya",
			"domain_names": [
				"iamspam.com",
				"iamnotspam.com"
			],
			"created_at": "2016-02-21T06:11:51 -11:00",
			"details": "GoosCorp",
			"shared_tickets": false,
			"tags": [
				"Jacobs",
				"Frank"
			]
		}
	]`)

var TestTickets = []byte(
	`[{
			"_id": "123fc8bc-31de-411e-92bf-a6d6b9dfa490",
			"url": "http://ivan.zendesk.com/api/v2/tickets/50dfc8bc-31de-411e-92bf-a6d6b9dfa490.json",
			"external_id": "8bc8bee7-2d98-4b69-b4a9-4f348ff41fa3",
			"created_at": "2016-03-08T09:44:54 -11:00",
			"type": "task",
			"subject": "A Problem in South Africa",
			"description": "Esse nisi occaecat pariatur veniam culpa dolore anim elit aliquip. Cupidatat mollit nulla consectetur ullamco tempor esse.",
			"priority": "high",
			"status": "hold",
			"submitter_id": 43,
			"assignee_id": 54,
			"organization_id": 111,
			"tags": [
				"Georgia",
				"Tennessee",
				"Mississippi",
				"Marshall Islands"
			],
			"has_incidents": true,
			"due_at": "2016-08-03T09:17:37 -10:00",
			"via": "voice"
		},
		{
			"_id": "456fc8bc-31de-411e-92bf-a6d6b9dfa490",
			"url": "http://vasya.zendesk.com/api/v2/tickets/50dfc8bc-31de-411e-92bf-a6d6b9dfa490.json",
			"external_id": "8bc8bee7-2d98-4b69-b4a9-4f348ff41fa3",
			"created_at": "2016-03-08T09:44:54 -11:00",
			"type": "task",
			"subject": "A Problem in Burgundy",
			"description": "Esse nisi occaecat pariatur veniam culpa dolore anim elit aliquip. Cupidatat mollit nulla consectetur ullamco tempor esse.",
			"priority": "Low",
			"status": "hold",
			"submitter_id": 43,
			"assignee_id": 54,
			"organization_id": 222,
			"tags": [
				"Georgia",
				"Tennessee",
				"Mississippi",
				"Marshall Islands"
			],
			"has_incidents": true,
			"due_at": "2016-08-03T09:17:37 -10:00",
			"via": "voice"
		}]`)

var _ = Describe("Search", func() {
	appFS := afero.NewMemMapFs()

	Describe("Filesystem tests", func() {
		BeforeEach(func() {
			afero.WriteFile(appFS, "/tmp/orgs.json", TestOrgs, 0644)
			afero.WriteFile(appFS, "/tmp/users.json", TestUsers, 0644)
			afero.WriteFile(appFS, "/tmp/tickets.json", TestTickets, 0644)

			Config.OrgsFilePath = "/tmp/orgs.json"
			Config.TicketsFilePath = "/tmp/tickets.json"
			Config.UsersFilePath = "/tmp/users.json"
		})

		Context("Basic search", func() {
			It("It should handle non-existing org file", func() {

				store, err := zfile.New(Config, appFS)
				searcher, err := New(Config, store)

				results, err := searcher.SearchOrgsByString("Name", "Petya")

				fmt.Printf("=> %+v\n", results.Organizations[0])
				fmt.Printf("=> %+v\n", results.Tickets)
				fmt.Printf("=> %+v\n", results.Users)
				fmt.Printf("=> %+v\n", err)

				Expect(results).ToNot(BeNil())
				Expect(err).ToNot(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("open /tmp/nonexistentorgs.json: file does not exist"))
				Expect(store).To(BeNil())
			})
		})
	})
})
