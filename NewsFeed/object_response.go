package NewsFeed

type SummaryResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	NewsFeedId string `json:"NewsFeedId"`
	Summary string `json:"Summary"`
}

// type ArticleResponse struct{
// 	Code int `json:"Code"`
// 	Message string `json:"Message"`
// 	NewsFeedId string `json:"NewsFeedId"`
// 	Article string `json:"Article"`
// }

type ArticleResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	Response NewsFeedArticleResponse `json:"Response"`
}
type NewsArticleResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	NewsData []NewsResponse `json:"NewsData"`
}

type NewsResponse struct{
	Id int64 `json:"Id"`
	Summary string `json:"Summary"`
	Sentiment string `json:"Sentiment"`
	Location string `json:"Location"`
	Event string `json:"Event"`
	Label string `json:"Label"`
	Title string `json:"Title"`
	ImageURL string `json:"ImageURL"`
	Flag string `json:"Flag"`
}

type NewsFeedArticleResponse struct{
	Id int64 `json:"Id`
	Article string `json:"Article"`
}