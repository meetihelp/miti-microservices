package Authentication

import(
	"net/http"
	// "io/ioutil"
	// "encoding/json"	
	// util "miti-microservices/Util"
	"github.com/jinzhu/gorm"
	sms "miti-microservices/Notification/SMS"
	// "log"
	"time"
	// "reflect"
	// "fmt"
)

const (
	MAXCOUNT = 5
	MAXMINUTE = 10
	MAXFAILCOUNT=5
	MAXRESENDCOUNT=5
	ONEDAY=1440
	NUM_OF_PREFERENCE=6
)

// func SendPreference(w http.ResponseWriter,preferenceCreationStatus int,code int){
// 	w.Header().Set("Content-Type", "application/json")
// 	msg:=util.GetMessageDecode(code)
// 	p:=&PreferenceContent{Code:code,Message:msg,Preference:preferenceCreationStatus}
// 	enc := json.NewEncoder(w)
// 	err:= enc.Encode(p)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }


func SendOTP(phone string,otp string)(*http.Response,error){
	return sms.SendSMS(phone,otp)
}

func OTPHelper(db *gorm.DB,userId string) (int,bool){
	otp,dbError:=GetOTPDetails(db,userId)
	duration:=CalculateDuration(otp.CreatedAt)
	if(duration>ONEDAY && !dbError){
		dbError=DeleteOTP(db,userId)
		if(!dbError){
			return 200,dbError	
		}
	}
	if(otp.FailCount>=MAXFAILCOUNT && !dbError){
		return 3000,dbError
	}
	if(otp.ResendCount>MAXRESENDCOUNT && !dbError){
		return 3001,dbError
	}
	
	// deliveryCount:=otp.DeliverCount
	// if((duration<MAXMINUTE || deliveryCount==0) && !dbError){
	// 	return 200,dbError
	// }
	if((duration<MAXMINUTE) && !dbError){
		return 1003,dbError
	}
	if(duration>MAXMINUTE && !dbError){
  		return 200,dbError
	}
	return 1003,dbError
}

func CalculateDuration(lastModified string) int{
	layout:="2006-01-02 15:04:05"
	t,_:=time.Parse(layout,lastModified)
	now:=time.Now()
	now,_=time.Parse(layout,now.Format("2006-01-02 15:04:05"))
	duration:=now.Sub(t)
	h:=duration.Minutes()
	h_int:=int(h)
	return h_int
}

// func SendTemporaryIdList(temporaryIdList TempUserList){
	
// }


