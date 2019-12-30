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

type NewsFeedReactionHeader struct{
	Cookie string `header:"Miti-Cookie"`
}

type NewsFeedReactionDS struct{
	NewsFeedId string `json:"NewsFeedId"`
	Reaction string `json:"Reaction"`
}

type GetNewsArticleDS struct{
	Id int64 `json:"Id"`
}