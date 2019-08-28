package Redis

import(
	// "github.com/garyburd/redigo/redis"
	// "github.com/go-redis/redis"
	CD "app/Model/CreateDatabase"
	"encoding/json"
	"fmt"
	// "strings"
)

func EnterChat(chat CD.Chat){
	conn:=getRedisDB()
	json,err:=json.Marshal(chat)
	if err!=nil{
		fmt.Println(err)
		return
	}
	// args:=[]interface{chat.Chat_id}
	// args=append(args,json)
	// _,e:=conn.Do("lpush",args...)
	// _,e:=conn.Do("LPUSH",chat.Chat_id,json)
	e:=conn.LPush(chat.Chat_id,json)
	// e:=conn.LPush("chat",json)
	if e!=nil{
		fmt.Println(e)
	}
}

func GetChatMessages(chat_id string,offset int,num_of_rows int)([]CD.Chat){
	fmt.Println(offset)
	fmt.Println(num_of_rows)
	conn:=getRedisDB()
	// s, err := redis.String(conn.Do("GET", chat_id))
	// args:=[]interface{chat_id}
	// args=append(args,0)
	// args=append(args,-1)
	// s, _ := conn.Do("lrange", chat_id,0 -1)
	// s,_:=conn.Do("lrange",args...)
	s:=conn.LRange(chat_id,0,-1)
	fmt.Println(s)
	// x,_:=s.([]byte)
	// str:=string([]byte(x))
	// fmt.Println(str)
	// chat:=make([] CD.Chat,0)
	// // err = json.Unmarshal([]byte(s), &chat)
	// // fmt.Println(chat)
	// // // c:=[]CD.Chat{}
	// // // c=append(c,chat)
	// // if err!=nil{
	// // 	return chat
	// // }
	// // return nil

	// arr:=strings.Split(str,"}")
	// arr=arr[:len(arr)-1]
	// // fmt.Println(arr)
	// for _,a:=range arr{
	// 	temp:=CD.Chat{}
	// 	a=a+"}"
	// 	fmt.Println(a)
	// 	err1:=json.Unmarshal([]byte(a),&temp)
	// 	chat=append(chat,temp)
	// 	if err1!=nil{
	// 		fmt.Println(err1)
	// 		return nil
	// 	}
	// }
	// return chat
	return nil
}