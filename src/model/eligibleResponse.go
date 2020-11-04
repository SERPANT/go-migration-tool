package model

type ErrorResponse struct {
	ResponseDesc string `json:"responseDesc"`
}

type EligibilityError struct {
	Response ErrorResponse `json:"response"`
}

type EligibleResponse struct {
	IsEligible bool             `json:"isEligible"`
	Error      EligibilityError `json:"error"`
}
