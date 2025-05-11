package valueobject

type Address struct {
	Country string
	City    string
	Street  string
	ZipCode string
}

func (a Address) String() string {
	return a.Country + ", " + a.City + ", " + a.Street + ", " + a.ZipCode
}
