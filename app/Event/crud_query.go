package Event

import(
	database "app/Database"
	util "app/Util"
)

func InsertEvent(event Event) string{
	db:=database.GetDB()
	event.EventId=util.GenerateToken()
	event.CreatedAt=util.GetTime()
	db.Create(&event)
	return event.EventId
}
