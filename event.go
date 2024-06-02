package main

type Event struct {
	Id               string `csv:"Id"`
	Region           string `csv:"Region"`
	Ward             string `csv:"Ward"`
	ONSCode          string `csv:"ONSCode"`
	ConstituencyName string `csv:"Constituency Name"`
	CouncilOns       string `csv:"Council Ons"`
	CouncilName      string `csv:"Council Name"`
	EventTitle       string `csv:"Event Title"`
	Description      string `csv:"Description"`
	ValidationStatus string `csv:"Validation Status"`
	RSVPs            string `csv:"RSVPs"`
	Category         string `csv:"Category"`
	Location         string `csv:"Location"`
	Postcode         string `csv:"Postcode"`
	StartTime        string `csv:"Start Time"`
	EndTime          string `csv:"End Time"`
	Creator          string `csv:"Creator"`
	URL              string `csv:"URL"`
}
