package Authentication

type PasswordChange struct{
	OldPassword string `validate:"required" json:"OldPassword"`
	NewPassword string `validate:"required" json:"NewPassword"`
}

type ForgetPasswordDS struct{
	Phone string `validate:"required" json:"Phone"`
}

type UpdateForgetPasswordDS struct{
	Password string `validate:"required" json:"Password"`
}

type SMSResponseCode struct{
	RequestID string `json:"request_id"`
	Type string `json:"type"`
}