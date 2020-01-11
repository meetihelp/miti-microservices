package Social

import(
	database "miti-microservices/Database"
	
)

func PoolStatusDB(userId string) PoolStatus{
	db:=database.GetDB()
	poolStatus:=PoolStatus{}
	db.Where("user_id=?",userId).Find(&poolStatus)
	return poolStatus
}


func EnterInPooL(userId string,pincode string,createdAt string,gender string,sex string){
	db:=database.GetDB()
	poolWait:=PoolWaiting{}
	poolWait.UserId=userId
	poolWait.Pincode=pincode
	poolWait.CreatedAt=createdAt
	poolWait.Gender=gender
	poolWait.Sex=sex
	_=db.Create(&poolWait).Error
}

func DeleteWaitPool(userId string) {
	db:=database.GetDB()
	db.Where("user_id=?",userId).Delete(&PoolWaiting{})
}

func DeletePool(userId string,areaCode string,gender string){
	db:=database.GetDB()
	pool:=Pool{}
	db.Where("user_id=?",userId).Find(&pool)
	if(pool.UserId!=""){
		status:=EnterInPoolFromWait(areaCode,gender,1)
		Complementary_gender:="Male"
		if(gender=="Male"){
			Complementary_gender="Female"
		}
		if(status=="NA"){
			DeleteFromPoolHelper(areaCode,Complementary_gender,1)
		}
	}
}

func EnterInPoolFromWait(areaCode string,gender string,number_of_person int) string{
	return "Ok"
}

func DeleteFromPoolHelper(areaCode string,gender string,number_of_person int){

}