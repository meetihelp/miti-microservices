package Util

type ErrorList struct{
	SessionError bool
	TemporarySessionError bool
	BodyReadError bool
	UnmarshallingError bool
	SanatizationError bool
	DatabaseError bool
	LogicError bool
}

func GetErrorList(list []bool) ErrorList{
	errorList:=ErrorList{}
	errorList.SessionError=list[0]
	errorList.TemporarySessionError=list[1]
	errorList.BodyReadError=list[2]
	errorList.UnmarshallingError=list[3]
	errorList.SanatizationError=list[4]
	errorList.DatabaseError=list[5]
	errorList.LogicError=false
	return errorList
}

func GetCode(errorList ErrorList) int{
	if(errorList.DatabaseError){
		return 1006
	}

	if(errorList.SessionError){
		return 1003
	}

	if(errorList.TemporarySessionError){
		return 1003
	}

	if(errorList.BodyReadError){
		return 1002
	}	

	if(errorList.UnmarshallingError){
		return 1001
	}
	if(errorList.SanatizationError){
		return 1002
	}
	
	return 200
}

func ErrorListStatus(errorList ErrorList) bool{
	status:=false
	status=status || errorList.SessionError
	status=status || errorList.TemporarySessionError
	status=status || errorList.BodyReadError
	status=status || errorList.UnmarshallingError
	status=status || errorList.SanatizationError
	status=status || errorList.DatabaseError
	status=status || errorList.LogicError
	return status
}