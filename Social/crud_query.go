package Social

import(
	database "miti-microservices/Database"
	"fmt"
	util "miti-microservices/Util"
)

func PoolStatusDB(userId string) PoolStatus{
	db:=database.GetDB()
	poolStatus:=PoolStatus{}
	err:=db.Where("user_id=?",userId).Find(&poolStatus).Error
	fmt.Print("PoolStatusDB:")
	fmt.Println(err)
	return poolStatus
}


func EnterInPooL(userId string,pincode string,createdAt string,gender string,sex string,requestId string) PoolStatus{
	db:=database.GetDB()
	poolWait:=PoolWaiting{}

	db.Where("user_id=? AND request_id=?",userId,requestId).Find(&poolWait)
	if(poolWait.UserId==""){
		poolWait.UserId=userId
		poolWait.Pincode=pincode
		poolWait.CreatedAt=createdAt
		poolWait.Gender=gender
		poolWait.Sex=sex
		err:=db.Create(&poolWait).Error
		fmt.Print("EnterInPooL DB1:")
		fmt.Println(err)
	}

	poolStatus:=PoolStatus{}
	db.Where("user_id=?",userId).Find(&poolStatus)
	if(poolStatus.UserId==""){
		poolStatus.UserId=userId
		poolStatus.Status="Waiting"
		poolStatus.CreatedAt=util.GetTime()
		err:=db.Create(&poolStatus).Error
		fmt.Print("EnterInPooL DB2:")
		fmt.Println(err)
	}

	return poolStatus
}

func EnterInGroupPooL(userId string,pincode string,interest string,createdAt string,gender string,sex string){
	db:=database.GetDB()
	poolWait:=GroupPoolWaiting{}
	poolWait.UserId=userId
	poolWait.Pincode=pincode
	poolWait.Interest=interest
	poolWait.CreatedAt=createdAt
	poolWait.Gender=gender
	poolWait.Sex=sex
	_=db.Create(&poolWait).Error
}
func DeleteWaitPool(userId string) {
	db:=database.GetDB()
	db.Where("user_id=?",userId).Delete(&PoolWaiting{})
}

func DeleteWaitGroupPool(userId string) {
	db:=database.GetDB()
	db.Where("user_id=?",userId).Delete(&GroupPoolWaiting{})
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

func DeleteGroupPool(userId string,areaCode string,gender string){
	db:=database.GetDB()
	pool:=Pool{}
	db.Where("user_id=?",userId).Find(&pool)
	if(pool.UserId!=""){
		status:=EnterInGroupPoolFromWait(areaCode,gender,1)
		Complementary_gender:="Male"
		if(gender=="Male"){
			Complementary_gender="Female"
		}
		if(status=="NA"){
			DeleteFromGroupPoolHelper(areaCode,Complementary_gender,1)
		}
	}
}
func EnterInPoolFromWait(areaCode string,gender string,number_of_person int) string{
	return "Ok"
	//Code for checking if user is available to replace the person who cancelled the pooling
	//and replace the person
}

func DeleteFromPoolHelper(areaCode string,gender string,number_of_person int){
	//Delete a user from the pool table and put him/her back to wait pool table
}

func EnterInGroupPoolFromWait(areaCode string,gender string,number_of_person int) string{
	return "Ok"
}

func DeleteFromGroupPoolHelper(areaCode string,gender string,number_of_person int){

}