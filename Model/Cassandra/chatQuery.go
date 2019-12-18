package Cassandra

import(
	"fmt"
	"log"
	CD "miti-microservices/Model/CreateDatabase"
	// util "miti-microservices/Utility"
)

func EnterChat(chat CD.Chat){
	db:=GetDB()
	err:=db.Query(`INSERT INTO chats(user_id,chat_id,message_id,message_type,message_content,created_at) VALUES(?,?,?,?,?,?)`,
						chat.User_id,chat.Chat_id,chat.Message_id,chat.Message_type,chat.Message_content,chat.CreatedAt).Exec()
	if err!=nil{
		log.Fatal(err)
	}

}

func GetChatMessages(chat_id string,offset int,num_of_rows int)([]CD.Chat){
	chat:=[]CD.Chat{}
	db:=GetDB()
	iter:=db.Query(`SELECT * FROM chats WHERE chat_id=?`,chat_id).Iter()
	log.Fatal(iter)
	for iter.Scan(&chat){
		fmt.Printf("gaurav")
	}
	return chat
}