package models

type event struct {
	ID string `json:"ID"`
	Title string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var Events = allEvents{

	{
		ID: "1",
		Title:"Hello from event",
		Description: "Hey from Description",
	},
}