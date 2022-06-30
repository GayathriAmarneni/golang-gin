package main

type HistoryObj struct {
	Tags     string
	Title    string
	Mediaid  string
	Duration float64
	Progress float64
}
type ExternalDataObj struct {
	History   []HistoryObj
	Favorites []string
}
type ResponseDataObj struct {
	Id            int32
	Email         string
	FirstName     string
	LastName      string
	Country       string
	RegDate       string
	LastLoginDate string
	LastUserIp    string
	ExternalId    string
	ExternalData  ExternalDataObj
}
type Customer struct {
	ResponseData ResponseDataObj
	Errors       []string
}

// `json: email`
// `json: lastName`
// `json: firstName`
// `json: regDate`
// `json: country`
// `json: lastLoginDate`
// `json: externalId`
// `json: externalData`
// `json: lastUserIp`
