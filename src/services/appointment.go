package services

import (
	"fmt"

	"svna/src/constants"
	"svna/src/db"
	"svna/src/model"
	"svna/src/util"
)

func FetchWithAllFields(count int, organizationName string) ([]model.Appointment, error) {

	var url string = constants.BaseURL + "msemr_appointmentemrs?$select=*&$filter=svna_health_insurance_company%20eq%20'" + organizationName + "'&$top=" + fmt.Sprint(count)

	util.Print("Calling url: " + url)

	data, err := util.Get(url)

	if err != nil {

		return nil, err
	}

	var appointmentList struct {
		Value []model.Appointment `json:"value"`
	}

	errorjson := util.ParseToObject(data, &appointmentList)

	if errorjson != nil {
		return nil, errorjson
	}

	if len(appointmentList.Value) == 0 {
		return nil, nil
	}

	return appointmentList.Value, nil

}

func FetchByOrganization(count int, organizationName string) ([]model.Appointment, error) {

	var url string = constants.BaseURL + "msemr_appointmentemrs?$select=svna_member_id,_msemr_actorpatient_value,svna_health_insurance_company&$filter=svna_health_insurance_company%20eq%20'" + organizationName + "'&$top=" + fmt.Sprint(count)

	util.Print("calling url: " + url)

	data, err := util.Get(url)

	if err != nil {

		return nil, err
	}

	var appointmentList struct {
		Value []model.Appointment `json:"value"`
	}

	errorjson := util.ParseToObject(data, &appointmentList)

	if errorjson != nil {
		return nil, errorjson
	}

	if len(appointmentList.Value) == 0 {
		return nil, nil
	}

	return appointmentList.Value, nil

}

func InsertToDataBase(appointmentObj *model.Appointment) {

	fmt.Println("Inserting into database")

	var query string = fmt.Sprintf(`INSERT INTO Appointment (member_id, patient_id, health_insurance_company, id) VALUES ('%s', '%s', '%s', '%s');`, appointmentObj.MemberId, appointmentObj.PatientId, appointmentObj.HealthINsuranceCompany, appointmentObj.Id)

	util.Print("Query: " + query)

	results, errDBRetrieve := db.DbConn.Query(query)

	if errDBRetrieve != nil {
		fmt.Println("Query Failed :  ", errDBRetrieve)
		return
	}

	defer results.Close()
}
