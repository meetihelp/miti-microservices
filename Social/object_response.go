package Social

type PoolStatusResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	MatchUserId string `json:"MatchUserId"`
	Status string `json:"Status"`
	CreatedAt string `json:"CreatedAt"`
	MatchTime string `json:"MatchTime"`
	IPIP int `json:"IPIP"`
}

type GetInPoolResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	IPIP int `json:"IPIP"`
	RequestId string `json:"RequestId"`
	PoolStatus PoolStatus `json:"PoolStatus"`
}

type GetInGroupPoolResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	CreatedAt string `json:"CreatedAt"`
	RequestId string `json:"RequestId"`
}

type CancelGroupPoolResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	RequestId string `json:"RequestId"`
	Interest string `json:"Interest"`	
}