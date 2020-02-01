package Social

//Pool Status
type PoolStatusHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

type PoolStatusRequest struct{
	Latitude string `json:"Latitude"`
	Longitude string `json:"Longitude"`
}


//Get In Pool
type GetInPoolHeader struct{
	Cookie string `header:"Miti-Cookie"`	
}

//Group Pool Status
type GroupPoolStatusHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

type GroupPoolStatusRequest struct{
	Latitude string `json:"Latitude"`
	Longitude string `json:"Longitude"`
}

//Last








// type GetInPoolRequest struct{
// 	RequestId string `json:"RequestId"`
// }

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