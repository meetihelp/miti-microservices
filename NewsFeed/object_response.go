package NewsFeed

//Get News Article
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
	ArticleURL string `json:"ArticeURL"`
}

type GetNewsFeedArticleResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	NewsData []NewsResponse `json:"NewsData"`
}

type GetNewsFeedArticleResponseHeader struct{
	ContentType string `json:"Content-Type"`
}










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


type NewsFeedArticleResponse struct{
	Id int64 `json:"Id`
	Article string `json:"Article"`
}

type UploadNewsFeedImageResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	ImageURL string `json:"ImageURL"`
}