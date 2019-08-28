package Utility


func get_message_decode(status_code int) string{
	message_decode:=make(map[int]string)

	message_decode[200]="Successfull"

	message_decode[1000]="Could not read  Body"
	message_decode[1001]="Could not Unmarshal  Data"
	message_decode[1002]="User Data Invalid"
	message_decode[1003]="Session does not exist"
	message_decode[1004]="User already verified"
	message_decode[1005]="User is not verified"

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
	message_decode[1501]="Invalid user id or password"

	//MESSAGE STATUS CODE FOR PROFILE CREATION

	//MESSAGE STATUS CODE FOR LOGOUT	

	//MESSAGE STATUS CODE FOR GETCHATDETAIL
	message_decode[7000]="Error in retrieving database"
	
	return message_decode[status_code]
}