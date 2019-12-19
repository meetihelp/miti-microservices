package NewsFeed

type GetNewsFeedSummaryHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

type GetNewsFeedSummaryDS struct{
	NewsFeedId string `json:"NewsFeedId"`
}

type GetNewsFeedArticleHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

type GetNewsFeedArticleDS struct{
	NewsFeedId string `json:"NewsFeedId"`
}