package NewsFeed

type SummaryResponse struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	NewsFeedId string `json:"NewsFeedId"`
	Summary string `json:"Summary"`
}