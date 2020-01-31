package Profile

//Profile Creation
type ProfileCreationResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
}

type ProfileCreationResponseHeader struct{
	ContentType string `json:"Content-Type"`
}

//Update IPIP 
type UpdateIPIPResponseContent struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
}

type UpdateIPIPResponseContentHeader struct{
	ContentType string `json:"Content-Type"`	
}

//Update Preference
type UpdatePreferenceResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
}
type UpdatePreferenceResponseHeader struct{
	ContentType string `json:"Content-Type"`		
}

//Get Profile
type ProfileResponse struct{
	UserId string `son:"UserId"`
	Name string `json:"Name"`
	DateOfBirth string `json:"DateOfBirth"`
	Job string `json:"Job"`
	ProfilePicID string `json:"ProfilePicID"`
	// ProfilePicURL string `json:"ProfilePicURL"`
	Gender string `json:"Gender"`
	Language string `json:"Language"`
	Country string `json:"Country"`
	Sex string `json:"Sex"`
	RelationshipStatus string `json:"RelationshipStatus"`
	InterestIndoorPassive1 string `json:"InterestIndoorPassive1"`
	InterestIndoorPassive2 string `json:"InterestIndoorPassive2"`
	InterestIndoorActive1 string `json:"InterestIndoorActive1"`
	InterestIndoorActive2 string `json:"InterestIndoorActive2"`	
	InterestOutdoorPassive1 string `json:"InterestOutdoorPassive1"`
	InterestOutdoorPassive2 string `json:"InterestOutdoorPassive2"`
	InterestOutdoorActive1 string `json:"InterestOutdoorActive1"`
	InterestOutdoorActive2 string `json:"InterestOutdoorActive2"`
	InterestOthers1 string `json:"InterestOthers1"`
	InterestOthers2 string `json:"InterestOthers2"`
	InterestIdeology1 string `json:"InterestIdeology1"`
	InterestIdeology2 string `json:"InterestIdeology2"`
	// ParentsAddress string `gorm:"type:varchar(10)"  json:"ParentsAddress"`
}

type GetProfileResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	ProfileResponse ProfileResponse `json:"ProfileResponse"` 
}

type GetProfileResponseHeader struct{
	ContentType string `json:"Content-Type"`
}

//Last

type SendQuestionContent struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	Question []Question `json:"Question"`
}


type CreateStatusResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	RequestId string `json:"RequestId"`
	CreatedAt string `json:"CreatedAt"`
}

type GetStatusResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	ChatId string `json:"ChatId"`
	StatusList []StatusResponse `json:"StatusList"`
}
type StatusResponse struct{
	UserId string `json:"UserId"`
	StatusContent string `json:"Status"`
	CreatedAt string `json:"CreatedAt"`
}

type GetCheckInterestResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	List []GetCheckInterestList `json:"List"`
}

type GetCheckInterestList struct{
	UserId string `json:"UserId"`
	Interest string `json:"Interest"`
	CreatedAt string `json:"CreatedAt"`
}

type CreatePrimaryTrustChainResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	UpdatedAt string `json:"UpdatedAt"`
	RequestId string `json:"RequestId"`
}

type DeletePrimaryTrustChainResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	UpdatedAt string `json:"UpdatedAt"`
	RequestId string `json:"RequestId"`
}

type CreateSecondaryTrustChainResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	CreatedAt string `json:"CreatedAt"`
	RequestId string `json:"RequestId"`
}

type DeleteSecondaryTrustChainResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	RequestId string `json:"RequestId"`
}