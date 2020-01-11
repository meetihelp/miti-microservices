package Social

type PoolStatusResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	MatchUserId string `json:"MatchUserId"`
	Status string `json:"Status"`
	CreatedAt string `json:"CreatedAt"`
	MatchTime string `json:"MatchTime"`
}