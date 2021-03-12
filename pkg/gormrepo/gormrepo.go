package gormrepo

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rajeshtatipaka/contacts/pkg/contact"
)

//Gormrepo implementation
type Gormrepo struct {
	filename string
	db       *gorm.DB
}

//NewGormrepo method
func NewGormrepo(fn string) (*Gormrepo, error) {
	db, err := gorm.Open("sqlite3", fn)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&contact.Contact{})

	return &Gormrepo{
		filename: fn,
		db:       db,
	}, nil
}

//Close Implementation
func (g *Gormrepo) Close() {
	g.db.Close()
}

//New Implementation
func (g *Gormrepo) New(c *contact.Contact) (*contact.Contact, error) {
	g.db.Create(c)
	return c, nil
}

//Update Implementation
func (g *Gormrepo) Update(c *contact.Contact) (*contact.Contact, error) {
	panic("not implemented") // TODO: Implement
}

//Delete Implementation
func (g *Gormrepo) Delete(c *contact.Contact) error {
	panic("not implemented") // TODO: Implement
}

//Get Implementation
func (g *Gormrepo) Get(c int) (*contact.Contact, error) {
	panic("not implemented") // TODO: Implement
}

//List Implementation
func (g *Gormrepo) List() ([]*contact.Contact, error) {
	var allcontacts []*contact.Contact
	g.db.Find(&allcontacts)
	return allcontacts, nil
}
