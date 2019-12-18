package Event

import(
	database "miti-microservices/Database"
	util "miti-microservices/Util"
)

func InsertEvent(event Event) string{
	db:=database.GetDB()
	event.EventId=util.GenerateToken()
	event.CreatedAt=util.GetTime()
	db.Create(&event)
	return event.EventId
}

func GetEventByIdDB(eventId string) (Event,string){
	db:=database.GetDB()
	event:=Event{}
	db.Where("event_id=?",eventId).Find(&event)
	if event.EventId==""{
		return event,"Error"
	}
	return event,"Ok"
}