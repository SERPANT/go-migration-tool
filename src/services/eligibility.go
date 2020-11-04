package services

import (
	"encoding/json"

	"svna/src/constants"
	"svna/src/model"
	"svna/src/util"
)

func CheckEligibility(appointment *model.Appointment, contact *model.Contact) (*model.InsuraceEiligible, error) {

	var payerId string = "00932"

	var eligilityRes model.EligibilityPost = model.EligibilityPost{
		PatientAddress:       contact.AddressLine1,
		PatientCity:          contact.AddressCity1,
		PatientDateOfBirth:   "01/06/1983",
		PatientDateofService: "10/20/2020",
		PatientFirstName:     contact.FirstName,
		PatientGender:        "",
		PatientLastName:      contact.LastName,
		PatientPolicyNumber:  appointment.MemberId,
		PatientState:         contact.AddressStateProvice,
		PatientZip:           "98115",
		PayerEligibilityID:   payerId,
		PayerID:              payerId,
	}

	jsonByte, _ := json.Marshal(eligilityRes)

	data, err := util.Post(constants.EligibilityUrl, string(jsonByte))

	if err != nil {
		return nil, err
	}

	var eligibilityResult model.EligibleResponse

	errorJson := util.ParseToObject(data, &eligibilityResult)

	if errorJson != nil {
		return nil, errorJson
	}

	var eligibleObj model.InsuraceEiligible = model.InsuraceEiligible{
		MemberId:               appointment.MemberId,
		HealthINsuranceCompany: appointment.HealthINsuranceCompany,
		AppointmentId:          appointment.Id,
		FirstName:              contact.FirstName,
		PatientId:              contact.Id,
		Lastname:               contact.LastName,
		BirthDate:              contact.BirthDate,
		AddressLine1:           contact.AddressLine1,
		AddressCity1:           contact.AddressCity1,
		AddressStateProvice:    contact.AddressStateProvice,
		IsEligible:             eligibilityResult.IsEligible,
	}

	return &eligibleObj, nil
}
