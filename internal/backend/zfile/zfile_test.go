package zfile_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/google/uuid"
	. "github.com/logingood/gooz/internal/backend/zfile"
	"github.com/logingood/gooz/internal/config"
	"github.com/spf13/afero"
)

var Config = &config.Config{}

var _ = Describe("Zfile", func() {
	appFS := afero.NewMemMapFs()

	Describe("Filesystem tests", func() {
		BeforeEach(func() {
			afero.WriteFile(appFS, "/tmp/orgs.json", []byte("orgs"), 0644)
			afero.WriteFile(appFS, "/tmp/users.json", []byte("users"), 0644)
			afero.WriteFile(appFS, "/tmp/tickets.json", []byte("tickets"), 0644)

			Config.OrgsFilePath = "/tmp/orgs.json"
			Config.TicketsFilePath = "/tmp/tickets.json"
			Config.UsersFilePath = "/tmp/users.json"
		})

		Context("Files do not exist", func() {
			It("It should handle non-existing org file", func() {

				Config.OrgsFilePath = "/tmp/nonexistentorgs.json"

				store, err := New(Config, appFS)
				Expect(err).ToNot(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("open /tmp/nonexistentorgs.json: file does not exist"))
				Expect(store).To(BeNil())
			})

			It("It should handle non-existing users file", func() {

				Config.UsersFilePath = "/tmp/nonexistentusers.json"

				store, err := New(Config, appFS)
				Expect(err).ToNot(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("open /tmp/nonexistentusers.json: file does not exist"))
				Expect(store).To(BeNil())
			})

			It("It should handle non-existing tickets file", func() {

				Config.TicketsFilePath = "/tmp/nonexistenttickets.json"

				store, err := New(Config, appFS)
				Expect(err).ToNot(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("open /tmp/nonexistenttickets.json: file does not exist"))
				Expect(store).To(BeNil())
			})

			It("It should handle non-existing directory file", func() {

				Config.TicketsFilePath = "/bad/directory/nonexistenttickets.json"

				store, err := New(Config, appFS)
				Expect(err).ToNot(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("open /bad/directory/nonexistenttickets.json: file does not exist"))
				Expect(store).To(BeNil())
			})
		})

		Context("Files exist", func() {
			It("It should return File Descriptors", func() {
				store, err := New(Config, appFS)

				Expect(err).To(BeNil())
				Expect(err).ToNot(HaveOccurred())
				Expect(store).ToNot(BeNil())
			})

			It("Should handle json parsing if descriptor(s) closed", func() {
				store, _ := New(Config, appFS)

				store.OrgsFD.Close()
				store.TicketsFD.Close()
				store.UsersFD.Close()

				orgs, err_orgs := store.ReadOrganizations()
				users, err_users := store.ReadUsers()
				tickets, err_tickets := store.ReadTickets()

				Expect(err_orgs).To(HaveOccurred())
				Expect(err_users).To(HaveOccurred())
				Expect(err_tickets).To(HaveOccurred())
				Expect(orgs).To(BeNil())
				Expect(users).To(BeNil())
				Expect(tickets).To(BeNil())
				Expect(err_orgs).To(MatchError("unexpected end of JSON input"))
				Expect(err_users).To(MatchError("unexpected end of JSON input"))
				Expect(err_tickets).To(MatchError("unexpected end of JSON input"))

			})

			It("Should handle malformed json", func() {
				store, _ := New(Config, appFS)

				orgs, err_orgs := store.ReadOrganizations()
				users, err_users := store.ReadUsers()
				tickets, err_tickets := store.ReadTickets()

				Expect(err_orgs).To(HaveOccurred())
				Expect(err_users).To(HaveOccurred())
				Expect(err_tickets).To(HaveOccurred())
				Expect(orgs).To(BeNil())
				Expect(users).To(BeNil())
				Expect(tickets).To(BeNil())
			})

			It("Should handle incorrect json data types", func() {
				store, _ := New(Config, appFS)
				TestUsers := []byte(`
					[{
						"_id": "blah",
						"url": "http://test.zendesk.com/api/v2/users/75.json",
						"last_login_at": 123
					}]`)

				TestOrgs := []byte(`[
					{
						"_id": 🧠,
						"url": "http://test.zendesk.com/api/v2/users/75.json",
						"tags": "I am not a list"
					}]`)

				TestTickets := []byte(`[{
					"_id": 2321321,
					"url": 123,
					"tags": 123
				}]`)

				afero.WriteFile(appFS, "/tmp/orgs.json", TestOrgs, 0644)
				afero.WriteFile(appFS, "/tmp/users.json", TestUsers, 0644)
				afero.WriteFile(appFS, "/tmp/tickets.json", TestTickets, 0644)

				_, err_orgs := store.ReadOrganizations()
				_, err_users := store.ReadUsers()
				_, err_tickets := store.ReadTickets()

				Expect(err_orgs).To(HaveOccurred())
				Expect(err_users).To(HaveOccurred())
				Expect(err_tickets).To(HaveOccurred())
				Expect(err_orgs).To(MatchError("invalid character 'ð' looking for beginning of value"))
				Expect(err_users).To(MatchError("json: cannot unmarshal string into Go struct field User._id of type int64"))
				Expect(err_tickets).To(MatchError("json: cannot unmarshal number into Go struct field Ticket._id of type uuid.UUID"))
			})

			It("Should bind users, tickets and organizations to json", func() {
				store, _ := New(Config, appFS)
				TestUsers := []byte(`[
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
						"organization_id": 3,
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
						"organization_id": 2,
						"tags": [
							"tag1",
							"tag2",
							"anothertag"
						],
						"suspended": false,
						"role": "agent"
					}]`)

				TestOrgs := []byte(`[
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

				TestTickets := []byte(
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
						"organization_id": 103,
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
						"organization_id": 103,
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

				afero.WriteFile(appFS, "/tmp/orgs.json", TestOrgs, 0644)
				afero.WriteFile(appFS, "/tmp/users.json", TestUsers, 0644)
				afero.WriteFile(appFS, "/tmp/tickets.json", TestTickets, 0644)

				orgs, err_orgs := store.ReadOrganizations()
				users, err_users := store.ReadUsers()
				tickets, err_tickets := store.ReadTickets()

				uuid1, _ := uuid.Parse("123fc8bc-31de-411e-92bf-a6d6b9dfa490")
				uuid2, _ := uuid.Parse("456fc8bc-31de-411e-92bf-a6d6b9dfa490")

				Expect(err_orgs).ToNot(HaveOccurred())
				Expect(err_users).ToNot(HaveOccurred())
				Expect(err_tickets).ToNot(HaveOccurred())
				Expect(orgs[0].Id).To(Equal(int64(111)))
				Expect(orgs[1].Id).To(Equal(int64(222)))
				Expect(users[0].Id).To(Equal(int64(1)))
				Expect(users[1].Id).To(Equal(int64(2)))
				Expect(tickets[0].Id).To(Equal(uuid1))
				Expect(tickets[1].Id).To(Equal(uuid2))
				Expect(tickets[0].DueAt.Unix()).To(Equal(int64(1470251857)))
				Expect(tickets[0].CreatedAt.Unix()).To(Equal(int64(1457469894)))
			})
		})
	})
})
