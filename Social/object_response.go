package Social

type PoolStatusResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	ChatId string `json:"ChatId"`
	MatchUserId string `json:"MatchUserId"`
	Status string `json:"Status"`
	CreatedAt string `json:"CreatedAt"`
	MatchTime string `json:"MatchTime"`
	IPIP int `json:"IPIP"`
}

type GroupPoolStatusResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	// ChatId string `json:"ChatId"`
	Interest []string `json:"Interest"`
	Status []GroupPoolStatusHelper `json:"Status"`
	// CreatedAt string `json:"CreatedAt"`
	// MatchTime string `json:"MatchTime"`
	IPIP int `json:"IPIP"`
}

type GetInPoolResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	IPIP int `json:"IPIP"`
	// RequestId string `json:"RequestId"`
	PoolStatus PoolStatusHelper `json:"PoolStatus"`
}

type PoolStatusHelper struct{
	ChatId string `json:"ChatId"`
	MatchUserId string `json:"MatchUserId"`
	Status string `json:"Status"`
}

type GetInGroupPoolResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	Interest string `json:"Interest"`
	Status GroupPoolStatusHelper `json:"Status"`
	CreatedAt string `json:"CreatedAt"`
	RequestId string `json:"RequestId"`
}

type GroupPoolStatusHelper struct{
	ChatId string `json:"ChatId"`
	Status string `json:"Status"`
	Interest string `json:"Interest"`
	CreatedAt string `json:"CreatedAt"`
}

type CancelGroupPoolResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	RequestId string `json:"RequestId"`
	Interest string `json:"Interest"`	
}