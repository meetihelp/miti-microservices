package Utility


func get_message_decode(status_code int) string{
	message_decode:=make(map[int]string)

	message_decode[100]="Could not read Register Body"
	message_decode[101]="Could not Unmarshal User Data"
	message_decode[102]="User Data Invalid"
	message_decode[103]="User Already Exists"
	message_decode[104]="User Registeres SuccessFully"

	message_decode[200]="Session does not exist"
	message_decode[201]="Verification Email Sent"
	message_decode[202]="Verification OTP sent"

	message_decode[300]="Could not read profile Creation Body"
	message_decode[301]="Could not Unmarshal profile creation data"
	message_decode[302]="Profile Creation Data Invalid"
	message_decode[303]="Profile Created SuccessFully"

	
	return message_decode[status_code]
}