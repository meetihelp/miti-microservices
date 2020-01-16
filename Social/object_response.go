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
	ChatId string `json:"ChatId"`
	Interest string `json:"Interest"`
	Status string `json:"Status"`
	CreatedAt string `json:"CreatedAt"`
	MatchTime string `json:"MatchTime"`
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
	CreatedAt string `json:"CreatedAt"`
	Interest string `json:"Interest"`
}

type CancelGroupPoolResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	RequestId string `json:"RequestId"`
	Interest string `json:"Interest"`	
}