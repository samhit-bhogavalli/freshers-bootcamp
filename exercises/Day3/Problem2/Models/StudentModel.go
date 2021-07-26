package Models

type Student struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	DOB       string `json:"dob"`
	Address   string `json:"address"`
	Subject   string `json:"subject"`
	Marks     int    `json:"marks"`
}
