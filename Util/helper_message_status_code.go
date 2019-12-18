package Util


func GetMessageDecode(status_code int) string{
	message_decode:=make(map[int]string)

	message_decode[200]="Successfull"

	//MESSAGE STATUS CODE FOR LOADING PAGE
	message_decode[300]="User already logded in"
	message_decode[2000]="New or Logged out user"
	message_decode[2001]="User not verified"
	message_decode[2002]="Profile Not Created"
	message_decode[2003]="Preference not completed"
	









	message_decode[1000]="Could not read  Body"
	message_decode[1001]="Could not Unmarshal  Data"
	message_decode[1002]="User Data Invalid"
	message_decode[1003]="Session does not exist"
	message_decode[1004]="User already verified"
	message_decode[1005]="User is not verified"
	message_decode[1006]="Database Problem"
	message_decode[1007]="InvalidURL"
	message_decode[1008]="Profile not created"

	//MESSAGE STATUS CODE FOR REGISTER
	message_decode[1101]="User Already Exists"

	//MESSAGE STATUS CODE FOR GENERATE EMAIL VERIFICATION
	message_decode[1201]="Email id does not exist"
	message_decode[1202]="Link sent more than limit"

	//MESSAGE STATUS CODE FOR GENERATE OTP
	message_decode[1301]="Mobile no does not exist"
	message_decode[1302]="OTP sent more than limit"

	//MESSAGE STATUS CODE FOR VERIFY EMAIL


	//MESSAGE STATUS CODE FOR VERIFY OTP
	message_decode[1401]="Wrong OTP"

	//MESSAGE STATUS CODE FOR LOGIN
	message_decode[1502]="Invalid User Id or Password"
	message_decode[1501]="No such User"

	//MESSAGE STATUS CODE FOR PROFILE CREATION
	message_decode[2001]="Cannot Access this Profile"
	//MESSAGE STATUS CODE FOR LOGOUT	

	//MESSAGE STATUS CODE FOR GETCHATDETAIL
	message_decode[7000]="Error in retrieving database"

	//MESSAGE STATUS CODE FOR GETEVENTBYID
	message_decode[8000]="Error in retrieving database"
	message_decode[8001]="Event Id does not exist"
	
	return message_decode[status_code]
}