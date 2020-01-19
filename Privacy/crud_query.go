package Privacy

import(
	database "miti-microservices/Database"
	util "miti-microservices/Util"
	"fmt"
)

func CreateBoard(userId string,date string,boardId string){
	db:=database.GetDB()
	board:=Board{}
	db.Where("user_id=? AND date=? AND board_id=?",userId,date,boardId).Find(&board)
	if(board.UserId==""){
		board.UserId=userId
		board.Date=date
		board.BoardId=boardId
		board.CreatedAt=util.GetTime()
		board.AccessType="Private"
		db.Create(&board)
	}
}
func GetBoardId(userId string,date string) string{
	db:=database.GetDB()
	board:=Board{}
	db.Where("user_id=? AND date=?",userId,date).Find(&board)
	if(board.UserId==""){
		board.UserId=userId
		board.Date=date
		board.CreatedAt=util.GetTime()
		board.BoardId=util.GenerateToken()
		board.AccessType="Private"
		db.Create(&board)
		return board.BoardId
	}

	return board.BoardId
}

func EnterBoardContent(userId string,boardId string,text string,imageId string,contentId string,requestId string,createdAt string) (string,string,int){
	db:=database.GetDB()
	boardContent:=BoardContent{}
	code:=200
	db.Where("user_id=? AND board_id=? AND request_id=?",userId,boardId,requestId).Find(&boardContent)
	if(boardContent.UserId==""){
		boardContent.UserId=userId
		boardContent.BoardId=boardId
		boardContent.ContentId=contentId
		boardContent.ContentText=text
		boardContent.ContentImageId=imageId
		boardContent.AccessType="Private"
		boardContent.CreatedAt=createdAt
		boardContent.RequestId=requestId
		err:=db.Create(&boardContent).Error
		if(err!=nil){
			code=1006
			fmt.Print("EnterBoardContent DB Error:")
			fmt.Println(err)
		}
		return createdAt,boardContent.ContentId,code
	}else{
		return boardContent.CreatedAt,boardContent.ContentId,code
	}
}

func UpdateBoardSharePolicy(userId string,boardId string,accessType string,requestId string)string{
	board:=Board{}
	db:=database.GetDB()
	db.Where("user_id=? AND board_id=?",userId,boardId).Find(&board)
	if(board.UserId==""){
		return ""
	}else if(board.RequestId==requestId){
		return board.UpdatedAt
	}else{
		board.RequestId=requestId
		board.UpdatedAt=util.GetTime()
		board.AccessType=accessType
		db.Where("user_id=? AND board_id=?",userId,boardId).Updates(board)
		return board.UpdatedAt
	}
}

func UpdateBoardContentSharePolicy(userId string,contentId string,boardId string,accessType string,accessRequestId string)string{
	boardContent:=BoardContent{}
	db:=database.GetDB()
	db.Where("user_id=? AND content_id=? AND board_id=?",userId,contentId,boardId).Find(&boardContent)
	if(boardContent.UserId==""){
		return ""
	}else if(boardContent.AccessRequestId==accessRequestId){
		return boardContent.AccessUpdatedAt
	}else{
		boardContent.AccessRequestId=accessRequestId
		boardContent.AccessUpdatedAt=util.GetTime()
		boardContent.AccessType=accessType
		db.Where("user_id=? AND content_id=?  AND board_id=?",userId,contentId,boardId).Updates(boardContent)
		return boardContent.AccessRequestId
	}
}

func GetBoardContentDB(userId string,createdAt string) []BoardContentList{
	db:=database.GetDB()
	boardContent:=[]BoardContent{}
	db.Order("created_at").Where("user_id=? AND created_at>?",userId,createdAt).Find(&boardContent)
	boardContentList:=make([]BoardContentList,0)
	for _,content:=range boardContent{
		listElement:=BoardContentList{}
		listElement.UserId=content.UserId
		listElement.ContentId=content.ContentId
		listElement.BoardId=content.BoardId
		listElement.AccessType=content.AccessType
		listElement.ContentText=content.ContentText
		listElement.ContentImageId=content.ContentImageId
		listElement.CreatedAt=content.CreatedAt
		boardContentList=append(boardContentList,listElement)
	}
	return boardContentList
}