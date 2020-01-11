package Social


type PoolStatusHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

type GetInPoolHeader struct{
	Cookie string `header:"Miti-Cookie"`	
}