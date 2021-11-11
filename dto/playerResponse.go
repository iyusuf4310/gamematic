package dto

type PlayerResponse struct {
	Id           string `json:"Id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	DateofBirth  string `json:"date_of_birth"`
	Gender       string `json:"gender"`
	PhoneNumber  string `json:"phone_number"`
	EmailAddress string `json:eEmail_address"`
	JerseNumber  string `json:"jerse_number"`
	Team         string `json:"team"`
	AddressResponse
}

type AddressResponse struct {
	Address1 string `json:"address_1"`
	Address2 string `json:"address_2"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zipcode  string `json:"zip_code"`
}
