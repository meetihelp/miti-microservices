package Social

//Pool Status
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

type PoolStatusResponseHeader struct{
	ContentType string `json:"Content-Type"`
}

type GetInPoolResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	IPIP int `json:"IPIP"`
	PoolStatus PoolStatusHelper `json:"PoolStatus"`
}

type GetInPoolResponseHeader struct{
	ContentType string `json:"Content-Type"`
}

type GroupPoolStatusResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	Interest []string `json:"Interest"`
	Status []GroupPoolStatusHelper `json:"Status"`
} 
type GroupPoolStatusResponseHeader struct{
	ContentType string `json:"Content-Type"`
}

type GetInGroupPoolResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	Interest string `json:"Interest"`
	Status GroupPoolStatusHelper `json:"Status"`
	CreatedAt string `json:"CreatedAt"`
	RequestId string `json:"RequestId"`
}

type GetInGroupPoolResponseHeader struct{
	ContentType string `json:"Content-Type"`
}



//Last





type PoolStatusHelper struct{
	ChatId string `json:"ChatId"`
	MatchUserId string `json:"MatchUserId"`
	Status string `json:"Status"`
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