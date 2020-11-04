package main

import (
	"svna/src/constants"
	"svna/src/db"
	"svna/src/services"
	"svna/src/util"
)

const numberOfRecord int = 1

func main() {
	db.SetupDatabase()

	fetchAppointmentData()

}

func fetchAppointmentData() {
	util.Print("Starting Application to check eligibility..........")

	list, err := services.FetchWithAllFields(numberOfRecord, constants.Regence)

	if err != nil {
		util.Error(err)

		return
	}

	for _, value := range list {

		util.Print("Sending contact fetch Request for: " + value.Id)

		contact, err := services.GetContactById(value.PatientId)

		if err != nil {
			util.Error(err)

			continue
		}

		util.Print(contact)

		insuranceEligibilityResult, err := services.CheckEligibility(&value, contact)

		if err != nil {
			util.Error(err)

			continue

		}

		services.InsertToDataBase(&value)

		util.Print(insuranceEligibilityResult)
	}
}
