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
	message_decode[304]="Session does not exist"

	message_decode[400]="Could not read Login Body"
	message_decode[401]="Could not Unmarshall user data"
	message_decode[402]="User data invalid"
	message_decode[403]="Invalid user id or password"
	message_decode[404]="User is not verified"
	message_decode[405]="User Login SuccessFully"
	
	return message_decode[status_code]
}