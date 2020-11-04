package model

type Contact struct {
	DataTag             string `json:"@odata.etag"`
	FirstName           string `json:"firstname"`
	Id                  string `json:"contactid"`
	LastName            string `json:"lastname"`
	BirthDate           string `json:"birthdate"`
	AddressLine1        string `json:"address1_line1"`
	AddressCity1        string `json:"address1_city"`
	AddressStateProvice string `json:"address1_stateorprovince"`
}
