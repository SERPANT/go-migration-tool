package services

import (
	"svna/src/constants"
	"svna/src/model"
	"svna/src/util"
)

func GetContactById(id string) (*model.Contact, error) {

	var url string = constants.BaseURL + constants.Contact + "(" + id + ")?$select=address1_line1,firstname,lastname,birthdate,address1_city,address1_stateorprovince"

	util.Print("Calling Url: " + url)

	data, err := util.Get(url)

	if err != nil {
		return nil, err
	}

	var contact model.Contact

	errjson := util.ParseToObject(data, &contact)

	if errjson != nil {
		return nil, errjson
	}

	return &contact, nil
}
