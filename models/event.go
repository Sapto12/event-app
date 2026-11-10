package models

import "time"

type Event struct {
	Id          int
	Name        string `binding:required`
	Description string `binding:required`
	Location    string `binding:required`
	DateTime    time.Time
	UserId int
}


var events []Event = []Event{}

//func untuk simpan avent
func (e Event) Save(){
	events = append(events, e)

}

//menampilkan tampilan semua event

func GetAllEvents() []Event {
	return events
}