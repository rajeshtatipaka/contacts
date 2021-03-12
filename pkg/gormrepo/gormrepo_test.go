package gormrepo_test

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rajeshtatipaka/contacts/pkg/contact"
	"github.com/rajeshtatipaka/contacts/pkg/gormrepo"
)

func TestGormrepo(t *testing.T) {
	g, err := gormrepo.NewGormrepo("testdb.db")
	defer g.Close()

	if err != nil {
		t.Fatalf("Could not open database: %v", err)
	}

	c1 := contact.Contact{
		Name: "First Contact",
		Age:  50,
		City: "Bangalore",
	}

	c2 := contact.Contact{
		Name: "Second Contact",
		Age:  25,
		City: "Mumbai",
	}

	g.New(&c1)
	g.New(&c2)

	allcontacts, err := g.List()

	for _, c := range allcontacts {
		t.Log(*c)
	}

	g.Close()
}
