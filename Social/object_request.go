package Social


type PoolStatusHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

type GetInPoolHeader struct{
	Cookie string `header:"Miti-Cookie"`	
}
// type GetInPoolRequest struct{
// 	RequestId string `json:"RequestId"`
// }

type GroupPoolStatusHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

type GroupPoolStatusRequest struct{
	Interest string `json:"Interest"`
}
type GetInGroupPoolHeader struct{
	Cookie string `header:"Miti-Cookie"`	
}

type GetInGroupPoolRequest struct{
	RequestId string `json:"RequestId"`
	Interest string `json:"Interest"`
}

type CancelPoolHeader struct{
	Cookie string `header:"Miti-Cookie"`	
}

type CancelGroupPoolHeader struct{
	Cookie string `header:"Miti-Cookie"`	
}

type CancelGroupPoolRequest struct{
	RequestId string `json:"RequestId"`
	Interest string `json:"Interest"`
}