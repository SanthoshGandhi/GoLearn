package Models

type User struct {
	UserID    string `json:"userid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Gender    string `json:"gender"`
	Address   string `json:"address"`
	DOB       string `json:"dob"`
	Country   string `json:"country"`
	Status    string `json:"status"`
	EmailId   string `json:"emailid"`
	Password  string `json:"password"`
}
