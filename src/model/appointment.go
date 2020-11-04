package model

type Appointment struct {
	MemberId               string `json:"svna_member_id"`
	PatientId              string `json:"_msemr_actorpatient_value"`
	HealthINsuranceCompany string `json:"svna_health_insurance_company"`
	Id                     string `json:"activityid"`
}
