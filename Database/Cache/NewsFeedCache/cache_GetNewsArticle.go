package NewsFeedCache

import(
	gocache "github.com/patrickmn/go-cache"
	"fmt"
	"time"
)

const(
	NEWSFEED_EXPIRY_TIME=1440	
	NEWSFEED_PURGE_TIME=1560
)

var cache *gocache.Cache

func init(){
	cache=gocache.New(NEWSFEED_EXPIRY_TIME*time.Minute,NEWSFEED_PURGE_TIME*time.Minute)
	fmt.Println("GetNewsArticle Cache Initialised")
}

func GetNewsArticleCache() *gocache.Cache{
	return cache
}