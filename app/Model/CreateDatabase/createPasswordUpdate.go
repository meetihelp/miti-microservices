package CreateDatabase

import()

type Password_change struct{
	Old_Password string `json:"old_password" validate:"required"`
	New_Password string `json:"new_password" validate:"required"`
}