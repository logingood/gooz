package cmd_test

import (
	"encoding/json"
	"fmt"

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
		"url": "http://test.zendesk.com/api/v2/users/75.json",
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

			_ = json.Unmarshal(TestOrgs, &data)
			for _, element := range data {
				for k, v := range element {
					Context(fmt.Sprintf("Test org %s", k), func() {
						It(fmt.Sprintf("should search organizations %s", k), func() {
							str, err := helpers.DetectTypeAndStringfy(v)
							results := SearchField(organizationsFilePath, k, str)
							Expect(len(results)).ToNot(BeNil())
							Expect(err).ToNot(HaveOccurred())
						})
					})
				}
			}

			_ = json.Unmarshal(TestUsers, &data)
			for _, element := range data {
				for k, v := range element {
					Context(fmt.Sprintf("Test user %s", k), func() {
						It(fmt.Sprintf("should search users %s", k), func() {
							str, err := helpers.DetectTypeAndStringfy(v)
							results := SearchField(usersFilePath, k, str)
							Expect(len(results)).ToNot(BeNil())
							Expect(err).ToNot(HaveOccurred())
						})
					})
				}
			}

			_ = json.Unmarshal(TestTickets, &data)
			for _, element := range data {
				for k, v := range element {
					Context(fmt.Sprintf("Test tickets %s", k), func() {
						It(fmt.Sprintf("should search tickets %s", k), func() {
							str, err := helpers.DetectTypeAndStringfy(v)
							results := SearchField(ticketsFilePath, k, str)
							Expect(len(results)).ToNot(BeNil())
							Expect(err).ToNot(HaveOccurred())
						})
					})
				}
			}
		})
	})
})
