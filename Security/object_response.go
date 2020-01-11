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