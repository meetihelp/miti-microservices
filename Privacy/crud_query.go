package Privacy

import(
	database "miti-microservices/Database"
	util "miti-microservices/Util"
)

func CreateBoard(userId string,date string,boardId string){
	db:=database.GetDB()
	board:=Board{}
	db.Where("user_id=? AND board_id=?",userId,boardId).Find(&board)
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

func EnterBoardContent(userId string,boardId string,text string,imageId string,contentId string,requestId string,createdAt string) (string,string){
	db:=database.GetDB()
	boardContent:=BoardContent{}
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
		db.Create(&boardContent)
		return createdAt,boardContent.ContentId
	}else{
		return boardContent.CreatedAt,boardContent.ContentId
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