package zfile

import (
	"encoding/json"
	"io/ioutil"

	"github.com/logingood/gooz/internal/config"
	"github.com/logingood/gooz/internal/schema"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

type Store struct {
	OrgsFD    afero.File
	UsersFD   afero.File
	TicketsFD afero.File
}

func New(goozConfig *config.Config, goozFS afero.Fs) (store *Store, err error) {

	orgsFD, err := goozFS.Open(goozConfig.OrgsFilePath)
	if err != nil {
		log.Errorf("Error has occured when opening a file %s, error: %s", goozConfig.OrgsFilePath, err)
		return nil, err
	}

	usersFD, err := goozFS.Open(goozConfig.UsersFilePath)
	if err != nil {
		log.Errorf("Error has occured when opening a file %s, error: %s", goozConfig.UsersFilePath, err)
		return nil, err
	}

	ticketsFD, err := goozFS.Open(goozConfig.TicketsFilePath)
	if err != nil {
		log.Errorf("Error has occured when opening a file %s, error: %s", goozConfig.TicketsFilePath, err)
		return nil, err
	}

	store = &Store{
		OrgsFD:    orgsFD,
		UsersFD:   usersFD,
		TicketsFD: ticketsFD,
	}

	return store, err
}

func (f *Store) ReadOrganizations() (orgs []*schema.Organization, err error) {
	byteData, err := ioutil.ReadAll(f.OrgsFD)

	if err != nil {
		log.Errorf("Can't read file content %s", err)
	}

	err = json.Unmarshal(byteData, &orgs)

	if err != nil {
		log.Errorf("Can't read organizations, invalid json %s", err)
		return nil, err
	}

	defer f.OrgsFD.Close()

	return orgs, err
}

func (f *Store) ReadTickets() (tickets []*schema.Ticket, err error) {
	byteData, err := ioutil.ReadAll(f.TicketsFD)

	if err != nil {
		log.Errorf("Can't read file content %s", err)
	}

	err = json.Unmarshal(byteData, &tickets)

	if err != nil {
		log.Errorf("Can't read tickets, invalid json %s", err)
		return nil, err
	}

	defer f.TicketsFD.Close()

	return tickets, err
}

func (f *Store) ReadUsers() (users []*schema.User, err error) {
	byteData, err := ioutil.ReadAll(f.UsersFD)

	if err != nil {
		log.Errorf("Can't read file content %s", err)
	}

	err = json.Unmarshal(byteData, &users)

	if err != nil {
		log.Errorf("Can't read users, invalid json %s", err)
		return nil, err
	}

	defer f.UsersFD.Close()

	return users, err
}
