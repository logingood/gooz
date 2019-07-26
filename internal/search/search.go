package search

import (
	"reflect"
	"regexp"

	"github.com/logingood/gooz/internal/backend"
	"github.com/logingood/gooz/internal/config"
	"github.com/logingood/gooz/internal/schema"
	log "github.com/sirupsen/logrus"
)

type Search struct {
	Users         []*schema.User
	Organizations []*schema.Organization
	Tickets       []*schema.Ticket
	store         backend.Store
}

func New(config *config.Config, store backend.Store) (searcher *Search, err error) {
	searcher = &Search{
		store: store,
	}

	searcher.Tickets, err = store.ReadTickets()
	if err != nil {
		log.Errorf("Error has occurred while read users %s", err)
	}

	searcher.Organizations, err = store.ReadOrganizations()
	if err != nil {
		log.Errorf("Error has occurred while read users %s", err)
	}

	searcher.Users, err = store.ReadUsers()
	if err != nil {
		log.Errorf("Error has occurred while read users %s", err)
	}

	return searcher, nil
}

func (s *Search) SearchOrgsByString(property string, pattern string) (results *Search, err error) {

	results = &Search{
		Organizations: make([]*schema.Organization, 0),
		Users:         make([]*schema.User, 0),
		Tickets:       make([]*schema.Ticket, 0),
	}

	for _, element := range s.Organizations {

		orgsReflection := reflect.Indirect(reflect.ValueOf(element))
		field := orgsReflection.FieldByName(property)

		match, err := regexp.MatchString(pattern, field.String())

		if err != nil {
			log.Errorf("Can't perform match, malformed patter %s", err)
			return results, err
		}

		if match {
			results.Organizations = append(results.Organizations, element)
			usersFound := results.SearchUsersByOrgId(element.Id)
			ticketsFound := results.SearchTicketsByOrgId(element.Id)
			results.Users = append(results.Users, usersFound...)
			results.Tickets = append(results.Tickets, ticketsFound...)
		}
	}

	return results, err
}

func (s *Search) SearchOrgsById(id int64) (results *schema.Organization) {
	for _, element := range s.Organizations {
		if element.Id == id {
			return element
		}
	}

	return nil
}

func (s *Search) SearchUsersById(id int64) (results *schema.User) {
	for _, element := range s.Users {
		if element.Id == id {
			return element
		}
	}

	return nil
}

func (s *Search) SearchUsersByOrgId(id int64) (results []*schema.User) {
	results = make([]*schema.User, 0)
	for _, element := range s.Users {
		if element.Id == id {
			results = append(results, element)
		}
	}

	return results
}

func (s *Search) SearchTicketsByOrgId(id int64) (results []*schema.Ticket) {
	results = make([]*schema.Ticket, 0)

	for _, element := range s.Tickets {
		if element.OrganizationId == id {
			results = append(results, element)
		}
	}

	return results
}

func (s *Search) SearchTicketsByAssigneId(id int64) (results []*schema.Ticket) {
	for _, element := range s.Organizations {
		if element.Id == id {
			return results
		}
	}

	return nil
}

func (s *Search) SearchTicketsByRequesterId(id int64) (results []*schema.Ticket) {
	for _, element := range s.Organizations {
		if element.Id == id {
			return results
		}
	}

	return nil
}
