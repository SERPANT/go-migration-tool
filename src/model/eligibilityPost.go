package model

type EligibilityPost struct {
	PatientAddress       string `json:"patientAddress"`
	PatientCity          string `json:"patientCity"`
	PatientDateOfBirth   string `json:"patientDateOfBirth"`
	PatientDateofService string `json:"patientDateofService"`
	PatientFirstName     string `json:"patientFirstName"`
	PatientGender        string `json:"patientGender"`
	PatientLastName      string `json:"patientLastName"`
	PatientPolicyNumber  string `json:"patientPolicyNumber"`
	PatientState         string `json:"patientState"`
	PatientZip           string `json:"patientZip"`
	PayerEligibilityID   string `json:"payerEligibilityID"`
	PayerID              string `json:"payerID"`
}
