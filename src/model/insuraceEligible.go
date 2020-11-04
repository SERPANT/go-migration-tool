package model

type InsuraceEiligible struct {
	MemberId               string
	HealthINsuranceCompany string
	AppointmentId          string
	DataTag                string
	FirstName              string
	PatientId              string
	Lastname               string
	BirthDate              string
	AddressLine1           string
	AddressCity1           string
	AddressStateProvice    string
	IsEligible             bool
}
