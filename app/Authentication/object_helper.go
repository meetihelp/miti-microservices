package Authentication

type PasswordChange struct{
	OldPassword string `validate:"required" json:"OldPassword"`
	NewPassword string `validate:"required" json:"NewPassword"`
}