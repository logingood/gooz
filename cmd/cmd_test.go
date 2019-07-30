package cmd_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/logingood/gooz/cmd"
	"github.com/logingood/gooz/internal/helpers"
	"github.com/spf13/afero"
)

var TestUsers = []byte(`[
	{
		"_id": 1,
		"url": "http://test.zendesk.com/api/v2/users/75.json",
		"external_id": "0db0c1da-8901-4dc3-a469-fe4b500d0fca",
		"name": "John Pumpkin",
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
		"organization_id": 111,
		"tags": [
			"tag5",
			"tag6",
			"anothertag"
		],
		"suspended": true,
		"role": "agent"
	},
	{
		"_id": 3,
		"url": "http://test.zendesk.com/api/v2/users/75.json",
		"external_id": "0db0c1ff-8901-4dc3-a469-fe4b500d0fca",
		"name": "Cobra Commander",
		"verified": true,
		"shared": true,
		"locale": "ru-ru",
		"timezone": "us minor outlying islands",
		"last_login_at": "2012-11-05t12:36:41 -11:00",
		"email": "rosannasimpson@domain.com",
		"phone": "8615-883-789",
		"signature": "don't worry be happy!",
		"organization_id": 111,
		"tags": [
			"tag1",
			"tag2",
			"anothertag"
		],
		"suspended": false,
		"role": "agent"
  },
	{
		"_id": 4,
		"external_id": "0db0c1ff-8901-4dc3-a469-fe4b500d0fca",
		"name": "Jean Luc Picard",
		"verified": true,
		"shared": "",
		"locale": "ru-ru",
		"timezone": "us minor outlying islands",
		"last_login_at": "2012-11-05t12:36:41 -11:00",
		"email": "rosannasimpson@domain.com",
		"phone": "8615-883-789",
		"signature": "don't worry be happy!",
		"organization_id": 111,
		"tags": [
			"tag1",
			"tag2",
			"anothertag"
		],
		"suspended": false,
		"role": "agent"
  },
	{
		"_id": 2,
		"url": "http://test.zendesk.com/api/v2/users/75.json",
		"external_id": "0db0c1ff-8901-4dc3-a469-fe4b500d0fca",
		"name": "John Smith",
		"alias": "Agent Smith",
		"created_at": "2016-06-07t09:18:00 -10:00",
		"active": false,
		"verified": true,
		"shared": true,
		"locale": "ru-ru",
		"timezone": "us minor outlying islands",
		"last_login_at": "2012-11-05t12:36:41 -11:00",
		"email": "rosannasimpson@domain.com",
		"phone": "8615-883-789",
		"signature": "don't worry be happy!",
		"organization_id": 111,
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
			"name": "Umbrella Corporation",
			"domain_names": [
				"goodwebsite.com",
				"web.site",
				"betterwebsite.com"
			],
			"created_at": "2016-02-21T06:11:51 -11:00",
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
			"name": "Hello Pty",
			"domain_names": [
				"iamspam.com",
				"web.site",
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
			"submitter_id": 2,
			"assignee_id": 1,
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
			"_id": "aaafc8bc-31de-411e-92bf-a6d6b9dfa490",
			"url": "http://moon.zendesk.com/api/v2/tickets/50dfc8bc-31de-411e-92bf-a6d6b9dfa490.json",
			"external_id": "8bb8bee7-2d98-4b69-b4a9-4f348ff41fa3",
			"created_at": "2016-03-08T09:44:54 -11:00",
			"type": "task",
			"subject": "NASA landed on the Moon 50 years",
			"description": "Esse nisi occaecat pariatur veniam culpa dolore anim elit aliquip. Cupidatat mollit nulla consectetur ullamco tempor esse.",
			"priority": "high",
			"status": "hold",
			"submitter_id": 2,
			"organization_id": 111,
			"tags": [
				"Bazz",
				"Aldrin",
				"Moon",
				"Rover"
			],
			"has_incidents": true,
			"due_at": "2016-08-03T09:17:37 -10:00",
			"via": "web"
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
			"submitter_id": 1,
			"assignee_id": 2,
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

var (
	usersFilePath         = "/tmp/users.json"
	organizationsFilePath = "/tmp/organizations.json"
	ticketsFilePath       = "/tmp/tickets.json"
)

var _ = Describe("Cmd", func() {
	appFS := afero.NewOsFs()

	Describe("Search tests", func() {
		afero.WriteFile(appFS, "/tmp/organizations.json", TestOrgs, 0644)
		afero.WriteFile(appFS, "/tmp/users.json", TestUsers, 0644)
		afero.WriteFile(appFS, "/tmp/tickets.json", TestTickets, 0644)

		// Thread unsafe tests
		Context("Test search", func() {
			data := []map[string]interface{}{}

			Context("Test search of organizations", func() {
				It("Should search organization by name", func() {
					_ = json.Unmarshal(TestOrgs, &data)
					strData, err := helpers.DetectTypeAndStringfy(data[0]["name"])

					results := SearchField(organizationsFilePath, "name", strData)
					Expect(len(results)).To(Equal(1))
					Expect(err).ToNot(HaveOccurred())
				})

				It("Should search organization by url", func() {
					_ = json.Unmarshal(TestOrgs, &data)
					strData, err := helpers.DetectTypeAndStringfy(data[0]["url"])

					results := SearchField(organizationsFilePath, "url", strData)
					Expect(len(results)).To(Equal(2))
					Expect(err).ToNot(HaveOccurred())
				})

				It("Should search organization by external_id", func() {
					_ = json.Unmarshal(TestOrgs, &data)
					strData, err := helpers.DetectTypeAndStringfy(data[0]["external_id"])

					results := SearchField(organizationsFilePath, "external_id", strData)
					Expect(len(results)).To(Equal(2))
					Expect(err).ToNot(HaveOccurred())
				})

				It("Should search organization by bool", func() {
					_ = json.Unmarshal(TestOrgs, &data)
					strData, err := helpers.DetectTypeAndStringfy(data[0]["shared_tickets"])

					results := SearchField(organizationsFilePath, "shared_tickets", strData)
					Expect(len(results)).To(Equal(2))
					Expect(err).ToNot(HaveOccurred())
				})

				It("Should search organization by id", func() {
					_ = json.Unmarshal(TestOrgs, &data)
					strData, err := helpers.DetectTypeAndStringfy(data[0]["_id"])

					results := SearchField(organizationsFilePath, "_id", strData)
					Expect(len(results)).To(Equal(1))
					Expect(err).ToNot(HaveOccurred())
				})

				It("Should search organization by empty field", func() {
					results := SearchField(organizationsFilePath, "details", "")
					Expect(len(results)).To(Equal(1))
				})

				It("Should search organization by domain names []string", func() {
					results := SearchField(organizationsFilePath, "domain_names", "web.site")
					Expect(len(results)).To(Equal(2))
				})
			})

			Context("Test search of users", func() {
				It("Should search users by name", func() {
					_ = json.Unmarshal(TestUsers, &data)
					strData, err := helpers.DetectTypeAndStringfy(data[0]["name"])

					results := SearchField(usersFilePath, "name", strData)
					Expect(len(results)).To(Equal(1))
					Expect(err).ToNot(HaveOccurred())
				})

				It("Should search users by bool", func() {
					_ = json.Unmarshal(TestUsers, &data)
					strData, err := helpers.DetectTypeAndStringfy(data[0]["active"])

					results := SearchField(usersFilePath, "active", strData)
					Expect(len(results)).To(Equal(2))
					Expect(err).ToNot(HaveOccurred())
				})

				It("Should search users by id", func() {
					_ = json.Unmarshal(TestUsers, &data)
					strData, err := helpers.DetectTypeAndStringfy(data[0]["_id"])

					results := SearchField(usersFilePath, "_id", strData)
					Expect(len(results)).To(Equal(1))
					Expect(err).ToNot(HaveOccurred())
				})

				It("Should search users by empty field", func() {
					results := SearchField(usersFilePath, "url", "")
					Expect(len(results)).To(Equal(1))
				})

				It("Should search users by []string (tags) field", func() {
					results := SearchField(usersFilePath, "tags", "anothertag")
					Expect(len(results)).To(Equal(4))
				})
			})

			Context("Test search of tickets", func() {
				It("Should search users by subject", func() {
					_ = json.Unmarshal(TestTickets, &data)
					strData, err := helpers.DetectTypeAndStringfy(data[0]["subject"])

					results := SearchField(ticketsFilePath, "subject", strData)
					Expect(len(results)).To(Equal(1))
					Expect(err).ToNot(HaveOccurred())
				})

				It("Should search tickets by bool", func() {
					_ = json.Unmarshal(TestTickets, &data)
					strData, err := helpers.DetectTypeAndStringfy(data[0]["has_incidents"])

					results := SearchField(ticketsFilePath, "has_incidents", strData)
					Expect(len(results)).To(Equal(3))
					Expect(err).ToNot(HaveOccurred())
				})

				It("Should search ticket by org_id", func() {
					_ = json.Unmarshal(TestTickets, &data)
					strData, err := helpers.DetectTypeAndStringfy(data[0]["organization_id"])

					results := SearchField(ticketsFilePath, "organization_id", strData)
					Expect(len(results)).To(Equal(2))
					Expect(err).ToNot(HaveOccurred())
				})

				It("Should search users by empty field", func() {
					results := SearchField(ticketsFilePath, "assignee_id", "")
					Expect(len(results)).To(Equal(1))
				})

				It("Should search users by []string (tags) field", func() {
					results := SearchField(ticketsFilePath, "tags", "Georgia")
					Expect(len(results)).To(Equal(2))
				})
			})
		})
	})
})
