package filerepo

import (
	"os"

	"github.com/rajeshtatipaka/contact/pkg/contact"
)

//Filerepo implements filerepo
type Filerepo struct {
	filepath string
	file     *os.File
	LastID   uint
	Contacts []contact.Contact
}

/*
func (f *Filerepo) New(c *contact.Contact) (*contact.Contact, error) {
	f.LastID += 1
	c.Id = f.LastID
	f.Contacts = append(f.Contacts, c)
	return c, nil
}

func (f *Filerepo) Update(c *contact.Contact) (*contact.Contact, error) {
	panic("not implemented") // TODO: Implement
}
func (f *Filerepo) Delete(c *contact.Contact) error {
	panic("not implemented") // TODO: Implement
}

func (f *Filerepo) Get(id int) (*contact.Contact, error) {
	panic("not implemented") // TODO: Implement
}

func (f *Filerepo) List() ([]*contact.Contact, error) {
	panic("not implemented") // TODO: Implement
}
*/
