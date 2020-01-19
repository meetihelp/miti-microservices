package Security

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

type AlertMessageResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
}

type GetPrimaryTrustChainResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	ChainName string `json:"ChainName"`
	ChainId string `json:"ChainId"`
	UpdatedAt string `json:"UpdatedAt"`
	PrimaryTrustChainList []PrimaryTrustChainList `json:"PrimaryTrustChainList"`
}

type PrimaryTrustChainList struct{
	Phone string `json:"Phone"`
	Name string `json:"Name"`
}