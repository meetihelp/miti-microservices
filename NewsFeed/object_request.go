package NewsFeed


//GetNewsFeedArticle
type GetNewsFeedArticleHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

type GetNewsArticleDS struct{
	// Id int64 `json:"Id"`
	Label string `json:"Label"`
}









type GetNewsFeedSummaryHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

type GetNewsFeedSummaryDS struct{
	NewsFeedId string `json:"NewsFeedId"`
}



type GetNewsFeedArticleDS struct{
	NewsFeedId string `json:"NewsFeedId"`
}

type NewsFeedReactionHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

type NewsFeedReactionDS struct{
	Id int64 `json:"Id"`
	Reaction string `json:"Reaction"`
}



type UploadNewsFeedImageHeader struct{
	Cookie string `header:"Miti-Cookie"`	
}

type UploadNewsFeedImageRequest struct{
	URL string `json:"URL"`
	Format string `json:"Format"`
}
