package Models

type Student struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	DOB       string `json:"dob"`
	Address   string `json:"address"`
}

type StudentMarks struct {
	StudentID uint    `json:"studentId"`
	Student   Student `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Subject   string  `json:"subject"`
	Marks     int     `json:"marks"`
}
