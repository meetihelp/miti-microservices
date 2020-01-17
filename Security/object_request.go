package Security

type CreatePrimaryTrustChainHeader struct{
	Cookie string `header:"Miti-Cookie"`	
}

type CreatePrimaryTrustChainRequest struct{
	Phone string `json:"Phone"`
	ChainId string `json:"ChainId"`
	Name string `json:"Name"`
	RequestId string `json:"RequestId"`
}

type DeletePrimaryTrustChainHeader struct{
	Cookie string `header:"Miti-Cookie"`	
}

type DeletePrimaryTrustChainRequest struct{
	Phone string `json:"Phone"`
	ChainId string `json:"ChainId"`
	RequestId string `json:"RequestId"`
	Name string `json:"Name"`
}

type CreateSecondaryTrustChainHeader struct{
	Cookie string `header:"Miti-Cookie"`	
}

type CreateSecondaryTrustChainRequest struct{
	RequestId string `json:"RequestId"`
	ChatId string `header:"ChatId"`
}


type DeleteSecondaryTrustChainHeader struct{
	Cookie string `header:"Miti-Cookie"`	
}

type DeleteSecondaryTrustChainRequest struct{
	ChatId string `json:"ChatId"`
	RequestId string `json:"RequestId"`
}

type AlertMessageHeader struct{
	Cookie string `header:"Miti-Cookie"`	
}

type AlertMessageRequest struct{
	Latitude string `json:"Latitude"`
	Longitude string `json:"Longitude"`
}