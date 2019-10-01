//show code1 OMIT
type Proponent struct {
	FirstName    string
	MiddleName   string
	LastName     string
	Street       string
	StreetNumber int
	Complement   string
	City         string
	State        string
}

//end show code1 OMIT
//show code2 OMIT
type Proponent struct {
	Name Name
	Addr Address
}

//end show code2 OMIT
//show code3 OMIT
type Name struct {
	Family Surname
	Given  GivenNames
}

type Surname struct {
	Family string
}

type GivenNames struct {
	Names []string
}

//end show code3 OMIT
//show code4 OMIT
type Address struct {
	Street       string
	StreetNumber int
	Complement   string
	District     string
	City         string
	State        string
}
//end show code4 OMIT