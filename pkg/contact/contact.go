package contact

//Contact is type
type Contact struct {
	Id    uint `json:"id" gorm:"primary_key"`
	Name  string
	Phone string
	Email string
	City  string
	Age   int
}

//ContactRepository is interface
type ContactRepository interface {
	New(*Contact) (*Contact, error)
	Update(*Contact) (*Contact, error)
	Delete(*Contact) error

	Get(int) (*Contact, error)
	List() ([]*Contact, error)
}
